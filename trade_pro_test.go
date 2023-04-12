package goctp

import (
	"fmt"
	"testing"
	"time"
)

func TestTradePro(t *testing.T) {
	trd := NewTradePro()

	trd.OnOrder = func(pOrder *CThostFtdcOrderField) {
		fmt.Println("--------- 委托 -----------")
		fmt.Printf("%+v\n", pOrder)
	}
	trd.OnTrade = func(pTrade *CThostFtdcTradeField) {
		fmt.Println("--------- 成交 -----------")
		fmt.Printf("%+v\n", pTrade)
	}

	logInfo, rsp := trd.Start(LoginConfig{
		// Front:    "tcp://180.168.146.187:10130", // 7*24
		Front:    "tcp://180.168.146.187:10202",
		Broker:   "9999",
		UserID:   "008107",
		Password: "1",
		AppID:    "",
		AuthCode: "0000000000000000",
	})
	if rsp.ErrorID != 0 {
		fmt.Printf("%+v\n", rsp)
		return
	}

	fmt.Println("------------ 登录 ------------")
	fmt.Printf("%+v\n", logInfo)

	fmt.Println("------------ 银行帐户 ------------")
	for k := range trd.AccountRegisters {
		trd.ReqFromBankToFutureByFuture(k, "bankPwd", 10)
		fmt.Printf("入金: %+v\n", rsp)
		trd.ReqFromFutureToBankByFuture(k, 10)
		fmt.Printf("出金: %+v\n", rsp)
	}

	time.Sleep(1 * time.Second)
	fmt.Println("------------ 持仓 ------------")
	ps := trd.ReqQryPosition()
	for _, v := range ps {
		fmt.Printf("%+v\n", v)
	}
	time.Sleep(1 * time.Second)
	fmt.Println("------------ 权益 ------------")
	as := trd.ReqQryTradingAccount()
	for _, v := range as {
		fmt.Printf("%+v\n", v)
	}

	fmt.Println("------------ 委托测试 ------------")
	id, rsp := trd.ReqOrderInsertLimit(THOST_FTDC_D_Buy, THOST_FTDC_OF_Open, "rb2403", 4200, 1)
	if rsp.ErrorID == 0 {
		fmt.Println("委托: ", id)
	} else {
		fmt.Println("委托错误: ", rsp.ErrorMsg.String())
	}
	select {}
}
