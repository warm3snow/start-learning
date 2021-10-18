package cryptotest

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"encoding/hex"
	"fmt"
	"math/big"
	"testing"

	"chainmaker.org/gotest/gmssl/sm2"

	txsm2 "chainmaker.org/gotest/tencentsm/sm2"

	hfx509 "github.com/Hyperledger-TWGC/tjfoc-gm/x509"

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

		//failed cases
		{"3044021feca4bf4fad8b925d385d40351250e208c399323a98e6397243cdbb20bc106e022100e7ed1c3d94f7b3fe065ee347cdaf6df60a23514f85f6e5b4b1ab9a129bfc2e27"},
		{"3044022100b5715695cc79c5ea52e9db9ed817c01419b9a82013dda8f2c7adefecba4d39da021fc82615deaba0f015dcc69b517fcf2535b7269474d76442b6d31c1ab3e6366a"},
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

		//failed cases
		{"3044021feca4bf4fad8b925d385d40351250e208c399323a98e6397243cdbb20bc106e022100e7ed1c3d94f7b3fe065ee347cdaf6df60a23514f85f6e5b4b1ab9a129bfc2e27"},
		{"3044022100b5715695cc79c5ea52e9db9ed817c01419b9a82013dda8f2c7adefecba4d39da021fc82615deaba0f015dcc69b517fcf2535b7269474d76442b6d31c1ab3e6366a"},
	}

	for i, testCase := range testCases {
		plain, err := hex.DecodeString(plainHex)
		assert.NoError(t, err)
		pubBytes, err := hex.DecodeString(pubHex)
		assert.NoError(t, err)
		sigBytes, err := hex.DecodeString(testCase.sigHex)
		assert.NoError(t, err)

		//get sm2.PublicKey
		pubKey, err := txsm2.UnmarshalPublicKey(pubBytes)
		assert.NoError(t, err)

		ok := pubKey.Verify(plain, sigBytes)
		assert.True(t, ok)
		if !ok {
			fmt.Printf("index = %d\nplainHex = %s\npubHex = %s\nsigHex=%s\n", i, plainHex, pubHex, testCase.sigHex)
		}
	}
}

//TestSansecHSM_SoftVerify sansec sm2 signature (CKM_SM3_SM2_APPID1_DER) verify using tjfoc sm2
func TestSansecHSMSignature_HFVerify(t *testing.T) {
	plainHex := "68656c6c6f20636861696e6d616b6572"
	pubHex := "3059301306072a8648ce3d020106082a811ccf5501822d03420004fb0d63136159d1ddbb5154291a1839d166ecefdadcc7d324069d971b562cf08cca80a82c0ec41e03acc2c4e60c83a4b86640850a6b7df3a705d4a1681830210b"
	testCases := []struct {
		sigHex string
	}{
		//success cases
		{"30450220648f89660672bbca8fa8e878c09cdfdcfacaede9c58650acf6de9de127c456f7022100a8a4dbfd72872c8af78c028cb889cb6eebd0c9e8af02bc5ac8624a10e326fdbb"},

		//failed cases
		{"3044021feca4bf4fad8b925d385d40351250e208c399323a98e6397243cdbb20bc106e022100e7ed1c3d94f7b3fe065ee347cdaf6df60a23514f85f6e5b4b1ab9a129bfc2e27"},
		{"3044022100b5715695cc79c5ea52e9db9ed817c01419b9a82013dda8f2c7adefecba4d39da021fc82615deaba0f015dcc69b517fcf2535b7269474d76442b6d31c1ab3e6366a"},
	}

	for i, testCase := range testCases {
		plain, err := hex.DecodeString(plainHex)
		assert.NoError(t, err)
		pubBytes, err := hex.DecodeString(pubHex)
		assert.NoError(t, err)
		sigBytes, err := hex.DecodeString(testCase.sigHex)
		assert.NoError(t, err)

		//get sm2.PublicKey
		pk, err := hfx509.ParsePKIXPublicKey(pubBytes)
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
func TestSansecHSMSignature_GMSSL_Verify(t *testing.T) {
	plainHex := "68656c6c6f20636861696e6d616b6572"
	pubHex := "3059301306072a8648ce3d020106082a811ccf5501822d03420004fb0d63136159d1ddbb5154291a1839d166ecefdadcc7d324069d971b562cf08cca80a82c0ec41e03acc2c4e60c83a4b86640850a6b7df3a705d4a1681830210b"
	testCases := []struct {
		sigHex string
	}{
		//success cases
		{"30450220648f89660672bbca8fa8e878c09cdfdcfacaede9c58650acf6de9de127c456f7022100a8a4dbfd72872c8af78c028cb889cb6eebd0c9e8af02bc5ac8624a10e326fdbb"},

		//failed cases
		{"3044021feca4bf4fad8b925d385d40351250e208c399323a98e6397243cdbb20bc106e022100e7ed1c3d94f7b3fe065ee347cdaf6df60a23514f85f6e5b4b1ab9a129bfc2e27"},
		{"3044022100b5715695cc79c5ea52e9db9ed817c01419b9a82013dda8f2c7adefecba4d39da021fc82615deaba0f015dcc69b517fcf2535b7269474d76442b6d31c1ab3e6366a"},
	}

	for i, testCase := range testCases {
		plain, err := hex.DecodeString(plainHex)
		assert.NoError(t, err)
		pubBytes, err := hex.DecodeString(pubHex)
		assert.NoError(t, err)
		sigBytes, err := hex.DecodeString(testCase.sigHex)
		assert.NoError(t, err)

		//get sm2.PublicKey
		sm2Key, err := sm2.UnmarshalPublicKey(pubBytes)
		assert.NoError(t, err)

		//verify
		ok, err := sm2Key.Verify(plain, sigBytes)
		assert.True(t, ok)
		if !ok {
			fmt.Printf("index = %d\nplainHex = %s\npubHex = %s\nsigHex=%s\n", i, plainHex, pubHex, testCase.sigHex)
		}
	}
}
