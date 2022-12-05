package main

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	ipfsFiles "github.com/ipfs/go-ipfs-files"
	ipfsApi "github.com/ipfs/go-ipfs-http-client"
	"github.com/ipfs/interface-go-ipfs-core/options"
)

const infruaUrl = "https://ipfs.infura.io:5001"
const ipfsGateway = "https://cortex.infura-ipfs.io/ipfs/"

// get the access url
func IpfsAccessUrl(cid string) string {
	return ipfsGateway + cid
}

type InfruaIpfs struct {
	ipfsClient *ipfsApi.HttpApi
}

func NewInfruaIpfs(apiId string, apiSecret string, proxyUrl string) *InfruaIpfs {
	proxy, err := url.Parse(proxyUrl)
	if err != nil {
		fmt.Println("fail to parse the proxy url")
		panic(err)
	}
	tr := &http.Transport{
		Proxy:           http.ProxyURL(proxy),
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	httpClient := &http.Client{
		Transport: tr,
	}
	httpApi, err := ipfsApi.NewURLApiWithClient(infruaUrl, httpClient)
	if err != nil {
		log.Fatal(err)
	}
	httpApi.Headers.Add("Authorization", "Basic "+basicAuth(apiId, apiSecret))
	return &InfruaIpfs{
		httpApi,
	}
}

func (ii *InfruaIpfs) UploadByteFile(fb []byte) (string, error) {
	return ii.uploadAndPin(bytes.NewReader(fb))
}

func (ii *InfruaIpfs) UploadPathFile(fpath string) (string, error) {
	content, err := os.Open(fpath)
	defer content.Close()
	if err != nil {
		panic(err)
	}
	return ii.uploadAndPin(content)
}
func (ii *InfruaIpfs) uploadAndPin(c io.Reader) (string, error) {
	result, err := ii.ipfsClient.Unixfs().Add(context.Background(), ipfsFiles.NewReaderFile(c), options.Unixfs.Pin(true))
	if err != nil {
		return "", err
	}
	return result.Cid().String(), nil
}

func basicAuth(apiId, apiSecret string) string {
	auth := apiId + ":" + apiSecret
	return base64.StdEncoding.EncodeToString([]byte(auth))
}
