package main

import (
	// "ecommerce/cmd"
	// "bytes"
	// "encoding/base64"
	"encoding/base64"
	"fmt"
)

type Shishir int

func main() {

	var s string 

	s = "b"

	byteArr := []byte(s)

	fmt.Println(byteArr)
	fmt.Println(s)
	
	enc := base64.URLEncoding

	enc = enc.WithPadding(base64.NoPadding)

	b64Str := enc.EncodeToString(byteArr)

	fmt.Println(b64Str)

	// base64.URLEncoding.WithPadding(base64.NoPadding)


	// cmd.Serv()
}
