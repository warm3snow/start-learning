package cryptotest

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"encoding/hex"
	"fmt"
	"math/big"
	"testing"

	"chainmaker.org/gotest/cryptotest/tencentsm/sm2"

	"github.com/stretchr/testify/assert"
	tjsm2 "github.com/tjfoc/gmsm/sm2"
	tjsm3 "github.com/tjfoc/gmsm/sm3"
	tjx509 "github.com/tjfoc/gmsm/x509"
)

//TestSansecHSM_SoftVerify sansec sm2 signature (CKM_SM3_SM2) verify using tjfoc sm2
func TestSansecHSM_SoftVerify(t *testing.T) {
	plain := []byte("hello chainmaker")

	pubHex := "04fb0d63136159d1ddbb5154291a1839d166ecefdadcc7d324069d971b562cf08cca80a82c0ec41e03acc2c4e60c83a4b86640850a6b7df3a705d4a1681830210b"
	rHex := "02f59073676d87391323c58033607b1460867143bc7b2646b7b112aa3dfb54c4"
	sHex := "8d47545718b18167f7c09fe283030040de7feb6b5bcfda06011757973c81a8d3"

	pubBytes, err := hex.DecodeString(pubHex)
	assert.NoError(t, err)
	x, y := elliptic.Unmarshal(tjsm2.P256Sm2(), pubBytes)

	r, err := hex.DecodeString(rHex)
	assert.NoError(t, err)
	s, err := hex.DecodeString(sHex)
	assert.NoError(t, err)
	pubKey := tjsm2.PublicKey{
		Curve: tjsm2.P256Sm2(),
		X:     x,
		Y:     y,
	}

	ok := tjsm2.Verify(&pubKey, tjsm3.Sm3Sum(plain), big.NewInt(0).SetBytes(r), big.NewInt(0).SetBytes(s))
	assert.True(t, ok)
}

//TestSansecHSM_SoftVerify sansec sm2 signature (CKM_SM3_SM2_APPID1_DER) verify using tjfoc sm2
func TestSansecHSMSignature_TjfocVerify(t *testing.T) {
	plainHex := "68656c6c6f20636861696e6d616b6572"
	pubHex := "3059301306072a8648ce3d020106082a811ccf5501822d03420004fb0d63136159d1ddbb5154291a1839d166ecefdadcc7d324069d971b562cf08cca80a82c0ec41e03acc2c4e60c83a4b86640850a6b7df3a705d4a1681830210b"
	testCases := []struct {
		sigHex string
	}{
		//success cases
		{"30450220648f89660672bbca8fa8e878c09cdfdcfacaede9c58650acf6de9de127c456f7022100a8a4dbfd72872c8af78c028cb889cb6eebd0c9e8af02bc5ac8624a10e326fdbb"},
		{"30450220632de5c9dc0e61f392c982f0371aadbbe101e12966d9514ef62939bae490444c022100da901964bcec2d22aea59ee98c63eb285efb885dd8d6f1996350f8190ab0d0f5"},
		{"304402203b5c4be928294161d3276ff2ef8538b185a6d3a12a4679605df90c7c91444aee02202e84e0a7a6f5f7d67a45f0f8a92a853f43d0803873b022bab005d14fc8079ee2"},

		//failed cases
		{"3044021feca4bf4fad8b925d385d40351250e208c399323a98e6397243cdbb20bc106e022100e7ed1c3d94f7b3fe065ee347cdaf6df60a23514f85f6e5b4b1ab9a129bfc2e27"},
	}

	for i, testCase := range testCases {
		plain, err := hex.DecodeString(plainHex)
		assert.NoError(t, err)
		pubBytes, err := hex.DecodeString(pubHex)
		assert.NoError(t, err)
		sigBytes, err := hex.DecodeString(testCase.sigHex)
		assert.NoError(t, err)

		//get sm2.PublicKey
		pk, err := tjx509.ParsePKIXPublicKey(pubBytes)
		assert.NoError(t, err)
		ecKey := pk.(*ecdsa.PublicKey)
		sm2Key := &tjsm2.PublicKey{
			tjsm2.P256Sm2(),
			ecKey.X,
			ecKey.Y,
		}

		//verify
		ok := sm2Key.Verify(plain, sigBytes)
		assert.True(t, ok)
		if !ok {
			fmt.Printf("index = %d\nplainHex = %s\npubHex = %s\nsigHex=%s\n", i, plainHex, pubHex, testCase.sigHex)
		}
	}
}

//TestSansecHSM_SoftVerify sansec sm2 signature (CKM_SM3_SM2_APPID1_DER) verify using tjfoc sm2
func TestSansecHSMSignature_TecentSMVerify(t *testing.T) {
	plainHex := "68656c6c6f20636861696e6d616b6572"
	pubHex := "3059301306072a8648ce3d020106082a811ccf5501822d03420004fb0d63136159d1ddbb5154291a1839d166ecefdadcc7d324069d971b562cf08cca80a82c0ec41e03acc2c4e60c83a4b86640850a6b7df3a705d4a1681830210b"
	testCases := []struct {
		sigHex string
	}{
		//success cases
		{"30450220648f89660672bbca8fa8e878c09cdfdcfacaede9c58650acf6de9de127c456f7022100a8a4dbfd72872c8af78c028cb889cb6eebd0c9e8af02bc5ac8624a10e326fdbb"},
		{"30450220632de5c9dc0e61f392c982f0371aadbbe101e12966d9514ef62939bae490444c022100da901964bcec2d22aea59ee98c63eb285efb885dd8d6f1996350f8190ab0d0f5"},
		{"304402203b5c4be928294161d3276ff2ef8538b185a6d3a12a4679605df90c7c91444aee02202e84e0a7a6f5f7d67a45f0f8a92a853f43d0803873b022bab005d14fc8079ee2"},

		//failed cases
		{"3044021feca4bf4fad8b925d385d40351250e208c399323a98e6397243cdbb20bc106e022100e7ed1c3d94f7b3fe065ee347cdaf6df60a23514f85f6e5b4b1ab9a129bfc2e27"},
	}

	for i, testCase := range testCases {
		plain, err := hex.DecodeString(plainHex)
		assert.NoError(t, err)
		pubBytes, err := hex.DecodeString(pubHex)
		assert.NoError(t, err)
		sigBytes, err := hex.DecodeString(testCase.sigHex)
		assert.NoError(t, err)

		//get sm2.PublicKey
		pubKey, err := sm2.UnmarshalPublicKey(pubBytes)
		assert.NoError(t, err)

		ok := pubKey.Verify(plain, sigBytes)
		assert.True(t, ok)
		if !ok {
			fmt.Printf("index = %d\nplainHex = %s\npubHex = %s\nsigHex=%s\n", i, plainHex, pubHex, testCase.sigHex)
		}
	}
}
