package crypto

//import (
//	"encoding/asn1"
//	"encoding/hex"
//	"fmt"
//	"log"
//	"math/big"
//	"os"
//	"testing"
//
//	bccrypto "chainmaker.org/chainmaker/common/v2/crypto"
//	"chainmaker.org/chainmaker/common/v2/crypto/asym"
//	"chainmaker.org/chainmaker/common/v2/crypto/kms"
//	"github.com/stretchr/testify/assert"
//)
//
//type Sig struct {
//	R *big.Int
//	S *big.Int
//}
//
//func TestGetHSMAdapter(t *testing.T) {
//	//os.Setenv("KMS_ADAPTER_LIB", "./tencentcloudkms/adapter.so")
//	sm2KeyId := "e2920cd5-5a02-11eb-840b-525400e8e6ea"
//	msg := "Valar morgulis."
//
//	os.Setenv("SecretId", "AKIDA9uPef95S0JOHQx0RzCF3qPilYUfrjNm")
//	os.Setenv("SecretKey", "VllDRRBfCTVHn46sNYUKSqysWhdR0K0T")
//	os.Setenv("ServerAddress", "kms.tencentcloudapi.com")
//	os.Setenv("ServerRegion", "ap-guangzhou")
//
//	adapter := kms.GetKMSAdapter()
//	assert.NotNil(t, adapter)
//
//	sk, err := adapter.NewPrivateKey(sm2KeyId, "SM2DSA", "")
//	assert.NoError(t, err)
//
//	for i := 0; i < 1000; i++ {
//		sig, err := sk.SignWithOpts([]byte(msg), &bccrypto.SignOpts{Hash: bccrypto.HASH_TYPE_SM3, UID: bccrypto.CRYPTO_DEFAULT_UID})
//		assert.NoError(t, err)
//
//		//if _, err := asn1.Unmarshal(sig, &Sig{}); err != nil {
//		//	fmt.Printf("ERROR: %d signature is invalid asn1-encoded, continue", i)
//		//	continue
//		//}
//
//		pk, err := adapter.NewPublicKey(sm2KeyId)
//		assert.NoError(t, err)
//		ok, err := pk.VerifyWithOpts([]byte(msg), sig, &bccrypto.SignOpts{Hash: bccrypto.HASH_TYPE_SM3, UID: bccrypto.CRYPTO_DEFAULT_UID})
//		assert.NoError(t, err)
//		assert.True(t, ok)
//
//		if err != nil {
//			skBytes, _ := sk.Bytes()
//			pkBytes, _ := pk.Bytes()
//			fmt.Printf("i = %d\nsk = %s\npk = %s\nmsg =%s\nsig = %s\n", i, hex.EncodeToString(skBytes),
//				hex.EncodeToString(pkBytes),
//				hex.EncodeToString([]byte(msg)),
//				hex.EncodeToString(sig))
//		}
//	}
//}
//
//func TestKMSVerify(t *testing.T) {
//	// 1. invalid signature
//	pkHex, sigHex, msgHex := "3059301306072a8648ce3d020106082a811ccf5501822d0342000415fcadb4ffdb5d014864ab6e7a3f4d30d6515584077be7559ab08a24a91811d902c07fc49e4b03b5dbd33ac6bbf5d6823baeb8e08c973225af45d544103a4e3c",
//		"30440220001bef09646f1a9e7a71c2d78f670082c1ae374e0389174e333381d7d6f2febb02201ab0d04b5fd86389d21b425a4ee934a9b2cc592b9a5152725b25448404b96a18",
//		"56616c6172206d6f7267756c69732e"
//
//	// 2. valid signature
//	//pk = 3059301306072a8648ce3d020106082a811ccf5501822d0342000415fcadb4ffdb5d014864ab6e7a3f4d30d6515584077be7559ab08a24a91811d902c07fc49e4b03b5dbd33ac6bbf5d6823baeb8e08c973225af45d544103a4e3c
//	//msg =56616c6172206d6f7267756c69732e
//	//sig = 3045022041ed0b0b3bef06e8957440207d07357b7f12021d72b7d196abf43120986402890221009943b41b27c837bcc0bd56da0690648c164f765feaa41716b1620a6000670843
//
//	pkBytes, err := hex.DecodeString(pkHex)
//	assert.NoError(t, err)
//	sigBytes, err := hex.DecodeString(sigHex)
//	assert.NoError(t, err)
//	msgBytes, err := hex.DecodeString(msgHex)
//
//	pk, err := asym.PublicKeyFromDER(pkBytes)
//	assert.NoError(t, err)
//
//	ok, err := pk.VerifyWithOpts(msgBytes, sigBytes, &bccrypto.SignOpts{Hash: bccrypto.HASH_TYPE_SM3, UID: bccrypto.CRYPTO_DEFAULT_UID})
//	assert.NoError(t, err)
//	assert.True(t, ok)
//}
//
//func TestSig(t *testing.T) {
//	sigHex := "3043021f1bef09646f1a9e7a71c2d78f670082c1ae374e0389174e333381d7d6f2febb02201ab0d04b5fd86389d21b425a4ee934a9b2cc592b9a5152725b25448404b96a18"
//	sigBytes, err := hex.DecodeString(sigHex)
//	assert.NoError(t, err)
//
//	log.Println("sigHex = ", sigHex)
//	log.Println("sigBytes = ", sigBytes)
//	log.Printf("sigBytes = %b", sigBytes)
//
//	sigStruct := Sig{}
//	_, err = asn1.Unmarshal(sigBytes, &sigStruct)
//	assert.NoError(t, err)
//}

//
//sigHex = "3043021f1bef09646f1a9e7a71c2d78f670082c1ae374e0389174e333381d7d6f2febb02201ab0d04b5fd86389d21b425a4ee934a9b2cc592b9a5152725b25448404b96a18"
//sigBytes, err = hex.DecodeString(sigHex)
//assert.NoError(t, err)
//
//log.Println("sigHex = ", sigHex)
//log.Println("sigBytes = ", sigBytes)
//log.Printf("sigBytes = %b", sigBytes)
//
//sigStruct = Sig{}
//_, err = asn1.Unmarshal(sigBytes, &sigStruct)
//assert.NoError(t, err)
