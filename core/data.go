package main

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
)

const (
	defaultDataUri = "http://share.cortexlabs.ai:7881/storage/"
)

var exist = struct{}{}

// badFile include some infohash that cant't get the data of it now.
// (no avaiable peer have these data now)
var badFile = map[string]struct{}{
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

func GetData(ih string, dataType string, baseDir string) bool {
	if _, ok := badFile[ih]; ok {
		return false
	}
	if dataType == "model" {
		GetModelData(ih, baseDir)
	} else {
		GetInputData(ih, baseDir)
	}
	return true
}

func GetModelData(ih string, baseDir string) {
	dataPath := filepath.Join(baseDir, ih, "data")
	os.MkdirAll(dataPath, 0777)
	paramUrl := defaultDataUri + ih + "/data/params"
	paramDir := filepath.Join(dataPath, "params")
	symbolUrl := defaultDataUri + ih + "/data/symbol"
	symbolDir := filepath.Join(dataPath, "symbol")
	storeDataFromWeb(paramUrl, paramDir)
	storeDataFromWeb(symbolUrl, symbolDir)
}

func GetInputData(ih string, baseDir string) {
	dataPath := filepath.Join(baseDir, ih)
	os.MkdirAll(dataPath, 0777)
	dataUrl := defaultDataUri + ih + "/data"
	dataDir := filepath.Join(dataPath, "data")
	storeDataFromWeb(dataUrl, dataDir)
}

func storeDataFromWeb(dataUrl string, dataDir string) {
	resp, err := http.Get(dataUrl)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	f, err := os.Create(dataDir)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	io.Copy(f, resp.Body)
}
