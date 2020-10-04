package test_quote

import (
	"fmt"
	"hf_go_ctp"
	"hf_go_ctp/go_ctp_lnx/quote"
)

// TestQuote 行情测试
func TestQuote(quoteFront, investorID, password, brokerID string, onTick hf_go_ctp.OnTickType) {
	q := quote.NewQuote()
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
