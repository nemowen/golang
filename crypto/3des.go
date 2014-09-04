package main

import (
	"crypto/des"
	"encoding/base64"
	"fmt"
)

const encodeStd = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

var coder = base64.NewEncoding(encodeStd)

func tripleDESDecode(key, src string) string {
	if len(key) != 16 && len(key) != 24 {
		panic("key length error, must be 16 or 24")
	}

	data, err := coder.DecodeString(src)
	if err != nil {
		panic(err)
	}

	tripleDESKey := make([]byte, 0, 24)
	if len(key) == 16 {
		tripleDESKey = append(tripleDESKey, key[:16]...)
		tripleDESKey = append(tripleDESKey, key[:8]...)
	} else {
		tripleDESKey = append(tripleDESKey, key[:]...)
	}

	td, err := des.NewTripleDESCipher(tripleDESKey)
	if err != nil {
		panic(err)
	}

	n := len(data) / td.BlockSize()
	var rb []byte
	for i := 0; i < n; i++ {
		dst := make([]byte, td.BlockSize())
		td.Decrypt(dst, data[i*8:(i+1)*8])
		rb = append(rb, dst[:]...)
	}

	lastValue := int(rb[len(rb)-1])
	return string(rb[0 : len(rb)-lastValue])
}

func tripleDESEncode(key, src string) string {
	var result string

	if len(key) != 16 && len(key) != 24 {
		panic("key length error, must be 16 or 24")
	}

	tripleDESKey := make([]byte, 0, 24)
	if len(key) == 16 {
		tripleDESKey = append(tripleDESKey, key[:16]...)
		tripleDESKey = append(tripleDESKey, key[:8]...)
	} else {
		tripleDESKey = append(tripleDESKey, key[:]...)
	}

	td, err := des.NewTripleDESCipher(tripleDESKey)
	if err != nil {
		panic(err)
	}

	mod := len(src) % td.BlockSize()
	v := td.BlockSize() - mod

	data := []byte(src)
	for i := 0; i < v; i++ {
		data = append(data, byte(v))
	}

	n := len(data) / td.BlockSize()
	var rb []byte
	for i := 0; i < n; i++ {
		dst := make([]byte, td.BlockSize())
		td.Encrypt(dst, data[i*8:(i+1)*8])
		rb = append(rb, dst[:]...)
	}

	result = coder.EncodeToString(rb)
	return result
}

func main() {
	encodeStr := tripleDESEncode("a1c2e3g4i5k6m7o8q9s0u1w2", "1234567")
	fmt.Println("encode:", encodeStr)
	decodeStr := tripleDESDecode("a1c2e3g4i5k6m7o8q9s0u1w2", encodeStr)
	fmt.Println("decode:", decodeStr)

	encodeStr = tripleDESEncode("a1c2e3g4i5k6m7o8", "123456789")
	fmt.Println("encode:", encodeStr)
	decodeStr = tripleDESDecode("a1c2e3g4i5k6m7o8", encodeStr)
	fmt.Println("decode:", decodeStr)
}
