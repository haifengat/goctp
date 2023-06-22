<!--
 * @Author: haifengat hubert28@qq.com
 * @Date: 2023-04-22 18:50:17
 * @LastEditors: haifengat hubert28@qq.com
 * @LastEditTime: 2023-04-22 18:50:17
 * @FilePath: /goctp/CHANGELOG.md
 * @Description: 更新记录
-->

# 更新记录

v2.0.6

- 优化: 调整 order 参数顺序
- 修复: 确认结算的 bug

v2.0.5

- 更新: 函数说明
- 修复: OnRtnOrder 时调用应用层的 bug
- 新增: Req 函数增加 int 返回值, 用以控制流控
- 更新: 出入金函数增加资金密码参数

v2.0.4-p3

- 优化: 判断是否为交易员登录(investors 中是否有 userid),简化登录过程
- 修复: AccountRegisters 变量未初始化
- 修复: 交易员模式提示没有确认结算的权限(响应中不判断是否有错误)
- 修复: 交易员模式时未查询到结果
  - ReqQryInvestor 去掉查询中的 investorid
  - ReqQryAccountregister 去掉查询中的 AccountID
- dateType.go 提示“值错误” 改为 string(e)+"非定义值"
- 修复: 出入金响应 bug

v2.0.4-p2

- 修复: Start 函数中 rsp 错误处理

v2.0.4-p1

- 新增: trade_pro OnOrder OnTrade 函数,为应用层提供响应
- 新增: trade_pro StartQuick 快速登录,不查询委托和成交
- 新增: quote_pro Start 登录

v2.0.4 20230409

- 更新: README.md
- 新增: trade_pro OnRtnTrade 处理
- 新增: trade_pro.go FAK FOK Market 委托类型
- 修复: 银转函数 bug
- 新增: trade_pro.go ReqQryTradingAccount ReqQryPositionDetail ReqQryPosition 函数
- 新增: 出入金函数
- 更新: String 函数只取 \x00 前的数据
- 升级: CTP v6.6.9
- 优化: trade_pro.go 保存登录过程中查询的数据
- 优化: trade_pro.go 登录过程中查银转的开户信息
- 增加: ReqOrderInsertLimit 限价委托指令
- 优化: trade.go,quote.go 函数提示
- 优化: struct.go 生成时 reserve ➡️ Reserve 避免出现 unused 警告
- 新增: trade_pro 封装登录函数
- 优化: 函数/变量名

v2.0.3

- 新增: trade_ext.go 撤单 ReqOrderAction
- 新增: trade_ext.go 订阅行情 ReqQryDepthMarketData
- 新增: trade_ext.go 银转相关 ReqQryAccountregister ReqQryTransferBank ReqFromBankToFutureByFuture ReqFromFutureToBankByFuture
- 新增: trade_ext.go 修改密码 ReqUpdateInvestorPwd ReqUpdateUserPwd
- 新增: trade_ext.go 确认结算 ReqSettlementInfoConfirm 查合约 ReqQryInstrument ReqQryClassifiedInstrument 委托 ReqOrderInsert
- 新增: data_type.go 枚举类型 ➡️ String 函数
- 新增: trade_ext.go ReqQryTrade ReqQryPosition ReqQryPositionDetail ReqQryTradingAccount
- 新增: trade_ext.go trade_ext_test.go 交易接口易用性封装
- 新增: quote_ext.go quote_ext_test.go 行情接口易用性封装
- 新增: teade_test.go
- 新增: quote_test.go
- 新增: datetype.go []byte ➡️ String 函数
- 新增: datetype.go double ➡️ StringFormat(6) 函数

v2.0.2

- 修复: trade fn 函数前加 t, set 前加 t, 以免与 quote 接口覆盖
- 新增: trade.go 交易接口
- 新增: trade.h trade.cpp 交易接口
- 新增: quote.go 封装
- 新增: demo/main.go 行情示例
- 新增: 完整封装 Quote 对应的 quote.h quote.cpp
- 新增: struct.go 生成器
- 优化: datatype 生成器, 模板中使用 $ 全局变量
- 新增: dataType.go 生成器
- 更新: 接口文档编码 gb2312->utf8

v2.0.1

- 新增: 响应函数(OnRspUserLogin) 参数 struct CThostFtdcRspInfoField
- 新增: 响应函数(OnRspUserLogin)
- 新增: 响应函数(OnFrontConnected)
- 新增: 调用函数(CreateFtdcMdSpi,RegisterSpi,SetOnFrontConnected)
- 新增: main.go 中 -L 和 -rpath 参数指向 lib 目录,与 make.sh 一致
- 新增: make.sh 编译 so
- 新增: 调用函数(CreateFtdcMdApi,RegisterFront,Init)
