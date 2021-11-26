# goctp

## 介绍

CTP 封装之 golang 版,支持 Windows Linux x64.
采用二次封装，将 C++封装成 C，并 export.
win lnx 封装逻辑相同: trade.go quote.go

## 运行

### 环境变量

```json
"tradeFront": "tcp://180.168.146.187:10130",
"quoteFront": "tcp://180.168.146.187:10131",
"loginInfo": "9999/008107/1/simnow_client_test/0000000000000000"
```

### 看穿式评测

```shell
# 替换接口库文件
cp CTPv6.3.19_20200423_cp/*.so lnx/
cp CTPv6.3.19_20200423_cp/*.dll win/
```

### 示例

```go
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"gitee.com/haifengat/goctp"
	// ctp "gitee.com/haifengat/goctp/lnx" // linx
	ctp "gitee.com/haifengat/goctp/win"  // windows
)

var (
	investorID, password, brokerID, appID, authCode string
	tradeFront, quoteFront, loginInfo               string
)

var t = ctp.NewTrade()
var q = ctp.NewQuote()

func init() {
	tradeFront = os.Getenv("tradeFront")
	quoteFront = os.Getenv("quoteFront")
	loginInfo = os.Getenv("loginInfo")
	fs := strings.Split(loginInfo, "/")
	brokerID, investorID, password, appID, authCode = fs[0], fs[1], fs[2], fs[3], fs[4]
	fmt.Println("tradeFront: ", tradeFront)
	fmt.Println("quoteFront: ", quoteFront)
	fmt.Printf("brokerID:%s\ninvestorID:%s\npassword:%s\nappID:%s\nauthCode:%s\n", brokerID, investorID, password, appID, authCode)
}

func testQuote() {
	q.RegOnFrontConnected(func() {
		fmt.Println("quote connected")
		q.ReqLogin(investorID, password, brokerID)
	})
	q.RegOnRspUserLogin(func(login *goctp.RspUserLoginField, info *goctp.RspInfoField) {
		fmt.Println("quote login:", info)
	})
	q.RegOnTick(func(tick *goctp.TickField) {
		bs, _ := json.Marshal(tick)
		fmt.Println(string(bs))
	})
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
		bs, _ := json.Marshal(login)
		fmt.Println("login: ", string(bs))
	})

	t.RegOnRtnOrder(func(field *goctp.OrderField) {
		bs, _ := json.Marshal(field)
		fmt.Println("OnRtnOrder:", string(bs))
	})
	t.RegOnRtnCancel(func(field *goctp.OrderField) {
		bs, _ := json.Marshal(field)
		fmt.Println("OnRtnCancel: ", string(bs))
	})
	t.RegOnErrRtnOrder(func(field *goctp.OrderField, info *goctp.RspInfoField) {
		bs, _ := json.Marshal(info)
		fmt.Println("OnErrRtnOrder: ", string(bs))
	})
	// 交易状态
	t.RegOnRtnInstrumentStatus(func(field *goctp.InstrumentStatus) {
		// fmt.Println(field)
	})
	// 断开
	t.RegOnFrontDisConnected(func(reason int) {
		fmt.Println("disconntcted: ", reason)
	})
	fmt.Println("connecting to trade " + tradeFront)
	t.ReqConnect(tradeFront)
}

func main() {
	go testQuote()
	go testTrade()
	for !t.IsLogin {
		time.Sleep(1 * time.Second)
	}

	time.Sleep(3 * time.Second)
	// 委托测试
	if true {
		t.ReqOrderInsert("rb2205", goctp.DirectionBuy, goctp.OffsetFlagOpen, 4000, 1)
	}
	// 合约
	if true {
		cnt := 0
		t.Instruments.Range(func(k, v interface{}) bool {
			cnt++
			return true
		})
		fmt.Println("instrument count:", cnt)
	}
	// 权益
	if true {
		bs, _ := json.Marshal(t.Account)
		fmt.Println(string(bs))
	}
	// 委托信息
	if true {
		t.Orders.Range(func(key, value interface{}) bool {
			fmt.Printf("%s: %v\n", key, value)
			return true
		})
	}
	// 成交信息
	if true {
		t.Trades.Range(func(key, value interface{}) bool {
			fmt.Printf("%s: %v\n", key, value)
			return true
		})
	}
	// 持仓
	if true {
		t.Positions.Range(func(key, value interface{}) bool {
			p := value.(*goctp.PositionField)
			fmt.Printf("%s: %s: 昨：%d,今：%d,总: %d, 可: %d\n", key, p.InstrumentID, p.YdPosition, p.TodayPosition, p.Position, p.Position-p.ShortFrozen)
			return true
		})
		time.Sleep(500 * time.Millisecond)
	}
	// 入金
	if false {
		t.RegOnRtnFromFutureToBank(func(field *goctp.TransferField) {
			fmt.Println(field)
		})
		t.ReqFutureToBank("", "", 30)
	}
	// 订阅合约
	if true {
		q.ReqSubscript("rb2205")
	}

	fmt.Scanf("exit: ")
	t.Release()
	q.Release()
}
```

## 版本切换

复制 6.5.1 版本以上官方库文件覆盖到 lnx win 下同名文件即可。

## QA

### operator delete(void\*, unsigned long)@CXXABI_1.3.9’未定义的引用

> 不同系统，不同版本的底层依赖不同
> 重新编译即可

```bash
cd lnx && g++ -shared -fPIC -Wl,-rpath . -o ./libctp_quote.so ../generate/quote.cpp  thostmduserapi_se.so && cd ..
cd lnx && g++ -shared -fPIC -Wl,-rpath . -o ./libctp_trade.so ../generate/trade.cpp  thosttraderapi_se.so && cd ..
```
