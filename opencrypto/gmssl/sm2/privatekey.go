package sm2

import (
	crypto2 "crypto"
	"encoding/hex"
	"math/big"

	"chainmaker.org/chainmaker/common/v2/crypto"
	"chainmaker.org/chainmaker/common/v2/crypto/hash"

	"chainmaker.org/gotest/opencrypto/gmssl/gmssl"
	"chainmaker.org/gotest/opencrypto/gmssl/sm3"
)

type PrivateKey struct {
	*gmssl.PrivateKey
	D *big.Int

	Pub PublicKey
}

func (sk *PrivateKey) Bytes() ([]byte, error) {
	return sk.D.Bytes(), nil
}

func (sk *PrivateKey) Type() crypto.KeyType {
	return crypto.SM2
}

func (sk *PrivateKey) String() (string, error) {
	return hex.EncodeToString(sk.D.Bytes()), nil
}

func (sk *PrivateKey) PublicKey() crypto.PublicKey {
	return &sk.Pub
}

func (sk *PrivateKey) Sign(msg []byte) ([]byte, error) {
	return sk.signWithSM3(msg, crypto.CRYPTO_DEFAULT_UID)
}

func (sk *PrivateKey) SignWithOpts(msg []byte, opts *crypto.SignOpts) ([]byte, error) {
	if opts == nil {
		return sk.Sign(msg)
	}
	if opts.Hash == crypto.HASH_TYPE_SM3 && sk.Type() == crypto.SM2 {
		uid := opts.UID
		if len(uid) == 0 {
			uid = crypto.CRYPTO_DEFAULT_UID
		}
		return sk.signWithSM3(msg, uid)
	}
	dgst, err := hash.Get(opts.Hash, msg)
	if err != nil {
		return nil, err
	}
	return sk.Sign(dgst)

}

func (sk *PrivateKey) ToStandardKey() crypto2.PrivateKey {
	return &signer{PrivateKey: *sk}
}

// PrivateKey implements bccrypto.PrivateKey
func (sk *PrivateKey) signWithSM3(msg []byte, uid string) ([]byte, error) {
	sm2zid, err := sk.ComputeSM2IDDigest(uid)
	if err != nil {
		return nil, err
	}

	sm3Hash := sm3.New()
	sm3Hash.Write(sm2zid)
	sm3Hash.Write(msg)
	dgst := sm3Hash.Sum(nil)
	return sk.PrivateKey.Sign("sm2sign", dgst, nil)
}

func (sk *PrivateKey) Decrypt(ciphertext []byte) ([]byte, error) {
	return sk.PrivateKey.Decrypt("sm2encrypt-with-sm3", ciphertext, nil)
}
