package basic

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
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

	log.Fatalf("failed to initialize pkcs11 handle, err = %s", err)
}
