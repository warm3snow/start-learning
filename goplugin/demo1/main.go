package main

import (
	"errors"
	"fmt"
	"plugin"

	"chainmaker.org/gotest/goplugin"
)

func main() {
	//LoadAndInvokeSthFromPlugin("../plugins/plugin1.so")
	LoadAndInvokeCryptoAdapterPlugin("../plugins/hsm_adapter.so")
}

func LoadAndInvokeCryptoAdapterPlugin(pluginPath string) error {
	p, _ := plugin.Open(pluginPath)
	s, _ := p.Lookup("Adapter")

	adapter, ok := s.(goplugin.IHSMAdapter)
	if !ok {
		return errors.New("Adapter struct not found")
	}

	ckm := adapter.GetSM3SM2CKM()
	fmt.Printf("ckm_sm3_sm2 = %0x\n", ckm)
	pubId, _ := adapter.GetSM2KeyId(1, true)
	priId, _ := adapter.GetSM2KeyId(1, false)
	fmt.Printf("pubId = %s, priId = %s\n", pubId, priId)

	keyIdx, need := adapter.GetSM2KeyAccessRight(1)
	fmt.Printf("access right, keyIdx = %d, need = %t\n", keyIdx, need)

	return nil
}

func LoadAndInvokeSthFromPlugin(pluginPath string) error {
	p, err := plugin.Open(pluginPath)
	if err != nil {
		return err
	}

	//variable
	v, err := p.Lookup("V")
	if err != nil {
		return err
	}
	*v.(*int) = 15

	//func
	f, err := p.Lookup("F")
	if err != nil {
		return err
	}
	f.(func())()

	//GetSM2KeyId
	getSm2KeyId, err := p.Lookup("GetSM2KeyId")
	if err != nil {
		return err
	}
	keyId := getSm2KeyId.(func(uint) string)(1)
	fmt.Printf("KeyId = %s\n", keyId)

	return nil
}
