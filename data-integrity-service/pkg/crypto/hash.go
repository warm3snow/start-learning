package crypto

import (
	"errors"

	"github.com/tjfoc/gmsm/sm3"
)

type SoftSM3 struct {
}

func (s SoftSM3) Hash(data []byte) ([]byte, error) {
	return sm3.Sm3Sum(data), nil
}

type HsmSM3 struct {
}

func (h HsmSM3) Hash(data []byte) ([]byte, error) {
	return nil, errors.New("Not supported")
}
