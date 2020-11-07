package main

import (
	"encoding/json"
	"fmt"
	"hf_go_ctp"
	ctp "hf_go_ctp/go_ctp_lnx"
	// ctp "hf_go_ctp/go_ctp_win"
)

var (
	tradeFront   = "tcp://180.168.146.187:10130" //10130
	quoteFront   = "tcp://180.168.146.187:10131" //10131
	brokerID     = "9999"
	investorID   = "008107"
	password     = "1"
	appID        = "simnow_client_test"
	authCode     = "0000000000000000"
	instrumentID = "rb2101"
)

var t = ctp.NewTrade()
var q = ctp.NewQuote()

func onTick(data *hf_go_ctp.TickField) {
	if bs, err := json.Marshal(data); err == nil {
		println(string(bs))
	} else {
		fmt.Print("ontick")
	}
}
func TestQuote() {
	q.RegOnFrontConnected(func() {
		fmt.Println("connected")
		q.ReqLogin(investorID, password, brokerID)
	})
	q.RegOnRspUserLogin(func(login *hf_go_ctp.RspUserLoginField, info *hf_go_ctp.RspInfoField) {
		fmt.Println("quote login:", info)
		q.ReqSubscript("rb2011")
	})
	q.RegOnTick(onTick)
	q.ReqConnect(quoteFront)
}

func TestTrade() {
	fmt.Println("connected to trade...")
	t.RegOnFrontConnected(func() {
		fmt.Println("connected")
		go t.ReqLogin(investorID, password, brokerID, appID, authCode)
	})
	t.RegOnRspUserLogin(func(login *hf_go_ctp.RspUserLoginField, info *hf_go_ctp.RspInfoField) {
		fmt.Println(info)
		fmt.Printf("login info: %v\n", *login)
		// t.ReqOrderInsertMarket("rb2011", hf_go_ctp.DirectionBuy, hf_go_ctp.OffsetFlagOpen, 1)
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

func main() {
	go TestQuote() // 不能同时测试交易
	go TestTrade()
	for {

	}
}
