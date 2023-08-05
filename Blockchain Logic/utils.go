package main

import "fmt"

func IntToHex(n int64) []byte {
	hexString := fmt.Sprintf("%x", n)
	return []byte(hexString)
}
