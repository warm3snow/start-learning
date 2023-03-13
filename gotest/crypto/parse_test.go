package crypto

//
//import (
//	"io/ioutil"
//	"testing"
//
//	"chainmaker.org/chainmaker/common/v2/crypto/asym"
//	"github.com/stretchr/testify/assert"
//)
//
//func TestParsePublicKey(t *testing.T) {
//	prvFile := "../testdata/admin1.key"
//	pubFile := "../testdata/admin1.pem"
//
//	pubPem1, err := ioutil.ReadFile(pubFile)
//	assert.NoError(t, err)
//
//	prvBytes, err := ioutil.ReadFile(prvFile)
//	assert.NoError(t, err)
//
//	prv, err := asym.PrivateKeyFromPEM(prvBytes, nil)
//	assert.NoError(t, err)
//
//	pubPem2, err := prv.PublicKey().String()
//	assert.NoError(t, err)
//
//	assert.Equal(t, string(pubPem1), pubPem2)
//}
