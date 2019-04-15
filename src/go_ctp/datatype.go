package go_ctp

// 交易所交易员代码类型
type TThostFtdcTraderIDType [21]byte

// 投资者代码类型
type TThostFtdcInvestorIDType [13]byte

// 经纪公司代码类型
type TThostFtdcBrokerIDType [11]byte

// 经纪公司简称类型
type TThostFtdcBrokerAbbrType [9]byte

// 经纪公司名称类型
type TThostFtdcBrokerNameType [81]byte

// 合约在交易所的代码类型
type TThostFtdcExchangeInstIDType [31]byte

// 报单引用类型
type TThostFtdcOrderRefType [13]byte

// 会员代码类型
type TThostFtdcParticipantIDType [11]byte

// 用户代码类型
type TThostFtdcUserIDType [16]byte

// 密码类型
type TThostFtdcPasswordType [41]byte

// 交易编码类型
type TThostFtdcClientIDType [11]byte

// 合约代码类型
type TThostFtdcInstrumentIDType [31]byte

// 合约标识码类型
type TThostFtdcInstrumentCodeType [31]byte

// 市场代码类型
type TThostFtdcMarketIDType [31]byte

// 产品名称类型
type TThostFtdcProductNameType [21]byte

// 交易所代码类型
type TThostFtdcExchangeIDType [9]byte

// 交易所名称类型
type TThostFtdcExchangeNameType [61]byte

// 交易所简称类型
type TThostFtdcExchangeAbbrType [9]byte

// 交易所标志类型
type TThostFtdcExchangeFlagType [2]byte

// Mac地址类型
type TThostFtdcMacAddressType [21]byte

// 系统编号类型
type TThostFtdcSystemIDType [21]byte

// 交易所属性类型
type TThostFtdcExchangePropertyType byte

// 正常
const TThostFtdcExchangePropertyType_Normal = '0'

// 根据成交生成报单
const TThostFtdcExchangePropertyType_GenOrderByTrade = '1'

// 日期类型
type TThostFtdcDateType [9]byte

// 时间类型
type TThostFtdcTimeType [9]byte

// 长时间类型
type TThostFtdcLongTimeType [13]byte

// 合约名称类型
type TThostFtdcInstrumentNameType [21]byte

// 结算组代码类型
type TThostFtdcSettlementGroupIDType [9]byte

// 报单编号类型
type TThostFtdcOrderSysIDType [21]byte

// 成交编号类型
type TThostFtdcTradeIDType [21]byte

// DB命令类型类型
type TThostFtdcCommandTypeType [65]byte

// IP地址类型
type TThostFtdcIPAddressType [16]byte

// 产品信息类型
type TThostFtdcProductInfoType [11]byte

// 协议信息类型
type TThostFtdcProtocolInfoType [11]byte

// 业务单元类型
type TThostFtdcBusinessUnitType [21]byte

// 出入金流水号类型
type TThostFtdcDepositSeqNoType [15]byte

// 证件号码类型
type TThostFtdcIdentifiedCardNoType [51]byte

// 证件类型类型
type TThostFtdcIdCardTypeType byte

// 组织机构代码
const TThostFtdcIdCardTypeType_EID = '0'

// 中国公民身份证
const TThostFtdcIdCardTypeType_IDCard = '1'

// 军官证
const TThostFtdcIdCardTypeType_OfficerIDCard = '2'

// 警官证
const TThostFtdcIdCardTypeType_PoliceIDCard = '3'

// 士兵证
const TThostFtdcIdCardTypeType_SoldierIDCard = '4'

// 户口簿
const TThostFtdcIdCardTypeType_HouseholdRegister = '5'

// 护照
const TThostFtdcIdCardTypeType_Passport = '6'

// 台胞证
const TThostFtdcIdCardTypeType_TaiwanCompatriotIDCard = '7'

// 回乡证
const TThostFtdcIdCardTypeType_HomeComingCard = '8'

// 营业执照号
const TThostFtdcIdCardTypeType_LicenseNo = '9'

// 营业执照号
const TThostFtdcIdCardTypeType_TaxNo = 'A'

// 港澳居民来往内地通行证
const TThostFtdcIdCardTypeType_HMMainlandTravelPermit = 'B'

// 台湾居民来往大陆通行证
const TThostFtdcIdCardTypeType_TwMainlandTravelPermit = 'C'

// 驾照
const TThostFtdcIdCardTypeType_DrivingLicense = 'D'

// 当地社保ID
const TThostFtdcIdCardTypeType_SocialID = 'F'

// 当地身份证
const TThostFtdcIdCardTypeType_LocalID = 'G'

// 商业登记证
const TThostFtdcIdCardTypeType_BusinessRegistration = 'H'

// 港澳永久性居民身份证
const TThostFtdcIdCardTypeType_HKMCIDCard = 'I'

// 人行开户许可证
const TThostFtdcIdCardTypeType_AccountsPermits = 'J'

// 其他证件
const TThostFtdcIdCardTypeType_OtherCard = 'x'

// 本地报单编号类型
type TThostFtdcOrderLocalIDType [13]byte

// 用户名称类型
type TThostFtdcUserNameType [81]byte

// 参与人名称类型
type TThostFtdcPartyNameType [81]byte

// 错误信息类型
type TThostFtdcErrorMsgType [81]byte

// 字段名类型
type TThostFtdcFieldNameType [2049]byte

// 字段内容类型
type TThostFtdcFieldContentType [2049]byte

// 系统名称类型
type TThostFtdcSystemNameType [41]byte

// 消息正文类型
type TThostFtdcContentType [501]byte

// 投资者范围类型
type TThostFtdcInvestorRangeType byte

// 所有
const TThostFtdcInvestorRangeType_All = '1'

// 投资者组
const TThostFtdcInvestorRangeType_Group = '2'

// 单一投资者
const TThostFtdcInvestorRangeType_Single = '3'

type TThostFtdcDepartmentRangeType byte

// 所有
const TThostFtdcDepartmentRangeType_All = '1'

// 组织架构
const TThostFtdcDepartmentRangeType_Group = '2'

// 单一投资者
const TThostFtdcDepartmentRangeType_Single = '3'

// 数据同步状态类型
type TThostFtdcDataSyncStatusType byte

// 未同步
const TThostFtdcDataSyncStatusType_Asynchronous = '1'

// 同步中
const TThostFtdcDataSyncStatusType_Synchronizing = '2'

// 已同步
const TThostFtdcDataSyncStatusType_Synchronized = '3'

// 经纪公司数据同步状态类型
type TThostFtdcBrokerDataSyncStatusType byte

// 已同步
const TThostFtdcBrokerDataSyncStatusType_Synchronized = '1'

// 同步中
const TThostFtdcBrokerDataSyncStatusType_Synchronizing = '2'

// 交易所连接状态类型
type TThostFtdcExchangeConnectStatusType byte

// 没有任何连接
const TThostFtdcExchangeConnectStatusType_NoConnection = '1'

// 已经发出合约查询请求
const TThostFtdcExchangeConnectStatusType_QryInstrumentSent = '2'

// 已经获取信息
const TThostFtdcExchangeConnectStatusType_GotInformation = '9'

// 交易所交易员连接状态类型
type TThostFtdcTraderConnectStatusType byte

// 没有任何连接
const TThostFtdcTraderConnectStatusType_NotConnected = '1'

// 已经连接
const TThostFtdcTraderConnectStatusType_Connected = '2'

// 已经发出合约查询请求
const TThostFtdcTraderConnectStatusType_QryInstrumentSent = '3'

// 订阅私有流
const TThostFtdcTraderConnectStatusType_SubPrivateFlow = '4'

// 功能代码类型
type TThostFtdcFunctionCodeType byte

// 数据异步化
const TThostFtdcFunctionCodeType_DataAsync = '1'

// 强制用户登出
const TThostFtdcFunctionCodeType_ForceUserLogout = '2'

// 变更管理用户口令
const TThostFtdcFunctionCodeType_UserPasswordUpdate = '3'

// 变更经纪公司口令
const TThostFtdcFunctionCodeType_BrokerPasswordUpdate = '4'

// 变更投资者口令
const TThostFtdcFunctionCodeType_InvestorPasswordUpdate = '5'

// 报单插入
const TThostFtdcFunctionCodeType_OrderInsert = '6'

// 报单操作
const TThostFtdcFunctionCodeType_OrderAction = '7'

// 同步系统数据
const TThostFtdcFunctionCodeType_SyncSystemData = '8'

// 同步经纪公司数据
const TThostFtdcFunctionCodeType_SyncBrokerData = '9'

// 批量同步经纪公司数据
const TThostFtdcFunctionCodeType_BachSyncBrokerData = 'A'

// 超级查询
const TThostFtdcFunctionCodeType_SuperQuery = 'B'

// 预埋报单插入
const TThostFtdcFunctionCodeType_ParkedOrderInsert = 'C'

// 预埋报单操作
const TThostFtdcFunctionCodeType_ParkedOrderAction = 'D'

// 同步动态令牌
const TThostFtdcFunctionCodeType_SyncOTP = 'E'

// 删除未知单
const TThostFtdcFunctionCodeType_DeleteOrder = 'F'

// 经纪公司功能代码类型
type TThostFtdcBrokerFunctionCodeType byte

// 强制用户登出
const TThostFtdcBrokerFunctionCodeType_ForceUserLogout = '1'

// 变更用户口令
const TThostFtdcBrokerFunctionCodeType_UserPasswordUpdate = '2'

// 同步经纪公司数据
const TThostFtdcBrokerFunctionCodeType_SyncBrokerData = '3'

// 批量同步经纪公司数据
const TThostFtdcBrokerFunctionCodeType_BachSyncBrokerData = '4'

// 报单插入
const TThostFtdcBrokerFunctionCodeType_OrderInsert = '5'

// 报单操作
const TThostFtdcBrokerFunctionCodeType_OrderAction = '6'

// 全部查询
const TThostFtdcBrokerFunctionCodeType_AllQuery = '7'

// 全部查询
const TThostFtdcBrokerFunctionCodeType_log = 'a'

// 基本查询：查询基础数据，如合约，交易所等常量
const TThostFtdcBrokerFunctionCodeType_BaseQry = 'b'

// 交易查询：如查成交，委托
const TThostFtdcBrokerFunctionCodeType_TradeQry = 'c'

// 交易功能：报单，撤单
const TThostFtdcBrokerFunctionCodeType_Trade = 'd'

// 银期转账
const TThostFtdcBrokerFunctionCodeType_Virement = 'e'

// 风险监控
const TThostFtdcBrokerFunctionCodeType_Risk = 'f'

// 风险监控
const TThostFtdcBrokerFunctionCodeType_Session = 'g'

// 风控通知控制
const TThostFtdcBrokerFunctionCodeType_RiskNoticeCtl = 'h'

// 风控通知发送
const TThostFtdcBrokerFunctionCodeType_RiskNotice = 'i'

// 察看经纪公司资金权限
const TThostFtdcBrokerFunctionCodeType_BrokerDeposit = 'j'

// 资金查询
const TThostFtdcBrokerFunctionCodeType_QueryFund = 'k'

// 报单查询
const TThostFtdcBrokerFunctionCodeType_QueryOrder = 'l'

// 成交查询
const TThostFtdcBrokerFunctionCodeType_QueryTrade = 'm'

// 持仓查询
const TThostFtdcBrokerFunctionCodeType_QueryPosition = 'n'

// 行情查询
const TThostFtdcBrokerFunctionCodeType_QueryMarketData = 'o'

// 用户事件查询
const TThostFtdcBrokerFunctionCodeType_QueryUserEvent = 'p'

// 风险通知查询
const TThostFtdcBrokerFunctionCodeType_QueryRiskNotify = 'q'

// 出入金查询
const TThostFtdcBrokerFunctionCodeType_QueryFundChange = 'r'

// 投资者信息查询
const TThostFtdcBrokerFunctionCodeType_QueryInvestor = 's'

// 交易编码查询
const TThostFtdcBrokerFunctionCodeType_QueryTradingCode = 't'

// 强平
const TThostFtdcBrokerFunctionCodeType_ForceClose = 'u'

// 压力测试
const TThostFtdcBrokerFunctionCodeType_PressTest = 'v'

// 权益反算
const TThostFtdcBrokerFunctionCodeType_RemainCalc = 'w'

// 净持仓保证金指标
const TThostFtdcBrokerFunctionCodeType_NetPositionInd = 'x'

// 风险预算
const TThostFtdcBrokerFunctionCodeType_RiskPredict = 'y'

// 数据导出
const TThostFtdcBrokerFunctionCodeType_DataExport = 'z'

// 风控指标设置
const TThostFtdcBrokerFunctionCodeType_RiskTargetSetup = 'A'

// 行情预警
const TThostFtdcBrokerFunctionCodeType_MarketDataWarn = 'B'

// 业务通知查询
const TThostFtdcBrokerFunctionCodeType_QryBizNotice = 'C'

// 业务通知模板设置
const TThostFtdcBrokerFunctionCodeType_CfgBizNotice = 'D'

// 同步动态令牌
const TThostFtdcBrokerFunctionCodeType_SyncOTP = 'E'

// 发送业务通知
const TThostFtdcBrokerFunctionCodeType_SendBizNotice = 'F'

// 风险级别标准设置
const TThostFtdcBrokerFunctionCodeType_CfgRiskLevelStd = 'G'

// 交易终端应急功能
const TThostFtdcBrokerFunctionCodeType_TbCommand = 'H'

// 删除未知单
const TThostFtdcBrokerFunctionCodeType_DeleteOrder = 'J'

// 预埋报单插入
const TThostFtdcBrokerFunctionCodeType_ParkedOrderInsert = 'K'

// 预埋报单操作
const TThostFtdcBrokerFunctionCodeType_ParkedOrderAction = 'L'

// 资金不够仍允许行权
const TThostFtdcBrokerFunctionCodeType_ExecOrderNoCheck = 'M'

// 指定
const TThostFtdcBrokerFunctionCodeType_Designate = 'N'

// 证券处置
const TThostFtdcBrokerFunctionCodeType_StockDisposal = 'O'

// 席位资金预警
const TThostFtdcBrokerFunctionCodeType_BrokerDepositWarn = 'Q'

// 备兑不足预警
const TThostFtdcBrokerFunctionCodeType_CoverWarn = 'S'

// 行权试算
const TThostFtdcBrokerFunctionCodeType_PreExecOrder = 'T'

// 行权交收风险
const TThostFtdcBrokerFunctionCodeType_ExecOrderRisk = 'P'

// 持仓限额预警
const TThostFtdcBrokerFunctionCodeType_PosiLimitWarn = 'U'

// 持仓限额查询
const TThostFtdcBrokerFunctionCodeType_QryPosiLimit = 'V'

// 银期签到签退
const TThostFtdcBrokerFunctionCodeType_FBSign = 'W'

// 银期签约解约
const TThostFtdcBrokerFunctionCodeType_FBAccount = 'X'

// 报单操作状态类型
type TThostFtdcOrderActionStatusType byte

// 已经提交
const TThostFtdcOrderActionStatusType_Submitted = 'a'

// 已经接受
const TThostFtdcOrderActionStatusType_Accepted = 'b'

// 已经被拒绝
const TThostFtdcOrderActionStatusType_Rejected = 'c'

// 报单状态类型
type TThostFtdcOrderStatusType byte

// 全部成交
const TThostFtdcOrderStatusType_AllTraded = '0'

// 部分成交还在队列中
const TThostFtdcOrderStatusType_PartTradedQueueing = '1'

// 部分成交不在队列中
const TThostFtdcOrderStatusType_PartTradedNotQueueing = '2'

// 未成交还在队列中
const TThostFtdcOrderStatusType_NoTradeQueueing = '3'

// 未成交不在队列中
const TThostFtdcOrderStatusType_NoTradeNotQueueing = '4'

// 撤单
const TThostFtdcOrderStatusType_Canceled = '5'

// 未知
const TThostFtdcOrderStatusType_Unknown = 'a'

// 尚未触发
const TThostFtdcOrderStatusType_NotTouched = 'b'

// 已触发
const TThostFtdcOrderStatusType_Touched = 'c'

// 报单提交状态类型
type TThostFtdcOrderSubmitStatusType byte

// 已经提交
const TThostFtdcOrderSubmitStatusType_InsertSubmitted = '0'

// 撤单已经提交
const TThostFtdcOrderSubmitStatusType_CancelSubmitted = '1'

// 修改已经提交
const TThostFtdcOrderSubmitStatusType_ModifySubmitted = '2'

// 已经接受
const TThostFtdcOrderSubmitStatusType_Accepted = '3'

// 报单已经被拒绝
const TThostFtdcOrderSubmitStatusType_InsertRejected = '4'

// 撤单已经被拒绝
const TThostFtdcOrderSubmitStatusType_CancelRejected = '5'

// 改单已经被拒绝
const TThostFtdcOrderSubmitStatusType_ModifyRejected = '6'

// 持仓日期类型
type TThostFtdcPositionDateType byte

// 今日持仓
const TThostFtdcPositionDateType_Today = '1'

// 历史持仓
const TThostFtdcPositionDateType_History = '2'

// 持仓日期类型类型
type TThostFtdcPositionDateTypeType byte

// 使用历史持仓
const TThostFtdcPositionDateTypeType_UseHistory = '1'

// 不使用历史持仓
const TThostFtdcPositionDateTypeType_NoUseHistory = '2'

// 交易角色类型
type TThostFtdcTradingRoleType byte

// 代理
const TThostFtdcTradingRoleType_Broker = '1'

// 自营
const TThostFtdcTradingRoleType_Host = '2'

// 做市商
const TThostFtdcTradingRoleType_Maker = '3'

// 产品类型类型
type TThostFtdcProductClassType byte

// 期货
const TThostFtdcProductClassType_Futures = '1'

// 期货期权
const TThostFtdcProductClassType_Options = '2'

// 组合
const TThostFtdcProductClassType_Combination = '3'

// 即期
const TThostFtdcProductClassType_Spot = '4'

// 期转现
const TThostFtdcProductClassType_EFP = '5'

// 现货期权
const TThostFtdcProductClassType_SpotOption = '6'

// 合约生命周期状态类型
type TThostFtdcInstLifePhaseType byte

// 未上市
const TThostFtdcInstLifePhaseType_NotStart = '0'

// 上市
const TThostFtdcInstLifePhaseType_Started = '1'

// 停牌
const TThostFtdcInstLifePhaseType_Pause = '2'

// 到期
const TThostFtdcInstLifePhaseType_Expired = '3'

// 买卖方向类型
type TThostFtdcDirectionType byte

// 买
const TThostFtdcDirectionType_Buy = '0'

// 卖
const TThostFtdcDirectionType_Sell = '1'

// 持仓类型类型
type TThostFtdcPositionTypeType byte

// 净持仓
const TThostFtdcPositionTypeType_Net = '1'

// 综合持仓
const TThostFtdcPositionTypeType_Gross = '2'

// 持仓多空方向类型
type TThostFtdcPosiDirectionType byte

// 净
const TThostFtdcPosiDirectionType_Net = '1'

// 多头
const TThostFtdcPosiDirectionType_Long = '2'

// 空头
const TThostFtdcPosiDirectionType_Short = '3'

// 系统结算状态类型
type TThostFtdcSysSettlementStatusType byte

// 不活跃
const TThostFtdcSysSettlementStatusType_NonActive = '1'

// 启动
const TThostFtdcSysSettlementStatusType_Startup = '2'

// 操作
const TThostFtdcSysSettlementStatusType_Operating = '3'

// 结算
const TThostFtdcSysSettlementStatusType_Settlement = '4'

// 结算完成
const TThostFtdcSysSettlementStatusType_SettlementFinished = '5'

