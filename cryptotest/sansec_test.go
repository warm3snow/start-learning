package cryptotest

//
//import (
//	"crypto/ecdsa"
//	"encoding/asn1"
//	"encoding/hex"
//	"fmt"
//	"math/big"
//	"testing"
//
//	tjsm2 "github.com/tjfoc/gmsm/sm2"
//
//	gmsslsm2 "chainmaker.org/chainmaker/common/v2/opencrypto/gmssl/sm2"
//
//	txsm2 "chainmaker.org/chainmaker/common/v2/opencrypto/tencentsm/sm2"
//
//	hfsm2 "github.com/Hyperledger-TWGC/tjfoc-gm/sm2"
//	hfx509 "github.com/Hyperledger-TWGC/tjfoc-gm/x509"
//
//	bcsm2 "chainmaker.org/chainmaker/common/v2/crypto/asym/sm2"
//
//	"github.com/stretchr/testify/assert"
//)
//
//var (
//	plainHex  = "56616c6172206d6f7267756c69732e"
//	pubHex    = "3059301306072a8648ce3d020106082a811ccf5501822d0342000415fcadb4ffdb5d014864ab6e7a3f4d30d6515584077be7559ab08a24a91811d902c07fc49e4b03b5dbd33ac6bbf5d6823baeb8e08c973225af45d544103a4e3c"
//	testCases = []struct {
//		sigHex string
//	}{
//		//success cases
//		//{"30450220648f89660672bbca8fa8e878c09cdfdcfacaede9c58650acf6de9de127c456f7022100a8a4dbfd72872c8af78c028cb889cb6eebd0c9e8af02bc5ac8624a10e326fdbb"},
//
//		//failed cases
//		//{"3044021feca4bf4fad8b925d385d40351250e208c399323a98e6397243cdbb20bc106e022100e7ed1c3d94f7b3fe065ee347cdaf6df60a23514f85f6e5b4b1ab9a129bfc2e27"},
//		//{"3045022000eca4bf4fad8b925d385d40351250e208c399323a98e6397243cdbb20bc106e022100e7ed1c3d94f7b3fe065ee347cdaf6df60a23514f85f6e5b4b1ab9a129bfc2e27"},
//
//		//{"3045022100b5715695cc79c5ea52e9db9ed817c01419b9a82013dda8f2c7adefecba4d39da022000c82615deaba0f015dcc69b517fcf2535b7269474d76442b6d31c1ab3e6366a"},
//		//{"3044022100161a10915a22b5724e067d693a2501cd1f028a610f1270d581898b334742f8022004107331a8808cecf784d0996222a06e6f824c9cc6d53ece7c8712de77a69e5d"},
//
//		//{"3045022000eca4bf4fad8b925d385d40351250e208c399323a98e6397243cdbb20bc106e022100e7ed1c3d94f7b3fe065ee347cdaf6df60a23514f85f6e5b4b1ab9a129bfc2e27"},
//		//{"3044022000161a10915a22b5724e067d693a2501cd1f028a610f1270d581898b334742f8022004107331a8808cecf784d0996222a06e6f824c9cc6d53ece7c8712de77a69e5d"},
//		//{"3045022000eca4bf4fad8b925d385d40351250e208c399323a98e6397243cdbb20bc106e022100e7ed1c3d94f7b3fe065ee347cdaf6df60a23514f85f6e5b4b1ab9a129bfc2e27"},
//		{"3044021fc69b65d90279b29eda8d9bca424b4ebde6bb9a0c8683e3a3ebb8268d82dc7f022100e52bf823bf59919bc75bc27ad296ccb9dfd3d2a9f1f94738d78efb368a2f570f"},
//		//{"30440220001bef09646f1a9e7a71c2d78f670082c1ae374e0389174e333381d7d6f2febb02201ab0d04b5fd86389d21b425a4ee934a9b2cc592b9a5152725b25448404b96a18"},
//	}
//)
//
//func TestASN1(t *testing.T) {
//	for _, testCase := range testCases {
//		sigBytes, err := hex.DecodeString(testCase.sigHex)
//		assert.NoError(t, err)
//		assert.NotNil(t, sigBytes)
//		fmt.Printf("sig: len=%d, hex=%s\n", len(sigBytes), testCase.sigHex)
//
//		var sigStruct bcsm2.Sig
//		_, err = asn1.Unmarshal(sigBytes, &sigStruct)
//		assert.NoError(t, err)
//		fmt.Printf("%+v\n", sigStruct)
//
//		one := new(big.Int).SetInt64(1)
//		if sigStruct.R.Cmp(one) < 0 {
//			sigStruct.R = one.Abs(sigStruct.R)
//		}
//
//		assert.True(t, sigStruct.R.Cmp(one) >= 0)
//		assert.True(t, sigStruct.S.Cmp(one) >= 0)
//	}
//}
//
////TestSansecHSM_SoftVerify sansec sm2 signature (CKM_SM3_SM2_APPID1_DER) verify using tjfoc sm2
//func TestSansecHSMSignature_Verify_tjfoc(t *testing.T) {
//	for i, testCase := range testCases {
//		plain, err := hex.DecodeString(plainHex)
//		assert.NoError(t, err)
//		pubBytes, err := hex.DecodeString(pubHex)
//		assert.NoError(t, err)
//		sigBytes, err := hex.DecodeString(testCase.sigHex)
//		assert.NoError(t, err)
//
//		//get sm2.PublicKey
//		pk, err := hfx509.ParsePKIXPublicKey(pubBytes)
//		assert.NoError(t, err)
//		ecKey := pk.(*ecdsa.PublicKey)
//		sm2Key := &tjsm2.PublicKey{
//			tjsm2.P256Sm2(),
//			ecKey.X,
//			ecKey.Y,
//		}
//
//		var sigStruct bcsm2.Sig
//		_, err = asn1.Unmarshal(sigBytes, &sigStruct)
//		assert.NoError(t, err)
//		fmt.Printf("%+v\n", sigStruct)
//
//		////one := new(big.Int).SetInt64(1)
//		////if sigStruct.R.Cmp(one) < 0 {
//		////	sigStruct.R = new(big.Int).Abs(sigStruct.R)
//		////}
//		////if sigStruct.S.Cmp(one) < 0 {
//		////	sigStruct.S = new(big.Int).Abs(sigStruct.S)
//		////}
//		//sigBytes, err = asn1.Marshal(sigStruct)
//		//assert.NoError(t, err)
//
//		//verify
//		ok := sm2Key.Verify(plain, sigBytes)
//		//ok := tjsm2.Sm2Verify(sm2Key, plain, []byte(bccrypto.CRYPTO_DEFAULT_UID), sigStruct.R, sigStruct.S)
//		assert.True(t, ok)
//		if !ok {
//			fmt.Printf("index = %d\nplainHex = %s\npubHex = %s\nsigHex=%s\n", i, plainHex, pubHex, testCase.sigHex)
//		}
//	}
//}
//
//func TestSansecHSMSignature_Verify_TencentSM(t *testing.T) {
//
//	for i, testCase := range testCases {
//		plain, err := hex.DecodeString(plainHex)
//		assert.NoError(t, err)
//		pubBytes, err := hex.DecodeString(pubHex)
//		assert.NoError(t, err)
//		sigBytes, err := hex.DecodeString(testCase.sigHex)
//		assert.NoError(t, err)
//
//		//get sm2.PublicKey
//		pubKey, err := txsm2.UnmarshalPublicKey(pubBytes)
//		assert.NoError(t, err)
//
//		ok, err := pubKey.Verify(plain, sigBytes)
//		assert.NoError(t, err)
//		assert.True(t, ok)
//		if !ok {
//			fmt.Printf("index = %d\nplainHex = %s\npubHex = %s\nsigHex=%s\n", i, plainHex, pubHex, testCase.sigHex)
//		}
//	}
//}
//
//func TestSansecHSMSignature_Verify_TWGC_gmsm(t *testing.T) {
//	for i, testCase := range testCases {
//		plain, err := hex.DecodeString(plainHex)
//		assert.NoError(t, err)
//		pubBytes, err := hex.DecodeString(pubHex)
//		assert.NoError(t, err)
//		sigBytes, err := hex.DecodeString(testCase.sigHex)
//		assert.NoError(t, err)
//
//		//get sm2.PublicKey
//		pk, err := hfx509.ParsePKIXPublicKey(pubBytes)
//		assert.NoError(t, err)
//		ecKey := pk.(*ecdsa.PublicKey)
//		sm2Key := &hfsm2.PublicKey{
//			hfsm2.P256Sm2(),
//			ecKey.X,
//			ecKey.Y,
//		}
//
//		//verify
//		ok := sm2Key.Verify(plain, sigBytes)
//		assert.True(t, ok)
//		if !ok {
//			fmt.Printf("index = %d\nplainHex = %s\npubHex = %s\nsigHex=%s\n", i, plainHex, pubHex, testCase.sigHex)
//		}
//	}
//}
//
//func TestSansecHSMSignature_Verify_Gmssl(t *testing.T) {
//	for i, testCase := range testCases {
//		plain, err := hex.DecodeString(plainHex)
//		assert.NoError(t, err)
//		pubBytes, err := hex.DecodeString(pubHex)
//		assert.NoError(t, err)
//		sigBytes, err := hex.DecodeString(testCase.sigHex)
//		assert.NoError(t, err)
//
//		//get sm2.PublicKey
//		sm2Key, err := gmsslsm2.UnmarshalPublicKey(pubBytes)
//		assert.NoError(t, err)
//
//		//verify
//		ok, err := sm2Key.Verify(plain, sigBytes)
//		assert.True(t, ok)
//		if !ok {
//			fmt.Printf("index = %d\nplainHex = %s\npubHex = %s\nsigHex=%s\n", i, plainHex, pubHex, testCase.sigHex)
//		}
//	}
//}
