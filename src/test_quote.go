package main

import (
	"encoding/json"
	"fmt"
	"hf_go_ctp/src/go_ctp"
)

var q *go_ctp.Quote

func main() {
	defer q.Release()
	q = go_ctp.NewQuote()
	q.RegOnFrontConnected(func() {
		fmt.Println("connected")
		q.ReqLogin("008107", "1", "9999")
	})
	q.RegOnRspUserLogin(func(data go_ctp.RspUserLoginField, info go_ctp.RspInfoField) {
		fmt.Println(info)
		fmt.Println(data.TradingDay)
		q.ReqSubscript("rb2001")
	})
	q.RegOnTick(func(data go_ctp.TickField) {
		bs, _err := json.Marshal(data)
		if _err == nil {
			println(string(bs))
		}
	})
	q.ReqConnect("tcp://180.168.146.187:10111")
	//q.Join()
	fmt.Scanln()
	//select {}
}
