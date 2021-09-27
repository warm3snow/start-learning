package sm2

import "C"
import (
	"bytes"
	"chainmaker.org/gotest/cryptotest/tencentsm/base"
	"crypto/elliptic"
	"errors"
	"fmt"
	"math/big"
)

type PrivateKey struct {
	PublicKey
	D    *big.Int
	Text []byte
}

type PublicKey struct {
	elliptic.Curve
	X, Y *big.Int
	Text []byte
	pool *ctxPool
}

func GenerateKeyPair() (*PrivateKey, *PublicKey, error) {
	var sk [SM2_PRIVATE_KEY_STR_LEN]byte
	var pk [SM2_PUBLIC_KEY_STR_LEN]byte

	var ctx base.SM2_ctx_t
	base.SM2InitCtx(&ctx)

	ret := base.GenerateKeyPair(&ctx, sk[:], pk[:])
	if ret != 0 {
		return nil, nil, fmt.Errorf("fail to generate SM2 key pair: internal error")
	}

	skD := new(big.Int)
	skD, ok := skD.SetString(string(sk[0:SM2_PRIVATE_KEY_SIZE]), 16)
	if !ok {
		return nil, nil, fmt.Errorf("fail to generate SM2 key pair: wrong private key")
	}

	pkX := new(big.Int)
	pkX, ok = pkX.SetString(string(pk[2:(SM2_PUBLIC_KEY_STR_LEN/2)]), 16)
	if !ok {
		return nil, nil, fmt.Errorf("fail to generate SM2 key pair: wrong public key")
	}
	pkY := new(big.Int)
	pkY, ok = pkY.SetString(string(pk[(SM2_PUBLIC_KEY_STR_LEN/2):SM2_PUBLIC_KEY_SIZE]), 16)
	if !ok {
		return nil, nil, fmt.Errorf("fail to generate SM2 key pair: wrong public key")
	}

	pkStruct := PublicKey{
		Curve: P256Sm2(),
		X:     pkX,
		Y:     pkY,
		Text:  pk[:],
		pool:  NewCtxPoolWithPubKey(pk[:]),
	}

	skStruct := PrivateKey{
		PublicKey: pkStruct,
		D:         skD,
		Text:      sk[:],
	}

	return &skStruct, &pkStruct, nil
}

func (pk *PublicKey) EncryptWithMode(msg []byte, mode base.SM2CipherMode) ([]byte, error) {
	if msg == nil {
		return nil, errors.New("SM2 encrypt: plaintext is null")
	}
	ctx := pk.pool.GetCtx()
	defer pk.pool.ReleaseCtx(ctx)

	var cipherLen int
	cipher := make([]byte, len(msg)+SM2_CIPHER_EXTRA_SIZE)
	pkByte := pk.Text
	ret := base.SM2EncryptWithMode(
		ctx,
		msg[:],
		len(msg),
		pkByte[:],
		SM2_PUBLIC_KEY_SIZE,
		cipher[:],
		&cipherLen,
		mode,
	)
	if ret != 0 {
		return nil, errors.New("SM2: fail to encrypt")
	}
	return cipher[0:cipherLen], nil
}

func (sk *PrivateKey) DecryptWithMode(cipher []byte, mode base.SM2CipherMode) ([]byte, error) {
	if cipher == nil {
		return nil, errors.New("SM2 decrypt: ciphertext is null")
	}
	ctx := sk.pool.GetCtx()
	defer sk.pool.ReleaseCtx(ctx)

	var plainLen int
	plain := make([]byte, len(cipher))
	skByte := sk.Text
	ret := base.SM2DecryptWithMode(
		ctx,
		cipher[:],
		len(cipher),
		skByte[:],
		SM2_PRIVATE_KEY_SIZE,
		plain[:],
		&plainLen,
		mode,
	)
	if ret != 0 {
		return nil, errors.New("SM2: fail to decrypt")
	}
	return plain[0:int(plainLen)], nil
}

func (pk *PublicKey) Encrypt(msg []byte) ([]byte, error) {
	return pk.EncryptWithMode(msg, SM2_CIPHER_MODE_C1C3C2_ASN1)
}

func (sk *PrivateKey) Decrypt(cipher []byte) ([]byte, error) {
	return sk.DecryptWithMode(cipher, SM2_CIPHER_MODE_C1C3C2_ASN1)
}

func (sk *PrivateKey) SignWithSM3WithMode(msg, id []byte, mode base.SM2SignMode) ([]byte, error) {
	ctx := sk.pool.GetCtx()
	defer sk.pool.ReleaseCtx(ctx)

	var sigLen int
	sig := make([]byte, SM2_SIGNATURE_MAX_SIZE)
	skByte := sk.Text
	pkByte := sk.PublicKey.Text
	ret := base.SM2SignWithMode(
		ctx,
		msg[:],
		len(msg),
		id[:],
		len(id),
		pkByte[:],
		SM2_PUBLIC_KEY_SIZE,
		skByte[:],
		SM2_PRIVATE_KEY_SIZE,
		sig[:],
		&sigLen,
		mode,
	)
	if ret != 0 {
		return nil, errors.New("SM2: fail to sign message")
	}
	return sig[0:sigLen], nil
}

