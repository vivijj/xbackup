package main

import (
	"fmt"
	"net/http"
	"testing"
)

func TestGetModelMeta(t *testing.T) {
	res := UpdateMetaInfo("../tmp/meta")
	fmt.Println("len is:", len(res))
	a := 0
	for _, v := range res {
		resp, err := http.Get("http://share.cortexlabs.ai:7881/storage/" + v.InfoHash)
		if err != nil {
			fmt.Println(err)
		}

		if resp.StatusCode == http.StatusNotFound {
			fmt.Println("unhealthy hash: ", v)
			continue
		} else {
			a += 1
		}
	}
	fmt.Println("healthy num: ", a)
}
