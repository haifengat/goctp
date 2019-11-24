package main

import "hf_go_ctp/src/generate"

func main() {
	generate.GenerateDataType()
	//generate.GenerateStruct()
	//generate.GenerateCtp("quote")
	//generate.GenerateCtp("trade")
	//q := new(generate.Quote)
	//q.Quote()
	//q.RegOnConnected(onCb)
	//q.RegOnUserLogin(onLogin)
	//q.RegisterFront("tcp://180.168.146.187:10000")
	//q.Init()
	//time.Sleep(time.Duration(3) * time.Second)
	//req := generate.CThostFtdcReqUserLoginField{}
	//copy(req.UserID[:], "007105")
	//copy(req.Password[:], "1")
	//copy(req.BrokerID[:], "9999")
	//q.ReqUserLogin(req)
	//time.Sleep(time.Duration(3) * time.Second)
	////q.SubscribeMarketData([]string{"rb1905"}, 1)
	//fmt.Scanln()
}