// 费率属性类型
type TThostFtdcRatioAttrType byte

// 交易费率
const TThostFtdcRatioAttrType_Trade = '0'

// 结算费率
const TThostFtdcRatioAttrType_Settlement = '1'

// 投机套保标志类型
type TThostFtdcHedgeFlagType byte

// 投机
const TThostFtdcHedgeFlagType_Speculation = '1'

// 套利
const TThostFtdcHedgeFlagType_Arbitrage = '2'

// 套保
const TThostFtdcHedgeFlagType_Hedge = '3'

// 做市商
const TThostFtdcHedgeFlagType_MarketMaker = '5'

type TThostFtdcBillHedgeFlagType byte

// 投机
const TThostFtdcBillHedgeFlagType_Speculation = '1'

// 套利
const TThostFtdcBillHedgeFlagType_Arbitrage = '2'

// 套保
const TThostFtdcBillHedgeFlagType_Hedge = '3'

// 交易编码类型类型
type TThostFtdcClientIDTypeType byte

// 投机
const TThostFtdcClientIDTypeType_Speculation = '1'

// 套利
const TThostFtdcClientIDTypeType_Arbitrage = '2'

// 套保
const TThostFtdcClientIDTypeType_Hedge = '3'

// 做市商
const TThostFtdcClientIDTypeType_MarketMaker = '5'

// 报单价格条件类型
type TThostFtdcOrderPriceTypeType byte

// 任意价
const TThostFtdcOrderPriceTypeType_AnyPrice = '1'

// 限价
const TThostFtdcOrderPriceTypeType_LimitPrice = '2'

// 最优价
const TThostFtdcOrderPriceTypeType_BestPrice = '3'

// 最新价
const TThostFtdcOrderPriceTypeType_LastPrice = '4'

// 最新价浮动上浮1个ticks
const TThostFtdcOrderPriceTypeType_LastPricePlusOneTicks = '5'

// 最新价浮动上浮2个ticks
const TThostFtdcOrderPriceTypeType_LastPricePlusTwoTicks = '6'

// 最新价浮动上浮3个ticks
const TThostFtdcOrderPriceTypeType_LastPricePlusThreeTicks = '7'

// 卖一价
const TThostFtdcOrderPriceTypeType_AskPrice1 = '8'

// 卖一价浮动上浮1个ticks
const TThostFtdcOrderPriceTypeType_AskPrice1PlusOneTicks = '9'

// 卖一价浮动上浮2个ticks
const TThostFtdcOrderPriceTypeType_AskPrice1PlusTwoTicks = 'A'

// 卖一价浮动上浮3个ticks
const TThostFtdcOrderPriceTypeType_AskPrice1PlusThreeTicks = 'B'

// 买一价
const TThostFtdcOrderPriceTypeType_BidPrice1 = 'C'

// 买一价浮动上浮1个ticks
const TThostFtdcOrderPriceTypeType_BidPrice1PlusOneTicks = 'D'

// 买一价浮动上浮2个ticks
const TThostFtdcOrderPriceTypeType_BidPrice1PlusTwoTicks = 'E'

// 买一价浮动上浮3个ticks
const TThostFtdcOrderPriceTypeType_BidPrice1PlusThreeTicks = 'F'

// 五档价
const TThostFtdcOrderPriceTypeType_FiveLevelPrice = 'G'

// 开平标志类型
type TThostFtdcOffsetFlagType byte

// 开仓
const TThostFtdcOffsetFlagType_Open = '0'

// 平仓
const TThostFtdcOffsetFlagType_Close = '1'

// 强平
const TThostFtdcOffsetFlagType_ForceClose = '2'

// 平今
const TThostFtdcOffsetFlagType_CloseToday = '3'

// 平昨
const TThostFtdcOffsetFlagType_CloseYesterday = '4'

// 强减
const TThostFtdcOffsetFlagType_ForceOff = '5'

// 本地强平
const TThostFtdcOffsetFlagType_LocalForceClose = '6'

// 强平原因类型
type TThostFtdcForceCloseReasonType byte

// 非强平
const TThostFtdcForceCloseReasonType_NotForceClose = '0'

// 资金不足
const TThostFtdcForceCloseReasonType_LackDeposit = '1'

// 客户超仓
const TThostFtdcForceCloseReasonType_ClientOverPositionLimit = '2'

// 会员超仓
const TThostFtdcForceCloseReasonType_MemberOverPositionLimit = '3'

// 持仓非整数倍
const TThostFtdcForceCloseReasonType_NotMultiple = '4'

// 违规
const TThostFtdcForceCloseReasonType_Violation = '5'

// 其它
const TThostFtdcForceCloseReasonType_Other = '6'

// 自然人临近交割
const TThostFtdcForceCloseReasonType_PersonDeliv = '7'

// 报单类型类型
type TThostFtdcOrderTypeType byte

// 正常
const TThostFtdcOrderTypeType_Normal = '0'

// 报价衍生
const TThostFtdcOrderTypeType_DeriveFromQuote = '1'

// 组合衍生
const TThostFtdcOrderTypeType_DeriveFromCombination = '2'

// 组合报单
const TThostFtdcOrderTypeType_Combination = '3'

// 条件单
const TThostFtdcOrderTypeType_ConditionalOrder = '4'

// 互换单
const TThostFtdcOrderTypeType_Swap = '5'

// 有效期类型类型
type TThostFtdcTimeConditionType byte

// 立即完成，否则撤销
const TThostFtdcTimeConditionType_IOC = '1'

// 本节有效
const TThostFtdcTimeConditionType_GFS = '2'

// 当日有效
const TThostFtdcTimeConditionType_GFD = '3'

// 指定日期前有效
const TThostFtdcTimeConditionType_GTD = '4'

// 撤销前有效
const TThostFtdcTimeConditionType_GTC = '5'

// 集合竞价有效
const TThostFtdcTimeConditionType_GFA = '6'

// 成交量类型类型
type TThostFtdcVolumeConditionType byte

// 任何数量
const TThostFtdcVolumeConditionType_AV = '1'

// 最小数量
const TThostFtdcVolumeConditionType_MV = '2'

// 全部数量
const TThostFtdcVolumeConditionType_CV = '3'

// 触发条件类型
type TThostFtdcContingentConditionType byte

// 立即
const TThostFtdcContingentConditionType_Immediately = '1'

// 止损
const TThostFtdcContingentConditionType_Touch = '2'

// 止赢
const TThostFtdcContingentConditionType_TouchProfit = '3'

// 预埋单
const TThostFtdcContingentConditionType_ParkedOrder = '4'

// 最新价大于条件价
const TThostFtdcContingentConditionType_LastPriceGreaterThanStopPrice = '5'

// 最新价大于等于条件价
const TThostFtdcContingentConditionType_LastPriceGreaterEqualStopPrice = '6'

// 最新价小于条件价
const TThostFtdcContingentConditionType_LastPriceLesserThanStopPrice = '7'

// 最新价小于等于条件价
const TThostFtdcContingentConditionType_LastPriceLesserEqualStopPrice = '8'

// 卖一价大于条件价
const TThostFtdcContingentConditionType_AskPriceGreaterThanStopPrice = '9'

// 卖一价大于等于条件价
const TThostFtdcContingentConditionType_AskPriceGreaterEqualStopPrice = 'A'

// 卖一价小于条件价
const TThostFtdcContingentConditionType_AskPriceLesserThanStopPrice = 'B'

// 卖一价小于等于条件价
const TThostFtdcContingentConditionType_AskPriceLesserEqualStopPrice = 'C'

// 买一价大于条件价
const TThostFtdcContingentConditionType_BidPriceGreaterThanStopPrice = 'D'

// 买一价大于等于条件价
const TThostFtdcContingentConditionType_BidPriceGreaterEqualStopPrice = 'E'

// 买一价小于条件价
const TThostFtdcContingentConditionType_BidPriceLesserThanStopPrice = 'F'

// 买一价小于等于条件价
const TThostFtdcContingentConditionType_BidPriceLesserEqualStopPrice = 'H'

// 操作标志类型
type TThostFtdcActionFlagType byte

// 删除
const TThostFtdcActionFlagType_Delete = '0'

// 修改
const TThostFtdcActionFlagType_Modify = '3'

// 交易权限类型
type TThostFtdcTradingRightType byte

// 可以交易
const TThostFtdcTradingRightType_Allow = '0'

// 只能平仓
const TThostFtdcTradingRightType_CloseOnly = '1'

// 不能交易
const TThostFtdcTradingRightType_Forbidden = '2'

// 报单来源类型
type TThostFtdcOrderSourceType byte

// 来自参与者
const TThostFtdcOrderSourceType_Participant = '0'

// 来自管理员
const TThostFtdcOrderSourceType_Administrator = '1'

// 成交类型类型
type TThostFtdcTradeTypeType byte

// 普通成交
const TThostFtdcTradeTypeType_Common = '0'

// 期权执行
const TThostFtdcTradeTypeType_OptionsExecution = '1'

// OTC成交
const TThostFtdcTradeTypeType_OTC = '2'

// 期转现衍生成交
const TThostFtdcTradeTypeType_EFPDerived = '3'

// 组合衍生成交
const TThostFtdcTradeTypeType_CombinationDerived = '4'

// 成交价来源类型
type TThostFtdcPriceSourceType byte

// 前成交价
const TThostFtdcPriceSourceType_LastPrice = '0'

// 买委托价
const TThostFtdcPriceSourceType_Buy = '1'

// 卖委托价
const TThostFtdcPriceSourceType_Sell = '2'

// 合约交易状态类型
type TThostFtdcInstrumentStatusType byte

// 开盘前
const TThostFtdcInstrumentStatusType_BeforeTrading = '0'

// 非交易
const TThostFtdcInstrumentStatusType_NoTrading = '1'

// 连续交易
const TThostFtdcInstrumentStatusType_Continous = '2'

// 集合竞价报单
const TThostFtdcInstrumentStatusType_AuctionOrdering = '3'

// 集合竞价价格平衡
const TThostFtdcInstrumentStatusType_AuctionBalance = '4'

// 集合竞价撮合
const TThostFtdcInstrumentStatusType_AuctionMatch = '5'

// 收盘
const TThostFtdcInstrumentStatusType_Closed = '6'

// 品种进入交易状态原因类型
type TThostFtdcInstStatusEnterReasonType byte

// 自动切换
const TThostFtdcInstStatusEnterReasonType_Automatic = '1'

// 手动切换
const TThostFtdcInstStatusEnterReasonType_Manual = '2'

// 熔断
const TThostFtdcInstStatusEnterReasonType_Fuse = '3'

// 组合开平标志类型
type TThostFtdcCombOffsetFlagType [5]byte

// 组合投机套保标志类型
type TThostFtdcCombHedgeFlagType [5]byte

// 序列编号类型
type TThostFtdcSequenceLabelType [2]byte

// 合同编号类型
type TThostFtdcContractCodeType [41]byte

// 市类型
type TThostFtdcCityType [51]byte

// 是否股民类型
type TThostFtdcIsStockType [11]byte

// 渠道类型
type TThostFtdcChannelType [51]byte

// 通讯地址类型
type TThostFtdcAddressType [101]byte

// 邮政编码类型
type TThostFtdcZipCodeType [7]byte

// 联系电话类型
type TThostFtdcTelephoneType [41]byte

// 传真类型
type TThostFtdcFaxType [41]byte

// 手机类型
type TThostFtdcMobileType [41]byte

// 电子邮件类型
type TThostFtdcEMailType [41]byte

// 备注类型
type TThostFtdcMemoType [161]byte

// 企业代码类型
type TThostFtdcCompanyCodeType [51]byte

// 网站地址类型
type TThostFtdcWebsiteType [51]byte

// 税务登记号类型
type TThostFtdcTaxNoType [31]byte

// 处理状态类型
type TThostFtdcBatchStatusType byte

// 未上传
const TThostFtdcBatchStatusType_NoUpload = '1'

// 已上传
const TThostFtdcBatchStatusType_Uploaded = '2'

// 审核失败
const TThostFtdcBatchStatusType_Failed = '3'

// 属性代码类型
type TThostFtdcPropertyIDType [33]byte

// 属性名称类型
type TThostFtdcPropertyNameType [65]byte

// 营业执照号类型
type TThostFtdcLicenseNoType [51]byte

// 经纪人代码类型
type TThostFtdcAgentIDType [13]byte

// 经纪人名称类型
type TThostFtdcAgentNameType [41]byte

// 经纪人组代码类型
type TThostFtdcAgentGroupIDType [13]byte

// 经纪人组名称类型
type TThostFtdcAgentGroupNameType [41]byte

// 按品种返还方式类型
type TThostFtdcReturnStyleType byte

// 按所有品种
const TThostFtdcReturnStyleType_All = '1'

// 按品种
const TThostFtdcReturnStyleType_ByProduct = '2'

// 返还模式类型
type TThostFtdcReturnPatternType byte

// 按成交手数
const TThostFtdcReturnPatternType_ByVolume = '1'

// 按留存手续费
const TThostFtdcReturnPatternType_ByFeeOnHand = '2'

// 返还级别类型
type TThostFtdcReturnLevelType byte

// 级别1
const TThostFtdcReturnLevelType_Level1 = '1'

// 级别2
const TThostFtdcReturnLevelType_Level2 = '2'

// 级别3
const TThostFtdcReturnLevelType_Level3 = '3'

// 级别4
const TThostFtdcReturnLevelType_Level4 = '4'

// 级别5
const TThostFtdcReturnLevelType_Level5 = '5'

// 级别6
const TThostFtdcReturnLevelType_Level6 = '6'

// 级别7
const TThostFtdcReturnLevelType_Level7 = '7'

// 级别8
const TThostFtdcReturnLevelType_Level8 = '8'

// 级别9
const TThostFtdcReturnLevelType_Level9 = '9'

// 返还标准类型
type TThostFtdcReturnStandardType byte

// 分阶段返还
const TThostFtdcReturnStandardType_ByPeriod = '1'

// 按某一标准
const TThostFtdcReturnStandardType_ByStandard = '2'

// 质押类型类型
type TThostFtdcMortgageTypeType byte

// 质出
const TThostFtdcMortgageTypeType_Out = '0'

// 质入
const TThostFtdcMortgageTypeType_In = '1'

// 投资者结算参数代码类型
type TThostFtdcInvestorSettlementParamIDType byte

// 质押比例
const TThostFtdcInvestorSettlementParamIDType_MortgageRatio = '4'

// 保证金算法
const TThostFtdcInvestorSettlementParamIDType_MarginWay = '5'

// 结算单结存是否包含质押
const TThostFtdcInvestorSettlementParamIDType_BillDeposit = '9'

// 交易所结算参数代码类型
type TThostFtdcExchangeSettlementParamIDType byte

// 质押比例
const TThostFtdcExchangeSettlementParamIDType_MortgageRatio = '1'

// 分项资金导入项
const TThostFtdcExchangeSettlementParamIDType_OtherFundItem = '2'

// 分项资金入交易所出入金
const TThostFtdcExchangeSettlementParamIDType_OtherFundImport = '3'

// 中金所开户最低可用金额
const TThostFtdcExchangeSettlementParamIDType_CFFEXMinPrepa = '6'

// 郑商所结算方式
const TThostFtdcExchangeSettlementParamIDType_CZCESettlementType = '7'

// 交易所交割手续费收取方式
const TThostFtdcExchangeSettlementParamIDType_ExchDelivFeeMode = '9'

// 投资者交割手续费收取方式
const TThostFtdcExchangeSettlementParamIDType_DelivFeeMode = '0'

// 郑商所组合持仓保证金收取方式
const TThostFtdcExchangeSettlementParamIDType_CZCEComMarginType = 'A'

// 大商所套利保证金是否优惠
const TThostFtdcExchangeSettlementParamIDType_DceComMarginType = 'B'

// 虚值期权保证金优惠比率
const TThostFtdcExchangeSettlementParamIDType_OptOutDisCountRate = 'a'

// 最低保障系数
const TThostFtdcExchangeSettlementParamIDType_OptMiniGuarantee = 'b'

// 系统参数代码类型
type TThostFtdcSystemParamIDType byte

// 投资者代码最小长度
const TThostFtdcSystemParamIDType_InvestorIDMinLength = '1'

// 投资者帐号代码最小长度
const TThostFtdcSystemParamIDType_AccountIDMinLength = '2'

// 投资者开户默认登录权限
const TThostFtdcSystemParamIDType_UserRightLogon = '3'

// 投资者交易结算单成交汇总方式
const TThostFtdcSystemParamIDType_SettlementBillTrade = '4'

// 统一开户更新交易编码方式
const TThostFtdcSystemParamIDType_TradingCode = '5'

// 结算是否判断存在未复核的出入金和分项资金
const TThostFtdcSystemParamIDType_CheckFund = '6'

// 是否启用手续费模板数据权限
const TThostFtdcSystemParamIDType_CommModelRight = '7'

// 是否启用保证金率模板数据权限
const TThostFtdcSystemParamIDType_MarginModelRight = '9'

// 是否规范用户才能激活
const TThostFtdcSystemParamIDType_IsStandardActive = '8'

// 上传的交易所结算文件路径
const TThostFtdcSystemParamIDType_UploadSettlementFile = 'U'

// 上报保证金监控中心文件路径
const TThostFtdcSystemParamIDType_DownloadCSRCFile = 'D'

// 生成的结算单文件路径
const TThostFtdcSystemParamIDType_SettlementBillFile = 'S'

// 证监会文件标识
const TThostFtdcSystemParamIDType_CSRCOthersFile = 'C'

// 投资者照片路径
const TThostFtdcSystemParamIDType_InvestorPhoto = 'P'

// 全结经纪公司上传文件路径
const TThostFtdcSystemParamIDType_CSRCData = 'R'

// 开户密码录入方式
const TThostFtdcSystemParamIDType_InvestorPwdModel = 'I'

// 投资者中金所结算文件下载路径
const TThostFtdcSystemParamIDType_CFFEXInvestorSettleFile = 'F'

// 投资者代码编码方式
const TThostFtdcSystemParamIDType_InvestorIDType = 'a'

// 休眠户最高权益
const TThostFtdcSystemParamIDType_FreezeMaxReMain = 'r'

// 手续费相关操作实时上场开关
const TThostFtdcSystemParamIDType_IsSync = 'A'

// 解除开仓权限限制
const TThostFtdcSystemParamIDType_RelieveOpenLimit = 'O'

// 是否规范用户才能休眠
const TThostFtdcSystemParamIDType_IsStandardFreeze = 'X'

// 郑商所是否开放所有品种套保交易
const TThostFtdcSystemParamIDType_CZCENormalProductHedge = 'B'

// 交易系统参数代码类型
type TThostFtdcTradeParamIDType byte

// 系统加密算法
const TThostFtdcTradeParamIDType_EncryptionStandard = 'E'

// 系统风险算法
const TThostFtdcTradeParamIDType_RiskMode = 'R'

// 系统风险算法是否全局 0-否 1-是
const TThostFtdcTradeParamIDType_RiskModeGlobal = 'G'

// 密码加密算法
const TThostFtdcTradeParamIDType_modeEncode = 'P'

// 价格小数位数参数
const TThostFtdcTradeParamIDType_tickMode = 'T'

// 用户最大会话数
const TThostFtdcTradeParamIDType_SingleUserSessionMaxNum = 'S'

// 最大连续登录失败数
const TThostFtdcTradeParamIDType_LoginFailMaxNum = 'L'

// 是否强制认证
const TThostFtdcTradeParamIDType_IsAuthForce = 'A'

// 是否冻结证券持仓
const TThostFtdcTradeParamIDType_IsPosiFreeze = 'F'

