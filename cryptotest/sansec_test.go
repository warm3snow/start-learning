package cryptotest

import (
	"crypto/elliptic"
	"encoding/hex"
	"github.com/stretchr/testify/assert"
	"github.com/tjfoc/gmsm/sm2"
	"github.com/tjfoc/gmsm/sm3"
	"math/big"
	"testing"
)

//TestSansecHSM_SoftVerify sansec sm2 signature (CKM_SM3_SM2) verify using tjfoc sm2
func TestSansecHSM_SoftVerify(t *testing.T) {
	plain := []byte("hello chainmaker")

	pubHex := "04fb0d63136159d1ddbb5154291a1839d166ecefdadcc7d324069d971b562cf08cca80a82c0ec41e03acc2c4e60c83a4b86640850a6b7df3a705d4a1681830210b"
	rHex := "02f59073676d87391323c58033607b1460867143bc7b2646b7b112aa3dfb54c4"
	sHex := "8d47545718b18167f7c09fe283030040de7feb6b5bcfda06011757973c81a8d3"

	pubBytes, err := hex.DecodeString(pubHex)
	assert.NoError(t, err)
	x, y := elliptic.Unmarshal(sm2.P256Sm2(), pubBytes)

	r, err := hex.DecodeString(rHex)
	assert.NoError(t, err)
	s, err := hex.DecodeString(sHex)
	assert.NoError(t, err)
	pubKey := sm2.PublicKey{
		Curve: sm2.P256Sm2(),
		X:     x,
		Y:     y,
	}

	ok := sm2.Verify(&pubKey, sm3.Sm3Sum(plain), big.NewInt(0).SetBytes(r), big.NewInt(0).SetBytes(s))
	assert.True(t, ok)

	//sigHex := "3045022002f59073676d87391323c58033607b1460867143bc7b2646b7b112aa3dfb54c40221008d47545718b18167f7c09fe283030040de7feb6b5bcfda06011757973c81a8d3"
	//sigBytes, err := hex.DecodeString(sigHex)
	//assert.NoError(t, err)
	//
	//ok = pubKey.Verify(sm3.Sm3Sum(plain), sigBytes)
	//assert.True(t, ok)
}
