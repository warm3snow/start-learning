package sm2

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"runtime"
	"sync"

	"chainmaker.org/gotest/opencrypto/tencentsm/base"

	"github.com/spf13/viper"
)

const (
	MAX_CTX_POOLSIZE  = 100
	INIT_CTX_POOLSIZE = 3
)

var (
	mutex         sync.Mutex
	pubCtxPoolMap map[string]*ctxPool
)

func init() {
	pubCtxPoolMap = make(map[string]*ctxPool, 100)
}

var errPoolClosed = errors.New("tencentsm ctx pool closed")

type ctxPool struct {
	m       sync.Mutex
	ctxChan chan *base.SM2_ctx_t
	closed  bool
	pubkey  []byte
}

func NewCtxPoolWithPubKey(pubkey []byte) *ctxPool {
	mutex.Lock()
	defer mutex.Unlock()
	//if exist
	dgst := sha256.Sum256(pubkey)
	pubHex := hex.EncodeToString(dgst[:20])
	if _, exist := pubCtxPoolMap[pubHex]; exist {
		return pubCtxPoolMap[pubHex]
	}

	//new
	maxPoolSize := MAX_CTX_POOLSIZE
	if viper.IsSet("common.tencentsm.ctx_pool_size.max") {
		maxPoolSize = viper.GetInt("common.tencentsm.ctx_pool_size.max")
		if maxPoolSize <= 0 {
			maxPoolSize = MAX_CTX_POOLSIZE
		}
	}

	pool := &ctxPool{
		ctxChan: make(chan *base.SM2_ctx_t, maxPoolSize),
		pubkey:  pubkey,
	}

	initPoolSize := INIT_CTX_POOLSIZE
	if viper.IsSet("common.tencentsm.ctx_pool_size.init") {
		initPoolSize = viper.GetInt("common.tencentsm.ctx_pool_size.init")
		if initPoolSize <= 0 {
			initPoolSize = INIT_CTX_POOLSIZE
		}
	}

	if initPoolSize > maxPoolSize {
		initPoolSize = maxPoolSize
	}

	var wg sync.WaitGroup
	n := runtime.NumCPU()
	groupSize := initPoolSize / n
	for i := 0; i < n; i++ {
		go func() {
			wg.Add(1)
			defer wg.Done()
			for j := 0; j < groupSize; j++ {
				var ctx base.SM2_ctx_t
				base.SM2InitCtxWithPubKey(&ctx, pubkey)
				pool.ctxChan <- &ctx
			}
		}()
	}
	if initPoolSize%n != 0 {
		go func() {
			wg.Add(1)
			defer wg.Done()
			for j := 0; j < initPoolSize-groupSize*n; j++ {
				var ctx base.SM2_ctx_t
				base.SM2InitCtxWithPubKey(&ctx, pubkey)
				pool.ctxChan <- &ctx
			}
		}()
	}
	wg.Wait()
	pubCtxPoolMap[pubHex] = pool

	return pool
}

func (c *ctxPool) GetCtx() *base.SM2_ctx_t {
	select {
	case r, _ := <-c.ctxChan:
		return r
	default:
		var ctx base.SM2_ctx_t
		base.SM2InitCtxWithPubKey(&ctx, c.pubkey)
		return &ctx
	}
}

func (c *ctxPool) ReleaseCtx(ctx *base.SM2_ctx_t) {
	c.m.Lock()
	defer c.m.Unlock()

	if c.closed {
		base.SM2FreeCtx(ctx)
		return
	}
	select {
	case c.ctxChan <- ctx:
	default:
		base.SM2FreeCtx(ctx)
	}
}

func (c *ctxPool) Close() {
	c.m.Lock()
	defer c.m.Unlock()
	c.closed = true
	close(c.ctxChan)

	for r := range c.ctxChan {
		base.SM2FreeCtx(r)
	}
}
