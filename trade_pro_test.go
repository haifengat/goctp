package goctp

import (
	"fmt"
	"testing"
)

func TestTradePro(t *testing.T) {
	trd := NewTradePro()
	logInfo, rsp := trd.Start(LoginConfig{
		Front:    "tcp://180.168.146.187:10130",
		Broker:   "9999",
		UserID:   "008107",
		Password: "1",
		AppID:    "simnow_client_test",
		AuthCode: "0000000000000000",
	})
	if rsp.ErrorID != 0 {
		fmt.Printf("%+v\n", rsp)
		return
	}

	fmt.Printf("登录响应: %+v\n", logInfo)
	fmt.Printf("权益: %+v\n", trd.accounts)

	id, rsp := trd.ReqOrderInsertLimit(THOST_FTDC_D_Buy, THOST_FTDC_OF_Open, "rb2403", 3200, 3)
	if rsp.ErrorID == 0 {
		fmt.Println("委托: ", id)
	} else {
		fmt.Println("委托错误: ", rsp.ErrorMsg.String())
	}
	fmt.Println(trd.AccountRegisters)
	for k := range trd.AccountRegisters {
		trd.ReqFromBankToFutureByFuture(k, 10)
	}
	rsp = trd.ReqFromBankToFutureByFuture("aaa", 10)
	fmt.Println("出入金: ", rsp)
}
