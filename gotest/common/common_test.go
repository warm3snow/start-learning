package common

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"fmt"
	"runtime"
	"testing"
)

func TestOS(t *testing.T) {
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
}

func TestSliceLen(t *testing.T) {
	var a []int
	fmt.Println("len(a) = ", len(a))
	fmt.Printf("a = %v\n", a)
	if a == nil {
		fmt.Println("a is nil")
	}

	b := make([]byte, 0)
	fmt.Println("len(b) = ", len(b))
	fmt.Printf("b = %v\n", b)
	if b == nil {
		fmt.Println("b is nil")
	}
}

func TestHashNil(t *testing.T) {
	t.Log(sha256.Sum256(nil))
	t.Log(sha256.Sum256(make([]byte, 0)))
}

func TestAESNil(t *testing.T) {
	//origData := []byte("Hello World") // 待加密的数据
	//origData := make([]byte, 0)
	var origData []byte = nil
	key := []byte("ABCDEFGHIJKLMNOP") // 加密的密钥
	// 分组秘钥
	// NewCipher该函数限制了输入k的长度必须为16, 24或者32
	block, _ := aes.NewCipher(key)
	blockSize := block.BlockSize()               // 获取秘钥块的长度
	origData = pkcs5Padding(origData, blockSize) // 补全码
	fmt.Println(origData)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize]) // 加密模式
	encrypted := make([]byte, len(origData))                    // 创建数组
	blockMode.CryptBlocks(encrypted, origData)                  // 加密

	fmt.Println(encrypted)
}

func pkcs5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}
