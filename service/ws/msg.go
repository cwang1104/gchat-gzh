package ws

import (
	"fmt"
	"time"
)

func MsgDeal(msg []byte) {
	msgString := string(msg)
	fmt.Println(msgString)
	time.Sleep(time.Second * 2)
	msgChan <- msgString
}
