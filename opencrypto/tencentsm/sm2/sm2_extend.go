package sm2

import (
	"crypto"
	"io"
)

//this is for crypto.Signer impl
func (priv *PrivateKey) Public() crypto.PublicKey {
	return &priv.PublicKey
}

func (priv *PrivateKey) Sign(rand io.Reader, msg []byte, opts crypto.SignerOpts) ([]byte, error) {
	return priv.sign(msg)
}
