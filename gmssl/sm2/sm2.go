package sm2

import (
	"crypto"
	"encoding/pem"
	"io"

	"chainmaker.org/gotest/gmssl/gmssl"

	"github.com/pkg/errors"
	tjx509 "github.com/tjfoc/gmsm/x509"
)

func GenerateKeyPair() (*PrivateKey, error) {
	sm2keygenargs := &gmssl.PkeyCtxParams{
		Keys:   []string{"ec_paramgen_curve", "ec_param_enc"},
		Values: []string{"sm2p256v1", "named_curve"},
	}
	sk, err := gmssl.GeneratePrivateKey("EC", sm2keygenargs, nil)
	if err != nil {
		return nil, err
	}
	skPem, err := sk.GetUnencryptedPEM()
	if err != nil {
		return nil, err
	}
	p, _ := pem.Decode([]byte(skPem))
	if p == nil {
		return nil, errors.New("invalid private key pem")
	}
	tjsk, err := tjx509.ParsePKCS8UnecryptedPrivateKey(p.Bytes)
	if err != nil {
		return nil, err
	}

	pkPem, err := sk.GetPublicKeyPEM()
	if err != nil {
		return nil, err
	}
	pk, err := gmssl.NewPublicKeyFromPEM(pkPem)
	if err != nil {
		return nil, err
	}

	pubKey := PublicKey{
		PublicKey: pk,
		Curve:     tjsk.Curve,
		X:         tjsk.X,
		Y:         tjsk.Y,
	}

	return &PrivateKey{PrivateKey: sk, D: tjsk.D, Pub: pubKey}, nil
}

type signer struct {
	PrivateKey
}

//this is for crypto.Signer impl
func (s *signer) Public() crypto.PublicKey {
	return s.PublicKey
}

func (s *signer) Sign(rand io.Reader, msg []byte, opts crypto.SignerOpts) ([]byte, error) {
	return s.PrivateKey.Sign(msg)
}
