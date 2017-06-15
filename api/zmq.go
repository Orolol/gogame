package main

import (
  "fmt"
	"github.com/zeromq/goczmq"
)


func ZMQPusherMSG() *goczmq.Channeler {
	fmt.Println("Init Pusher")
	push := goczmq.NewDealerChanneler("tcp://127.0.0.1:31337")

	return push
}

func ZMQReader(queueCreation chan [][]byte) {
	fmt.Printf("Init Reader")
	pull := goczmq.NewRouterChanneler("tcp://127.0.0.1:31338")
	for msg := range pull.RecvChan {
		fmt.Println("Recieving new game msg in ZMQ !! TYPE : ", string(msg[1]))
		queueCreation <- msg
	}
}
var ZMQPusher = ZMQPusherMSG()
