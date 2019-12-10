package main

import (
	"encoding/json"
	"fmt"
	. "hf_go_ctp/common"
	//. "hf_go_ctp/go_ctp_win"
	. "hf_go_ctp/go_ctp_lnx" // linux使用
	"os"
	"os/signal"
	"syscall"
)

var (
	tradeFront   = "tcp://180.168.146.187:10101"
	quoteFront   = "tcp://180.168.146.187:10111"
	brokerID     = "9999"
	investorID   = "008105"
	password     = "1"
	appID        = "simnow_client_test"
	authCode     = "0000000000000000"
	instrumentID = "rb2001"
	// 交易登录是否成功
	chanTradeLogged = make(chan bool, 1)
	t               = NewTrade() // 放在函数里 linux 会出现登录后断开的情况
	q               = NewQuote()
)

func exit() {
	signals := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	// `signal.Notify` registers the given channel to
	// receive notifications of the specified signals.
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	// This goroutine executes a blocking receive for
	// signals. When it gets one it'll print it out
	// and then notify the program that it can finish.
	go func() {
		sig := <-signals
		fmt.Println(sig)
		done <- true
	}()
	// The program will wait here until it gets the
	// expected signal (as indicated by the goroutine
	// above sending a value on `done`) and then exit.
	<-done
	fmt.Println("exiting")
}

func testTrade() {
	t := NewTrade()
	fmt.Println("connected to trade...")
	t.RegOnFrontConnected(func() {
		fmt.Println("connected")
		t.ReqLogin(investorID, password, brokerID, appID, authCode)
	})
	t.RegOnRspUserLogin(func(login *RspUserLoginField, info *RspInfoField) {
		fmt.Println(info)
		fmt.Printf("login info: %v\n", *login)
		//id := t.ReqOrderInsertMarket("rb2001", go_ctp.DirectionBuy, go_ctp.OffsetFlagOpen, 1)
		//_ = t.ReqOrderInsert("rb2001", DirectionBuy, OffsetFlagOpen, 3000, 1)
		chanTradeLogged <- true
	})
	t.RegOnRtnOrder(func(field *OrderField) {
		fmt.Printf("%v\n", field)
	})
	t.RegOnErrRtnOrder(func(field *OrderField, info *RspInfoField) {
		fmt.Printf("%v\n", info)
	})
	t.ReqConnect(tradeFront)
}

func testQuote() {
	q.RegOnFrontConnected(func() {
		fmt.Println("connected")
		q.ReqLogin(investorID, password, brokerID)
	})
	q.RegOnRspUserLogin(func(login *RspUserLoginField, info *RspInfoField) {
		fmt.Println(info)
		q.ReqSubscript([]string{instrumentID})
	})
	q.RegOnTick(func(data *TickField) {
		bs, _err := json.Marshal(data)
		if _err == nil {
			println(string(bs))
		}
	})
	q.ReqConnect(quoteFront)
}

func main() {
	defer t.Release()
	defer q.Release()
	go testTrade()
	<-chanTradeLogged
	testQuote() // 不能同时测试交易
	exit()
}
