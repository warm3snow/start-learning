package crypto

type IHash interface {
	Hash(data []byte) ([]byte, error)
}

type IEnc interface {
	Encrypt(data []byte, mode string) ([]byte, error)
}
