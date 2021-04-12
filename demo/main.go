package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"gitee.com/haifengat/goctp"
	// ctp "gitee.com/haifengat/goctp/lnx"
	ctp "gitee.com/haifengat/goctp/win"
)

var (
	instrumentID = "rb2105"
	tradeFront   = "tcp://180.168.146.187:10101"
	quoteFront   = "tcp://180.166.132.68:41213"
	loginInfo    = "9999/008105/1/simnow_client_test/0000000000000000"

	investorID, password, brokerID, appID, authCode string
)

var t = ctp.NewTrade()
var q = ctp.NewQuote()

func init() {
	fs := strings.Split(loginInfo, "|")
	brokerID, investorID, password, appID, authCode = fs[0], fs[1], fs[2], fs[3], fs[4]
}

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
		q.ReqSubscript(instrumentID)
	})
	q.RegOnTick(onTick)
	fmt.Println("quote connecting " + quoteFront)
	q.ReqConnect(quoteFront)
}

func testTrade() {
	t.RegOnFrontConnected(func() {
		fmt.Println("trade connected")
		go t.ReqLogin(investorID, password, brokerID, appID, authCode)
	})
	t.RegOnRspUserLogin(func(login *goctp.RspUserLoginField, info *goctp.RspInfoField) {
		fmt.Println(info)
		fmt.Printf("trade login info: %v\n", *login)
	})
	t.RegOnRtnOrder(func(field *goctp.OrderField) {
		// fmt.Printf("%v\n", field)
		t.Orders.Range(func(key, value interface{}) bool {
			fmt.Print(key)
			return true
		})
	})
	t.RegOnErrRtnOrder(func(field *goctp.OrderField, info *goctp.RspInfoField) {
		fmt.Printf("%v\n", info)
	})
	t.RegOnRtnInstrumentStatus(func(field *goctp.InstrumentStatus) {

	})
	fmt.Println("connected to trade " + tradeFront)
	t.ReqConnect(tradeFront)
}

func main() {
	go testQuote() // 不能同时测试交易
	go testTrade()
	for !t.IsLogin {
		time.Sleep(10 * time.Second)
	}
	cnt := 0
	// t.ReqOrderInsertMarket("rb2105", goctp.DirectionBuy, goctp.OffsetFlagOpen, 1)
	time.Sleep(3 * time.Second)
	// key := t.ReqOrderInsert("rb2105", goctp.DirectionBuy, goctp.OffsetFlagOpen, 3000, 1)
	// print(key)
	t.Instruments.Range(func(k, v interface{}) bool {
		// fmt.Printf("%v", v)
		cnt++
		return true
	})
	print("instrument count:", cnt)
	for {

	}
}
