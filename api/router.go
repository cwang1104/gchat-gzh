package api

import (
	"github.com/gin-gonic/gin"
)

func router() *gin.Engine {
	r := gin.Default()
	//r.Use(utils.Cors())

	r.GET("/wx", WxCheckSign)
	r.POST("/wx", WxMsgPost)

	return r
}
