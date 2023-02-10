package service

import (
	xmll "encoding/xml"
	"fmt"
	"log"
	"time"
)

type xml struct {
	//名字和收到的xml一致，否则无法获取值
	ToUserName   string `xml:"ToUserName"`
	FromUserName string `xml:"FromUserName"`
	CreateTime   int64  `xml:"CreateTime"`
	MsgType      string `xml:"MsgType"`
	Content      string `xml:"Content"`

	//XmlName xml.Name `xml:"xml"`
	//XmlName xml.Name `xml:"xml"`
}

func GetXml() {

	repText := xml{
		//因为是发送，所以把接收到的两个值换一下就行了
		ToUserName:   "touser",
		FromUserName: "from",
		CreateTime:   time.Now().Unix(),
		MsgType:      "text",
		Content:      "content",
	}
	//格式化为xml
	msg, err := xmll.Marshal(&repText)
	if err != nil {
		log.Println("error:拼接xml失败")
		return
	}

	fmt.Println(string(msg))
}
