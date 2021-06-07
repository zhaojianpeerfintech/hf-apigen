package router

import (
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/peerfintech/hf-apigen/apiserver/handler"
)

// Router 全局路由
var router *gin.Engine
var onceCreateRouter sync.Once

//GetRouter 获取路由
func GetRouter() *gin.Engine {
	onceCreateRouter.Do(func() {
		router = createRouter()
	})

	return router
}

func createRouter() *gin.Engine {
	r := gin.Default()

	v := r.Group("/v1")
	{
		v.POST("/invoke", handler.InvokeData)
		v.GET("/query", handler.QueryData)
	}
	return r
}
