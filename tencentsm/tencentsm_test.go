package tencentsm

import (
	"crypto/rand"
	"encoding/hex"
	"testing"

	tjx509 "github.com/tjfoc/gmsm/x509"

	tsm2 "chainmaker.org/gotest/tencentsm/sm2"
	"github.com/stretchr/testify/assert"
	"github.com/tjfoc/gmsm/sm2"
)

var (
	msg = []byte("hello gotest")
	id  = []byte("1234567812345678")
)

func TestSignAndVerify_10000(t *testing.T) {
	total := 10
	for i := 0; i < total; i++ {
		// generate key pair using tj sm
		sk, pk, err := tsm2.GenerateKeyPair()
		if err != nil {
			t.Fatal(err)
		}

		if len(pk.X.Bytes()) != 32 || len(pk.Y.Bytes()) != 32 {
			testTsmTjfocSTD(t, sk, pk)
		}
	}
}

func testTsmTjfocSTD(t *testing.T, sk *tsm2.PrivateKey, pk *tsm2.PublicKey) {
	// tencentsm sign and verify
	sig1, err := sk.SignWithSM3(msg, id)
	assert.NoError(t, err)
	ok := pk.VerifyWithSM3(msg, id, sig1)
	assert.True(t, ok)
	t.Log("tencent report:")
	t.Logf("signature[%d] = %s\n", len(sig1), hex.EncodeToString(sig1))
	t.Logf("publickey: x[%d] = %s,  y[%d] = %s\n", len(pk.X.Bytes()), hex.EncodeToString(pk.X.Bytes()), len(pk.Y.Bytes()), hex.EncodeToString(pk.Y.Bytes()))

	//convert tencent to tjfoc
	priBytes, err := tsm2.MarshalPrivateKey(sk)
	assert.NoError(t, err)
	assert.NotNil(t, priBytes)

	pri, err := tjx509.ParsePKCS8UnecryptedPrivateKey(priBytes)
	assert.NoError(t, err)
	assert.NotNil(t, pri)

	pub := &sm2.PublicKey{
		Curve: sm2.P256Sm2(),
		X:     pk.X,
		Y:     pk.Y,
	}
	assert.NotNil(t, pub)

	// tjfoc sign and verify
	sig2, err := pri.Sign(rand.Reader, msg, nil)
	assert.NoError(t, err)
	ok = pub.Verify(msg, sig2)
	assert.True(t, ok)
	t.Log("tjfoc report:")
	t.Logf("signature[%d] = %s\n", len(sig1), hex.EncodeToString(sig2))
	t.Logf("publickey: x[%d] = %s,  y[%d] = %s\n", len(pub.X.Bytes()), hex.EncodeToString(pub.X.Bytes()), len(pub.Y.Bytes()), hex.EncodeToString(pub.Y.Bytes()))

	// tencentsm verify tjfoc signature
	ok = pk.VerifyWithSM3(msg, id, sig2)
	assert.True(t, ok)

	// tjfoc verify tencentsm
	ok = pub.Verify(msg, sig1)
	assert.True(t, ok)
}
