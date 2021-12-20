package sm3

import "C"
import (
	"hash"

	"chainmaker.org/gotest/opencrypto/tencentsm/base"
)

type digest struct {
	ctx *base.SM3_ctx_t
}

func New() hash.Hash {
	d := new(digest)
	var ctx base.SM3_ctx_t
	base.SM3Init(&ctx)
	d.ctx = &ctx
	return d
}

func NewDigestCtx() *digest {
	d := new(digest)
	var ctx base.SM3_ctx_t
	base.SM3Init(&ctx)
	d.ctx = &ctx
	return d
}

func (d *digest) BlockSize() int {
	return base.SM3_BLOCK_SIZE
}

func (d *digest) Size() int {
	return base.SM3_DIGEST_LENGTH
}

func (d *digest) Reset() {
	var ctx base.SM3_ctx_t
	base.SM3Init(&ctx)
	d.ctx = &ctx
}

func (d *digest) Write(p []byte) (int, error) {
	if p == nil {
		return 0, nil
	}
	base.SM3Update(d.ctx, p[:], len(p))
	return len(p), nil
}

func (d *digest) Sum(in []byte) []byte {
	dgst := make([]byte, base.SM3_DIGEST_LENGTH)
	_, _ = d.Write(in)
	base.SM3Final(d.ctx, dgst[:])
	return dgst
}
