package main

import (
	"testing"
)

func TestIpfs(t *testing.T) {
	c := NewInfruaIpfs("2HbuLC3ALCiiAxRyIKjWnPc10Bt", "4b9a85430f30e2cf21b6bea4f88e33a5")
	// s := map[string]string{
	// 	"aa": "cc",
	// 	"bb": "dd",
	// }
	// b, err := json.Marshal(s)
	// if err != nil {
	// 	fmt.Println("fail marshal")
	// 	panic(err)
	// }
	// c.UploadByteFile(b)

	c.UploadPathFile("./abc.tar")
}
