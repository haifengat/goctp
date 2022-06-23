package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"gitee.com/haifengat/goctp"
	ctp "gitee.com/haifengat/goctp/lnx"
	// ctp "gitee.com/haifengat/goctp/win"
)

/*appid:simnow_client_test
authcode:0000000000000000*/
var (
	userID     = "008105"
	password   = "1"
	brokerID   = "9999"
	appID      = "simnow_client_test"
	authCode   = "0000000000000000"
	tradeFront = "tcp://180.168.146.187:10202"
	quoteFront = "tcp://180.168.146.187:10212"
)

var t = ctp.NewTradeByUser("zhaoyan")
var q = ctp.NewQuote()

func init() {
	if tmp := os.Getenv("userID"); tmp != "" {
		userID = tmp
	}
	if tmp := os.Getenv("password"); tmp != "" {
		password = tmp
	}
	if tmp := os.Getenv("brokerID"); tmp != "" {
		brokerID = tmp
	}
	if tmp := os.Getenv("appID"); tmp != "" {
		appID = tmp
	}
	if tmp := os.Getenv("authCode"); tmp != "" {
		authCode = tmp
	}
	if tmp := os.Getenv("tradeFront"); tmp != "" {
		tradeFront = tmp
	}
	if tmp := os.Getenv("quoteFront"); tmp != "" {
		quoteFront = tmp
	}
	fmt.Println("tradeFront: ", tradeFront)
	fmt.Println("quoteFront: ", quoteFront)
	fmt.Printf("brokerID:%s\nuserID:%s\npassword:%s\nappID:%s\nauthCode:%s\n", brokerID, userID, password, appID, authCode)
}

func testQuote() {
	q.RegOnFrontConnected(func() {
		fmt.Println("quote connected")
		q.ReqLogin(userID, password, brokerID)
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
		go t.ReqLogin(userID, password, brokerID, appID, authCode)
	})

	t.RegOnRspUserLogin(func(login *goctp.RspUserLoginField, info *goctp.RspInfoField) {
		fmt.Println(info)
		bs, _ := json.Marshal(login)
		fmt.Println("login: ", string(bs))
		fmt.Println(strings.Join(t.Investors, ","))
	})

	t.RegOnRtnOrder(func(field *goctp.OrderField) {
		bs, _ := json.Marshal(field)
		fmt.Println("OnRtnOrder:", string(bs))
	})
	t.RegOnRtnTrade(func(field *goctp.TradeField) {
		bs, _ := json.Marshal(field)
		fmt.Println("OnRtnTrade:", string(bs))
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
		fmt.Println("traded disconnected: ", reason)
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
		t.ReqOrderInsert("rb2210", goctp.DirectionBuy, goctp.OffsetFlagClose, 2850, 2)
		t.ReqOrderInsertByUser("00200008", "rb2210", goctp.DirectionBuy, goctp.OffsetFlagOpen, 2850, 2)
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
	if false {
		bs, _ := json.Marshal(t.Account)
		fmt.Println(string(bs))
	}
	// 委托信息
	if false {
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
	if false {
		q.ReqSubscript("rb2210")
	}

	time.Sleep(3 * time.Second)
	t.Release()
	q.Release()
	sig := make(chan os.Signal, 1)
	fmt.Println(<-sig)
}
