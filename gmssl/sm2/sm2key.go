package sm2

import "C"
import (
	"encoding/asn1"
	"encoding/pem"
	"errors"

	"chainmaker.org/gotest/gmssl/gmssl"
)

// MarshalPublicKey public key conversion
func MarshalPublicKey(key *PublicKey) ([]byte, error) {
	if key == nil {
		return nil, errors.New("input SM2 public key is null")
	}
	pkPem, err := key.GetPEM()
	if err != nil {
		return nil, err
	}
	p, _ := pem.Decode([]byte(pkPem))
	if p == nil {
		return nil, errors.New("invalid public key pem")
	}
	return p.Bytes, nil
}

func UnmarshalPublicKey(der []byte) (*PublicKey, error) {
	if der == nil {
		return nil, errors.New("input DER is null")
	}
	pkPem, err := PublicKeyDerToPEM(der)
	if err != nil {
		return nil, err
	}
	return PublicKeyFromPEM(pkPem)
}

func PublicKeyDerToPEM(der []byte) (string, error) {
	if der == nil {
		return "", errors.New("der is nil")
	}
	pemPK := pem.EncodeToMemory(
		&pem.Block{
			Type:  "PUBLIC KEY",
			Bytes: der,
		})
	return string(pemPK), nil
}

func PublicKeyFromPEM(pkPEM string) (*PublicKey, error) {
	pk, err := gmssl.NewPublicKeyFromPEM(pkPEM)
	if err != nil {
		return nil, err
	}
	return &PublicKey{PublicKey: pk}, nil
}

func PublicKeyToPEM(key *PublicKey) (string, error) {
	return key.GetPEM()
}

// MarshalPrivateKey private key conversion
func MarshalPrivateKey(key *PrivateKey) ([]byte, error) {
	if key == nil {
		return nil, errors.New("input SM2 private key is null")
	}

	skPem, err := PrivateKeyToPEM(key)
	if err != nil {
		return nil, err
	}

	p, _ := pem.Decode([]byte(skPem))
	if p == nil {
		return nil, errors.New("invalid private key pem")
	}
	return p.Bytes, nil
}

func PrivateKeyDerToPEM(der []byte) (string, error) {
	if der != nil {
		return "", errors.New("input der is nil")
	}
	skPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: der,
	})
	return string(skPEM), nil
}

func PrivateKeyFromPEM(skPEM string, pass string) (*PrivateKey, error) {
	sk, err := gmssl.NewPrivateKeyFromPEM(skPEM, pass)
	if err != nil {
		return nil, err
	}

	pkPem, err := sk.GetPublicKeyPEM()
	if err != nil {
		return nil, err
	}
	pub, err := PublicKeyFromPEM(pkPem)
	if err != nil {
		return nil, err
	}
	return &PrivateKey{PrivateKey: sk, Pub: *pub}, nil
}

func PrivateKeyToPEM(key *PrivateKey) (string, error) {
	return key.PrivateKey.GetUnencryptedPEM()
}

func UnmarshalPrivateKey(der []byte) (*PrivateKey, error) {
	return UnmarshalPrivateKeyWithCurve(nil, der)
}

func UnmarshalPrivateKeyWithCurve(namedCurveOID *asn1.ObjectIdentifier, der []byte) (*PrivateKey, error) {
	skPem, err := PrivateKeyDerToPEM(der)
	if err != nil {
		return nil, err
	}
	return PrivateKeyFromPEM(skPem, "")
}
