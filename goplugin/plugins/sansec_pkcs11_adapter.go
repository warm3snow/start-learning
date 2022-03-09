package main

import (
	"errors"
	"fmt"

	"chainmaker.org/gotest/goplugin"
)

//go build -buildmode=plugin -o plugin1.so plugin1.go

var _ goplugin.IHSMAdapter = (*hsmAdapter)(nil)

var Adapter hsmAdapter

type hsmAdapter struct {
}

func (a hsmAdapter) GetSM2KeyId(keyIdex int, isPrivate bool) (string, error) {
	return fmt.Sprintf("SM2SignKey%d", keyIdex), nil
}

func (a hsmAdapter) GetRSAKeyId(keyIdex int, isPrivate bool) (string, error) {
	return fmt.Sprintf("RSASignKey%d", keyIdex), nil
}

func (a hsmAdapter) GetECCKeyId(keyIdex int, isPrivate bool) (string, error) {
	return "", errors.New("not implemented")
}

func (a hsmAdapter) GetSM4KeyIdex(keyIdex int) (string, error) {
	return fmt.Sprintf("MasterKey%d", keyIdex), nil
}

func (a hsmAdapter) GetAESKeyId(keyIdex int) (string, error) {
	return fmt.Sprintf("MasterKey%d", keyIdex), nil
}

func (a hsmAdapter) GetSM3SM2CKM() uint {
	return 0x80000000 + 4 + 0x00000100
}

func (a hsmAdapter) GetSM2KeyAccessRight(keyIdex int) (newKeyIdex int, need bool) {
	return keyIdex + 10000, true
}

func (a hsmAdapter) GetSM4KeyAccessRight(keyIdex int) (newKeyIdex int, need bool) {
	return keyIdex, true
}

func (a hsmAdapter) GetRSAKeyAccessRight(keyIdex int) (newKeyIdex int, need bool) {
	return keyIdex + 0, true
}

func (a hsmAdapter) GetAESKeyAccessRight(keyIdex int) (newKeyIdex int, need bool) {
	return keyIdex, true
}
