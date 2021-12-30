package sansectest

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	bccrypto "chainmaker.org/chainmaker/common/v2/crypto"
	"chainmaker.org/chainmaker/common/v2/crypto/pkcs11"
)

var (
	lib              = "./libupkcs11.dylib"
	label            = "Sansec HSM"
	password         = "1234"
	sessionCacheSize = 10
	hashStr          = "SHA1"

	plain = []byte("chainmaker")
)

func TestNewSM4Key(t *testing.T) {
	var err error
	p11, err := pkcs11.New(lib, label, password, sessionCacheSize, hashStr)
	if err != nil || p11 == nil {
		fmt.Printf("Init pkcs11 handle fail, err = %s\n", err)
		os.Exit(1)
	}

	sk, err := pkcs11.NewSM4Key(p11, []byte("MasterKey1"))
	assert.NoError(t, err)

	cipherText, err := sk.Encrypt(plain)
	assert.NoError(t, err)
	assert.NotNil(t, cipherText)

	plainText, err := sk.Decrypt(cipherText)
	assert.NoError(t, err)
	assert.NotNil(t, plainText)
	assert.Equal(t, plain, plainText)
}

func TestGenerateSM4Key(t *testing.T) {
	var err error
	p11, err := pkcs11.New(lib, label, password, sessionCacheSize, hashStr)
	if err != nil || p11 == nil {
		fmt.Printf("Init pkcs11 handle fail, err = %s\n", err)
		os.Exit(1)
	}

	sk, err := pkcs11.GenSecretKey(p11, "SM4EncKey1", bccrypto.SM4, 16)

	assert.NoError(t, err)

	cipherText, err := sk.Encrypt(plain)
	assert.NoError(t, err)
	assert.NotNil(t, cipherText)

	plainText, err := sk.Decrypt(cipherText)
	assert.NoError(t, err)
	assert.NotNil(t, plainText)
	assert.Equal(t, plain, plainText)
}