// 是否限仓
const TThostFtdcTradeParamIDType_IsPosiLimit = 'M'

// 郑商所询价时间间隔
const TThostFtdcTradeParamIDType_ForQuoteTimeInterval = 'Q'

// 是否期货限仓
const TThostFtdcTradeParamIDType_IsFuturePosiLimit = 'B'

// 是否期货下单频率限制
const TThostFtdcTradeParamIDType_IsFutureOrderFreq = 'C'

// 行权冻结是否计算盈利
const TThostFtdcTradeParamIDType_IsExecOrderProfit = 'H'

// 银期开户是否验证开户银行卡号是否是预留银行账户
const TThostFtdcTradeParamIDType_IsCheckBankAcc = 'I'

// 弱密码最后修改日期
const TThostFtdcTradeParamIDType_PasswordDeadLine = 'J'

// 强密码校验
const TThostFtdcTradeParamIDType_IsStrongPassword = 'K'

// 自有资金质押比
const TThostFtdcTradeParamIDType_BalanceMorgage = 'a'

// 最小密码长度
const TThostFtdcTradeParamIDType_MinPwdLen = 'O'

// IP当日最大登陆失败次数
const TThostFtdcTradeParamIDType_LoginFailMaxNumForIP = 'U'

// 密码有效期
const TThostFtdcTradeParamIDType_PasswordPeriod = 'V'

// 参数代码值类型
type TThostFtdcSettlementParamValueType [256]byte

// 计数器代码类型
type TThostFtdcCounterIDType [33]byte

// 投资者分组名称类型
type TThostFtdcInvestorGroupNameType [41]byte

// 牌号类型
type TThostFtdcBrandCodeType [257]byte

// 仓库类型
type TThostFtdcWarehouseType [257]byte

// 产期类型
type TThostFtdcProductDateType [41]byte

// 等级类型
type TThostFtdcGradeType [41]byte

// 类别类型
type TThostFtdcClassifyType [41]byte

// 货位类型
type TThostFtdcPositionType [41]byte

// 产地类型
type TThostFtdcYieldlyType [41]byte

// 公定重量类型
type TThostFtdcWeightType [41]byte

// 文件标识类型
type TThostFtdcFileIDType byte

// 资金数据
const TThostFtdcFileIDType_SettlementFund = 'F'

// 成交数据
const TThostFtdcFileIDType_Trade = 'T'

// 投资者持仓数据
const TThostFtdcFileIDType_InvestorPosition = 'P'

// 投资者分项资金数据
const TThostFtdcFileIDType_SubEntryFund = 'O'

// 组合持仓数据
const TThostFtdcFileIDType_CZCECombinationPos = 'C'

// 上报保证金监控中心数据
const TThostFtdcFileIDType_CSRCData = 'R'

// 郑商所平仓了结数据
const TThostFtdcFileIDType_CZCEClose = 'L'

// 郑商所非平仓了结数据
const TThostFtdcFileIDType_CZCENoClose = 'N'

// 持仓明细数据
const TThostFtdcFileIDType_PositionDtl = 'D'

// 期权执行文件
const TThostFtdcFileIDType_OptionStrike = 'S'

// 结算价比对文件
const TThostFtdcFileIDType_SettlementPriceComparison = 'M'

// 上期所非持仓变动明细
const TThostFtdcFileIDType_NonTradePosChange = 'B'

// 文件名称类型
type TThostFtdcFileNameType [257]byte

// 文件上传类型类型
type TThostFtdcFileTypeType byte

// 结算
const TThostFtdcFileTypeType_Settlement = '0'

// 核对
const TThostFtdcFileTypeType_Check = '1'

// 文件格式类型
type TThostFtdcFileFormatType byte

// 文本文件(.txt)
const TThostFtdcFileFormatType_Txt = '0'

// 压缩文件(.zip)
const TThostFtdcFileFormatType_Zip = '1'

// DBF文件(.dbf)
const TThostFtdcFileFormatType_DBF = '2'

// 文件状态类型
type TThostFtdcFileUploadStatusType byte

// 上传成功
const TThostFtdcFileUploadStatusType_SucceedUpload = '1'

// 上传失败
const TThostFtdcFileUploadStatusType_FailedUpload = '2'

// 导入成功
const TThostFtdcFileUploadStatusType_SucceedLoad = '3'

// 导入部分成功
const TThostFtdcFileUploadStatusType_PartSucceedLoad = '4'

// 导入失败
const TThostFtdcFileUploadStatusType_FailedLoad = '5'

// 移仓方向类型
type TThostFtdcTransferDirectionType byte

// 移出
const TThostFtdcTransferDirectionType_Out = '0'

// 移入
const TThostFtdcTransferDirectionType_In = '1'

// 上传文件类型类型
type TThostFtdcUploadModeType [21]byte

// 投资者帐号类型
type TThostFtdcAccountIDType [13]byte

// 银行统一标识类型类型
type TThostFtdcBankFlagType [4]byte

// 银行账户类型
type TThostFtdcBankAccountType [41]byte

// 银行账户的开户人名称类型
type TThostFtdcOpenNameType [61]byte

// 银行账户的开户行类型
type TThostFtdcOpenBankType [101]byte

// 银行名称类型
type TThostFtdcBankNameType [101]byte

// 发布路径类型
type TThostFtdcPublishPathType [257]byte

// 操作员代码类型
type TThostFtdcOperatorIDType [65]byte

// 月份提前数组类型
type TThostFtdcAdvanceMonthArrayType [13]byte

// 日期表达式类型
type TThostFtdcDateExprType [1025]byte

// 合约代码表达式类型
type TThostFtdcInstrumentIDExprType [41]byte

// 合约名称表达式类型
type TThostFtdcInstrumentNameExprType [41]byte

// 特殊的创建规则类型
type TThostFtdcSpecialCreateRuleType byte

// 没有特殊创建规则
const TThostFtdcSpecialCreateRuleType_NoSpecialRule = '0'

// 不包含春节
const TThostFtdcSpecialCreateRuleType_NoSpringFestival = '1'

// 挂牌基准价类型类型
type TThostFtdcBasisPriceTypeType byte

// 上一合约结算价
const TThostFtdcBasisPriceTypeType_LastSettlement = '1'

// 上一合约收盘价
const TThostFtdcBasisPriceTypeType_LaseClose = '2'

// 产品生命周期状态类型
type TThostFtdcProductLifePhaseType byte

// 活跃
const TThostFtdcProductLifePhaseType_Active = '1'

// 不活跃
const TThostFtdcProductLifePhaseType_NonActive = '2'

// 注销
const TThostFtdcProductLifePhaseType_Canceled = '3'

// 交割方式类型
type TThostFtdcDeliveryModeType byte

// 现金交割
const TThostFtdcDeliveryModeType_CashDeliv = '1'

// 实物交割
const TThostFtdcDeliveryModeType_CommodityDeliv = '2'

// 日志级别类型
type TThostFtdcLogLevelType [33]byte

// 存储过程名称类型
type TThostFtdcProcessNameType [257]byte

// 操作摘要类型
type TThostFtdcOperationMemoType [1025]byte

// 出入金类型类型
type TThostFtdcFundIOTypeType byte

// 出入金
const TThostFtdcFundIOTypeType_FundIO = '1'

// 银期转帐
const TThostFtdcFundIOTypeType_Transfer = '2'

// 银期换汇
const TThostFtdcFundIOTypeType_SwapCurrency = '3'

// 资金类型类型
type TThostFtdcFundTypeType byte

// 银行存款
const TThostFtdcFundTypeType_Deposite = '1'

// 分项资金
const TThostFtdcFundTypeType_ItemFund = '2'

// 公司调整
const TThostFtdcFundTypeType_Company = '3'

// 资金内转
const TThostFtdcFundTypeType_InnerTransfer = '4'

// 出入金方向类型
type TThostFtdcFundDirectionType byte

// 入金
const TThostFtdcFundDirectionType_In = '1'

// 出金
const TThostFtdcFundDirectionType_Out = '2'

// 资金状态类型
type TThostFtdcFundStatusType byte

// 已录入
const TThostFtdcFundStatusType_Record = '1'

// 已复核
const TThostFtdcFundStatusType_Check = '2'

// 已冲销
const TThostFtdcFundStatusType_Charge = '3'

// 票据号类型
type TThostFtdcBillNoType [15]byte

// 票据名称类型
type TThostFtdcBillNameType [33]byte

// 发布状态类型
type TThostFtdcPublishStatusType byte

// 未发布
const TThostFtdcPublishStatusType_None = '1'

// 正在发布
const TThostFtdcPublishStatusType_Publishing = '2'

// 已发布
const TThostFtdcPublishStatusType_Published = '3'

// 枚举值代码类型
type TThostFtdcEnumValueIDType [65]byte

// 枚举值类型类型
type TThostFtdcEnumValueTypeType [33]byte

// 枚举值名称类型
type TThostFtdcEnumValueLabelType [65]byte

// 枚举值结果类型
type TThostFtdcEnumValueResultType [33]byte

// 系统状态类型
type TThostFtdcSystemStatusType byte

// 不活跃
const TThostFtdcSystemStatusType_NonActive = '1'

// 启动
const TThostFtdcSystemStatusType_Startup = '2'

// 交易开始初始化
const TThostFtdcSystemStatusType_Initialize = '3'

// 交易完成初始化
const TThostFtdcSystemStatusType_Initialized = '4'

// 收市开始
const TThostFtdcSystemStatusType_Close = '5'

// 收市完成
const TThostFtdcSystemStatusType_Closed = '6'

// 结算
const TThostFtdcSystemStatusType_Settlement = '7'

// 结算状态类型
type TThostFtdcSettlementStatusType byte

// 初始
const TThostFtdcSettlementStatusType_Initialize = '0'

// 结算中
const TThostFtdcSettlementStatusType_Settlementing = '1'

// 已结算
const TThostFtdcSettlementStatusType_Settlemented = '2'

// 结算完成
const TThostFtdcSettlementStatusType_Finished = '3'

// 限定值类型类型
type TThostFtdcRangeIntTypeType [33]byte

// 限定值下限类型
type TThostFtdcRangeIntFromType [33]byte

// 限定值上限类型
type TThostFtdcRangeIntToType [33]byte
type TThostFtdcFunctionIDType [25]byte

// 功能编码类型
type TThostFtdcFunctionValueCodeType [257]byte

// 功能名称类型
type TThostFtdcFunctionNameType [65]byte

// 角色编号类型
type TThostFtdcRoleIDType [11]byte

// 角色名称类型
type TThostFtdcRoleNameType [41]byte

// 描述类型
type TThostFtdcDescriptionType [401]byte

// 组合编号类型
type TThostFtdcCombineIDType [25]byte

// 组合类型类型
type TThostFtdcCombineTypeType [25]byte

// 投资者类型类型
type TThostFtdcInvestorTypeType byte

// 自然人
const TThostFtdcInvestorTypeType_Person = '0'

// 法人
const TThostFtdcInvestorTypeType_Company = '1'

// 投资基金
const TThostFtdcInvestorTypeType_Fund = '2'

// 特殊法人
const TThostFtdcInvestorTypeType_SpecialOrgan = '3'

// 资管户
const TThostFtdcInvestorTypeType_Asset = '4'

// 经纪公司类型类型
type TThostFtdcBrokerTypeType byte

// 交易会员
const TThostFtdcBrokerTypeType_Trade = '0'

// 交易结算会员
const TThostFtdcBrokerTypeType_TradeSettle = '1'

// 风险等级类型
type TThostFtdcRiskLevelType byte

// 低风险客户
const TThostFtdcRiskLevelType_Low = '1'

// 普通客户
const TThostFtdcRiskLevelType_Normal = '2'

// 关注客户
const TThostFtdcRiskLevelType_Focus = '3'

// 风险客户
const TThostFtdcRiskLevelType_Risk = '4'

// 手续费收取方式类型
type TThostFtdcFeeAcceptStyleType byte

// 按交易收取
const TThostFtdcFeeAcceptStyleType_ByTrade = '1'

// 按交割收取
const TThostFtdcFeeAcceptStyleType_ByDeliv = '2'

// 不收
const TThostFtdcFeeAcceptStyleType_None = '3'

// 按指定手续费收取
const TThostFtdcFeeAcceptStyleType_FixFee = '4'

// 密码类型类型
type TThostFtdcPasswordTypeType byte

// 交易密码
const TThostFtdcPasswordTypeType_Trade = '1'

// 资金密码
const TThostFtdcPasswordTypeType_Account = '2'

// 盈亏算法类型
type TThostFtdcAlgorithmType byte

// 浮盈浮亏都计算
const TThostFtdcAlgorithmType_All = '1'

// 浮盈不计，浮亏计
const TThostFtdcAlgorithmType_OnlyLost = '2'

// 浮盈计，浮亏不计
const TThostFtdcAlgorithmType_OnlyGain = '3'

// 浮盈浮亏都不计算
const TThostFtdcAlgorithmType_None = '4'

// 是否包含平仓盈利类型
type TThostFtdcIncludeCloseProfitType byte

// 包含平仓盈利
const TThostFtdcIncludeCloseProfitType_Include = '0'

// 不包含平仓盈利
const TThostFtdcIncludeCloseProfitType_NotInclude = '2'

// 是否受可提比例限制类型
type TThostFtdcAllWithoutTradeType byte

// 无仓无成交不受可提比例限制
const TThostFtdcAllWithoutTradeType_Enable = '0'

// 受可提比例限制
const TThostFtdcAllWithoutTradeType_Disable = '2'

// 无仓不受可提比例限制
const TThostFtdcAllWithoutTradeType_NoHoldEnable = '3'

// 盈亏算法说明类型
type TThostFtdcCommentType [31]byte

// 版本号类型
type TThostFtdcVersionType [4]byte

// 交易代码类型
type TThostFtdcTradeCodeType [7]byte

// 交易日期类型
type TThostFtdcTradeDateType [9]byte

// 交易时间类型
type TThostFtdcTradeTimeType [9]byte

// 发起方流水号类型
type TThostFtdcTradeSerialType [9]byte

// 期货公司代码类型
type TThostFtdcFutureIDType [11]byte

// 银行代码类型
type TThostFtdcBankIDType [4]byte

// 银行分中心代码类型
type TThostFtdcBankBrchIDType [5]byte

// 分中心代码类型
type TThostFtdcBankBranchIDType [11]byte

// 交易柜员类型
type TThostFtdcOperNoType [17]byte

// 渠道标志类型
type TThostFtdcDeviceIDType [3]byte

// 记录数类型
type TThostFtdcRecordNumType [7]byte

// 期货资金账号类型
type TThostFtdcFutureAccountType [22]byte

// 资金密码核对标志类型
type TThostFtdcFuturePwdFlagType byte

// 不核对
const TThostFtdcFuturePwdFlagType_UnCheck = '0'

// 核对
const TThostFtdcFuturePwdFlagType_Check = '1'

// 银期转账类型类型
type TThostFtdcTransferTypeType byte

// 银行转期货
const TThostFtdcTransferTypeType_BankToFuture = '0'

// 期货转银行
const TThostFtdcTransferTypeType_FutureToBank = '1'

// 期货资金密码类型
type TThostFtdcFutureAccPwdType [17]byte

// 币种类型
type TThostFtdcCurrencyCodeType [4]byte

// 响应代码类型
type TThostFtdcRetCodeType [5]byte

// 响应信息类型
type TThostFtdcRetInfoType [129]byte

// 银行总余额类型
type TThostFtdcTradeAmtType [20]byte

// 银行可用余额类型
type TThostFtdcUseAmtType [20]byte

// 银行可取余额类型
type TThostFtdcFetchAmtType [20]byte

// 转账有效标志类型
type TThostFtdcTransferValidFlagType byte

// 无效或失败
const TThostFtdcTransferValidFlagType_Invalid = '0'

// 有效
const TThostFtdcTransferValidFlagType_Valid = '1'

// 冲正
const TThostFtdcTransferValidFlagType_Reverse = '2'

type TThostFtdcCertCodeType [21]byte

// 事由类型
type TThostFtdcReasonType byte

// 错单
const TThostFtdcReasonType_CD = '0'

// 资金在途
const TThostFtdcReasonType_ZT = '1'

// 其它
const TThostFtdcReasonType_QT = '2'

// 资金项目编号类型
type TThostFtdcFundProjectIDType [5]byte

// 性别类型
type TThostFtdcSexType byte

// 未知
const TThostFtdcSexType_None = '0'

// 男
const TThostFtdcSexType_Man = '1'

// 女
const TThostFtdcSexType_Woman = '2'

// 职业类型
type TThostFtdcProfessionType [101]byte

// 国籍类型
type TThostFtdcNationalType [31]byte

// 省类型
type TThostFtdcProvinceType [51]byte

// 区类型
type TThostFtdcRegionType [16]byte

// 国家类型
type TThostFtdcCountryType [16]byte

// 营业执照类型
type TThostFtdcLicenseNOType [33]byte

// 企业性质类型
type TThostFtdcCompanyTypeType [16]byte

// 经营范围类型
type TThostFtdcBusinessScopeType [1001]byte

// 注册资本币种类型
type TThostFtdcCapitalCurrencyType [4]byte

// 用户类型类型
type TThostFtdcUserTypeType byte

// 投资者
const TThostFtdcUserTypeType_Investor = '0'

// 操作员
const TThostFtdcUserTypeType_Operator = '1'

// 管理员
const TThostFtdcUserTypeType_SuperUser = '2'

// 营业部编号类型
type TThostFtdcBranchIDType [9]byte

// 费率类型类型
type TThostFtdcRateTypeType byte

// 保证金率
const TThostFtdcRateTypeType_MarginRate = '2'

// 通知类型类型
type TThostFtdcNoteTypeType byte

// 交易结算单
const TThostFtdcNoteTypeType_TradeSettleBill = '1'

// 交易结算月报
const TThostFtdcNoteTypeType_TradeSettleMonth = '2'

// 追加保证金通知书
const TThostFtdcNoteTypeType_CallMarginNotes = '3'

// 强行平仓通知书
const TThostFtdcNoteTypeType_ForceCloseNotes = '4'

// 成交通知书
const TThostFtdcNoteTypeType_TradeNotes = '5'

// 交割通知书
const TThostFtdcNoteTypeType_DelivNotes = '6'

// 结算单方式类型
type TThostFtdcSettlementStyleType byte

// 逐日盯市
const TThostFtdcSettlementStyleType_Day = '1'

// 逐笔对冲
const TThostFtdcSettlementStyleType_Volume = '2'

// 域名类型
type TThostFtdcBrokerDNSType [256]byte

// 语句类型
type TThostFtdcSentenceType [501]byte

// 结算单类型类型
type TThostFtdcSettlementBillTypeType byte

// 日报
const TThostFtdcSettlementBillTypeType_Day = '0'

// 月报
const TThostFtdcSettlementBillTypeType_Month = '1'

// 客户权限类型类型
type TThostFtdcUserRightTypeType byte

// 登录
const TThostFtdcUserRightTypeType_Logon = '1'

// 银期转帐
const TThostFtdcUserRightTypeType_Transfer = '2'

// 邮寄结算单
const TThostFtdcUserRightTypeType_EMail = '3'

// 传真结算单
const TThostFtdcUserRightTypeType_Fax = '4'

// 条件单
const TThostFtdcUserRightTypeType_ConditionOrder = '5'

// 保证金价格类型类型
type TThostFtdcMarginPriceTypeType byte

// 昨结算价
const TThostFtdcMarginPriceTypeType_PreSettlementPrice = '1'

// 最新价
const TThostFtdcMarginPriceTypeType_SettlementPrice = '2'

