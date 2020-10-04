package main

import (
	"encoding/json"
	"fmt"
	"hf_go_ctp"
	"hf_go_ctp/demo/test_quote"
	"hf_go_ctp/go_ctp_lnx/trade"
	// "hf_go_ctp/go_ctp_win"
)

var (
	tradeFront   = "tcp://180.168.146.187:10130" //10130
	quoteFront   = "tcp://180.168.146.187:10131" //10131
	brokerID     = "9999"
	investorID   = "008105"
	password     = "1"
	appID        = "simnow_client_test"
	authCode     = "0000000000000000"
	instrumentID = "rb2001"
)

var t = trade.NewTrade()

func onTick(data *hf_go_ctp.TickField) {
	if bs, err := json.Marshal(data); err == nil {
		println(string(bs))
	} else {
		fmt.Print("ontick")
	}
}

// export LD_LIBRARY_PATH=/tmp/src/gitee.com/haifengat/hf_go_ctp/demo/lib64/:$LD_LIBRARY_PATH
func main() {
	go test_quote.TestQuote(quoteFront, investorID, password, brokerID, onTick) // 不能同时测试交易
	go TestTrade(tradeFront, investorID, password, brokerID, appID, authCode)
	for {

	}
}
