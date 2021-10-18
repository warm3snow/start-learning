package sm2

import (
	"crypto"
	"crypto/rand"
	"testing"

	tjsm2 "github.com/tjfoc/gmsm/sm2"
	tjx509 "github.com/tjfoc/gmsm/x509"

	"github.com/stretchr/testify/assert"
)

var (
	msg = []byte("hello gmssl")
)

func TestGenerateKeyPair(t *testing.T) {
	priv, err := GenerateKeyPair()
	assert.NoError(t, err)
	assert.NotNil(t, priv)

	sig, err := priv.Sign(msg)
	assert.NoError(t, err)

	ok, err := priv.Pub.Verify(msg, sig)
	assert.NoError(t, err)
	assert.True(t, ok)
}

func BenchmarkPrivateKey_Sign(b *testing.B) {
	priv, err := GenerateKeyPair()
	assert.NoError(b, err)
	assert.NotNil(b, priv)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		priv.Sign(msg)
	}
}

func BenchmarkPublicKey_Verify(b *testing.B) {
	priv, err := GenerateKeyPair()
	assert.NoError(b, err)
	assert.NotNil(b, priv)

	sig, err := priv.Sign(msg)
	assert.NoError(b, err)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		priv.Pub.Verify(msg, sig)
	}
}

func TestSigner_Sign(t *testing.T) {
	priv, _ := GenerateKeyPair()
	sig, err := priv.ToStandardKey().(crypto.Signer).Sign(rand.Reader, msg, nil)
	assert.NoError(t, err)

	pass, err := priv.PublicKey().Verify(msg, sig)
	assert.NoError(t, err)
	assert.True(t, pass)
}

func TestSM2Std(t *testing.T) {
	//gmssl sign, tjfoc verify
	priv, _ := GenerateKeyPair()
	sig, err := priv.Sign(msg)
	assert.NoError(t, err)

	pkBytes, _ := MarshalPublicKey(&priv.Pub)
	tjpk, err := tjx509.ParseSm2PublicKey(pkBytes)
	assert.NoError(t, err)

	pass := tjpk.Verify(msg, sig)
	assert.True(t, pass)

	//tjfoc sign, gmssl verify
	tjsk, _ := tjsm2.GenerateKey(rand.Reader)
	sig, _ = tjsk.Sign(rand.Reader, msg, nil)

	pkBytes, _ = tjx509.MarshalSm2PublicKey(&tjsk.PublicKey)
	gmsslpk, err := UnmarshalPublicKey(pkBytes)
	assert.NoError(t, err)

	pass, err = gmsslpk.Verify(msg, sig)
	assert.NoError(t, err)
	assert.True(t, pass)
}
