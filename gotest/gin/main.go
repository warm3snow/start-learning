/**
 * @Author: xueyanghan
 * @File: main.go
 * @Version: 1.0.0
 * @Description: TODO
 * @Date: 2023/3/23 10:19
 */
package main

import (
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	store := cookie.NewStore([]byte("secret"))
}

// CreateTBaasGin
func CreateTBaasGin() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = ioutil.Discard
	engine := gin.Default()
	engine.Use(recover.PanicRecover())
	engine.Use(mlog.InfoLog())
	engine.GET(middleutils.URL_METRICS, gin.WrapH(promhttp.Handler()))
	engine.GET(middleutils.URL_CONFIG_UPDATE_NOTIFY, ConfigCenterUpdate)
	engine.GET(middleutils.URL_HEART_BEAT, HeartBeat)
	return engine
}

// SetRespGetterFactory
func SetRespGetterFactory(factory middleutils.RespGetterFactory) {
	middleutils.SetRespGetterFactory(factory)
}

// ConfigCenterUpdate
func ConfigCenterUpdate(c *gin.Context) {
	configcenter.UpdateEventNotify()
	c.JSON(http.StatusOK, "")
}

// HeartBeat
func HeartBeat(c *gin.Context) {
	c.String(http.StatusOK, "SUCCESS")
}