func (pk *PublicKey) VerifyWithSM3WithMode(msg, id, sig []byte, mode base.SM2SignMode) bool {
	ctx := pk.pool.GetCtx()
	defer pk.pool.ReleaseCtx(ctx)

	pkByte := pk.Text
	ret := base.SM2VerifyWithMode(
		ctx,
		msg[:],
		len(msg),
		id[:],
		len(id),
		sig[:],
		len(sig),
		pkByte[:],
		SM2_PUBLIC_KEY_SIZE,
		mode,
	)
	if ret != 0 {
		return false
	}
	return true
}

func (sk *PrivateKey) SignWithSM3(msg, id []byte) ([]byte, error) {
	return sk.SignWithSM3WithMode(msg, id, SM2_SIGNATURE_MODE_RS_ASN1)
}

func (pk *PublicKey) VerifyWithSM3(msg, id, sig []byte) bool {
	return pk.VerifyWithSM3WithMode(msg, id, sig, SM2_SIGNATURE_MODE_RS_ASN1)
}

func (sk *PrivateKey) SignWithMode(msg []byte, id []byte, mode base.SM2SignMode) ([]byte, error) {
	ctx := sk.pool.GetCtx()
	defer sk.pool.ReleaseCtx(ctx)

	sig := make([]byte, SM2_SIGNATURE_MAX_SIZE)
	sigLen := len(sig)

	skByte := sk.Text
	pkByte := sk.PublicKey.Text

	ret := base.SM2SignWithMode(
		ctx,
		msg,
		len(msg),
		id,
		len(id),
		pkByte,
		len(pkByte),
		skByte,
		len(skByte),
		sig,
		&sigLen,
		mode)
	if ret != 0 {
		return nil, errors.New("SM2: fail to sign message")
	}
	return sig[0:int(sigLen)], nil
}

func (pk *PublicKey) VerifyWithMode(msg, sig []byte, id []byte, mode base.SM2SignMode) bool {
	ctx := pk.pool.GetCtx()
	defer pk.pool.ReleaseCtx(ctx)

	pkByte := pk.Text
	ret := base.SM2VerifyWithMode(
		ctx,
		msg,
		len(msg),
		id,
		len(id),
		sig,
		len(sig),
		pkByte,
		len(pkByte),
		mode)
	if ret != 0 {
		return false
	}
	return true
}

func (sk *PrivateKey) sign(dgst []byte) ([]byte, error) {
	return sk.SignWithMode(dgst, []byte(SM2_DEFAULT_USER_ID), SM2_SIGNATURE_MODE_RS_ASN1)
}

func (pk *PublicKey) Verify(dgst, sig []byte) bool {
	return pk.VerifyWithMode(dgst, sig, []byte(SM2_DEFAULT_USER_ID), SM2_SIGNATURE_MODE_RS_ASN1)
}

//CFCA证书若签名为31位，会补0，go本身是不补，长度写31
//兼容 去掉补0，长度改为31
func GetSignatureFromCFCA(signature []byte) []byte {
	dataLength := len(signature)
	dataIndex := 2 //当前下标，初始值为循环数据开始的位置

	//格式为 类型(1)+总长度(1)+[类型(1)+长度(1)+数据]
	//数据字节数为长度对应的大小，一般为32
	var signBuffer bytes.Buffer
	signBuffer.Write(signature[0:dataIndex])
	currentCount := signature[1]  //结构体总长度，用于减去补0后，总长度同样需要减
	currentDataCount := byte('0') //循环中有效数据实际长度
	dataCount := 0                //用于循环中记录每个数据的长度
	zeroCount := 0                //用于循环中记录出现的补0的个数
	for dataIndex+2 < dataLength {
		signBuffer.WriteByte(signature[dataIndex])
		dataCount = int(signature[dataIndex+1])
		if dataIndex+dataCount+2 > dataLength {
			signBuffer.Write(signature[dataIndex+1:])
			break
		}
		//只对长度为32字节的处理，如33字节表示正数但最高位为0需补符号，属于正常
		if 0 == signature[dataIndex+2] && 0 == signature[dataIndex+3]&0x80 {
			currentDataCount = signature[dataIndex+1] - 1
			zeroCount = 1
			//判断是否补多个0
			for {
				if 0 == signature[dataIndex+2+zeroCount] && 0 == signature[dataIndex+3+zeroCount]&0x80 {
					currentDataCount -= 1
					zeroCount += 1
				} else {
					break
				}
			}
			signBuffer.WriteByte(currentDataCount)
			signBuffer.Write(signature[dataIndex+2+zeroCount : dataIndex+2+dataCount])
			currentCount -= signature[dataIndex+1] - currentDataCount
		} else {
			signBuffer.Write(signature[dataIndex+1 : dataIndex+dataCount+2])
		}

		dataIndex += dataCount + 2
	}

	signature = signBuffer.Bytes()

	if 0 < signature[1]-currentCount {
		signature[1] = currentCount
	}

	return signature
}
