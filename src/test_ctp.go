package main

import (
	"encoding/json"
	"fmt"
	"hf_go_ctp/src/go_ctp"
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
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()
	// The program will wait here until it gets the
	// expected signal (as indicated by the goroutine
	// above sending a value on `done`) and then exit.
	<-done
	fmt.Println("exiting")
}

func testTrade() *go_ctp.Trade {
	fmt.Println("connected to trade...")
	t := go_ctp.NewTrade()
	t.RegOnFrontConnected(func() {
		fmt.Println("connected")
		t.ReqLogin(investorID, password, brokerID, appID, authCode)
	})
	t.RegOnRspUserLogin(func(login *go_ctp.RspUserLoginField, info *go_ctp.RspInfoField) {
		fmt.Println(info)
		fmt.Printf("login info: %#v\n", *login)
		//id := t.ReqOrderInsertMarket("rb2001", go_ctp.DirectionBuy, go_ctp.OffsetFlagOpen, 1)
		id := t.ReqOrderInsert("rb2001", go_ctp.DirectionBuy, go_ctp.OffsetFlagOpen, 3000, 1)
		fmt.Println(id)
	})
	t.RegOnRtnOrder(func(field *go_ctp.OrderField) {
		fmt.Printf("%#v\n", field)
	})
	t.RegOnErrRtnOrder(func(field *go_ctp.OrderField, info *go_ctp.RspInfoField) {
		fmt.Printf("%#v\n", info)
	})
	t.ReqConnect(tradeFront)
	return t
}

func testQuote() {
	q := go_ctp.NewQuote()
	// 自动关闭
	defer q.Release()
	q.RegOnFrontConnected(func() {
		fmt.Println("connected")
		q.ReqLogin(investorID, password, brokerID)
	})
	q.RegOnRspUserLogin(func(login *go_ctp.RspUserLoginField, info *go_ctp.RspInfoField) {
		fmt.Println(info)
		q.ReqSubscript(instrumentID)
	})
	q.RegOnTick(func(data go_ctp.TickField) {
		bs, _err := json.Marshal(data)
		if _err == nil {
			println(string(bs))
		}
	})
	q.ReqConnect(quoteFront) // mini 8
}

func main() {
	go func() { testTrade() }()
	exit()
}