// 成交均价
const TThostFtdcMarginPriceTypeType_AveragePrice = '3'

// 开仓价
const TThostFtdcMarginPriceTypeType_OpenPrice = '4'

// 结算单生成状态类型
type TThostFtdcBillGenStatusType byte

// 未生成
const TThostFtdcBillGenStatusType_None = '0'

// 生成中
const TThostFtdcBillGenStatusType_NoGenerated = '1'

// 已生成
const TThostFtdcBillGenStatusType_Generated = '2'

// 算法类型类型
type TThostFtdcAlgoTypeType byte

// 持仓处理算法
const TThostFtdcAlgoTypeType_HandlePositionAlgo = '1'

// 寻找保证金率算法
const TThostFtdcAlgoTypeType_FindMarginRateAlgo = '2'

// 持仓处理算法编号类型
type TThostFtdcHandlePositionAlgoIDType byte

// 基本
const TThostFtdcHandlePositionAlgoIDType_Base = '1'

// 大连商品交易所
const TThostFtdcHandlePositionAlgoIDType_DCE = '2'

// 郑州商品交易所
const TThostFtdcHandlePositionAlgoIDType_CZCE = '3'

// 寻找保证金率算法编号类型
type TThostFtdcFindMarginRateAlgoIDType byte

// 基本
const TThostFtdcFindMarginRateAlgoIDType_Base = '1'

// 大连商品交易所
const TThostFtdcFindMarginRateAlgoIDType_DCE = '2'

// 郑州商品交易所
const TThostFtdcFindMarginRateAlgoIDType_CZCE = '3'

// 资金处理算法编号类型
type TThostFtdcHandleTradingAccountAlgoIDType byte

// 基本
const TThostFtdcHandleTradingAccountAlgoIDType_Base = '1'

// 大连商品交易所
const TThostFtdcHandleTradingAccountAlgoIDType_DCE = '2'

// 郑州商品交易所
const TThostFtdcHandleTradingAccountAlgoIDType_CZCE = '3'

// 联系人类型类型
type TThostFtdcPersonTypeType byte

// 指定下单人
const TThostFtdcPersonTypeType_Order = '1'

// 开户授权人
const TThostFtdcPersonTypeType_Open = '2'

// 资金调拨人
const TThostFtdcPersonTypeType_Fund = '3'

// 结算单确认人
const TThostFtdcPersonTypeType_Settlement = '4'

// 法人
const TThostFtdcPersonTypeType_Company = '5'

// 法人代表
const TThostFtdcPersonTypeType_Corporation = '6'

// 投资者联系人
const TThostFtdcPersonTypeType_LinkMan = '7'

// 分户管理资产负责人
const TThostFtdcPersonTypeType_Ledger = '8'

// 托（保）管人
const TThostFtdcPersonTypeType_Trustee = '9'

// 托（保）管机构法人代表
const TThostFtdcPersonTypeType_TrusteeCorporation = 'A'

// 托（保）管机构开户授权人
const TThostFtdcPersonTypeType_TrusteeOpen = 'B'

// 托（保）管机构联系人
const TThostFtdcPersonTypeType_TrusteeContact = 'C'

// 境外自然人参考证件
const TThostFtdcPersonTypeType_ForeignerRefer = 'D'

// 法人代表参考证件
const TThostFtdcPersonTypeType_CorporationRefer = 'E'

// 查询范围类型
type TThostFtdcQueryInvestorRangeType byte

// 所有
const TThostFtdcQueryInvestorRangeType_All = '1'

// 查询分类
const TThostFtdcQueryInvestorRangeType_Group = '2'

// 单一投资者
const TThostFtdcQueryInvestorRangeType_Single = '3'

// 投资者风险状态类型
type TThostFtdcInvestorRiskStatusType byte

// 正常
const TThostFtdcInvestorRiskStatusType_Normal = '1'

// 警告
const TThostFtdcInvestorRiskStatusType_Warn = '2'

// 追保
const TThostFtdcInvestorRiskStatusType_Call = '3'

// 强平
const TThostFtdcInvestorRiskStatusType_Force = '4'

// 异常
const TThostFtdcInvestorRiskStatusType_Exception = '5'

// 结算账户类型
type TThostFtdcClearAccountType [33]byte
type TThostFtdcOrganNOType [6]byte

// 结算账户联行号类型
type TThostFtdcClearbarchIDType [6]byte

// 用户事件类型类型
type TThostFtdcUserEventTypeType byte

// 登录
const TThostFtdcUserEventTypeType_Login = '1'

// 登出
const TThostFtdcUserEventTypeType_Logout = '2'

// 交易成功
const TThostFtdcUserEventTypeType_Trading = '3'

// 交易失败
const TThostFtdcUserEventTypeType_TradingError = '4'

// 修改密码
const TThostFtdcUserEventTypeType_UpdatePassword = '5'

// 客户端认证
const TThostFtdcUserEventTypeType_Authenticate = '6'

// 其他
const TThostFtdcUserEventTypeType_Other = '9'

// 用户事件信息类型
type TThostFtdcUserEventInfoType [1025]byte

// 平仓方式类型
type TThostFtdcCloseStyleType byte

// 先开先平
const TThostFtdcCloseStyleType_Close = '0'

// 先平今再平昨
const TThostFtdcCloseStyleType_CloseToday = '1'

// 统计方式类型
type TThostFtdcStatModeType byte

// ----
const TThostFtdcStatModeType_Non = '0'

// 按合约统计
const TThostFtdcStatModeType_Instrument = '1'

// 按产品统计
const TThostFtdcStatModeType_Product = '2'

// 按投资者统计
const TThostFtdcStatModeType_Investor = '3'

// 预埋单状态类型
type TThostFtdcParkedOrderStatusType byte

// 未发送
const TThostFtdcParkedOrderStatusType_NotSend = '1'

// 已发送
const TThostFtdcParkedOrderStatusType_Send = '2'

// 已删除
const TThostFtdcParkedOrderStatusType_Deleted = '3'

// 预埋报单编号类型
type TThostFtdcParkedOrderIDType [13]byte

// 预埋撤单编号类型
type TThostFtdcParkedOrderActionIDType [13]byte
type TThostFtdcVirDealStatusType byte

// 正在处理
const TThostFtdcVirDealStatusType_Dealing = '1'

// 处理成功
const TThostFtdcVirDealStatusType_DeaclSucceed = '2'

// 原有系统代码类型
type TThostFtdcOrgSystemIDType byte

// 综合交易平台
const TThostFtdcOrgSystemIDType_Standard = '0'

// 易盛系统
const TThostFtdcOrgSystemIDType_ESunny = '1'

// 金仕达V6系统
const TThostFtdcOrgSystemIDType_KingStarV6 = '2'

// 交易状态类型
type TThostFtdcVirTradeStatusType byte

// 正常处理中
const TThostFtdcVirTradeStatusType_NaturalDeal = '0'

// 成功结束
const TThostFtdcVirTradeStatusType_SucceedEnd = '1'

// 失败结束
const TThostFtdcVirTradeStatusType_FailedEND = '2'

// 异常中
const TThostFtdcVirTradeStatusType_Exception = '3'

// 已人工异常处理
const TThostFtdcVirTradeStatusType_ManualDeal = '4'

// 通讯异常 ，请人工处理
const TThostFtdcVirTradeStatusType_MesException = '5'

// 系统出错，请人工处理
const TThostFtdcVirTradeStatusType_SysException = '6'

// 银行帐户类型类型
type TThostFtdcVirBankAccTypeType byte

// 存折
const TThostFtdcVirBankAccTypeType_BankBook = '1'

// 储蓄卡
const TThostFtdcVirBankAccTypeType_BankCard = '2'

// 信用卡
const TThostFtdcVirBankAccTypeType_CreditCard = '3'

type TThostFtdcVirementStatusType byte

// 正常
const TThostFtdcVirementStatusType_Natural = '0'

// 销户
const TThostFtdcVirementStatusType_Canceled = '9'

// 有效标志类型
type TThostFtdcVirementAvailAbilityType byte

// 未确认
const TThostFtdcVirementAvailAbilityType_NoAvailAbility = '0'

// 有效
const TThostFtdcVirementAvailAbilityType_AvailAbility = '1'

// 冲正
const TThostFtdcVirementAvailAbilityType_Repeal = '2'

type TThostFtdcVirementTradeCodeType byte

// 银行发起银行资金转期货
const TThostFtdcVirementTradeCodeType_BankBankToFuture = "102001"

// 银行发起期货资金转银行
const TThostFtdcVirementTradeCodeType_BankFutureToBank = "102002"

// 期货发起银行资金转期货
const TThostFtdcVirementTradeCodeType_FutureBankToFuture = "202001"

// 期货发起期货资金转银行
const TThostFtdcVirementTradeCodeType_FutureFutureToBank = "202002"

// 影像类型名称类型
type TThostFtdcPhotoTypeNameType [41]byte

// 影像类型代码类型
type TThostFtdcPhotoTypeIDType [5]byte

// 影像名称类型
type TThostFtdcPhotoNameType [161]byte

// 交易报告类型标识类型
type TThostFtdcReportTypeIDType [3]byte

// 交易特征代码类型
type TThostFtdcCharacterIDType [5]byte

// 参数代码类型
type TThostFtdcAMLParamIDType [21]byte
type TThostFtdcAMLInvestorTypeType [3]byte
type TThostFtdcAMLIdCardTypeType [3]byte

// 资金进出方向类型
type TThostFtdcAMLTradeDirectType [3]byte

// 资金进出方式类型
type TThostFtdcAMLTradeModelType [3]byte

// 客户身份证件/证明文件类型类型
type TThostFtdcAMLCustomerCardTypeType [81]byte

// 金融机构网点名称类型
type TThostFtdcAMLInstitutionNameType [65]byte

// 金融机构网点所在地区行政区划代码类型
type TThostFtdcAMLDistrictIDType [7]byte

// 金融机构网点与大额交易的关系类型
type TThostFtdcAMLRelationShipType [3]byte

// 金融机构网点代码类型类型
type TThostFtdcAMLInstitutionTypeType [3]byte

// 金融机构网点代码类型
type TThostFtdcAMLInstitutionIDType [13]byte

// 账户类型类型
type TThostFtdcAMLAccountTypeType [5]byte

// 交易方式类型
type TThostFtdcAMLTradingTypeType [7]byte

// 涉外收支交易分类与代码类型
type TThostFtdcAMLTransactClassType [7]byte

// 资金收付标识类型
type TThostFtdcAMLCapitalIOType [3]byte

// 交易地点类型
type TThostFtdcAMLSiteType [10]byte

// 资金用途类型
type TThostFtdcAMLCapitalPurposeType [129]byte

// 报文类型类型
type TThostFtdcAMLReportTypeType [2]byte

// 编号类型
type TThostFtdcAMLSerialNoType [5]byte

// 状态类型
type TThostFtdcAMLStatusType [2]byte

// Aml生成方式类型
type TThostFtdcAMLGenStatusType byte

// 程序生成
const TThostFtdcAMLGenStatusType_Program = '0'

// 人工生成
const TThostFtdcAMLGenStatusType_HandWork = '1'

// 业务标识号类型
type TThostFtdcAMLSeqCodeType [65]byte

// AML文件名类型
type TThostFtdcAMLFileNameType [257]byte

// 密钥类型(保证金监管)类型
type TThostFtdcCFMMCKeyType [21]byte

// 令牌类型(保证金监管)类型
type TThostFtdcCFMMCTokenType [21]byte

// 动态密钥类别(保证金监管)类型
type TThostFtdcCFMMCKeyKindType byte

// 主动请求更新
const TThostFtdcCFMMCKeyKindType_REQUEST = 'R'

// CFMMC自动更新
const TThostFtdcCFMMCKeyKindType_AUTO = 'A'

// CFMMC手动更新
const TThostFtdcCFMMCKeyKindType_MANUAL = 'M'

// 报文名称类型
type TThostFtdcAMLReportNameType [81]byte

// 个人姓名类型
type TThostFtdcIndividualNameType [51]byte

// 币种代码类型
type TThostFtdcCurrencyIDType [4]byte

// 客户编号类型
type TThostFtdcCustNumberType [36]byte

// 机构编码类型
type TThostFtdcOrganCodeType [36]byte

// 机构名称类型
type TThostFtdcOrganNameType [71]byte

// 上级机构编码,即期货公司总部、银行总行类型
type TThostFtdcSuperOrganCodeType [12]byte

// 分支机构类型
type TThostFtdcSubBranchIDType [31]byte

// 分支机构名称类型
type TThostFtdcSubBranchNameType [71]byte

// 机构网点号类型
type TThostFtdcBranchNetCodeType [31]byte

// 机构网点名称类型
type TThostFtdcBranchNetNameType [71]byte

// 机构标识类型
type TThostFtdcOrganFlagType [2]byte

// 银行对期货公司的编码类型
type TThostFtdcBankCodingForFutureType [33]byte

// 银行对返回码的定义类型
type TThostFtdcBankReturnCodeType [7]byte

// 银期转帐平台对返回码的定义类型
type TThostFtdcPlateReturnCodeType [5]byte

// 银行分支机构编码类型
type TThostFtdcBankSubBranchIDType [31]byte

// 期货分支机构编码类型
type TThostFtdcFutureBranchIDType [31]byte

// 返回代码类型
type TThostFtdcReturnCodeType [7]byte

// 操作员类型
type TThostFtdcOperatorCodeType [17]byte

// 机构结算帐户机构号类型
type TThostFtdcClearDepIDType [6]byte

// 机构结算帐户联行号类型
type TThostFtdcClearBrchIDType [6]byte

// 机构结算帐户名称类型
type TThostFtdcClearNameType [71]byte

// 银行帐户名称类型
type TThostFtdcBankAccountNameType [71]byte

// 机构投资人账号机构号类型
type TThostFtdcInvDepIDType [6]byte

// 机构投资人联行号类型
type TThostFtdcInvBrchIDType [6]byte

// 信息格式版本类型
type TThostFtdcMessageFormatVersionType [36]byte

// 摘要类型
type TThostFtdcDigestType [36]byte

// 认证数据类型
type TThostFtdcAuthenticDataType [129]byte

// 密钥类型
type TThostFtdcPasswordKeyType [129]byte

// 期货帐户名称类型
type TThostFtdcFutureAccountNameType [129]byte
type TThostFtdcMobilePhoneType [21]byte

// 期货公司主密钥类型
type TThostFtdcFutureMainKeyType [129]byte

// 期货公司工作密钥类型
type TThostFtdcFutureWorkKeyType [129]byte

// 期货公司传输密钥类型
type TThostFtdcFutureTransKeyType [129]byte

// 银行主密钥类型
type TThostFtdcBankMainKeyType [129]byte

// 银行工作密钥类型
type TThostFtdcBankWorkKeyType [129]byte

// 银行传输密钥类型
type TThostFtdcBankTransKeyType [129]byte

// 银行服务器描述信息类型
type TThostFtdcBankServerDescriptionType [129]byte

// 附加信息类型
type TThostFtdcAddInfoType [129]byte

// 返回码描述类型
type TThostFtdcDescrInfoForReturnCodeType [129]byte

// 国家代码类型
type TThostFtdcCountryCodeType [21]byte

// 银行流水号类型
type TThostFtdcBankSerialType [13]byte
type TThostFtdcCertificationTypeType byte

// 身份证
const TThostFtdcCertificationTypeType_IDCard = '0'

// 护照
const TThostFtdcCertificationTypeType_Passport = '1'

// 军官证
const TThostFtdcCertificationTypeType_OfficerIDCard = '2'

// 士兵证
const TThostFtdcCertificationTypeType_SoldierIDCard = '3'

// 回乡证
const TThostFtdcCertificationTypeType_HomeComingCard = '4'

// 户口簿
const TThostFtdcCertificationTypeType_HouseholdRegister = '5'

// 营业执照号
const TThostFtdcCertificationTypeType_LicenseNo = '6'

// 组织机构代码证
const TThostFtdcCertificationTypeType_InstitutionCodeCard = '7'

// 临时营业执照号
const TThostFtdcCertificationTypeType_TempLicenseNo = '8'

// 民办非企业登记证书
const TThostFtdcCertificationTypeType_NoEnterpriseLicenseNo = '9'

// 其他证件
const TThostFtdcCertificationTypeType_OtherCard = 'x'

// 主管部门批文
const TThostFtdcCertificationTypeType_SuperDepAgree = 'a'

// 文件业务功能类型
type TThostFtdcFileBusinessCodeType byte

// 其他
const TThostFtdcFileBusinessCodeType_Others = '0'

// 转账交易明细对账
const TThostFtdcFileBusinessCodeType_TransferDetails = '1'

// 客户账户状态对账
const TThostFtdcFileBusinessCodeType_CustAccStatus = '2'

// 账户类交易明细对账
const TThostFtdcFileBusinessCodeType_AccountTradeDetails = '3'

// 期货账户信息变更明细对账
const TThostFtdcFileBusinessCodeType_FutureAccountChangeInfoDetails = '4'

// 客户资金台账余额明细对账
const TThostFtdcFileBusinessCodeType_CustMoneyDetail = '5'

// 客户销户结息明细对账
const TThostFtdcFileBusinessCodeType_CustCancelAccountInfo = '6'

// 客户资金余额对账结果
const TThostFtdcFileBusinessCodeType_CustMoneyResult = '7'

// 其它对账异常结果文件
const TThostFtdcFileBusinessCodeType_OthersExceptionResult = '8'

// 客户结息净额明细
const TThostFtdcFileBusinessCodeType_CustInterestNetMoneyDetails = '9'

// 客户资金交收明细
const TThostFtdcFileBusinessCodeType_CustMoneySendAndReceiveDetails = 'a'

// 法人存管银行资金交收汇总
const TThostFtdcFileBusinessCodeType_CorporationMoneyTotal = 'b'

// 主体间资金交收汇总
const TThostFtdcFileBusinessCodeType_MainbodyMoneyTotal = 'c'

// 总分平衡监管数据
const TThostFtdcFileBusinessCodeType_MainPartMonitorData = 'd'

// 存管银行备付金余额
const TThostFtdcFileBusinessCodeType_PreparationMoney = 'e'

// 协办存管银行资金监管数据
const TThostFtdcFileBusinessCodeType_BankMoneyMonitorData = 'f'

// 汇钞标志类型
type TThostFtdcCashExchangeCodeType byte

// 汇
const TThostFtdcCashExchangeCodeType_Exchange = '1'

// 钞
const TThostFtdcCashExchangeCodeType_Cash = '2'

// 是或否标识类型
type TThostFtdcYesNoIndicatorType byte

// 是
const TThostFtdcYesNoIndicatorType_Yes = '0'

// 否
const TThostFtdcYesNoIndicatorType_No = '1'

// 余额类型类型
type TThostFtdcBanlanceTypeType byte

// 当前余额
const TThostFtdcBanlanceTypeType_CurrentMoney = '0'

// 可用余额
const TThostFtdcBanlanceTypeType_UsableMoney = '1'

// 可取余额
const TThostFtdcBanlanceTypeType_FetchableMoney = '2'

// 冻结余额
const TThostFtdcBanlanceTypeType_FreezeMoney = '3'

type TThostFtdcGenderType byte

// 未知状态
const TThostFtdcGenderType_Unknown = '0'

// 男
const TThostFtdcGenderType_Male = '1'

// 女
const TThostFtdcGenderType_Female = '2'

// 费用支付标志类型
type TThostFtdcFeePayFlagType byte

// 由受益方支付费用
const TThostFtdcFeePayFlagType_BEN = '0'

// 由发送方支付费用
const TThostFtdcFeePayFlagType_OUR = '1'

