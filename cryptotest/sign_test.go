package cryptotest

//
//import (
//	"encoding/pem"
//	"io/ioutil"
//	"testing"
//
//	"chainmaker.org/chainmaker/common/v2/crypto"
//
//	"chainmaker.org/chainmaker/common/v2/crypto/asym"
//	"chainmaker.org/chainmaker/common/v2/crypto/x509"
//	"github.com/stretchr/testify/assert"
//)
//
//var (
//	msg = []byte("hello")
//)
//
//func TestSign(t *testing.T) {
//	certPem, err := ioutil.ReadFile("./tmp/admin1.sign.crt")
//	assert.NoError(t, err)
//	priPem, err := ioutil.ReadFile("./tmp/admin1.sign.key")
//	assert.NoError(t, err)
//
//	privateKey, err := asym.PrivateKeyFromPEM(priPem, nil)
//	assert.NoError(t, err)
//
//	block, _ := pem.Decode(certPem)
//	cert, err := x509.ParseCertificate(block.Bytes)
//	assert.NoError(t, err)
//	hashalgo, err := x509.GetHashFromSignatureAlgorithm(cert.SignatureAlgorithm)
//	assert.NoError(t, err)
//
//	sig, err := privateKey.SignWithOpts(msg, &crypto.SignOpts{
//		Hash: hashalgo,
//		UID:  crypto.CRYPTO_DEFAULT_UID,
//	})
//	assert.NoError(t, err)
//
//	ok, err := cert.PublicKey.VerifyWithOpts(msg, sig, &crypto.SignOpts{
//		Hash: hashalgo,
//		UID:  crypto.CRYPTO_DEFAULT_UID,
//	})
//	assert.NoError(t, err)
//	assert.True(t, ok)
//}
