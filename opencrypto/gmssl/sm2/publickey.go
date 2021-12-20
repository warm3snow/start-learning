package sm2

import (
	"crypto"
	"crypto/elliptic"
	"fmt"
	"math/big"

	bccrypto "chainmaker.org/chainmaker/common/v2/crypto"
	"chainmaker.org/chainmaker/common/v2/crypto/hash"

	"chainmaker.org/gotest/opencrypto/gmssl/gmssl"
	"chainmaker.org/gotest/opencrypto/gmssl/sm3"
)

type PublicKey struct {
	*gmssl.PublicKey

	elliptic.Curve
	X, Y *big.Int
}

// PublicKey implements bccyrpto.PublicKey
var _ bccrypto.PublicKey = (*PublicKey)(nil)

func (pk *PublicKey) verifyWithSM3(msg, sig []byte, uid string) bool {
	sm2zid, _ := pk.ComputeSM2IDDigest(uid)

	sm3Hash := sm3.New()
	sm3Hash.Write(sm2zid)
	sm3Hash.Write(msg)
	dgst := sm3Hash.Sum(nil)
	if err := pk.PublicKey.Verify("sm2sign", dgst, sig, nil); err != nil {
		return false
	}
	return true
}

func (sk *PublicKey) Encrypt(plaintext []byte) ([]byte, error) {
	return sk.PublicKey.Encrypt("sm2encrypt-with-sm3", plaintext, nil)
}

func (pk *PublicKey) Bytes() ([]byte, error) {
	return elliptic.Marshal(pk.Curve, pk.X, pk.Y), nil
}

func (pk *PublicKey) Type() bccrypto.KeyType {
	return bccrypto.SM2
}

func (pk *PublicKey) String() (string, error) {
	return pk.GetPEM()
}

func (pk *PublicKey) Verify(data []byte, sig []byte) (bool, error) {
	return pk.verifyWithSM3(data, sig, bccrypto.CRYPTO_DEFAULT_UID), nil
}

func (pk *PublicKey) VerifyWithOpts(msg []byte, sig []byte, opts *bccrypto.SignOpts) (bool, error) {
	if opts == nil {
		return pk.Verify(msg, sig)
	}
	if opts.Hash == bccrypto.HASH_TYPE_SM3 && pk.Type() == bccrypto.SM2 {
		uid := opts.UID
		if len(uid) == 0 {
			uid = bccrypto.CRYPTO_DEFAULT_UID
		}

		if sig == nil {
			return false, fmt.Errorf("nil signature")
		}
		return pk.verifyWithSM3(msg, sig, uid), nil
	}
	dgst, err := hash.Get(opts.Hash, msg)
	if err != nil {
		return false, err
	}
	return pk.Verify(dgst, sig)
}

func (pk *PublicKey) ToStandardKey() crypto.PublicKey {
	return pk.PublicKey
}