// 由发送方支付发起的费用，受益方支付接受的费用
const TThostFtdcFeePayFlagType_SHA = '2'

// 密钥类型类型
type TThostFtdcPassWordKeyTypeType byte

// 交换密钥
const TThostFtdcPassWordKeyTypeType_ExchangeKey = '0'

// 密码密钥
const TThostFtdcPassWordKeyTypeType_PassWordKey = '1'

// MAC密钥
const TThostFtdcPassWordKeyTypeType_MACKey = '2'

// 报文密钥
const TThostFtdcPassWordKeyTypeType_MessageKey = '3'

type TThostFtdcFBTPassWordTypeType byte

// 查询
const TThostFtdcFBTPassWordTypeType_Query = '0'

// 取款
const TThostFtdcFBTPassWordTypeType_Fetch = '1'

// 转帐
const TThostFtdcFBTPassWordTypeType_Transfer = '2'

// 交易
const TThostFtdcFBTPassWordTypeType_Trade = '3'

// 加密方式类型
type TThostFtdcFBTEncryModeType byte

// 不加密
const TThostFtdcFBTEncryModeType_NoEncry = '0'

// DES
const TThostFtdcFBTEncryModeType_DES = '1'

// 3DES
const TThostFtdcFBTEncryModeType_3DES = '2'

// 银行冲正标志类型
type TThostFtdcBankRepealFlagType byte

// 银行无需自动冲正
const TThostFtdcBankRepealFlagType_BankNotNeedRepeal = '0'

// 银行待自动冲正
const TThostFtdcBankRepealFlagType_BankWaitingRepeal = '1'

// 银行已自动冲正
const TThostFtdcBankRepealFlagType_BankBeenRepealed = '2'

// 期商冲正标志类型
type TThostFtdcBrokerRepealFlagType byte

// 期商无需自动冲正
const TThostFtdcBrokerRepealFlagType_BrokerNotNeedRepeal = '0'

// 期商待自动冲正
const TThostFtdcBrokerRepealFlagType_BrokerWaitingRepeal = '1'

// 期商已自动冲正
const TThostFtdcBrokerRepealFlagType_BrokerBeenRepealed = '2'

// 机构类别类型
type TThostFtdcInstitutionTypeType byte

// 银行
const TThostFtdcInstitutionTypeType_Bank = '0'

// 期商
const TThostFtdcInstitutionTypeType_Future = '1'

// 券商
const TThostFtdcInstitutionTypeType_Store = '2'

// 最后分片标志类型
type TThostFtdcLastFragmentType byte

// 是最后分片
const TThostFtdcLastFragmentType_Yes = '0'

// 不是最后分片
const TThostFtdcLastFragmentType_No = '1'

// 银行账户状态类型
type TThostFtdcBankAccStatusType byte

// 正常
const TThostFtdcBankAccStatusType_Normal = '0'

// 冻结
const TThostFtdcBankAccStatusType_Freeze = '1'

// 挂失
const TThostFtdcBankAccStatusType_ReportLoss = '2'

// 资金账户状态类型
type TThostFtdcMoneyAccountStatusType byte

// 正常
const TThostFtdcMoneyAccountStatusType_Normal = '0'

// 销户
const TThostFtdcMoneyAccountStatusType_Cancel = '1'

// 存管状态类型
type TThostFtdcManageStatusType byte

// 指定存管
const TThostFtdcManageStatusType_Point = '0'

// 预指定
const TThostFtdcManageStatusType_PrePoint = '1'

// 撤销指定
const TThostFtdcManageStatusType_CancelPoint = '2'

// 应用系统类型类型
type TThostFtdcSystemTypeType byte

// 银期转帐
const TThostFtdcSystemTypeType_FutureBankTransfer = '0'

// 银证转帐
const TThostFtdcSystemTypeType_StockBankTransfer = '1'

// 第三方存管
const TThostFtdcSystemTypeType_TheThirdPartStore = '2'

// 银期转帐划转结果标志类型
type TThostFtdcTxnEndFlagType byte

// 正常处理中
const TThostFtdcTxnEndFlagType_NormalProcessing = '0'

// 成功结束
const TThostFtdcTxnEndFlagType_Success = '1'

// 失败结束
const TThostFtdcTxnEndFlagType_Failed = '2'

// 异常中
const TThostFtdcTxnEndFlagType_Abnormal = '3'

// 已人工异常处理
const TThostFtdcTxnEndFlagType_ManualProcessedForException = '4'

// 通讯异常 ，请人工处理
const TThostFtdcTxnEndFlagType_CommuFailedNeedManualProcess = '5'

// 系统出错，请人工处理
const TThostFtdcTxnEndFlagType_SysErrorNeedManualProcess = '6'

// 银期转帐服务处理状态类型
type TThostFtdcProcessStatusType byte

// 未处理
const TThostFtdcProcessStatusType_NotProcess = '0'

// 开始处理
const TThostFtdcProcessStatusType_StartProcess = '1'

// 处理完成
const TThostFtdcProcessStatusType_Finished = '2'

// 客户类型类型
type TThostFtdcCustTypeType byte

// 自然人
const TThostFtdcCustTypeType_Person = '0'

// 机构户
const TThostFtdcCustTypeType_Institution = '1'

// 银期转帐方向类型
type TThostFtdcFBTTransferDirectionType byte

// 入金，银行转期货
const TThostFtdcFBTTransferDirectionType_FromBankToFuture = '1'

// 出金，期货转银行
const TThostFtdcFBTTransferDirectionType_FromFutureToBank = '2'

// 开销户类别类型
type TThostFtdcOpenOrDestroyType byte

// 开户
const TThostFtdcOpenOrDestroyType_Open = '1'

// 销户
const TThostFtdcOpenOrDestroyType_Destroy = '0'

type TThostFtdcAvailabilityFlagType byte

// 未确认
const TThostFtdcAvailabilityFlagType_Invalid = '0'

// 有效
const TThostFtdcAvailabilityFlagType_Valid = '1'

// 冲正
const TThostFtdcAvailabilityFlagType_Repeal = '2'

// 机构类型类型
type TThostFtdcOrganTypeType byte

// 银行代理
const TThostFtdcOrganTypeType_Bank = '1'

// 交易前置
const TThostFtdcOrganTypeType_Future = '2'

// 银期转帐平台管理
const TThostFtdcOrganTypeType_PlateForm = '9'

// 机构级别类型
type TThostFtdcOrganLevelType byte

// 银行总行或期商总部
const TThostFtdcOrganLevelType_HeadQuarters = '1'

// 银行分中心或期货公司营业部
const TThostFtdcOrganLevelType_Branch = '2'

// 协议类型类型
type TThostFtdcProtocalIDType byte

// 期商协议
const TThostFtdcProtocalIDType_FutureProtocal = '0'

// 工行协议
const TThostFtdcProtocalIDType_ICBCProtocal = '1'

// 农行协议
const TThostFtdcProtocalIDType_ABCProtocal = '2'

// 中国银行协议
const TThostFtdcProtocalIDType_CBCProtocal = '3'

// 建行协议
const TThostFtdcProtocalIDType_CCBProtocal = '4'

// 交行协议
const TThostFtdcProtocalIDType_BOCOMProtocal = '5'

// 银期转帐平台协议
const TThostFtdcProtocalIDType_FBTPlateFormProtocal = 'X'

// 套接字连接方式类型
type TThostFtdcConnectModeType byte

// 短连接
const TThostFtdcConnectModeType_ShortConnect = '0'

// 长连接
const TThostFtdcConnectModeType_LongConnect = '1'

// 套接字通信方式类型
type TThostFtdcSyncModeType byte

// 异步
const TThostFtdcSyncModeType_ASync = '0'

// 同步
const TThostFtdcSyncModeType_Sync = '1'

// 银行帐号类型类型
type TThostFtdcBankAccTypeType byte

// 银行存折
const TThostFtdcBankAccTypeType_BankBook = '1'

// 储蓄卡
const TThostFtdcBankAccTypeType_SavingCard = '2'

// 信用卡
const TThostFtdcBankAccTypeType_CreditCard = '3'

// 期货公司帐号类型类型
type TThostFtdcFutureAccTypeType byte

// 银行存折
const TThostFtdcFutureAccTypeType_BankBook = '1'

// 储蓄卡
const TThostFtdcFutureAccTypeType_SavingCard = '2'

// 信用卡
const TThostFtdcFutureAccTypeType_CreditCard = '3'

// 接入机构状态类型
type TThostFtdcOrganStatusType byte

// 启用
const TThostFtdcOrganStatusType_Ready = '0'

// 签到
const TThostFtdcOrganStatusType_CheckIn = '1'

// 签退
const TThostFtdcOrganStatusType_CheckOut = '2'

// 对帐文件到达
const TThostFtdcOrganStatusType_CheckFileArrived = '3'

// 对帐
const TThostFtdcOrganStatusType_CheckDetail = '4'

// 日终清理
const TThostFtdcOrganStatusType_DayEndClean = '5'

// 注销
const TThostFtdcOrganStatusType_Invalid = '9'

// 建行收费模式类型
type TThostFtdcCCBFeeModeType byte

// 按金额扣收
const TThostFtdcCCBFeeModeType_ByAmount = '1'

// 按月扣收
const TThostFtdcCCBFeeModeType_ByMonth = '2'

// 通讯API类型类型
type TThostFtdcCommApiTypeType byte

// 客户端
const TThostFtdcCommApiTypeType_Client = '1'

// 服务端
const TThostFtdcCommApiTypeType_Server = '2'

// 交易系统的UserApi
const TThostFtdcCommApiTypeType_UserApi = '3'

// 服务名类型
type TThostFtdcServiceNameType [61]byte

// 连接状态类型
type TThostFtdcLinkStatusType byte

// 已经连接
const TThostFtdcLinkStatusType_Connected = '1'

// 没有连接
const TThostFtdcLinkStatusType_Disconnected = '2'

// 密码核对标志类型
type TThostFtdcPwdFlagType byte

// 不核对
const TThostFtdcPwdFlagType_NoCheck = '0'

// 明文核对
const TThostFtdcPwdFlagType_BlankCheck = '1'

// 密文核对
const TThostFtdcPwdFlagType_EncryptCheck = '2'

// 期货帐号类型类型
type TThostFtdcSecuAccTypeType byte

// 资金帐号
const TThostFtdcSecuAccTypeType_AccountID = '1'

// 资金卡号
const TThostFtdcSecuAccTypeType_CardID = '2'

// 上海股东帐号
const TThostFtdcSecuAccTypeType_SHStockholderID = '3'

// 深圳股东帐号
const TThostFtdcSecuAccTypeType_SZStockholderID = '4'

// 转账交易状态类型
type TThostFtdcTransferStatusType byte

// 正常
const TThostFtdcTransferStatusType_Normal = '0'

// 被冲正
const TThostFtdcTransferStatusType_Repealed = '1'

// 发起方类型
type TThostFtdcSponsorTypeType byte

// 期商
const TThostFtdcSponsorTypeType_Broker = '0'

// 银行
const TThostFtdcSponsorTypeType_Bank = '1'

// 请求响应类别类型
type TThostFtdcReqRspTypeType byte

// 请求
const TThostFtdcReqRspTypeType_Request = '0'

// 响应
const TThostFtdcReqRspTypeType_Response = '1'

// 银期转帐用户事件类型类型
type TThostFtdcFBTUserEventTypeType byte

// 签到
const TThostFtdcFBTUserEventTypeType_SignIn = '0'

// 银行转期货
const TThostFtdcFBTUserEventTypeType_FromBankToFuture = '1'

// 期货转银行
const TThostFtdcFBTUserEventTypeType_FromFutureToBank = '2'

// 开户
const TThostFtdcFBTUserEventTypeType_OpenAccount = '3'

// 销户
const TThostFtdcFBTUserEventTypeType_CancelAccount = '4'

// 变更银行账户
const TThostFtdcFBTUserEventTypeType_ChangeAccount = '5'

// 冲正银行转期货
const TThostFtdcFBTUserEventTypeType_RepealFromBankToFuture = '6'

// 冲正期货转银行
const TThostFtdcFBTUserEventTypeType_RepealFromFutureToBank = '7'

// 查询银行账户
const TThostFtdcFBTUserEventTypeType_QueryBankAccount = '8'

// 查询期货账户
const TThostFtdcFBTUserEventTypeType_QueryFutureAccount = '9'

// 签退
const TThostFtdcFBTUserEventTypeType_SignOut = 'A'

// 密钥同步
const TThostFtdcFBTUserEventTypeType_SyncKey = 'B'

// 预约开户
const TThostFtdcFBTUserEventTypeType_ReserveOpenAccount = 'C'

// 撤销预约开户
const TThostFtdcFBTUserEventTypeType_CancelReserveOpenAccount = 'D'

// 预约开户确认
const TThostFtdcFBTUserEventTypeType_ReserveOpenAccountConfirm = 'E'

// 其他
const TThostFtdcFBTUserEventTypeType_Other = 'Z'

// 银行自己的编码类型
type TThostFtdcBankIDByBankType [21]byte

// 银行操作员号类型
type TThostFtdcBankOperNoType [4]byte

// 银行客户号类型
type TThostFtdcBankCustNoType [21]byte

// FBT表名类型
type TThostFtdcTableNameType [61]byte

// FBT表操作主键名类型
type TThostFtdcPKNameType [201]byte

// FBT表操作主键值类型
type TThostFtdcPKValueType [501]byte

// 记录操作类型类型
type TThostFtdcDBOperationType byte

// 插入
const TThostFtdcDBOperationType_Insert = '0'

// 更新
const TThostFtdcDBOperationType_Update = '1'

// 删除
const TThostFtdcDBOperationType_Delete = '2'

// 同步标记类型
type TThostFtdcSyncFlagType byte

// 已同步
const TThostFtdcSyncFlagType_Yes = '0'

// 未同步
const TThostFtdcSyncFlagType_No = '1'

// 同步目标编号类型
type TThostFtdcTargetIDType [4]byte

// 同步类型类型
type TThostFtdcSyncTypeType byte

// 一次同步
const TThostFtdcSyncTypeType_OneOffSync = '0'

// 定时同步
const TThostFtdcSyncTypeType_TimerSync = '1'

// 定时完全同步
const TThostFtdcSyncTypeType_TimerFullSync = '2'

// 各种换汇时间类型
type TThostFtdcFBETimeType [7]byte

// 换汇银行行号类型
type TThostFtdcFBEBankNoType [13]byte

// 换汇凭证号类型
type TThostFtdcFBECertNoType [13]byte

// 换汇方向类型
type TThostFtdcExDirectionType byte

// 结汇
const TThostFtdcExDirectionType_Settlement = '0'

// 售汇
const TThostFtdcExDirectionType_Sale = '1'

// 换汇银行账户类型
type TThostFtdcFBEBankAccountType [33]byte

// 换汇银行账户名类型
type TThostFtdcFBEBankAccountNameType [61]byte

// 换汇业务类型类型
type TThostFtdcFBEBusinessTypeType [3]byte

// 换汇附言类型
type TThostFtdcFBEPostScriptType [61]byte

// 换汇备注类型
type TThostFtdcFBERemarkType [71]byte

// 换汇成功标志类型
type TThostFtdcFBEResultFlagType byte

// 成功
const TThostFtdcFBEResultFlagType_Success = '0'

// 账户余额不足
const TThostFtdcFBEResultFlagType_InsufficientBalance = '1'

// 交易结果未知
const TThostFtdcFBEResultFlagType_UnknownTrading = '8'

// 失败
const TThostFtdcFBEResultFlagType_Fail = 'x'

// 换汇返回信息类型
type TThostFtdcFBERtnMsgType [61]byte

// 换汇扩展信息类型
type TThostFtdcFBEExtendMsgType [61]byte

// 换汇记账流水号类型
type TThostFtdcFBEBusinessSerialType [31]byte

// 换汇流水号类型
type TThostFtdcFBESystemSerialType [21]byte

// 换汇交易状态类型
type TThostFtdcFBEExchStatusType byte

// 正常
const TThostFtdcFBEExchStatusType_Normal = '0'

// 交易重发
const TThostFtdcFBEExchStatusType_ReExchange = '1'

// 换汇文件标志类型
type TThostFtdcFBEFileFlagType byte

// 数据包
const TThostFtdcFBEFileFlagType_DataPackage = '0'

// 文件
const TThostFtdcFBEFileFlagType_File = '1'

// 换汇已交易标志类型
type TThostFtdcFBEAlreadyTradeType byte

// 未交易
const TThostFtdcFBEAlreadyTradeType_NotTrade = '0'

// 已交易
const TThostFtdcFBEAlreadyTradeType_Trade = '1'

// 换汇账户开户行类型
type TThostFtdcFBEOpenBankType [61]byte

// 银期换汇用户事件类型类型
type TThostFtdcFBEUserEventTypeType byte

// 签到
const TThostFtdcFBEUserEventTypeType_SignIn = '0'

// 换汇
const TThostFtdcFBEUserEventTypeType_Exchange = '1'

// 换汇重发
const TThostFtdcFBEUserEventTypeType_ReExchange = '2'

// 银行账户查询
const TThostFtdcFBEUserEventTypeType_QueryBankAccount = '3'

// 换汇明细查询
const TThostFtdcFBEUserEventTypeType_QueryExchDetial = '4'

// 换汇汇总查询
const TThostFtdcFBEUserEventTypeType_QueryExchSummary = '5'

// 换汇汇率查询
const TThostFtdcFBEUserEventTypeType_QueryExchRate = '6'

// 对账文件通知
const TThostFtdcFBEUserEventTypeType_CheckBankAccount = '7'

// 签退
const TThostFtdcFBEUserEventTypeType_SignOut = '8'

// 其他
const TThostFtdcFBEUserEventTypeType_Other = 'Z'

// 换汇相关文件名类型
type TThostFtdcFBEFileNameType [21]byte

// 换汇批次号类型
type TThostFtdcFBEBatchSerialType [21]byte

// 换汇发送标志类型
type TThostFtdcFBEReqFlagType byte

// 未处理
const TThostFtdcFBEReqFlagType_UnProcessed = '0'

// 等待发送
const TThostFtdcFBEReqFlagType_WaitSend = '1'

// 发送成功
const TThostFtdcFBEReqFlagType_SendSuccess = '2'

// 发送失败
const TThostFtdcFBEReqFlagType_SendFailed = '3'

// 等待重发
const TThostFtdcFBEReqFlagType_WaitReSend = '4'

// 风险通知类型类型
type TThostFtdcNotifyClassType byte

// 正常
const TThostFtdcNotifyClassType_NOERROR = '0'

// 警示
const TThostFtdcNotifyClassType_Warn = '1'

// 追保
const TThostFtdcNotifyClassType_Call = '2'

// 强平
const TThostFtdcNotifyClassType_Force = '3'

// 穿仓
const TThostFtdcNotifyClassType_CHUANCANG = '4'

// 异常
const TThostFtdcNotifyClassType_Exception = '5'

// 客户风险通知消息类型
type TThostFtdcRiskNofityInfoType [257]byte

// 强平场景编号类型
type TThostFtdcForceCloseSceneIdType [24]byte

// 强平单类型类型
type TThostFtdcForceCloseTypeType byte

// 手工强平
const TThostFtdcForceCloseTypeType_Manual = '0'

// 单一投资者辅助强平
const TThostFtdcForceCloseTypeType_Single = '1'

// 批量投资者辅助强平
const TThostFtdcForceCloseTypeType_Group = '2'

// 多个产品代码,用+分隔,如cu+zn类型
type TThostFtdcInstrumentIDsType [101]byte

// 风险通知途径类型
type TThostFtdcRiskNotifyMethodType byte

// 系统通知
const TThostFtdcRiskNotifyMethodType_System = '0'

