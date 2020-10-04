# hf_go_ctp

#### 介绍
CTP封装之golang版,支持Windows x86/x64 Linux x64.
采用二次封装，将C++封装成C，并export所用函数.

#### 软件架构
代码与C#/PYTHON版本逻辑相同


#### 使用说明

##### linux 使用
* 编译goctp 具体操作见go_ctp_swig_lns/readme.md

使用了sync.waitGroup判断登录是否成功:
1. 基础查询完成
2. OnRtnOrder旧数据完成
3. 响应中相比CTP,增加了OnCancel,以处理委托被撤单.

测试代码
```go
package main

import (
	"encoding/json"
	"fmt"
	"hf_go_ctp/src/go_ctp"
	"os"
	"os/signal"
	"syscall"
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
		t.ReqLogin("008105", "1", "9999", "simnow_client_test", "0000000000000000")
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
	t.ReqConnect("tcp://180.168.146.187:10101")
	return t
}

func testQuote() {
	q := go_ctp.NewQuote()
	// 自动关闭
	defer q.Release()
	q.RegOnFrontConnected(func() {
		fmt.Println("connected")
		q.ReqLogin("008107", "1", "9999")
	})
	q.RegOnRspUserLogin(func(login *go_ctp.RspUserLoginField, info *go_ctp.RspInfoField) {
		fmt.Println(info)
		q.ReqSubscript("rb2001")
	})
	q.RegOnTick(func(data go_ctp.TickField) {
		bs, _err := json.Marshal(data)
		if _err == nil {
			println(string(bs))
		}
	})
	q.ReqConnect("tcp://180.168.146.187:10111")
}

func main() {
	go func() { testTrade() }()
	exit()
}
```