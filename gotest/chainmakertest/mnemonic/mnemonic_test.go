package mnemonic

import (
	"fmt"
	"testing"

	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"

	"github.com/brianium/mnemonic"
	"github.com/brianium/mnemonic/entropy"
)

func TestMnemonic1(t *testing.T) {
	// Generate a mnemonic for memorization or user-friendly seeds
	entropy, _ := bip39.NewEntropy(256)
	mnemonic, _ := bip39.NewMnemonic(entropy)

	// Generate a Bip32 HD wallet for the mnemonic and a user supplied password
	seed := bip39.NewSeed(mnemonic, "Secret Passphrase")

	masterKey, _ := bip32.NewMasterKey(seed)
	publicKey := masterKey.PublicKey()

	// Display mnemonic and keys
	fmt.Println("Mnemonic: ", mnemonic)
	fmt.Println("Master private key: ", masterKey)
	fmt.Println("Master public key: ", publicKey)
}

func Test2(t *testing.T) {
	// generate some entropy from a hex string
	ent, _ := entropy.FromHex("8197a4a47f0425faeaa69deebc05ca29c0a5b5cc76ceacc0")
	// generate a Mnemonic in Japanese with the generated entropy
	jp, _ := mnemonic.New(ent, mnemonic.Japanese)
	fmt.Println(jp.Sentence())

	// generate some random 256 bit entropy
	rnd, _ := entropy.Random(256)
	// generate a Mnemonic in Spanish with the generated entropy
	sp, _ := mnemonic.New(rnd, mnemonic.Spanish)
	fmt.Println(sp.Sentence())

	cp, _ := mnemonic.New(rnd, mnemonic.ChineseSimplified)
	fmt.Println(cp.Words)
}

func TestWords(t *testing.T) {
	fmt.Println("english:", len(english))
	fmt.Println("chinese s:", len(chineseSimplified))
	fmt.Println("chinese t:", len(chineseTraditional))
	fmt.Println("japanese:", len(japanese))
}
