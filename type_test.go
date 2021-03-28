package sign

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestType(t *testing.T) {
	msg := []byte("abc")
	fmt.Printf("16进制:%X\n", msg)
	fmt.Println(msg)
	msg2, _ := hex.DecodeString("abcd")

	fmt.Printf("16进制:%X\n", msg2)
	fmt.Println(msg2)
}

//
