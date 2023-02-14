package api

import (
	"github.com/gin-gonic/gin"
)

func router() *gin.Engine {
	r := gin.Default()
	//r.Use(utils.Cors())

	//r.GET("/wx", WxCheckSign)
	r.POST("/wx", WxMsgPost)

	//r.GET("/wx/ws", MessageAIWs)

	r.GET("/ai", MessageAIWs)
	return r
}
