/**
 * @Author: xueyanghan
 * @File: middleware.go
 * @Version: 1.0.0
 * @Description: TODO
 * @Date: 2023/3/23 14:11
 */
package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

const (
	REQUEST_SUCCESS = 0
)

// RequestGetter
type RequestGetter struct {
	RequestID string `json:"RequestId" form:"RequestId"`
	Token     string `json:"token" form:"token"`
	Module    string `json:"module" form:"module"`
}

// DefaultRespGetter
type DefaultRespGetter struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// GetCode
func (p *DefaultRespGetter) GetCode() int {
	return p.Code
}

var defaultRespGetterFactory = func() RespGetter {
	return new(DefaultRespGetter)
}

func init() {
	respGetterFactory = defaultRespGetterFactory
}

// SetRespGetterFactory
func SetRespGetterFactory(factory RespGetterFactory) {
	respGetterFactory = factory
}

// GetRespGetterFactory
func GetRespGetterFactory() RespGetterFactory {
	return respGetterFactory
}

//RespGetterFactory
type RespGetterFactory func() RespGetter

var respGetterFactory RespGetterFactory

// RespGetter
type RespGetter interface {
	GetCode() int
}

// BodyLogWriter
type BodyLogWriter struct {
	gin.ResponseWriter
	BodyBuf *bytes.Buffer
}

// Write
func (w BodyLogWriter) Write(b []byte) (int, error) {
	//memory copy here!
	w.BodyBuf.Write(b)
	return w.ResponseWriter.Write(b)
}

const (
	URL_METRICS              = "/metrics"
	URL_CONFIG_UPDATE_NOTIFY = "/tbaas/sdk/configuration/reload"
	URL_HEART_BEAT           = "/heartbeat"
)

var IgnorePaths []string

func init() {
	IgnorePaths = []string{
		URL_METRICS,
		URL_CONFIG_UPDATE_NOTIFY,
		URL_HEART_BEAT,
	}
}

var (
	TotalCounterVec = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "request_count",
			Help: "Total number of HTTP requests made",
		},
		[]string{"module", "operation"},
	)
	ReqDurationVec = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: "request_latency",
		Help: "record request latency",
	}, []string{"module", "operation"})
	ReqLogicErrorVec = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "request_error_count",
		Help: "Total request error count of the host",
	}, []string{"module", "operation", "code"})
	ReqSystemErrorVec = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "server_error",
		Help: "Total error count of the request",
	}, []string{"module", "operation"})
)

func init() {
	prometheus.MustRegister(
		TotalCounterVec,
		ReqDurationVec,
		ReqLogicErrorVec,
		ReqSystemErrorVec,
	)
}