// 短信通知
const TThostFtdcRiskNotifyMethodType_SMS = '1'

// 邮件通知
const TThostFtdcRiskNotifyMethodType_EMail = '2'

// 人工通知
const TThostFtdcRiskNotifyMethodType_Manual = '3'

// 风险通知状态类型
type TThostFtdcRiskNotifyStatusType byte

// 未生成
const TThostFtdcRiskNotifyStatusType_NotGen = '0'

// 已生成未发送
const TThostFtdcRiskNotifyStatusType_Generated = '1'

// 发送失败
const TThostFtdcRiskNotifyStatusType_SendError = '2'

// 已发送未接收
const TThostFtdcRiskNotifyStatusType_SendOk = '3'

// 已接收未确认
const TThostFtdcRiskNotifyStatusType_Received = '4'

// 已确认
const TThostFtdcRiskNotifyStatusType_Confirmed = '5'

// 风控用户操作事件类型
type TThostFtdcRiskUserEventType byte

// 导出数据
const TThostFtdcRiskUserEventType_ExportData = '0'

// 参数名类型
type TThostFtdcParamNameType [41]byte

// 参数值类型
type TThostFtdcParamValueType [41]byte

// 条件单索引条件类型
type TThostFtdcConditionalOrderSortTypeType byte

// 使用最新价升序
const TThostFtdcConditionalOrderSortTypeType_LastPriceAsc = '0'

// 使用最新价降序
const TThostFtdcConditionalOrderSortTypeType_LastPriceDesc = '1'

// 使用卖价升序
const TThostFtdcConditionalOrderSortTypeType_AskPriceAsc = '2'

// 使用卖价降序
const TThostFtdcConditionalOrderSortTypeType_AskPriceDesc = '3'

// 使用买价升序
const TThostFtdcConditionalOrderSortTypeType_BidPriceAsc = '4'

// 使用买价降序
const TThostFtdcConditionalOrderSortTypeType_BidPriceDesc = '5'

// 报送状态类型
type TThostFtdcSendTypeType byte

// 未发送
const TThostFtdcSendTypeType_NoSend = '0'

// 已发送
const TThostFtdcSendTypeType_Sended = '1'

// 已生成
const TThostFtdcSendTypeType_Generated = '2'

// 报送失败
const TThostFtdcSendTypeType_SendFail = '3'

// 接收成功
const TThostFtdcSendTypeType_Success = '4'

// 接收失败
const TThostFtdcSendTypeType_Fail = '5'

// 取消报送
const TThostFtdcSendTypeType_Cancel = '6'

// 交易编码状态类型
type TThostFtdcClientIDStatusType byte

// 未申请
const TThostFtdcClientIDStatusType_NoApply = '1'

// 已提交申请
const TThostFtdcClientIDStatusType_Submited = '2'

// 已发送申请
const TThostFtdcClientIDStatusType_Sended = '3'

// 完成
const TThostFtdcClientIDStatusType_Success = '4'

// 拒绝
const TThostFtdcClientIDStatusType_Refuse = '5'

// 已撤销编码
const TThostFtdcClientIDStatusType_Cancel = '6'

// 行业编码类型
type TThostFtdcIndustryIDType [17]byte

// 特有信息编号类型
type TThostFtdcQuestionIDType [5]byte

// 特有信息说明类型
type TThostFtdcQuestionContentType [41]byte

// 选项编号类型
type TThostFtdcOptionIDType [13]byte

// 选项说明类型
type TThostFtdcOptionContentType [61]byte

// 特有信息类型类型
type TThostFtdcQuestionTypeType byte

// 单选
const TThostFtdcQuestionTypeType_Radio = '1'

// 多选
const TThostFtdcQuestionTypeType_Option = '2'

// 填空
const TThostFtdcQuestionTypeType_Blank = '3'

// 业务流水号类型
type TThostFtdcProcessIDType [33]byte

// 流程状态类型
type TThostFtdcUOAProcessStatusType [3]byte

// 流程功能类型类型
type TThostFtdcProcessTypeType [3]byte

// 业务类型类型
type TThostFtdcBusinessTypeType byte

// 请求
const TThostFtdcBusinessTypeType_Request = '1'

// 应答
const TThostFtdcBusinessTypeType_Response = '2'

// 通知
const TThostFtdcBusinessTypeType_Notice = '3'

// 监控中心返回码类型
type TThostFtdcCfmmcReturnCodeType byte

// 成功
const TThostFtdcCfmmcReturnCodeType_Success = '0'

// 该客户已经有流程在处理中
const TThostFtdcCfmmcReturnCodeType_Working = '1'

// 监控中客户资料检查失败
const TThostFtdcCfmmcReturnCodeType_InfoFail = '2'

// 监控中实名制检查失败
const TThostFtdcCfmmcReturnCodeType_IDCardFail = '3'

// 其他错误
const TThostFtdcCfmmcReturnCodeType_OtherFail = '4'

type TThostFtdcClientTypeType byte

// 所有
const TThostFtdcClientTypeType_All = '0'

// 个人
const TThostFtdcClientTypeType_Person = '1'

// 单位
const TThostFtdcClientTypeType_Company = '2'

// 其他
const TThostFtdcClientTypeType_Other = '3'

// 特殊法人
const TThostFtdcClientTypeType_SpecialOrgan = '4'

// 资管户
const TThostFtdcClientTypeType_Asset = '5'

// 交易所编号类型
type TThostFtdcExchangeIDTypeType byte

// 上海期货交易所
const TThostFtdcExchangeIDTypeType_SHFE = 'S'

// 郑州商品交易所
const TThostFtdcExchangeIDTypeType_CZCE = 'Z'

// 大连商品交易所
const TThostFtdcExchangeIDTypeType_DCE = 'D'

// 中国金融期货交易所
const TThostFtdcExchangeIDTypeType_CFFEX = 'J'

// 上海国际能源交易中心股份有限公司
const TThostFtdcExchangeIDTypeType_INE = 'N'

type TThostFtdcExClientIDTypeType byte

// 套保
const TThostFtdcExClientIDTypeType_Hedge = '1'

// 套利
const TThostFtdcExClientIDTypeType_Arbitrage = '2'

// 投机
const TThostFtdcExClientIDTypeType_Speculation = '3'

// 客户分类码类型
type TThostFtdcClientClassifyType [11]byte

// 单位性质类型
type TThostFtdcUOAOrganTypeType [11]byte
type TThostFtdcUOACountryCodeType [11]byte

// 区号类型
type TThostFtdcAreaCodeType [11]byte

// 监控中心为客户分配的代码类型
type TThostFtdcFuturesIDType [21]byte
type TThostFtdcCffmcDateType [11]byte
type TThostFtdcCffmcTimeType [11]byte

// 组织机构代码类型
type TThostFtdcNocIDType [21]byte

// 更新状态类型
type TThostFtdcUpdateFlagType byte

// 未更新
const TThostFtdcUpdateFlagType_NoUpdate = '0'

// 更新全部信息成功
const TThostFtdcUpdateFlagType_Success = '1'

// 更新全部信息失败
const TThostFtdcUpdateFlagType_Fail = '2'

// 更新交易编码成功
const TThostFtdcUpdateFlagType_TCSuccess = '3'

// 更新交易编码失败
const TThostFtdcUpdateFlagType_TCFail = '4'

// 已丢弃
const TThostFtdcUpdateFlagType_Cancel = '5'

// 申请动作类型
type TThostFtdcApplyOperateIDType byte

// 开户
const TThostFtdcApplyOperateIDType_OpenInvestor = '1'

// 修改身份信息
const TThostFtdcApplyOperateIDType_ModifyIDCard = '2'

// 修改一般信息
const TThostFtdcApplyOperateIDType_ModifyNoIDCard = '3'

// 申请交易编码
const TThostFtdcApplyOperateIDType_ApplyTradingCode = '4'

// 撤销交易编码
const TThostFtdcApplyOperateIDType_CancelTradingCode = '5'

// 销户
const TThostFtdcApplyOperateIDType_CancelInvestor = '6'

// 账户休眠
const TThostFtdcApplyOperateIDType_FreezeAccount = '8'

// 激活休眠账户
const TThostFtdcApplyOperateIDType_ActiveFreezeAccount = '9'

// 申请状态类型
type TThostFtdcApplyStatusIDType byte

// 未补全
const TThostFtdcApplyStatusIDType_NoComplete = '1'

// 已提交
const TThostFtdcApplyStatusIDType_Submited = '2'

// 已审核
const TThostFtdcApplyStatusIDType_Checked = '3'

// 已拒绝
const TThostFtdcApplyStatusIDType_Refused = '4'

// 已删除
const TThostFtdcApplyStatusIDType_Deleted = '5'

// 发送方式类型
type TThostFtdcSendMethodType byte

// 文件发送
const TThostFtdcSendMethodType_ByAPI = '1'

// 电子发送
const TThostFtdcSendMethodType_ByFile = '2'

// 业务操作类型类型
type TThostFtdcEventTypeType [33]byte

// 操作方法类型
type TThostFtdcEventModeType byte

// 增加
const TThostFtdcEventModeType_ADD = '1'

// 修改
const TThostFtdcEventModeType_UPDATE = '2'

// 删除
const TThostFtdcEventModeType_DELETE = '3'

// 复核
const TThostFtdcEventModeType_CHECK = '4'

// 复制
const TThostFtdcEventModeType_COPY = '5'

// 注销
const TThostFtdcEventModeType_CANCEL = '6'

// 冲销
const TThostFtdcEventModeType_Reverse = '7'

// 统一开户申请自动发送类型
type TThostFtdcUOAAutoSendType byte

// 自动发送并接收
const TThostFtdcUOAAutoSendType_ASR = '1'

// 自动发送，不自动接收
const TThostFtdcUOAAutoSendType_ASNR = '2'

// 不自动发送，自动接收
const TThostFtdcUOAAutoSendType_NSAR = '3'

// 不自动发送，也不自动接收
const TThostFtdcUOAAutoSendType_NSR = '4'

// 流程ID类型
type TThostFtdcFlowIDType byte

// 投资者对应投资者组设置
const TThostFtdcFlowIDType_InvestorGroupFlow = '1'

// 投资者手续费率设置
const TThostFtdcFlowIDType_InvestorRate = '2'

// 投资者手续费率模板关系设置
const TThostFtdcFlowIDType_InvestorCommRateModel = '3'

// 复核级别类型
type TThostFtdcCheckLevelType byte

// 零级复核
const TThostFtdcCheckLevelType_Zero = '0'

// 一级复核
const TThostFtdcCheckLevelType_One = '1'

// 二级复核
const TThostFtdcCheckLevelType_Two = '2'

type TThostFtdcCheckStatusType byte

// 未复核
const TThostFtdcCheckStatusType_Init = '0'

// 复核中
const TThostFtdcCheckStatusType_Checking = '1'

// 已复核
const TThostFtdcCheckStatusType_Checked = '2'

// 拒绝
const TThostFtdcCheckStatusType_Refuse = '3'

// 作废
const TThostFtdcCheckStatusType_Cancel = '4'

// 生效状态类型
type TThostFtdcUsedStatusType byte

// 未生效
const TThostFtdcUsedStatusType_Unused = '0'

// 已生效
const TThostFtdcUsedStatusType_Used = '1'

// 生效失败
const TThostFtdcUsedStatusType_Fail = '2'

// 模型名称类型
type TThostFtdcRateTemplateNameType [61]byte

// 用于查询的投资属性字段类型
type TThostFtdcPropertyStringType [2049]byte

// 账户来源类型
type TThostFtdcBankAcountOriginType byte

// 手工录入
const TThostFtdcBankAcountOriginType_ByAccProperty = '0'

// 银期转账
const TThostFtdcBankAcountOriginType_ByFBTransfer = '1'

// 结算单月报成交汇总方式类型
type TThostFtdcMonthBillTradeSumType byte

// 同日同合约
const TThostFtdcMonthBillTradeSumType_ByInstrument = '0'

// 同日同合约同价格
const TThostFtdcMonthBillTradeSumType_ByDayInsPrc = '1'

// 同合约
const TThostFtdcMonthBillTradeSumType_ByDayIns = '2'

// 银期交易代码枚举类型
type TThostFtdcFBTTradeCodeEnumType byte

// 银行发起银行转期货
const TThostFtdcFBTTradeCodeEnumType_BankLaunchBankToBroker = "102001"

// 期货发起银行转期货
const TThostFtdcFBTTradeCodeEnumType_BrokerLaunchBankToBroker = "202001"

// 银行发起期货转银行
const TThostFtdcFBTTradeCodeEnumType_BankLaunchBrokerToBank = "102002"

// 期货发起期货转银行
const TThostFtdcFBTTradeCodeEnumType_BrokerLaunchBrokerToBank = "202002"

// 模型代码类型
type TThostFtdcRateTemplateIDType [9]byte

// 风险度类型
type TThostFtdcRiskRateType [21]byte

// 号段规则名称类型
type TThostFtdcInvestorIDRuleNameType [61]byte

// 号段规则表达式类型
type TThostFtdcInvestorIDRuleExprType [513]byte

// 令牌密钥类型
type TThostFtdcAuthKeyType [41]byte

// 序列号类型
type TThostFtdcSerialNumberType [17]byte

// 动态令牌类型类型
type TThostFtdcOTPTypeType byte

// 无动态令牌
const TThostFtdcOTPTypeType_NONE = '0'

// 时间令牌
const TThostFtdcOTPTypeType_TOTP = '1'

// 动态令牌提供商类型
type TThostFtdcOTPVendorsIDType [2]byte

// 动态令牌提供商名称类型
type TThostFtdcOTPVendorsNameType [61]byte

// 动态令牌状态类型
type TThostFtdcOTPStatusType byte

// 未使用
const TThostFtdcOTPStatusType_Unused = '0'

// 已使用
const TThostFtdcOTPStatusType_Used = '1'

// 注销
const TThostFtdcOTPStatusType_Disuse = '2'

// 经济公司用户类型类型
type TThostFtdcBrokerUserTypeType byte

// 投资者
const TThostFtdcBrokerUserTypeType_Investor = '1'

// 操作员
const TThostFtdcBrokerUserTypeType_BrokerUser = '2'

// 期货类型类型
type TThostFtdcFutureTypeType byte

// 商品期货
const TThostFtdcFutureTypeType_Commodity = '1'

// 金融期货
const TThostFtdcFutureTypeType_Financial = '2'

// 资金管理操作类型类型
type TThostFtdcFundEventTypeType byte

// 转账限额
const TThostFtdcFundEventTypeType_Restriction = '0'

// 当日转账限额
const TThostFtdcFundEventTypeType_TodayRestriction = '1'

// 期商流水
const TThostFtdcFundEventTypeType_Transfer = '2'

// 资金冻结
const TThostFtdcFundEventTypeType_Credit = '3'

// 投资者可提资金比例
const TThostFtdcFundEventTypeType_InvestorWithdrawAlm = '4'

// 单个银行帐户转账限额
const TThostFtdcFundEventTypeType_BankRestriction = '5'

// 银期签约账户
const TThostFtdcFundEventTypeType_Accountregister = '6'

// 交易所出入金
const TThostFtdcFundEventTypeType_ExchangeFundIO = '7'

// 投资者出入金
const TThostFtdcFundEventTypeType_InvestorFundIO = '8'

// 资金账户来源类型
type TThostFtdcAccountSourceTypeType byte

// 银期同步
const TThostFtdcAccountSourceTypeType_FBTransfer = '0'

// 手工录入
const TThostFtdcAccountSourceTypeType_ManualEntry = '1'

// 交易编码来源类型
type TThostFtdcCodeSourceTypeType byte

// 统一开户(已规范)
const TThostFtdcCodeSourceTypeType_UnifyAccount = '0'

// 手工录入(未规范)
const TThostFtdcCodeSourceTypeType_ManualEntry = '1'

// 操作员范围类型
type TThostFtdcUserRangeType byte

// 所有
const TThostFtdcUserRangeType_All = '0'

// 单一操作员
const TThostFtdcUserRangeType_Single = '1'

// 时间跨度类型
type TThostFtdcTimeSpanType [9]byte

// 动态令牌导入批次编号类型
type TThostFtdcImportSequenceIDType [17]byte

// 交易统计表按客户统计方式类型
type TThostFtdcByGroupType byte

// 按投资者统计
const TThostFtdcByGroupType_Investor = '2'

// 按类统计
const TThostFtdcByGroupType_Group = '1'

// 交易统计表按范围统计方式类型
type TThostFtdcTradeSumStatModeType byte

// 按合约统计
const TThostFtdcTradeSumStatModeType_Instrument = '1'

// 按产品统计
const TThostFtdcTradeSumStatModeType_Product = '2'

// 按交易所统计
const TThostFtdcTradeSumStatModeType_Exchange = '3'

// 产品标识类型
type TThostFtdcUserProductIDType [33]byte
type TThostFtdcUserProductNameType [65]byte

// 产品说明类型
type TThostFtdcUserProductMemoType [129]byte

// 新增或变更标志类型
type TThostFtdcCSRCCancelFlagType [2]byte
type TThostFtdcCSRCDateType [11]byte

// 客户名称类型
type TThostFtdcCSRCInvestorNameType [201]byte
type TThostFtdcCSRCOpenInvestorNameType [101]byte

// 客户代码类型
type TThostFtdcCSRCInvestorIDType [13]byte
type TThostFtdcCSRCIdentifiedCardNoType [51]byte
type TThostFtdcCSRCClientIDType [11]byte

// 银行标识类型
type TThostFtdcCSRCBankFlagType [3]byte
type TThostFtdcCSRCBankAccountType [23]byte

// 开户人类型
type TThostFtdcCSRCOpenNameType [401]byte

// 说明类型
type TThostFtdcCSRCMemoType [101]byte
type TThostFtdcCSRCTimeType [11]byte

// 成交流水号类型
type TThostFtdcCSRCTradeIDType [21]byte
type TThostFtdcCSRCExchangeInstIDType [31]byte

// 质押品名称类型
type TThostFtdcCSRCMortgageNameType [7]byte
type TThostFtdcCSRCReasonType [3]byte

// 是否为非结算会员类型
type TThostFtdcIsSettlementType [2]byte

// 期权类型类型
type TThostFtdcCSRCOptionsTypeType [2]byte

// 标的品种类型
type TThostFtdcCSRCTargetProductIDType [3]byte

// 标的合约类型
type TThostFtdcCSRCTargetInstrIDType [31]byte

// 手续费率模板名称类型
type TThostFtdcCommModelNameType [161]byte

// 手续费率模板备注类型
type TThostFtdcCommModelMemoType [1025]byte

// 日期表达式设置类型类型
type TThostFtdcExprSetModeType byte

// 相对已有规则设置
const TThostFtdcExprSetModeType_Relative = '1'

// 典型设置
const TThostFtdcExprSetModeType_Typical = '2'

type TThostFtdcRateInvestorRangeType byte

// 公司标准
const TThostFtdcRateInvestorRangeType_All = '1'

// 模板
const TThostFtdcRateInvestorRangeType_Model = '2'

// 单一投资者
const TThostFtdcRateInvestorRangeType_Single = '3'

// 代理经纪公司代码类型
type TThostFtdcAgentBrokerIDType [13]byte

// 交易中心名称类型
type TThostFtdcDRIdentityNameType [65]byte

// DBLink标识号类型
type TThostFtdcDBLinkIDType [31]byte

// 主次用系统数据同步状态类型
type TThostFtdcSyncDataStatusType byte

// 未同步
const TThostFtdcSyncDataStatusType_Initialize = '0'

// 同步中
const TThostFtdcSyncDataStatusType_Settlementing = '1'

