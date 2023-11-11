<div><img src='https://img.shields.io/badge/Author-海风-orange.svg'></img></div><div align="right"><img src='https://gitee.com/haifengat/goctp/badge/star.svg?theme=dark' alt='star'></img><img src='https://gitee.com/haifengat/goctp/badge/fork.svg?theme=dark' alt='star'></img></div>

![](https://gitee.com/haifengat/goctp/widgets/widget_card.svg?colors=4183c4,ffffff,ffffff,e3e9ed,666666,9b9b9b)

## 介绍

CTP golang 封装, 封装分三个层次:

- 基础封装 Trade
  - [trade](trade_test.go)
  - 函数用法与官方一致
- 简易封装 TradeExt
  - [trade_ext](trade_ext_test.go)
  - 常用请求函数参数由 struct 变为基础数据类型
- 易用封装 TradePro
  - [TradePro](trade_pro_test.go)
  - 登录过程,包括基础数据查询
  - 委托与成交业务处理
  - 多种委托类型: 限价,市价,FOK,FAK
  - 权益与持仓查询
  - 出入金&改密码
  - 交易员模式(UserID 交易员, InvestorID 投资者)
- 多帐号测试
  - [多帐号](demo/main.go)

## TradePro

### 登录过程中查询的信息

```go
// 投资者 key:InvestorID
Investors map[string]CThostFtdcInvestorField
// 合约 key: InstrumentID
Instruments map[string]CThostFtdcInstrumentField
// 委托 key: OrderLocalID
Orders map[string]CThostFtdcOrderField
// 成交 key: OrderLocalID
Trades map[string][]CThostFtdcTradeField
// 银行开户信息
AccountRegisters map[string]CThostFtdcAccountregisterField
```

### 函数

- Start 登录
- ReqQryPosition 查持仓
- ReqQryPositionDetail 查持仓明细
- ReqQryTradingAccount 查持仓
- ReqFromBankToFutureByFuture 入金
- ReqFromFutureToBankByFuture 出金

#### 委托

- ReqOrderInsertLimit 限价
- ReqOrderInsertFAK 部成或撤单
- ReqOrderInsertFOK 全成或撤单
- ReqOrderInsertMarket 市价单

### 交易员模式

> 普通模式中 InvestorID 与 UserID 相同
> 交易员模式中 UserID 作为交易员帐号登录后,可处理多个 InvestorID 投资者帐号的业务
