package crypto

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestP256Key(t *testing.T) {
	pri, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)

	pubBytes := elliptic.Marshal(elliptic.P256(), pri.X, pri.Y)

	pubHex := hex.EncodeToString(pubBytes)

	fmt.Printf("pub[%d] = %s\n", len(pubHex), pubHex)

	xb, yb := pri.X.Bytes(), pri.Y.Bytes()

	fmt.Printf("Xb[%d] = %s, Yb[%d] = %s\n", len(xb), hex.EncodeToString(xb), len(yb), hex.EncodeToString(yb))

	digest := sha256.Sum256([]byte("hello world"))
	sig, err := pri.Sign(rand.Reader, digest[:], nil)
	assert.NoError(t, err)

	ok := ecdsa.VerifyASN1(&pri.PublicKey, digest[:], sig)
	assert.True(t, ok)

	ok = ecdsa.VerifyASN1(&ecdsa.PublicKey{
		elliptic.P256(),
		big.NewInt(0).SetBytes(xb),
		big.NewInt(0).SetBytes(yb),
	}, digest[:], sig)
	assert.True(t, ok)
}
func TestECDSAEcc(t *testing.T) {
	pri, _ := rsa.GenerateKey(rand.Reader, 1024)
	encryptedBytes, err := rsa.EncryptOAEP(
		sha256.New(),
		rand.Reader,
		&pri.PublicKey,
		[]byte("测试哈哈哈"), //需要加密的字符串
		nil)
	assert.NoError(t, err)
	fmt.Printf("enc1: %s\n", hex.EncodeToString(encryptedBytes))

	encryptedBytes, err = rsa.EncryptOAEP(
		sha256.New(),
		rand.Reader,
		&pri.PublicKey,
		[]byte("测试哈哈哈"), //需要加密的字符串
		nil)
	assert.NoError(t, err)
	fmt.Printf("enc1: %s\n", hex.EncodeToString(encryptedBytes))

	decryptedBytes, err := pri.Decrypt(nil, encryptedBytes, &rsa.OAEPOptions{Hash: crypto.SHA256})
	if err != nil {
		panic(err)
	}
	fmt.Println("decrypted message: ", string(decryptedBytes))
}
