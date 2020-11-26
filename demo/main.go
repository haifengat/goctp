package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/haifengat/goctp"
	ctp "github.com/haifengat/goctp/lnx"
	// ctp "github.com/haifengat/goctp/win"
)

var (
	tradeFront   = "tcp://180.168.146.187:10101" //10130
	quoteFront   = "tcp://180.168.146.187:10111" //10131
	brokerID     = "9999"
	investorID   = "008107"
	password     = "1"
	appID        = "simnow_client_test"
	authCode     = "0000000000000000"
	instrumentID = "rb2101"
)

var t = ctp.NewTrade()
var q = ctp.NewQuote()

func onTick(data *goctp.TickField) {
	if bs, err := json.Marshal(data); err == nil {
		println(string(bs))
	} else {
		fmt.Print("ontick")
	}
}

func testQuote() {
	q.RegOnFrontConnected(func() {
		fmt.Println("quote connected")
		q.ReqLogin(investorID, password, brokerID)
	})
	q.RegOnRspUserLogin(func(login *goctp.RspUserLoginField, info *goctp.RspInfoField) {
		fmt.Println("quote login:", info)
		q.ReqSubscript("rb2011")
	})
	q.RegOnTick(onTick)
	fmt.Print("quote connecting...")
	q.ReqConnect(quoteFront)
}

func testTrade() {
	fmt.Println("connected to trade...")
	t.RegOnFrontConnected(func() {
		fmt.Println("connected")
		go t.ReqLogin(investorID, password, brokerID, appID, authCode)
	})
	t.RegOnRspUserLogin(func(login *goctp.RspUserLoginField, info *goctp.RspInfoField) {
		fmt.Println(info)
		fmt.Printf("trade login info: %v\n", *login)
		// t.ReqOrderInsertMarket("rb2011", goctp.DirectionBuy, goctp.OffsetFlagOpen, 1)
		// _ = t.ReqOrderInsert("rb2001", DirectionBuy, OffsetFlagOpen, 3000, 1)
	})
	t.RegOnRtnOrder(func(field *goctp.OrderField) {
		fmt.Printf("%v\n", field)
	})
	t.RegOnErrRtnOrder(func(field *goctp.OrderField, info *goctp.RspInfoField) {
		fmt.Printf("%v\n", info)
	})
	t.RegOnRtnInstrumentStatus(func(field *goctp.InstrumentStatus) {
		fmt.Printf("%v\n", field)
	})
	t.ReqConnect(tradeFront)
}

func main() {
	go testQuote() // 不能同时测试交易
	go testTrade()
	for !t.IsLogin {
		time.Sleep(10 * time.Second)
	}
	t.Instruments.Range(func(k, v interface{}) bool {
		fmt.Printf("%v", v)
		return true
	})
	for {

	}
}
