package main

import (
	"fmt"
	"log"
	"github.com/soygul/nbusy-server/gcm/ccs"
)

func main() {
	config := GetConfig()
	fmt.Println(config)
	ccsClient, err := ccs.New(config.GCM.SenderID, config.GCM.APIKey, config.App.Env == "development")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully logged in to GCM.")

	msgCh := make(chan map[string]interface{})
	errCh := make(chan error)

	go ccsClient.Recv(msgCh, errCh)

	ccsMessage := ccs.NewMessage("GCM_TEST_REG_ID")
	ccsMessage.SetData("hello", "world")
	ccsMessage.CollapseKey = ""
	ccsMessage.TimeToLive = 0
	ccsMessage.DelayWhileIdle = true
	ccsClient.Send(ccsMessage)

	fmt.Print("NBusy messege server started.")

	for {
		select {
		case err := <-errCh:
			fmt.Println("err:", err)
		case msg := <-msgCh:
			fmt.Println("msg:", msg)
		}
	}
}
