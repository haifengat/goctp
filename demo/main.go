package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"gitee.com/haifengat/goctp"
	ctp "gitee.com/haifengat/goctp/lnx"
	// ctp "gitee.com/haifengat/goctp/win"
)

var (
	instrumentID = "rb2205"
	tradeFront   = "tcp://180.168.146.187:10201"
	quoteFront   = "tcp://180.168.146.187:10211"
	loginInfo    = "9999/008107/1/simnow_client_test/0000000000000000"

	investorID, password, brokerID, appID, authCode string
)

var t = ctp.NewTrade()
var q = ctp.NewQuote()

func init() {
	fs := strings.Split(loginInfo, "/")
	brokerID, investorID, password, appID, authCode = fs[0], fs[1], fs[2], fs[3], fs[4]
}

func onTick(data *goctp.TickField) {
	if bs, err := json.Marshal(data); err == nil {
		println("tick:" + string(bs))
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
		fmt.Println("orderKey:", field.OrderRef, "|", field.OrderSysID, "|", field.StatusMsg)
	})
	t.RegOnRtnCancel(func(field *goctp.OrderField) {
		fmt.Println("cancel: orderKey:", field.OrderSysID, "|", field.StatusMsg)
	})
	t.RegOnErrRtnOrder(func(field *goctp.OrderField, info *goctp.RspInfoField) {
		fmt.Printf("errRtnOrder: %v\n", info)
	})
	// 交易状态
	t.RegOnRtnInstrumentStatus(func(field *goctp.InstrumentStatus) {
		// fmt.Println(field)
	})
	// 断开
	t.RegOnFrontDisConnected(func(reason int) {
		fmt.Printf("disconntcted: %v\n", reason)
	})
	fmt.Println("connecting to trade " + tradeFront)
	t.ReqConnect(tradeFront)
}

func main() {
	go testQuote() // 不能同时测试交易
	go testTrade()
	for !t.IsLogin {
		time.Sleep(10 * time.Second)
	}

	// t0 := time.Now()
	// for i := 0; i < 800; i++ {
	// 	t.ReqOrderInsert("rb2205", goctp.DirectionBuy, goctp.OffsetFlagOpen, 5300, 1)
	// 	time.Sleep(1 * time.Millisecond)
	// }
	// ms := time.Since(t0).Milliseconds()
	// time.Sleep(5 * time.Second)
	// fmt.Println(ms, "ms") // 200-1083ms|300-993ms|500-1165ms|800-959ms

	// cnt := 0
	// t.Instruments.Range(func(k, v interface{}) bool {
	// 	cnt++
	// 	return true
	// })
	// print("instrument count:", cnt)

	// 持仓
	for {
		t.Positions.Range(func(key, value interface{}) bool {
			// fmt.Printf("%s:%v\n", key, value)
			p := value.(*goctp.PositionField)
			fmt.Printf("%s: %s: 昨：%d,今：%d,总: %d\n", key, p.InstrumentID, p.YdPosition, p.TodayPosition, p.Position)
			return true
		})
		break
		// time.Sleep(3 * time.Second)
	}

	// t.Trades.Range(func(key, value interface{}) bool {
	// 	fmt.Printf("%s: %v\n", key, value)
	// 	return true
	// })

	// t.RegOnRtnFromFutureToBank(func(field *goctp.TransferField) {
	// 	fmt.Print(field)
	// })
	// t.ReqFutureToBank("", "", 30)

	fmt.Scanf("input: ")
	t.Release()
	q.Release()
}
