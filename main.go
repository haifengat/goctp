package main

import (
	"fmt"
	"go_ctp"
	"time"
)

func main() {
	//go_ctp.GenerateDataType()
	//go_ctp.GenerateStruct()
	//go_ctp.GenerateMd()
	go_ctp.InitQuote()
	//go_ctp.RegisterSpi(0)
	addr := "tcp://180.168.146.187:10000"
	go_ctp.RegisterFront(addr)
	go_ctp.Init()
	time.Sleep(time.Duration(3) * time.Second)
	id := go_ctp.TThostFtdcUserIDType{'0', '0', '7', '1', '0', '5'}
	pwd := go_ctp.TThostFtdcPasswordType{'1'}
	br := go_ctp.TThostFtdcBrokerIDType{'9', '9', '9', '9'}
	req := go_ctp.CThostFtdcReqUserLoginField{
		UserID:   id,
		Password: pwd,
		BrokerID: br,
	}
	go_ctp.ReqUserLogin(req, 1)
	time.Sleep(time.Duration(3) * time.Second)
	go_ctp.SubscribeMarketData([]string{"cu1905"}, 1)
	fmt.Scanln()
}
