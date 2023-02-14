package ws

import (
	"fmt"
	"gchat-gzh/cgtp"
	"gchat-gzh/pkg/logger"
	"time"
)

func MsgDeal(msg []byte, id string) {
	msgString := string(msg)
	if msgString != "ping" {
		fmt.Println(msgString)
		time.Sleep(time.Second * 2)
		resp, err := cgtp.NewChatGtp().GetResponse(msgString)
		if err != nil {
			msgString = "有点小错误"
			logger.Log.Error("get resp error")
		}
		msgString = resp
		var msgS = make([]string, 2)
		msgS[0] = id
		msgS[1] = msgString
		msgChan <- msgS
	}
}
