package main

import (
	"fmt"
	"log"
)

func init() {
	log.Println("plugin1 init")
}

var V int

func F() {
	fmt.Printf("plugin1: public integer variable V =%d\n", V)
}

func GetSM2KeyId(keyIndex uint) string {
	return fmt.Sprintf("SM2SignKey%d", keyIndex)
}
