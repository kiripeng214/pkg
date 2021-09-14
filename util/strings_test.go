package util

import (
	"fmt"
	"testing"
)

func TestBytes2String(t *testing.T) {
	b := []byte{49, 44, 50, 44, 51}
	fmt.Println(Bytes2String(b))
}

func TestString2Bytes(t *testing.T) {
	fmt.Println(String2Bytes("1,2,3"))
}
