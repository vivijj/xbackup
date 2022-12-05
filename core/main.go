package main

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/umbracle/ethgo"
	"github.com/umbracle/ethgo/contract"
	"github.com/umbracle/ethgo/jsonrpc"
	"github.com/umbracle/ethgo/wallet"
	"github.com/urfave/cli/v2"
)

// TODO: move to config
const baseDir = "../tmp"

var (
	infruaId     string
	infruaSecret string
	httpProxy    string
	contractAddr string
	prvKey       string
	web3Rpc      string
)

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	infruaId = os.Getenv("INFRUA_ID")
	infruaSecret = os.Getenv("INFRUA_SECRET")
	httpProxy = os.Getenv("HTTP_PROXY")
	contractAddr = os.Getenv("CONTRACT_ADDR")
	prvKey = os.Getenv("PRV_KEY")
	web3Rpc = os.Getenv("WEB3_RPC")
}

var (
	metaDir    = filepath.Join(baseDir, "meta")
	dataDir    = filepath.Join(baseDir, "data")
	archiveDir = filepath.Join(baseDir, "archive")
	picDir     = filepath.Join(baseDir, "pic")
)

func PrepareDir() {
	os.MkdirAll(metaDir, 0777)
	os.MkdirAll(dataDir, 0777)
	os.MkdirAll(archiveDir, 0777)
	os.MkdirAll(picDir, 0777)
}

func main() {
	loadEnv()
	PrepareDir()
	app := &cli.App{
		Name:  "xbackup",
		Usage: "backup the cortex model and input data",
		Commands: []*cli.Command{
			{
				Name:    "download",
				Aliases: []string{"d"},
				Usage:   "update the data meta info, download and archive the data",
				Action: func(cCtx *cli.Context) error {
					DownloadData()
					return nil
				},
			},
			{
				Name:    "genpic",
				Aliases: []string{"g"},
				Usage:   "genrate picture for data with infohash by stable diffusion v2",
				Action: func(cCtx *cli.Context) error {
					GeneratePicForData()
					return nil
				},
			},
			{
				Name:    "upload",
				Aliases: []string{"u"},
				Usage:   "upload the archive data and pic to platform(e.g. ipfs...)",
				Action: func(cCtx *cli.Context) error {
					UploadDataAndPicIpfs()
					return nil
				},
			},
			{
				Name:    "mint",
				Aliases: []string{"m"},
				Usage:   "upload the nft meta to ipfs and mint the nft token",
				Action: func(cCtx *cli.Context) error {
					MintNft()
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

// DownloadData will get the newlist model and input, and download the
// data that are not exist in the data dir and archive them.
func DownloadData() {
	res := UpdateMetaInfo(metaDir)
	if len(res) == 0 {
		fmt.Println("No new model data for download!!!")
		return
	}
	for _, v := range res {
		fmt.Println("Get data for ", v.Type, "of ih: ", v.InfoHash)
		if isIhGood := GetData(v.InfoHash, v.Type, dataDir); isIhGood {
			CreateTarFile(archiveDir, filepath.Join(dataDir, v.InfoHash))
		}
	}
	fmt.Printf("Finish download and archive %v data./n", len(res))
}

// GeneratePicForData will call python code of stable diffusion to generate the result
func GeneratePicForData() {
	ihs := ListInfoHash(metaDir)
	for _, ih := range ihs {
		if _, err := os.Stat(filepath.Join(picDir, ih+".png")); err != nil {
			fmt.Println("genrate pic for: ", ih)
			args := []string{
				"../stablediffusion/scripts/txt2img.py",
				"--prompt",
				ih,
				"--ckpt",
				"../stablediffusion/768-v-ema.ckpt",
				"--config",
				"../stablediffusion/configs/stable-diffusion/v2-inference-v.yaml",
				"--outdir",
				picDir,
			}
			cmd := exec.Command("python", args...)
			var stdout, stderr bytes.Buffer
			cmd.Stdout = &stdout
			cmd.Stderr = &stderr
			err := cmd.Run()
			if err != nil {
				log.Fatalf("cmd.Run() failed with %s\n", err)
			}
			outStr, errStr := stdout.String(), stdout.String()
			fmt.Printf("out:\n%s\nerr:\n%s\n", outStr, errStr)
			fmt.Println("finish generate pic!!!")
		}
	}
}

func UploadDataAndPicIpfs() {
	infrua := NewInfruaIpfs(infruaId, infruaSecret, httpProxy)
	ihs := ListInfoHash(metaDir)
	for _, v := range ihs {
		shouldUpdate := false
		meta := LoadMetadata(v, metaDir)
		if meta.ExternalUrl == nil {
			meta.ExternalUrl = make(map[string]string)
		}
		if _, ok := meta.ExternalUrl["ipfs"]; !ok {
			fmt.Println("Upload ", v, "to ipfs")
			cid, err := infrua.UploadPathFile(filepath.Join(archiveDir, v+".tar"))
			if err != nil {
				panic(err)
			}
			meta.ExternalUrl["ipfs"] = IpfsAccessUrl(cid)
			shouldUpdate = true
		}
		if _, ok := meta.ExternalUrl["hfs"]; !ok {
			meta.ExternalUrl["hfs"] = hfsAccessUrl(v)
			shouldUpdate = true
		}
		if meta.Image == "" {
			cid, err := infrua.UploadPathFile(filepath.Join(picDir, v+".png"))
			if err != nil {
				panic(err)
			}
			meta.Image = IpfsAccessUrl(cid)
			shouldUpdate = true
		}
		if shouldUpdate {
			UpdateMetadata(v, meta, metaDir)
		}

	}

}

func MintNft() {
	infruac := NewInfruaIpfs(infruaId, infruaSecret, httpProxy)
	client, err := jsonrpc.NewClient(web3Rpc)
	if err != nil {
		panic(err)
	}
	prvKeyBytes, err := hex.DecodeString(prvKey)
	if err != nil {
		panic(err)
	}
	accountKey, err := wallet.NewWalletFromPrivKey([]byte(prvKeyBytes))
	if err != nil {
		panic(err)
	}
	xContract := NewXbackup(ethgo.HexToAddress(contractAddr), contract.WithJsonRPC(client.Eth()), contract.WithSender(accountKey))

	ihs := ListInfoHash(metaDir)
	for _, v := range ihs {
		fmt.Println("start deal with the ih: ", v)
		meta := LoadMetadata(v, metaDir)
		res, err := xContract.Infohash2tokenid(v)
		if err != nil {
			panic(err)
		}
		if res.BitLen() == 0 {
			fmt.Println("This model isn't minted yet, minting...")
			content, err := json.Marshal(meta)
			if err != nil {
				panic(err)
			}
			cid, err := infruac.UploadByteFile(content)
			fmt.Println("finish upload nft meta json!!!")
			metaUrl := IpfsAccessUrl(cid)
			if err != nil {
				panic(err)
			}
			urlJson, err := json.Marshal(meta.ExternalUrl)
			if err != nil {
				fmt.Println("marshal json fail")
				panic(err)
			}
			txn, err := xContract.MintTokenForDataAuthor(v, metaUrl, string(urlJson), ethgo.HexToAddress(meta.Author))
			if err != nil {
				panic(err)
			}
			err = txn.Do()
			if err != nil {
				panic(err)
			}
			fmt.Println("wait for mint finish ...")
			txn.Wait()
			fmt.Println("finish mint for ih: ", v)
		} else {
			fmt.Println("this data already have nft.")
		}
	}

}
