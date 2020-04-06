// 安装：go get github.com/ThreeDotsLabs/watermill
// 作用：watermill一个异步消息解决方案，支持消息重传，消息保存， 内置多种订阅-发布实现，支持kafka和rabbitmq等
// 参考链接：go get github.com/ThreeDotsLabs/watermill

package watermilluse

import (
	"context"
	"fmt"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
	"time"
)
var gochan *gochannel.GoChannel
type WatermillUse struct{
	gochan *gochannel.GoChannel
}
func NewWatermillUse() *WatermillUse{
	wm:=&WatermillUse{}
	if gochan==nil{
		gochan = gochannel.NewGoChannel(
			gochannel.Config{},
			watermill.NewStdLogger(false,false),
		)
	}
	wm.gochan = gochan
	return wm
}

//消息订阅
func (wm *WatermillUse) Subscribe(subId int,topic string){
	messages,err:=wm.gochan.Subscribe(context.Background(),topic)
	if err!=nil{
		fmt.Printf("编号%d 订阅失败,%s \n",subId,err)
		return
	}
	go func() {
		for msg:=range messages{
			fmt.Printf("编号%d 收到消息：%s \n",subId,string(msg.Payload))
			msg.Ack()
		}
	}()
}

//消息发布
func (wm *WatermillUse) Publisher(topic string){
	msg:=message.NewMessage(watermill.NewUUID(),[]byte("hello guohe"))
	err:=wm.gochan.Publish(topic,msg)
	if err!=nil{
		fmt.Println(err)
	}
}
func RunPub(){
	wm:=NewWatermillUse()
	for{
		wm.Publisher("test")
		time.Sleep(time.Second)
	}
}

func RunSub(){
	wm:=NewWatermillUse()
	wm.Subscribe(1000,"test")
	wm.Subscribe(1001,"test")
	wm.Subscribe(1002,"test")
}

