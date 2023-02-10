package gzh

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"sort"
)

func CheckSign(signature, timestamp, nonce, token string) bool {
	arr := []string{timestamp, nonce, token}
	sort.Strings(arr)
	fmt.Println(arr)
	var sha1String string = ""
	for _, v := range arr {
		sha1String += v
	}
	//sha1String := strings.Join(arr, "")
	h := sha1.New()
	h.Write([]byte(sha1String))
	sha1String = hex.EncodeToString(h.Sum([]byte("")))
	fmt.Println("sha1:", sha1String)
	return sha1String == signature
}
