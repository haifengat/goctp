package main

import (
	"fmt"
	"os"
	"time"

	"gitee.com/haifengat/goctp"
	// ctp "gitee.com/haifengat/goctp/lnx"
	ctp "gitee.com/haifengat/goctp/win"
)

/*
appid:simnow_client_test
authcode:0000000000000000
*/
var (
	userID   = "008107"
	password = "1"
	brokerID = "9999"
	appID    = "simnow_client_test"
	authCode = "0000000000000000"
	// tradeFront = "tcp://180.168.146.187:10201"
	// quoteFront = "tcp://180.168.146.187:10211"
	// tradeFront = "tcp://180.168.146.187:10202"
	// quoteFront = "tcp://180.168.146.187:10212"
	tradeFront = "tcp://180.168.146.187:10130"
	quoteFront = "tcp://180.168.146.187:10131"
)

var t *ctp.Trade
var q *ctp.Quote

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

	t = ctp.NewTrade()
	// t.SetQuick() // quick 模式, 处理指定帐号
	q = ctp.NewQuote()
}

func releaseQuote() {
	q.Release()
	q = ctp.NewQuote()
}

func testQuote() {
	chConnected := make(chan bool)
	q.RegOnFrontConnected(func() {
		chConnected <- true
		fmt.Println("quote connected")
		q.ReqLogin(userID, password, brokerID)
	})
	q.RegOnRspUserLogin(func(login *goctp.RspUserLoginField, info *goctp.RspInfoField) {
		fmt.Printf("quote login: %+v\n", info)
		// 只有第一个合约有效
		go q.ReqSubMarketData("rb2403")
	})
	q.RegOnTick(func(tick *goctp.TickField) {
		fmt.Printf("%+v\n", tick)
	})
	q.RegOnFrontDisConnected(func(reason int) {
		fmt.Println("quote disconected ", reason)
		if reason != 0 {
			releaseQuote()
		}
	})
	fmt.Println("connecting to quote " + quoteFront)
	q.ReqConnect(quoteFront)
	go func() {
		// 连接超时设置
		select {
		case <-chConnected:
		case <-time.After(10 * time.Second):
			fmt.Println("连接超时")
			releaseQuote()
		}
	}()
}

func releaseTrade() {
	t.Release()
	t = ctp.NewTrade()
}

func testTrade() {
	chConnected := make(chan bool)
	t.RegOnFrontConnected(func() {
		chConnected <- true
		fmt.Println("trade connected")
		go t.ReqLogin(userID, password, brokerID, appID, authCode)
	})

	t.RegOnRspUserLogin(func(login *goctp.RspUserLoginField, info *goctp.RspInfoField) {
		fmt.Printf("%+v\n", info)
		if info.ErrorID == 7 { // 密码错误
			go func() {
				time.Sleep(1 * time.Minute) // 一分钟试一次
				t.ReqLogin(userID, password, brokerID, appID, authCode)
			}()
		} else if info.ErrorID != 0 {
			go releaseTrade()
		} else {
			fmt.Printf("login: %+v\n", login)
			// fmt.Println("investors: ", t.Investors)
		}
	})

	t.RegOnRtnOrder(func(field *goctp.OrderField) {
		// fmt.Printf("OnRtnOrder: %+v\n", field)
	})
	t.RegOnRtnTrade(func(field *goctp.TradeField) {
		// fmt.Printf("OnRtnTrade: %+v\n", field)
	})
	t.RegOnRtnCancel(func(field *goctp.OrderField) {
		// fmt.Printf("OnRtnCancel: %+v\n", field)
	})
	t.RegOnErrRtnOrder(func(field *goctp.OrderField, info *goctp.RspInfoField) {
		// fmt.Printf("OnErrRtnOrder: %+v\n", field)
	})
	// 交易状态
	t.RegOnRtnInstrumentStatus(func(field *goctp.InstrumentStatus) {
		// fmt.Println(field)
	})
	// 断开
	t.RegOnFrontDisConnected(func(reason int) {
		fmt.Println("trade disconnected ", reason)
		if reason != 0 {
			releaseTrade()
		}
	})
	fmt.Println("connecting to trade " + tradeFront)
	t.ReqConnect(tradeFront)
	go func() {
		// 连接超时设置
		select {
		case <-chConnected:
		case <-time.After(10 * time.Second):
			fmt.Println("连接超时")
			releaseTrade()
		}
	}()
}

