package main

import (
	"fmt"
	"go_ctp"
)

func cbOn(id int, msg string) {
	fmt.Printf("%d %s", id, msg)
	fmt.Println(msg)
}

func main() {
	//go_ctp.GenerateDataType()
	//go_ctp.GenerateStruct()
	go_ctp.GenerateMd()
	//
	//q := new(go_ctp.Quote)
	//q.Quote()
	//q.RegisterFront("tcp://180.168.146.187:10000")
	//q.Init()
	//time.Sleep(time.Duration(3) * time.Second)
	//req := go_ctp.CThostFtdcReqUserLoginField{}
	//copy(req.UserID[:], "007105")
	//copy(req.Password[:], "1")
	//copy(req.BrokerID[:], "9999")
	//q.ReqUserLogin(req)
	//time.Sleep(time.Duration(3) * time.Second)
	//q.SubscribeMarketData([]string{"rb1905"}, 1)
	fmt.Scanln()
}