// 已同步
const TThostFtdcSyncDataStatusType_Settlemented = '2'

// 成交来源类型
type TThostFtdcTradeSourceType byte

// 来自交易所普通回报
const TThostFtdcTradeSourceType_NORMAL = '0'

// 来自查询
const TThostFtdcTradeSourceType_QUERY = '1'

// 产品合约统计方式类型
type TThostFtdcFlexStatModeType byte

// 产品统计
const TThostFtdcFlexStatModeType_Product = '1'

// 交易所统计
const TThostFtdcFlexStatModeType_Exchange = '2'

// 统计所有
const TThostFtdcFlexStatModeType_All = '3'

// 投资者范围统计方式类型
type TThostFtdcByInvestorRangeType byte

// 属性统计
const TThostFtdcByInvestorRangeType_Property = '1'

// 统计所有
const TThostFtdcByInvestorRangeType_All = '2'

type TThostFtdcSRiskRateType [21]byte
type TThostFtdcPropertyInvestorRangeType byte

// 所有
const TThostFtdcPropertyInvestorRangeType_All = '1'

// 投资者属性
const TThostFtdcPropertyInvestorRangeType_Property = '2'

// 单一投资者
const TThostFtdcPropertyInvestorRangeType_Single = '3'

type TThostFtdcFileStatusType byte

// 未生成
const TThostFtdcFileStatusType_NoCreate = '0'

// 已生成
const TThostFtdcFileStatusType_Created = '1'

// 生成失败
const TThostFtdcFileStatusType_Failed = '2'

// 文件生成方式类型
type TThostFtdcFileGenStyleType byte

// 下发
const TThostFtdcFileGenStyleType_FileTransmit = '0'

// 生成
const TThostFtdcFileGenStyleType_FileGen = '1'

// 系统日志操作方法类型
type TThostFtdcSysOperModeType byte

// 增加
const TThostFtdcSysOperModeType_Add = '1'

// 修改
const TThostFtdcSysOperModeType_Update = '2'

// 删除
const TThostFtdcSysOperModeType_Delete = '3'

// 复制
const TThostFtdcSysOperModeType_Copy = '4'

// 激活
const TThostFtdcSysOperModeType_AcTive = '5'

// 注销
const TThostFtdcSysOperModeType_CanCel = '6'

// 重置
const TThostFtdcSysOperModeType_ReSet = '7'

// 系统日志操作类型类型
type TThostFtdcSysOperTypeType byte

// 修改操作员密码
const TThostFtdcSysOperTypeType_UpdatePassword = '0'

// 操作员组织架构关系
const TThostFtdcSysOperTypeType_UserDepartment = '1'

// 角色管理
const TThostFtdcSysOperTypeType_RoleManager = '2'

// 角色功能设置
const TThostFtdcSysOperTypeType_RoleFunction = '3'

// 基础参数设置
const TThostFtdcSysOperTypeType_BaseParam = '4'

// 设置操作员
const TThostFtdcSysOperTypeType_SetUserID = '5'

// 用户角色设置
const TThostFtdcSysOperTypeType_SetUserRole = '6'

// 用户IP限制
const TThostFtdcSysOperTypeType_UserIpRestriction = '7'

// 组织架构管理
const TThostFtdcSysOperTypeType_DepartmentManager = '8'

// 组织架构向查询分类复制
const TThostFtdcSysOperTypeType_DepartmentCopy = '9'

// 交易编码管理
const TThostFtdcSysOperTypeType_Tradingcode = 'A'

// 投资者状态维护
const TThostFtdcSysOperTypeType_InvestorStatus = 'B'

// 投资者权限管理
const TThostFtdcSysOperTypeType_InvestorAuthority = 'C'

// 属性设置
const TThostFtdcSysOperTypeType_PropertySet = 'D'

// 重置投资者密码
const TThostFtdcSysOperTypeType_ReSetInvestorPasswd = 'E'

// 投资者个性信息维护
const TThostFtdcSysOperTypeType_InvestorPersonalityInfo = 'F'

// 上报数据查询类型类型
type TThostFtdcCSRCDataQueyTypeType byte

// 查询当前交易日报送的数据
const TThostFtdcCSRCDataQueyTypeType_Current = '0'

// 查询历史报送的代理经纪公司的数据
const TThostFtdcCSRCDataQueyTypeType_History = '1'

// 休眠状态类型
type TThostFtdcFreezeStatusType byte

// 活跃
const TThostFtdcFreezeStatusType_Normal = '1'

// 休眠
const TThostFtdcFreezeStatusType_Freeze = '0'

// 规范状态类型
type TThostFtdcStandardStatusType byte

// 已规范
const TThostFtdcStandardStatusType_Standard = '0'

// 未规范
const TThostFtdcStandardStatusType_NonStandard = '1'

type TThostFtdcCSRCFreezeStatusType [2]byte

// 配置类型类型
type TThostFtdcRightParamTypeType byte

// 休眠户
const TThostFtdcRightParamTypeType_Freeze = '1'

// 激活休眠户
const TThostFtdcRightParamTypeType_FreezeActive = '2'

// 开仓权限限制
const TThostFtdcRightParamTypeType_OpenLimit = '3'

// 解除开仓权限限制
const TThostFtdcRightParamTypeType_RelieveOpenLimit = '4'

// 模板代码类型
type TThostFtdcRightTemplateIDType [9]byte

// 模板名称类型
type TThostFtdcRightTemplateNameType [61]byte

// 反洗钱审核表数据状态类型
type TThostFtdcDataStatusType byte

// 正常
const TThostFtdcDataStatusType_Normal = '0'

// 已删除
const TThostFtdcDataStatusType_Deleted = '1'

// 审核状态类型
type TThostFtdcAMLCheckStatusType byte

// 未复核
const TThostFtdcAMLCheckStatusType_Init = '0'

// 复核中
const TThostFtdcAMLCheckStatusType_Checking = '1'

// 已复核
const TThostFtdcAMLCheckStatusType_Checked = '2'

// 拒绝上报
const TThostFtdcAMLCheckStatusType_RefuseReport = '3'

// 日期类型类型
type TThostFtdcAmlDateTypeType byte

// 检查日期
const TThostFtdcAmlDateTypeType_DrawDay = '0'

// 发生日期
const TThostFtdcAmlDateTypeType_TouchDay = '1'

// 审核级别类型
type TThostFtdcAmlCheckLevelType byte

// 零级审核
const TThostFtdcAmlCheckLevelType_CheckLevel0 = '0'

// 一级审核
const TThostFtdcAmlCheckLevelType_CheckLevel1 = '1'

// 二级审核
const TThostFtdcAmlCheckLevelType_CheckLevel2 = '2'

// 三级审核
const TThostFtdcAmlCheckLevelType_CheckLevel3 = '3'

// 反洗钱数据抽取审核流程类型
type TThostFtdcAmlCheckFlowType [2]byte

// 数据类型类型
type TThostFtdcDataTypeType [129]byte

// 导出文件类型类型
type TThostFtdcExportFileTypeType byte

// CSV
const TThostFtdcExportFileTypeType_CSV = '0'

// Excel
const TThostFtdcExportFileTypeType_EXCEL = '1'

// DBF
const TThostFtdcExportFileTypeType_DBF = '2'

// 结算配置类型类型
type TThostFtdcSettleManagerTypeType byte

// 结算前准备
const TThostFtdcSettleManagerTypeType_Before = '1'

// 结算
const TThostFtdcSettleManagerTypeType_Settlement = '2'

// 结算后核对
const TThostFtdcSettleManagerTypeType_After = '3'

// 结算后处理
const TThostFtdcSettleManagerTypeType_Settlemented = '4'

// 结算配置代码类型
type TThostFtdcSettleManagerIDType [33]byte

// 结算配置名称类型
type TThostFtdcSettleManagerNameType [129]byte

// 结算配置等级类型
type TThostFtdcSettleManagerLevelType byte

// 必要
const TThostFtdcSettleManagerLevelType_Must = '1'

// 警告
const TThostFtdcSettleManagerLevelType_Alarm = '2'

// 提示
const TThostFtdcSettleManagerLevelType_Prompt = '3'

// 不检查
const TThostFtdcSettleManagerLevelType_Ignore = '4'

// 模块分组类型
type TThostFtdcSettleManagerGroupType byte

// 交易所核对
const TThostFtdcSettleManagerGroupType_Exhcange = '1'

// 内部核对
const TThostFtdcSettleManagerGroupType_ASP = '2'

// 上报数据核对
const TThostFtdcSettleManagerGroupType_CSRC = '3'

// 核对结果说明类型
type TThostFtdcCheckResultMemoType [1025]byte

// 功能链接类型
type TThostFtdcFunctionUrlType [1025]byte

// 客户端认证信息类型
type TThostFtdcAuthInfoType [129]byte

// 客户端认证码类型
type TThostFtdcAuthCodeType [17]byte

// 保值额度使用类型类型
type TThostFtdcLimitUseTypeType byte

// 可重复使用
const TThostFtdcLimitUseTypeType_Repeatable = '1'

// 不可重复使用
const TThostFtdcLimitUseTypeType_Unrepeatable = '2'

// 数据来源类型
type TThostFtdcDataResourceType byte

// 本系统
const TThostFtdcDataResourceType_Settle = '1'

// 交易所
const TThostFtdcDataResourceType_Exchange = '2'

// 报送数据
const TThostFtdcDataResourceType_CSRC = '3'

// 保证金类型类型
type TThostFtdcMarginTypeType byte

// 交易所保证金率
const TThostFtdcMarginTypeType_ExchMarginRate = '0'

// 投资者保证金率
const TThostFtdcMarginTypeType_InstrMarginRate = '1'

// 投资者交易保证金率
const TThostFtdcMarginTypeType_InstrMarginRateTrade = '2'

// 生效类型类型
type TThostFtdcActiveTypeType byte

// 仅当日生效
const TThostFtdcActiveTypeType_Intraday = '1'

// 长期生效
const TThostFtdcActiveTypeType_Long = '2'

// 冲突保证金率类型类型
type TThostFtdcMarginRateTypeType byte

// 交易所保证金率
const TThostFtdcMarginRateTypeType_Exchange = '1'

// 投资者保证金率
const TThostFtdcMarginRateTypeType_Investor = '2'

// 投资者交易保证金率
const TThostFtdcMarginRateTypeType_InvestorTrade = '3'

// 备份数据状态类型
type TThostFtdcBackUpStatusType byte

// 未生成备份数据
const TThostFtdcBackUpStatusType_UnBak = '0'

// 备份数据生成中
const TThostFtdcBackUpStatusType_BakUp = '1'

// 已生成备份数据
const TThostFtdcBackUpStatusType_BakUped = '2'

// 备份数据失败
const TThostFtdcBackUpStatusType_BakFail = '3'

// 结算初始化状态类型
type TThostFtdcInitSettlementType byte

// 结算初始化未开始
const TThostFtdcInitSettlementType_UnInitialize = '0'

// 结算初始化中
const TThostFtdcInitSettlementType_Initialize = '1'

// 结算初始化完成
const TThostFtdcInitSettlementType_Initialized = '2'

// 报表数据生成状态类型
type TThostFtdcReportStatusType byte

// 未生成报表数据
const TThostFtdcReportStatusType_NoCreate = '0'

// 报表数据生成中
const TThostFtdcReportStatusType_Create = '1'

// 已生成报表数据
const TThostFtdcReportStatusType_Created = '2'

// 生成报表数据失败
const TThostFtdcReportStatusType_CreateFail = '3'

// 数据归档状态类型
type TThostFtdcSaveStatusType byte

// 归档未完成
const TThostFtdcSaveStatusType_UnSaveData = '0'

// 归档完成
const TThostFtdcSaveStatusType_SaveDatad = '1'

// 结算确认数据归档状态类型
type TThostFtdcSettArchiveStatusType byte

// 未归档数据
const TThostFtdcSettArchiveStatusType_UnArchived = '0'

// 数据归档中
const TThostFtdcSettArchiveStatusType_Archiving = '1'

// 已归档数据
const TThostFtdcSettArchiveStatusType_Archived = '2'

// 归档数据失败
const TThostFtdcSettArchiveStatusType_ArchiveFail = '3'

// CTP交易系统类型类型
type TThostFtdcCTPTypeType byte

// 未知类型
const TThostFtdcCTPTypeType_Unkown = '0'

// 主中心
const TThostFtdcCTPTypeType_MainCenter = '1'

// 备中心
const TThostFtdcCTPTypeType_BackUp = '2'

// 工具代码类型
type TThostFtdcToolIDType [9]byte

// 工具名称类型
type TThostFtdcToolNameType [81]byte

// 平仓处理类型类型
type TThostFtdcCloseDealTypeType byte

// 正常
const TThostFtdcCloseDealTypeType_Normal = '0'

// 投机平仓优先
const TThostFtdcCloseDealTypeType_SpecFirst = '1'

// 货币质押资金可用范围类型
type TThostFtdcMortgageFundUseRangeType byte

// 不能使用
const TThostFtdcMortgageFundUseRangeType_None = '0'

// 用于保证金
const TThostFtdcMortgageFundUseRangeType_Margin = '1'

// 用于手续费、盈亏、保证金
const TThostFtdcMortgageFundUseRangeType_All = '2'

// 特殊产品类型类型
type TThostFtdcSpecProductTypeType byte

// 郑商所套保产品
const TThostFtdcSpecProductTypeType_CzceHedge = '1'

// 货币质押产品
const TThostFtdcSpecProductTypeType_IneForeignCurrency = '2'

// 大连短线开平仓产品
const TThostFtdcSpecProductTypeType_DceOpenClose = '3'

// 货币质押类型类型
type TThostFtdcFundMortgageTypeType byte

// 质押
const TThostFtdcFundMortgageTypeType_Mortgage = '1'

// 解质
const TThostFtdcFundMortgageTypeType_Redemption = '2'

// 投资者账户结算参数代码类型
type TThostFtdcAccountSettlementParamIDType byte

// 基础保证金
const TThostFtdcAccountSettlementParamIDType_BaseMargin = '1'

// 最低权益标准
const TThostFtdcAccountSettlementParamIDType_LowestInterest = '2'

// 币种名称类型
type TThostFtdcCurrencyNameType [31]byte

// 币种符号类型
type TThostFtdcCurrencySignType [4]byte

// 货币质押方向类型
type TThostFtdcFundMortDirectionType byte

// 货币质入
const TThostFtdcFundMortDirectionType_In = '1'

// 货币质出
const TThostFtdcFundMortDirectionType_Out = '2'

// 换汇类别类型
type TThostFtdcBusinessClassType byte

// 盈利
const TThostFtdcBusinessClassType_Profit = '0'

// 亏损
const TThostFtdcBusinessClassType_Loss = '1'

// 其他
const TThostFtdcBusinessClassType_Other = 'Z'

// 换汇数据来源类型
type TThostFtdcSwapSourceTypeType byte

// 手工
const TThostFtdcSwapSourceTypeType_Manual = '0'

// 自动生成
const TThostFtdcSwapSourceTypeType_Automatic = '1'

// 换汇类型类型
type TThostFtdcCurrExDirectionType byte

// 结汇
const TThostFtdcCurrExDirectionType_Settlement = '0'

// 售汇
const TThostFtdcCurrExDirectionType_Sale = '1'

type TThostFtdcCurrencySwapStatusType byte

// 已录入
const TThostFtdcCurrencySwapStatusType_Entry = '1'

// 已审核
const TThostFtdcCurrencySwapStatusType_Approve = '2'

// 已拒绝
const TThostFtdcCurrencySwapStatusType_Refuse = '3'

// 已撤销
const TThostFtdcCurrencySwapStatusType_Revoke = '4'

// 已发送
const TThostFtdcCurrencySwapStatusType_Send = '5'

// 换汇成功
const TThostFtdcCurrencySwapStatusType_Success = '6'

// 换汇失败
const TThostFtdcCurrencySwapStatusType_Failure = '7'

// 凭证号类型
type TThostFtdcCurrExchCertNoType [13]byte

// 批次号类型
type TThostFtdcBatchSerialNoType [21]byte
type TThostFtdcReqFlagType byte

// 未发送
const TThostFtdcReqFlagType_NoSend = '0'

// 发送成功
const TThostFtdcReqFlagType_SendSuccess = '1'

// 发送失败
const TThostFtdcReqFlagType_SendFailed = '2'

// 等待重发
const TThostFtdcReqFlagType_WaitReSend = '3'

// 换汇返回成功标志类型
type TThostFtdcResFlagType byte

// 成功
const TThostFtdcResFlagType_Success = '0'

// 账户余额不足
const TThostFtdcResFlagType_InsuffiCient = '1'

// 交易结果未知
const TThostFtdcResFlagType_UnKnown = '8'

// 换汇页面控制类型
type TThostFtdcPageControlType [2]byte

// 换汇需确认信息类型
type TThostFtdcCurrencySwapMemoType [101]byte

// 修改状态类型
type TThostFtdcExStatusType byte

// 修改前
const TThostFtdcExStatusType_Before = '0'

// 修改后
const TThostFtdcExStatusType_After = '1'

// 开户客户地域类型
type TThostFtdcClientRegionType byte

// 国内客户
const TThostFtdcClientRegionType_Domestic = '1'

// 港澳台客户
const TThostFtdcClientRegionType_GMT = '2'

// 国外客户
const TThostFtdcClientRegionType_Foreign = '3'

// 工作单位类型
type TThostFtdcWorkPlaceType [101]byte

// 经营期限类型
type TThostFtdcBusinessPeriodType [21]byte

// 网址类型
type TThostFtdcWebSiteType [101]byte

// 统一开户证件类型类型
type TThostFtdcUOAIdCardTypeType [3]byte

// 开户模式类型
type TThostFtdcClientModeType [3]byte

// 投资者全称类型
type TThostFtdcInvestorFullNameType [101]byte

// 境外中介机构ID类型
type TThostFtdcUOABrokerIDType [11]byte
type TThostFtdcUOAZipCodeType [11]byte

// 电子邮箱类型
type TThostFtdcUOAEMailType [101]byte

// 城市类型
type TThostFtdcOldCityType [41]byte

// 法人代表证件号码类型
type TThostFtdcCorporateIdentifiedCardNoType [101]byte

// 是否有董事会类型
type TThostFtdcHasBoardType byte

// 没有
const TThostFtdcHasBoardType_No = '0'

// 有
const TThostFtdcHasBoardType_Yes = '1'

// 启动模式类型
type TThostFtdcStartModeType byte

// 正常
const TThostFtdcStartModeType_Normal = '1'

// 应急
const TThostFtdcStartModeType_Emerge = '2'

// 恢复
const TThostFtdcStartModeType_Restore = '3'

// 模型类型类型
type TThostFtdcTemplateTypeType byte

// 全量
const TThostFtdcTemplateTypeType_Full = '1'

// 增量
const TThostFtdcTemplateTypeType_Increment = '2'

// 备份
const TThostFtdcTemplateTypeType_BackUp = '3'

// 登录模式类型
type TThostFtdcLoginModeType byte

// 交易
const TThostFtdcLoginModeType_Trade = '0'

// 转账
const TThostFtdcLoginModeType_Transfer = '1'

// 日历提示类型类型
type TThostFtdcPromptTypeType byte

// 合约上下市
const TThostFtdcPromptTypeType_Instrument = '1'

// 保证金分段生效
const TThostFtdcPromptTypeType_Margin = '2'

// 分户管理资产编码类型
type TThostFtdcLedgerManageIDType [51]byte

// 投资品种类型
type TThostFtdcInvestVarietyType [101]byte

// 账户类别类型
type TThostFtdcBankAccountTypeType [2]byte