func run724() {
	// 应用启动
	testQuote()
	testTrade()
	tick := time.NewTicker(1 * time.Minute)
	for range tick.C {
		hhmm := time.Now().Local().Format("15:04")
		if t.IsLogin { // 收盘后退出
			isContinue := false
			t.InstrumentStatuss.Range(func(key, value interface{}) bool {
				if value.(*goctp.InstrumentStatus).InstrumentStatus == goctp.InstrumentStatusContinous {
					isContinue = true
					return false
				}
				return true
			})
			if time.Now().Local().Minute()%15 == 0 { // 15 分钟显示一次
				fmt.Println(hhmm, " 交易中: ", isContinue)
			}
			if !isContinue {
				hour := time.Now().Local().Hour()
				if hour <= 3 || hour == 15 { // 夜盘结束&当日收盘
					fmt.Println("release at ", time.Now().Local())
					releaseTrade()
				}
			}
		} else { // 定时连接
			if time.Now().Local().Minute()%15 == 0 { // 15 分钟显示一次
				fmt.Println(hhmm, " 接口关闭")
			}
			if hhmm == "08:40" || hhmm == "20:40" {
				fmt.Println("relogin at ", hhmm)
				testQuote()
				testTrade()
			}
		}
	}
}

func main() {
	go run724()

	for t == nil || !t.IsLogin {
		time.Sleep(1 * time.Second)
	}

	// 委托测试
	if true {
		// t.ReqOrderInsert("rb2210", goctp.DirectionBuy, goctp.OffsetFlagClose, 2850, 2)
		// t.ReqOrderInsertByUser("00200008", "rb2210", goctp.DirectionBuy, goctp.OffsetFlagOpen, 2850, 2)
	}
	// 合约
	if false {
		cnt := 0
		t.Instruments.Range(func(k, v interface{}) bool {
			cnt++
			return true
		})
		fmt.Println("instrument count:", cnt)
	}
	// 权益
	if false {
		fmt.Printf("Account: %+v\n", t.Account)
	}
	// 委托信息
	if false {
		t.Orders.Range(func(key, value interface{}) bool {
			fmt.Printf("%s: %+v\n", key, value)
			return true
		})
	}
	// 成交信息
	if false {
		t.Trades.Range(func(key, value interface{}) bool {
			fmt.Printf("%s: %+v\n", key, value)
			return true
		})
	}
	// 入金
	if false {
		t.RegOnRtnFromFutureToBank(func(field *goctp.TransferField) {
			fmt.Printf("入金: %+v\n", field)
		})
		t.ReqFutureToBank("", "", 30)
	}
	// 订阅合约
	if false {
		q.ReqSubMarketData("rb2409")
	}

	// 行情
	if false {
		time.Sleep(5 * time.Second)
		q.Ticks.Range(func(key, value interface{}) bool {
			fmt.Printf("%+v", value)
			return true
		})
	}
	// 权益
	if false {
		for { // 检查持仓/权益查询是否生效
			for k, v := range t.UserAccounts {
				fmt.Printf("%s 权益: %+v\n", k, v)
			}
			// fmt.Printf("%+v\n", t.UserAccounts)
			time.Sleep(3 * time.Second)
		}
	}
	// 持仓
	if false {
		for k, v := range t.UserPositions {
			v.Range(func(key, value interface{}) bool {
				p := value.(*goctp.PositionField)
				fmt.Printf("%s || %s: %s: 昨：%d,今：%d,总: %d, 可: %d\n", k, key, p.InstrumentID, p.YdPosition, p.TodayPosition, p.Position, p.Position-p.ShortFrozen)
				return true
			})
		}
	}

	// 交易所状态
	if false {
		t.InstrumentStatuss.Range(func(key, value interface{}) bool {
			fmt.Printf("%+v\n", *value.(*goctp.InstrumentStatus))
			return true
		})
	}
	select {}
}
