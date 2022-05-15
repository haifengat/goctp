package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"gitee.com/haifengat/goctp"
	ctp "gitee.com/haifengat/goctp/lnx"
	// ctp "gitee.com/haifengat/goctp/win"
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
		if tick.InstrumentID == "" {
			fmt.Print("行情版本错误")
			return
		}
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
	if false {
		t.ReqOrderInsert("rb2205", goctp.DirectionBuy, goctp.OffsetFlagClose, 4300, 2)
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
			if p.Position == 0 {
				return true
			}
			if p.PositionDirection == goctp.PosiDirectionLong {
				fmt.Printf("%s: %s: 昨：%d,今：%d,总: %d, 可: %d\n", key, p.InstrumentID, p.YdPosition, p.TodayPosition, p.Position, p.Position-p.ShortFrozen)
			} else {
				fmt.Printf("%s: %s: 昨：%d,今：%d,总: %d, 可: %d\n", key, p.InstrumentID, p.YdPosition, p.TodayPosition, p.Position, p.Position-p.LongFrozen)
			}
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
		q.ReqSubscript("ag2205")
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println(<-sig)
	t.Release()
	q.Release()
}
