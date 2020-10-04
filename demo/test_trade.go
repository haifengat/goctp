package main

import (
	"fmt"
	"hf_go_ctp"
)

func OnConnect() {
	fmt.Println("connected")
	t.ReqLogin(investorID, password, brokerID, appID, authCode)
}

// TestTrade 交易测试
func TestTrade(tradeFront, investorID, password, brokerID, appID, authCode string) {
	fmt.Println("connected to trade...")
	t.RegOnFrontConnected(OnConnect)
	t.RegOnRspUserLogin(func(login *hf_go_ctp.RspUserLoginField, info *hf_go_ctp.RspInfoField) {
		fmt.Println(info)
		fmt.Printf("login info: %v\n", *login)
		t.ReqOrderInsertMarket("rb2011", hf_go_ctp.DirectionBuy, hf_go_ctp.OffsetFlagOpen, 1)
		// _ = t.ReqOrderInsert("rb2001", DirectionBuy, OffsetFlagOpen, 3000, 1)
	})
	t.RegOnRtnOrder(func(field *hf_go_ctp.OrderField) {
		fmt.Printf("%v\n", field)
	})
	t.RegOnErrRtnOrder(func(field *hf_go_ctp.OrderField, info *hf_go_ctp.RspInfoField) {
		fmt.Printf("%v\n", info)
	})
	t.ReqConnect(tradeFront)
}
