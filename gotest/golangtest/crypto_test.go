package golangtest

import (
	"fmt"
	"testing"

	"golang.org/x/crypto/cryptobyte"
)

func TestCryptoByte(t *testing.T) {
	// Writing more data that can be expressed by the length prefix results
	// in an error from Bytes().

	tooLarge := make([]byte, 256)

	var b cryptobyte.Builder
	b.AddUint8LengthPrefixed(func(b *cryptobyte.Builder) {
		b.AddBytes(tooLarge)
	})

	result, err := b.Bytes()
	fmt.Printf("len=%d err=%s\n", len(result), err)
}
