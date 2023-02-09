package api

import (
	"encoding/xml"
	"fmt"
	"gchat-gzh/gzh"
	"gchat-gzh/pkg/logger"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

const (
	token = "ddsagsddesdxzdf223"
)

func WxCheckSign(c *gin.Context) {
	signature := c.Query("signature")
	timestamp := c.Query("timestamp")
	nonce := c.Query("nonce")
	echostr := c.Query("echostr")

	fmt.Println(signature)
	fmt.Println(timestamp)
	fmt.Println(nonce)
	fmt.Println(echostr)

	ok := gzh.CheckSign(signature, timestamp, nonce, token)
	if !ok {
		logger.Log.Error("wx api check failed")
		return
	}
	logger.Log.Info("wx api check success")
	c.Writer.WriteString(echostr)
}

type WxTextMsg struct {
	//名字和收到的xml一致，否则无法获取值
	ToUserName   string `xml:"ToUserName"`
	FromUserName string `xml:"FromUserName"`
	CreateTime   int64  `xml:"CreateTime"`
	MsgType      string `xml:"MsgType"`
	Content      string `xml:"Content"`
	MsgId        int64  `xml:"MsgId"`
}
type WxRepTextMsg struct {
	ToUserName   string
	FromUserName string
	CreateTime   int64
	MsgType      string
	Content      string
	//下面这个字段是指定外层标签名为<xml></xml>
	//不添加这个的话外层标签就是结构体的名字<WXRepTextMsg></WXRepTextMsg>
	XMLName xml.Name `xml:"xml"`
}

func WxMsgPost(c *gin.Context) {
	var req WxTextMsg
	if err := c.ShouldBindXML(&req); err != nil {
		logger.Log.Error("get post body failed", err)
		return
	}
	logger.Log.Info("get post body success,msgID = ", req.MsgId)

	repText := WxRepTextMsg{
		//因为是发送，所以把接收到的两个值换一下就行了
		ToUserName:   req.FromUserName,
		FromUserName: req.ToUserName,
		CreateTime:   time.Now().Unix(),
		MsgType:      "text",
		Content:      fmt.Sprintf("接收的消息是:%s 回复时间：%s", req.Content, time.Now().Format("2006-01-02 15:04:05")),
	}
	//格式化为xml
	msg, err := xml.Marshal(&repText)
	if err != nil {
		log.Println("error:拼接xml失败")
		return
	}
	//此处因为msg就是byte所以直接用Write
	c.Writer.Write(msg)

	//c.XML(http.StatusOK, gin.H{
	//	"xml": repText,
	//})
}
