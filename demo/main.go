package main

import (
	"fmt"
	"os"
	"time"

	"gitee.com/haifengat/goctp"
	ctp "gitee.com/haifengat/goctp/lnx"
	// ctp "gitee.com/haifengat/goctp/win"
)

/*appid:simnow_client_test
authcode:0000000000000000*/
var (
	userID     = "008107"
	password   = "1"
	brokerID   = "9999"
	appID      = "simnow_client_test"
	authCode   = "0000000000000000"
	tradeFront = "tcp://180.168.146.187:10202"
	quoteFront = "tcp://180.168.146.187:10212"
	// tradeFront = "tcp://180.168.146.187:10130"
	// quoteFront = "tcp://180.168.146.187:10131"
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
}

func testQuote() {
	q = ctp.NewQuote()
	q.RegOnFrontConnected(func() {
		fmt.Println("quote connected")
		q.ReqLogin(userID, password, brokerID)
	})
	q.RegOnRspUserLogin(func(login *goctp.RspUserLoginField, info *goctp.RspInfoField) {
		fmt.Printf("quote login: %+v\n", info)
	})
	q.RegOnTick(func(tick *goctp.TickField) {
		fmt.Printf("%+v", tick)
	})
	q.RegOnFrontDisConnected(func(reason int) {
		fmt.Print("quote disconected ", reason)
	})
	fmt.Println("connecting to quote " + quoteFront)
	q.ReqConnect(quoteFront)
}

func testTrade() {
	ctp.SetQuick() // quick 模式, 处理指定帐号
	t = ctp.NewTrade()
	t.RegOnFrontConnected(func() {
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
			go t.Release()
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
		// t.Release() // 不要在此处 release.  未正常连接后返回4097错误, release会报错: signal: segmentation fault
	})
	fmt.Println("connecting to trade " + tradeFront)
	t.ReqConnect(tradeFront)
}

func main() {
	testQuote()
	testTrade()

	for !t.IsLogin {
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
		q.ReqSubscript("rb2210")
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
	if true {
		for { // 检查持仓/权益查询是否生效
			// for k, v := range t.UserAccounts {
			// 	fmt.Printf("%s 权益: %+v\n", k, v)
			// }
			fmt.Printf("%+v\n", t.UserAccounts["00200008"])
			time.Sleep(3 * time.Second)
		}
	}
	// 持仓
	if true {
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

	// 测试 release
	time.Sleep(3 * time.Second)
	t.Release()
	q.Release()
	select {}
}
