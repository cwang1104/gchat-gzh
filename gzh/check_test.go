package gzh

import (
	"fmt"
	"testing"
)

const (
	signature1 = "7a91d8600d01a97748df5b6c712e4164809cd68b"
	signature2 = "e3dbd17619181b0ec7e238e814b7edb218654b6e"
	time       = "1675917956"
	nonce      = "1356339128"
	echostr    = "1435162263045896893"
	token      = "ddsagsddesdxzdf223"
)

func TestCheckSign(t *testing.T) {
	a := CheckSign(signature1, time, nonce, token)
	fmt.Println(a)
}
