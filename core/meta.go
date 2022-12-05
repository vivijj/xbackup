package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
)

const (
	defaultQueryAPi = "https://cerebro.cortexlabs.ai/mysql"
	queryStep       = 100
)

var badInfoHash = map[string]struct{}{
	"3edcb8a793887d92db12d53124955681d5c20a43": exist,
	"533675955a632610e569fecf9b43b6a2c6a299aa": exist,
	"6b09ec425ce5bc796f06884df7662eeb4e34faf9": exist,
	"6dc44c788ab1e35c3949ad0880cd07cb18fd82aa": exist,
	"0fb3dedf8d753670e90f2dd7c37c77610d07b62d": exist,
	"e17538ab36aeeab524e8a8b085a38871e12c28f9": exist,
	"b1590ba813c125b47ccd03a34ea0908668bc709c": exist,
	"10a1347cff5e3e8174d6b074164d0f4787882738": exist,
	"7dbdd3bf5eb641e89071b37ff42b6c5a4464da34": exist,
	"934427da068cc446aa0f84d26fce8a15d3135e5b": exist,
	"7c6ff2ae3b9dfc654c5dec61cc950b70a16431c4": exist,
}

// DataMeta include basic info about model data or input data for
// us to process.
type DataMeta struct {
	InfoHash    string            `json:"info_hash"`
	Author      string            `json:"author"`
	Description string            `json:"description"`
	Type        string            `json:"type"`
	Image       string            `json:"image"`
	ExternalUrl map[string]string `json:"external_url"`
}

// both the
func ReadyForMint(meta DataMeta) bool {
	return meta.Image != "" && meta.ExternalUrl != nil
}

type ApiModel struct {
	Id   string
	Tag  string
	Meta struct {
		Hash    string
		Address string
		Type    string
	}
}

type ApiInput struct {
	Id   string
	Tag  string
	Meta struct {
		Hash    string
		Address string
		Type    string
	}
}

type CerebroApiModel struct {
	Length int
	Result []ApiModel
}

type CerebroApiInput struct {
	Length int
	Result []ApiInput
}

func ListInfoHash(metadir string) []string {
	metadataFiles, err := os.ReadDir(metadir)
	if err != nil {
		panic(err)
	}
	ihs := make([]string, 0, len(metadataFiles))
	for _, f := range metadataFiles {
		ih := strings.TrimSuffix(f.Name(), ".json")
		if _, ok := badInfoHash[ih]; !ok {
			ihs = append(ihs, ih)
		}
	}
	return ihs
}

func LoadMetadata(ih string, metadir string) (dm DataMeta) {
	metaFile := filepath.Join(metadir, ih+".json")
	content, err := os.ReadFile(metaFile)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(content, &dm)
	if err != nil {
		panic(err)
	}
	return
}

func UpdateMetadata(ih string, md DataMeta, metadir string) {
	content, err := json.Marshal(md)
	if err != nil {
		panic(err)
	}
	metaFile := filepath.Join(metadir, ih+".json")
	err = os.WriteFile(metaFile, content, 0666)
	if err != nil {
		panic(err)
	}
}

// UpdateMetaInfo return the meta of new data that we have not download yet.
func UpdateMetaInfo(metaDir string) []DataMeta {
	allMeta := GetAllInputData()
	allMeta = append(allMeta, GetAllModelMeta()...)
	res := make([]DataMeta, 0, len(allMeta))
	for _, v := range allMeta {
		metaFile := filepath.Join(metaDir, v.InfoHash+".json")
		if _, err := os.Stat(metaFile); err != nil {
			data, err := json.Marshal(v)
			if err != nil {
				panic(err)
			}
			os.WriteFile(metaFile, data, 0777)
			res = append(res, v)
		}
	}
	return res
}

// Note: we should seprete the getting of model data and input data,
// due to that the input data in cerebro absent for creator.

func GetAllModelMeta() []DataMeta {
	// var modelsRes []
	var metaRes []DataMeta
	for i := 0; ; i++ {
		respBody := doGetReq(map[string]string{
			"begin": strconv.Itoa(i * queryStep),
			"end":   strconv.Itoa((i + 1) * queryStep),
			"type":  "modules",
		})
		var res CerebroApiModel
		err := json.Unmarshal(respBody, &res)
		if err != nil {
			panic(err)
		}
		for _, v := range res.Result {
			metaRes = append(metaRes, DataMeta{
				InfoHash:    strings.TrimPrefix(v.Meta.Hash, "0x"),
				Author:      v.Meta.Address,
				Description: v.Tag,
				Type:        v.Meta.Type,
			})
		}
		// if the result is less than the step, there is no more data
		if len(res.Result) < queryStep {
			break
		}

	}
	return metaRes
}

func GetAllInputData() []DataMeta {
	wg := &sync.WaitGroup{}
	mux := &sync.Mutex{}
	var inputRes []DataMeta

	for i := 0; ; i++ {
		respBody := doGetReq(map[string]string{
			"begin": strconv.Itoa(i * queryStep),
			"end":   strconv.Itoa((i + 1) * queryStep),
			"type":  "inputs",
		})
		var res CerebroApiInput
		err := json.Unmarshal(respBody, &res)
		if err != nil {
			panic(err)
		}
		for _, v := range res.Result {
			wg.Add(1)
			go func(data ApiInput) {
				meta := DataMeta{
					InfoHash:    strings.TrimPrefix(data.Meta.Hash, "0x"),
					Author:      data.Meta.Address,
					Description: data.Tag,
					Type:        data.Meta.Type,
				}
				if meta.Author == "" {
					meta.Author = GetInputCreator(data.Id)
				}
				mux.Lock()
				inputRes = append(inputRes, meta)
				mux.Unlock()
				wg.Done()
			}(v)
		}
		wg.Wait()
		if len(res.Result) < queryStep {
			break
		}
	}
	return inputRes
}

func GetInputCreator(id string) string {
	// now, each input data address only have 2 tx now
	respBody := doGetReq(map[string]string{
		"addressId": id,
		"begin":     "0",
		"end":       "3",
		"type":      "addrTX",
	})
	var res struct {
		TxCount int
		Length  int
		Result  []struct{ From string }
	}
	err := json.Unmarshal(respBody, &res)
	if err != nil {
		panic(err)
	}
	return res.Result[len(res.Result)-1].From
}

func doGetReq(params map[string]string) []byte {
	req, err := http.NewRequest("GET", defaultQueryAPi, nil)
	if err != nil {
		log.Fatal(err)
	}
	q := req.URL.Query()
	for k, v := range params {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	body, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		panic(err)
	}

	return body

}
