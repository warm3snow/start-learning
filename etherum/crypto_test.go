package etherum

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"testing"

	crypto2 "chainmaker.org/chainmaker/common/v2/crypto"
	"chainmaker.org/chainmaker/common/v2/crypto/asym"
	"github.com/stretchr/testify/assert"

	"github.com/ethereum/go-ethereum/crypto"
)

func TestKey(t *testing.T) {
	priKeyHash := "796c823671b118258b53ef6056fd1f9fc96d125600f348f75f397b2000267fe8"
	// 创建私钥对象，上面私钥没有钱哦
	priKey, err := crypto.HexToECDSA(priKeyHash)
	if err != nil {
		panic(err)
	}
	priKeyBytes := crypto.FromECDSA(priKey)
	fmt.Printf("私钥为: %s", hex.EncodeToString(priKeyBytes))

	pubKey := priKey.Public().(*ecdsa.PublicKey)
	// 获取公钥并去除头部0x04
	pubKeyBytes := crypto.FromECDSAPub(pubKey)[1:]
	fmt.Printf("公钥为: %s", hex.EncodeToString(pubKeyBytes))

	// 获取地址
	addr := crypto.PubkeyToAddress(*pubKey)
	fmt.Printf("地址为: %s", addr.Hex())
}

func TestEthKey(t *testing.T) {
	keyHex := "53aa146fbaed5198fd750200a62b7f7445073d24b592baffd9c1e05e4b142fa1"

	priv, err := asym.PrivateKeyFromPEM([]byte(keyHex), nil)
	assert.NoError(t, err)

	fmt.Println(crypto2.KeyType2NameMap[priv.Type()])

	keyBytes, err := ioutil.ReadFile("user.key")
	assert.NoError(t, err)

	priv, err = asym.PrivateKeyFromPEM(keyBytes, nil)
	assert.NoError(t, err)

	fmt.Println(crypto2.KeyType2NameMap[priv.Type()])
}