// 开户银行类型
type TThostFtdcLedgerManageBankType [101]byte

// 开户营业部类型
type TThostFtdcCffexDepartmentNameType [101]byte

// 营业部代码类型
type TThostFtdcCffexDepartmentCodeType [9]byte

// 是否有托管人类型
type TThostFtdcHasTrusteeType byte

// 有
const TThostFtdcHasTrusteeType_Yes = '1'

// 没有
const TThostFtdcHasTrusteeType_No = '0'

type TThostFtdcCSRCMemo1Type [41]byte

// 代理资产管理业务的期货公司全称类型
type TThostFtdcAssetmgrCFullNameType [101]byte

// 资产管理业务批文号类型
type TThostFtdcAssetmgrApprovalNOType [51]byte

// 资产管理业务负责人姓名类型
type TThostFtdcAssetmgrMgrNameType [401]byte
type TThostFtdcAmTypeType byte

// 银行
const TThostFtdcAmTypeType_Bank = '1'

// 证券公司
const TThostFtdcAmTypeType_Securities = '2'

// 基金公司
const TThostFtdcAmTypeType_Fund = '3'

// 保险公司
const TThostFtdcAmTypeType_Insurance = '4'

// 信托公司
const TThostFtdcAmTypeType_Trust = '5'

// 其他
const TThostFtdcAmTypeType_Other = '9'

type TThostFtdcCSRCAmTypeType [5]byte
type TThostFtdcCSRCFundIOTypeType byte

// 出入金
const TThostFtdcCSRCFundIOTypeType_FundIO = '0'

// 银期换汇
const TThostFtdcCSRCFundIOTypeType_SwapCurrency = '1'

// 结算账户类型类型
type TThostFtdcCusAccountTypeType byte

// 期货结算账户
const TThostFtdcCusAccountTypeType_Futures = '1'

// 纯期货资管业务下的资管结算账户
const TThostFtdcCusAccountTypeType_AssetmgrFuture = '2'

// 综合类资管业务下的期货资管托管账户
const TThostFtdcCusAccountTypeType_AssetmgrTrustee = '3'

// 综合类资管业务下的资金中转账户
const TThostFtdcCusAccountTypeType_AssetmgrTransfer = '4'

type TThostFtdcCSRCNationalType [4]byte

// 二级代理ID类型
type TThostFtdcCSRCSecAgentIDType [11]byte

// 通知语言类型类型
type TThostFtdcLanguageTypeType byte

// 中文
const TThostFtdcLanguageTypeType_Chinese = '1'

// 英文
const TThostFtdcLanguageTypeType_English = '2'

// 投资账户类型
type TThostFtdcAmAccountType [23]byte

// 资产管理客户类型类型
type TThostFtdcAssetmgrClientTypeType byte

// 个人资管客户
const TThostFtdcAssetmgrClientTypeType_Person = '1'

// 单位资管客户
const TThostFtdcAssetmgrClientTypeType_Organ = '2'

// 特殊单位资管客户
const TThostFtdcAssetmgrClientTypeType_SpecialOrgan = '4'

// 投资类型类型
type TThostFtdcAssetmgrTypeType byte

// 期货类
const TThostFtdcAssetmgrTypeType_Futures = '3'

// 综合类
const TThostFtdcAssetmgrTypeType_SpecialOrgan = '4'

// 计量单位类型
type TThostFtdcUOMType [11]byte

// 上期所合约生命周期状态类型
type TThostFtdcSHFEInstLifePhaseType [3]byte
type TThostFtdcSHFEProductClassType [11]byte

// 价格小数位类型
type TThostFtdcPriceDecimalType [2]byte

// 平值期权标志类型
type TThostFtdcInTheMoneyFlagType [2]byte

// 合约比较类型类型
type TThostFtdcCheckInstrTypeType byte

// 合约交易所不存在
const TThostFtdcCheckInstrTypeType_HasExch = '0'

// 合约本系统不存在
const TThostFtdcCheckInstrTypeType_HasATP = '1'

// 合约比较不一致
const TThostFtdcCheckInstrTypeType_HasDiff = '2'

// 交割类型类型
type TThostFtdcDeliveryTypeType byte

// 手工交割
const TThostFtdcDeliveryTypeType_HandDeliv = '1'

// 到期交割
const TThostFtdcDeliveryTypeType_PersonDeliv = '2'

// 大额单边保证金算法类型
type TThostFtdcMaxMarginSideAlgorithmType byte

// 不使用大额单边保证金算法
const TThostFtdcMaxMarginSideAlgorithmType_NO = '0'

// 使用大额单边保证金算法
const TThostFtdcMaxMarginSideAlgorithmType_YES = '1'

type TThostFtdcDAClientTypeType byte

// 自然人
const TThostFtdcDAClientTypeType_Person = '0'

// 法人
const TThostFtdcDAClientTypeType_Company = '1'

// 其他
const TThostFtdcDAClientTypeType_Other = '2'

// 套利合约代码类型
type TThostFtdcCombinInstrIDType [61]byte

// 各腿结算价类型
type TThostFtdcCombinSettlePriceType [61]byte
type TThostFtdcUOAAssetmgrTypeType byte

// 期货类
const TThostFtdcUOAAssetmgrTypeType_Futures = '1'

// 综合类
const TThostFtdcUOAAssetmgrTypeType_SpecialOrgan = '2'

type TThostFtdcDirectionEnType byte

// Buy
const TThostFtdcDirectionEnType_Buy = '0'

// Sell
const TThostFtdcDirectionEnType_Sell = '1'

type TThostFtdcOffsetFlagEnType byte

// Position Opening
const TThostFtdcOffsetFlagEnType_Open = '0'

// Position Close
const TThostFtdcOffsetFlagEnType_Close = '1'

// Forced Liquidation
const TThostFtdcOffsetFlagEnType_ForceClose = '2'

// Close Today
const TThostFtdcOffsetFlagEnType_CloseToday = '3'

// Close Prev.
const TThostFtdcOffsetFlagEnType_CloseYesterday = '4'

// Forced Reduction
const TThostFtdcOffsetFlagEnType_ForceOff = '5'

// Local Forced Liquidation
const TThostFtdcOffsetFlagEnType_LocalForceClose = '6'

type TThostFtdcHedgeFlagEnType byte

// Speculation
const TThostFtdcHedgeFlagEnType_Speculation = '1'

// Arbitrage
const TThostFtdcHedgeFlagEnType_Arbitrage = '2'

// Hedge
const TThostFtdcHedgeFlagEnType_Hedge = '3'

type TThostFtdcFundIOTypeEnType byte

// TFtdcFundIOTypeEnType是一个出入金类型类型
const TThostFtdcFundIOTypeEnType_FundIO = '1'

// Bank-Futures Transfer
const TThostFtdcFundIOTypeEnType_Transfer = '2'

// Bank-Futures FX Exchange
const TThostFtdcFundIOTypeEnType_SwapCurrency = '3'

type TThostFtdcFundTypeEnType byte

// Bank Deposit
const TThostFtdcFundTypeEnType_Deposite = '1'

// Bank Deposit
const TThostFtdcFundTypeEnType_ItemFund = '2'

// Brokerage Adj
const TThostFtdcFundTypeEnType_Company = '3'

// Internal Transfer
const TThostFtdcFundTypeEnType_InnerTransfer = '4'

type TThostFtdcFundDirectionEnType byte

// Deposit
const TThostFtdcFundDirectionEnType_In = '1'

// Withdrawal
const TThostFtdcFundDirectionEnType_Out = '2'

type TThostFtdcFundMortDirectionEnType byte

// Pledge
const TThostFtdcFundMortDirectionEnType_In = '1'

// Redemption
const TThostFtdcFundMortDirectionEnType_Out = '2'

// 换汇业务种类类型
type TThostFtdcSwapBusinessTypeType [3]byte
type TThostFtdcOptionsTypeType byte

// 看涨
const TThostFtdcOptionsTypeType_CallOptions = '1'

// 看跌
const TThostFtdcOptionsTypeType_PutOptions = '2'

// 执行方式类型
type TThostFtdcStrikeModeType byte

// 欧式
const TThostFtdcStrikeModeType_Continental = '0'

// 美式
const TThostFtdcStrikeModeType_American = '1'

// 百慕大
const TThostFtdcStrikeModeType_Bermuda = '2'

// 执行类型类型
type TThostFtdcStrikeTypeType byte

// 自身对冲
const TThostFtdcStrikeTypeType_Hedge = '0'

// 匹配执行
const TThostFtdcStrikeTypeType_Match = '1'

// 中金所期权放弃执行申请类型类型
type TThostFtdcApplyTypeType byte

// 不执行数量
const TThostFtdcApplyTypeType_NotStrikeNum = '4'

// 放弃执行申请数据来源类型
type TThostFtdcGiveUpDataSourceType byte

// 系统生成
const TThostFtdcGiveUpDataSourceType_Gen = '0'

// 手工添加
const TThostFtdcGiveUpDataSourceType_Hand = '1'

// 执行宣告系统编号类型
type TThostFtdcExecOrderSysIDType [21]byte

// 执行结果类型
type TThostFtdcExecResultType byte

// 没有执行
const TThostFtdcExecResultType_NoExec = 'n'

// 已经取消
const TThostFtdcExecResultType_Canceled = 'c'

// 执行成功
const TThostFtdcExecResultType_OK = '0'

// 期权持仓不够
const TThostFtdcExecResultType_NoPosition = '1'

// 资金不够
const TThostFtdcExecResultType_NoDeposit = '2'

// 会员不存在
const TThostFtdcExecResultType_NoParticipant = '3'

// 客户不存在
const TThostFtdcExecResultType_NoClient = '4'

// 合约不存在
const TThostFtdcExecResultType_NoInstrument = '6'

// 没有执行权限
const TThostFtdcExecResultType_NoRight = '7'

// 不合理的数量
const TThostFtdcExecResultType_InvalidVolume = '8'

// 没有足够的历史成交
const TThostFtdcExecResultType_NoEnoughHistoryTrade = '9'

// 未知
const TThostFtdcExecResultType_Unknown = 'a'

// 执行时间类型
type TThostFtdcStrikeTimeType [13]byte
type TThostFtdcCombinationTypeType byte

// 期货组合
const TThostFtdcCombinationTypeType_Future = '0'

// 垂直价差BUL
const TThostFtdcCombinationTypeType_BUL = '1'

// 垂直价差BER
const TThostFtdcCombinationTypeType_BER = '2'

// 跨式组合
const TThostFtdcCombinationTypeType_STD = '3'

// 宽跨式组合
const TThostFtdcCombinationTypeType_STG = '4'

// 备兑组合
const TThostFtdcCombinationTypeType_PRT = '5'

// 时间价差组合
const TThostFtdcCombinationTypeType_CLD = '6'

// 期权权利金价格类型类型
type TThostFtdcOptionRoyaltyPriceTypeType byte

// 昨结算价
const TThostFtdcOptionRoyaltyPriceTypeType_PreSettlementPrice = '1'

// 开仓价
const TThostFtdcOptionRoyaltyPriceTypeType_OpenPrice = '4'

// 最新价与昨结算价较大值
const TThostFtdcOptionRoyaltyPriceTypeType_MaxPreSettlementPrice = '5'

// 权益算法类型
type TThostFtdcBalanceAlgorithmType byte

// 不计算期权市值盈亏
const TThostFtdcBalanceAlgorithmType_Default = '1'

// 计算期权市值亏损
const TThostFtdcBalanceAlgorithmType_IncludeOptValLost = '2'

type TThostFtdcActionTypeType byte

// 执行
const TThostFtdcActionTypeType_Exec = '1'

// 放弃
const TThostFtdcActionTypeType_Abandon = '2'

// 询价状态类型
type TThostFtdcForQuoteStatusType byte

// 已经提交
const TThostFtdcForQuoteStatusType_Submitted = 'a'

// 已经接受
const TThostFtdcForQuoteStatusType_Accepted = 'b'

// 已经被拒绝
const TThostFtdcForQuoteStatusType_Rejected = 'c'

// 取值方式类型
type TThostFtdcValueMethodType byte

// 按绝对值
const TThostFtdcValueMethodType_Absolute = '0'

// 按比率
const TThostFtdcValueMethodType_Ratio = '1'

// 期权行权后是否保留期货头寸的标记类型
type TThostFtdcExecOrderPositionFlagType byte

// 保留
const TThostFtdcExecOrderPositionFlagType_Reserve = '0'

// 不保留
const TThostFtdcExecOrderPositionFlagType_UnReserve = '1'

// 期权行权后生成的头寸是否自动平仓类型
type TThostFtdcExecOrderCloseFlagType byte

// 自动平仓
const TThostFtdcExecOrderCloseFlagType_AutoClose = '0'

// 免于自动平仓
const TThostFtdcExecOrderCloseFlagType_NotToClose = '1'

type TThostFtdcProductTypeType byte

// 期货
const TThostFtdcProductTypeType_Futures = '1'

// 期权
const TThostFtdcProductTypeType_Options = '2'

// 郑商所结算文件名类型
type TThostFtdcCZCEUploadFileNameType byte

// TFtdcCZCEUploadFileNameType是一个郑商所结算文件名类型
const TThostFtdcCZCEUploadFileNameType_O = 'O'

// TFtdcCZCEUploadFileNameType是一个郑商所结算文件名类型
const TThostFtdcCZCEUploadFileNameType_T = 'T'

// TFtdcCZCEUploadFileNameType是一个郑商所结算文件名类型
const TThostFtdcCZCEUploadFileNameType_P = 'P'

// TFtdcCZCEUploadFileNameType是一个郑商所结算文件名类型
const TThostFtdcCZCEUploadFileNameType_N = 'N'

// TFtdcCZCEUploadFileNameType是一个郑商所结算文件名类型
const TThostFtdcCZCEUploadFileNameType_L = 'L'

// TFtdcCZCEUploadFileNameType是一个郑商所结算文件名类型
const TThostFtdcCZCEUploadFileNameType_F = 'F'

// TFtdcCZCEUploadFileNameType是一个郑商所结算文件名类型
const TThostFtdcCZCEUploadFileNameType_C = 'C'

// TFtdcCZCEUploadFileNameType是一个郑商所结算文件名类型
const TThostFtdcCZCEUploadFileNameType_M = 'M'

// 大商所结算文件名类型
type TThostFtdcDCEUploadFileNameType byte

// TFtdcDCEUploadFileNameType是一个大商所结算文件名类型
const TThostFtdcDCEUploadFileNameType_O = 'O'

// TFtdcDCEUploadFileNameType是一个大商所结算文件名类型
const TThostFtdcDCEUploadFileNameType_T = 'T'

// TFtdcDCEUploadFileNameType是一个大商所结算文件名类型
const TThostFtdcDCEUploadFileNameType_P = 'P'

// TFtdcDCEUploadFileNameType是一个大商所结算文件名类型
const TThostFtdcDCEUploadFileNameType_F = 'F'

// TFtdcDCEUploadFileNameType是一个大商所结算文件名类型
const TThostFtdcDCEUploadFileNameType_C = 'C'

// TFtdcDCEUploadFileNameType是一个大商所结算文件名类型
const TThostFtdcDCEUploadFileNameType_D = 'D'

// TFtdcDCEUploadFileNameType是一个大商所结算文件名类型
const TThostFtdcDCEUploadFileNameType_M = 'M'

// TFtdcDCEUploadFileNameType是一个大商所结算文件名类型
const TThostFtdcDCEUploadFileNameType_S = 'S'

// 上期所结算文件名类型
type TThostFtdcSHFEUploadFileNameType byte

// TFtdcSHFEUploadFileNameType是一个上期所结算文件名类型
const TThostFtdcSHFEUploadFileNameType_O = 'O'

// TFtdcSHFEUploadFileNameType是一个上期所结算文件名类型
const TThostFtdcSHFEUploadFileNameType_T = 'T'

// TFtdcSHFEUploadFileNameType是一个上期所结算文件名类型
const TThostFtdcSHFEUploadFileNameType_P = 'P'

// TFtdcSHFEUploadFileNameType是一个上期所结算文件名类型
const TThostFtdcSHFEUploadFileNameType_F = 'F'

// 中金所结算文件名类型
type TThostFtdcCFFEXUploadFileNameType byte

// TFtdcCFFEXUploadFileNameType是一个中金所结算文件名类型
const TThostFtdcCFFEXUploadFileNameType_T = 'T'

// TFtdcCFFEXUploadFileNameType是一个中金所结算文件名类型
const TThostFtdcCFFEXUploadFileNameType_P = 'P'

// TFtdcCFFEXUploadFileNameType是一个中金所结算文件名类型
const TThostFtdcCFFEXUploadFileNameType_F = 'F'

// TFtdcCFFEXUploadFileNameType是一个中金所结算文件名类型
const TThostFtdcCFFEXUploadFileNameType_S = 'S'

// 组合指令方向类型
type TThostFtdcCombDirectionType byte

// 申请组合
const TThostFtdcCombDirectionType_Comb = '0'

// 申请拆分
const TThostFtdcCombDirectionType_UnComb = '1'

// 行权偏移类型类型
type TThostFtdcStrikeOffsetTypeType byte

// 实值额
const TThostFtdcStrikeOffsetTypeType_RealValue = '1'

// 盈利额
const TThostFtdcStrikeOffsetTypeType_ProfitValue = '2'

// 实值比例
const TThostFtdcStrikeOffsetTypeType_RealRatio = '3'

// 盈利比例
const TThostFtdcStrikeOffsetTypeType_ProfitRatio = '4'

// 预约开户状态类型
type TThostFtdcReserveOpenAccStasType byte

// 等待处理中
const TThostFtdcReserveOpenAccStasType_Processing = '0'

// 已撤销
const TThostFtdcReserveOpenAccStasType_Cancelled = '1'

// 已开户
const TThostFtdcReserveOpenAccStasType_Opened = '2'

// 无效请求
const TThostFtdcReserveOpenAccStasType_Invalid = '3'

// 登录备注类型
type TThostFtdcLoginRemarkType [36]byte

// 投资单元代码类型
type TThostFtdcInvestUnitIDType [17]byte

// 公告类型类型
type TThostFtdcNewsTypeType [3]byte

// 紧急程度类型
type TThostFtdcNewsUrgencyType byte

// 消息摘要类型
type TThostFtdcAbstractType [81]byte

// 消息来源类型
type TThostFtdcComeFromType [21]byte

// WEB地址类型
type TThostFtdcURLLinkType [201]byte

// 长个人姓名类型
type TThostFtdcLongIndividualNameType [161]byte

// 长换汇银行账户名类型
type TThostFtdcLongFBEBankAccountNameType [161]byte

// 日期时间类型
type TThostFtdcDateTimeType [17]byte

// 弱密码来源类型
type TThostFtdcWeakPasswordSourceType byte

// 弱密码库
const TThostFtdcWeakPasswordSourceType_Lib = '1'

// 手工录入
const TThostFtdcWeakPasswordSourceType_Manual = '2'

// 随机串类型
type TThostFtdcRandomStringType [17]byte

// 期权行权的头寸是否自对冲类型
type TThostFtdcOptSelfCloseFlagType byte

// 自对冲期权仓位
const TThostFtdcOptSelfCloseFlagType_CloseSelfOptionPosition = '1'

// 保留期权仓位
const TThostFtdcOptSelfCloseFlagType_ReserveOptionPosition = '2'

// 自对冲卖方履约后的期货仓位
const TThostFtdcOptSelfCloseFlagType_SellCloseSelfFuturePosition = '3'

type TThostFtdcBizTypeType byte

// 期货
const TThostFtdcBizTypeType_Future = '1'

// 证券
const TThostFtdcBizTypeType_Stock = '2'
