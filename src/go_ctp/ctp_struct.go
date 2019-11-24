package go_ctp

// 信息分发
type tCThostFtdcDisseminationField struct {
	// 序列系列号
	SequenceSeries tTThostFtdcSequenceSeriesType
	// 序列号
	SequenceNo tTThostFtdcSequenceNoType
}

// 用户登录请求
type tCThostFtdcReqUserLoginField struct {
	// 交易日
	TradingDay tTThostFtdcDateType
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 用户代码
	UserID tTThostFtdcUserIDType
	// 密码
	Password tTThostFtdcPasswordType
	// 用户端产品信息
	UserProductInfo tTThostFtdcProductInfoType
	// 接口端产品信息
	InterfaceProductInfo tTThostFtdcProductInfoType
	// 协议信息
	ProtocolInfo tTThostFtdcProtocolInfoType
	// Mac地址
	MacAddress tTThostFtdcMacAddressType
	// 动态密码
	OneTimePassword tTThostFtdcPasswordType
	// 终端IP地址
	ClientIPAddress tTThostFtdcIPAddressType
	// 登录备注
	LoginRemark tTThostFtdcLoginRemarkType
	// 终端IP端口
	ClientIPPort tTThostFtdcIPPortType
}

// 用户登录应答
type tCThostFtdcRspUserLoginField struct {
	// 交易日
	TradingDay tTThostFtdcDateType
	// 登录成功时间
	LoginTime tTThostFtdcTimeType
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 用户代码
	UserID tTThostFtdcUserIDType
	// 交易系统名称
	SystemName tTThostFtdcSystemNameType
	// 前置编号
	FrontID tTThostFtdcFrontIDType
	// 会话编号
	SessionID tTThostFtdcSessionIDType
	// 最大报单引用
	MaxOrderRef tTThostFtdcOrderRefType
	// 上期所时间
	SHFETime tTThostFtdcTimeType
	// 大商所时间
	DCETime tTThostFtdcTimeType
	// 郑商所时间
	CZCETime tTThostFtdcTimeType
	// 中金所时间
	FFEXTime tTThostFtdcTimeType
	// 能源中心时间
	INETime tTThostFtdcTimeType
}

// 用户登出请求
type tCThostFtdcUserLogoutField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 用户代码
	UserID tTThostFtdcUserIDType
}

// 强制交易员退出
type tCThostFtdcForceUserLogoutField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 用户代码
	UserID tTThostFtdcUserIDType
}

// 客户端认证请求
type tCThostFtdcReqAuthenticateField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 用户代码
	UserID tTThostFtdcUserIDType
	// 用户端产品信息
	UserProductInfo tTThostFtdcProductInfoType
	// 认证码
	AuthCode tTThostFtdcAuthCodeType
	// App代码
	AppID tTThostFtdcAppIDType
}

// 客户端认证响应
type tCThostFtdcRspAuthenticateField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 用户代码
	UserID tTThostFtdcUserIDType
	// 用户端产品信息
	UserProductInfo tTThostFtdcProductInfoType
	// App代码
	AppID tTThostFtdcAppIDType
	// App类型
	AppType tTThostFtdcAppTypeType
}

// 客户端认证信息
type tCThostFtdcAuthenticationInfoField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 用户代码
	UserID tTThostFtdcUserIDType
	// 用户端产品信息
	UserProductInfo tTThostFtdcProductInfoType
	// 认证信息
	AuthInfo tTThostFtdcAuthInfoType
	// 是否为认证结果
	IsResult tTThostFtdcBoolType
	// App代码
	AppID tTThostFtdcAppIDType
	// App类型
	AppType tTThostFtdcAppTypeType
}

// 用户登录应答2
type tCThostFtdcRspUserLogin2Field struct {
	// 交易日
	TradingDay tTThostFtdcDateType
	// 登录成功时间
	LoginTime tTThostFtdcTimeType
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 用户代码
	UserID tTThostFtdcUserIDType
	// 交易系统名称
	SystemName tTThostFtdcSystemNameType
	// 前置编号
	FrontID tTThostFtdcFrontIDType
	// 会话编号
	SessionID tTThostFtdcSessionIDType
	// 最大报单引用
	MaxOrderRef tTThostFtdcOrderRefType
	// 上期所时间
	SHFETime tTThostFtdcTimeType
	// 大商所时间
	DCETime tTThostFtdcTimeType
	// 郑商所时间
	CZCETime tTThostFtdcTimeType
	// 中金所时间
	FFEXTime tTThostFtdcTimeType
	// 能源中心时间
	INETime tTThostFtdcTimeType
	// 随机串
	RandomString tTThostFtdcRandomStringType
}

// 银期转帐报文头
type tCThostFtdcTransferHeaderField struct {
	// 版本号，常量，1.0
	Version tTThostFtdcVersionType
	// 交易代码，必填
	TradeCode tTThostFtdcTradeCodeType
	// 交易日期，必填，格式：yyyymmdd
	TradeDate tTThostFtdcTradeDateType
	// 交易时间，必填，格式：hhmmss
	TradeTime tTThostFtdcTradeTimeType
	// 发起方流水号，N/A
	TradeSerial tTThostFtdcTradeSerialType
	// 期货公司代码，必填
	FutureID tTThostFtdcFutureIDType
	// 银行代码，根据查询银行得到，必填
	BankID tTThostFtdcBankIDType
	// 银行分中心代码，根据查询银行得到，必填
	BankBrchID tTThostFtdcBankBrchIDType
	// 操作员，N/A
	OperNo tTThostFtdcOperNoType
	// 交易设备类型，N/A
	DeviceID tTThostFtdcDeviceIDType
	// 记录数，N/A
	RecordNum tTThostFtdcRecordNumType
	// 会话编号，N/A
	SessionID tTThostFtdcSessionIDType
	// 请求编号，N/A
	RequestID tTThostFtdcRequestIDType
}

// 银行资金转期货请求，TradeCode=202001
type tCThostFtdcTransferBankToFutureReqField struct {
	// 期货资金账户
	FutureAccount tTThostFtdcAccountIDType
	// 密码标志
	FuturePwdFlag tTThostFtdcFuturePwdFlagType
	// 密码
	FutureAccPwd tTThostFtdcFutureAccPwdType
	// 转账金额
	TradeAmt tTThostFtdcMoneyType
	// 客户手续费
	CustFee tTThostFtdcMoneyType
	// 币种：RMB-人民币 USD-美圆 HKD-港元
	CurrencyCode tTThostFtdcCurrencyCodeType
}

// 银行资金转期货请求响应
type tCThostFtdcTransferBankToFutureRspField struct {
	// 响应代码
	RetCode tTThostFtdcRetCodeType
	// 响应信息
	RetInfo tTThostFtdcRetInfoType
	// 资金账户
	FutureAccount tTThostFtdcAccountIDType
	// 转帐金额
	TradeAmt tTThostFtdcMoneyType
	// 应收客户手续费
	CustFee tTThostFtdcMoneyType
	// 币种
	CurrencyCode tTThostFtdcCurrencyCodeType
}

// 期货资金转银行请求，TradeCode=202002
type tCThostFtdcTransferFutureToBankReqField struct {
	// 期货资金账户
	FutureAccount tTThostFtdcAccountIDType
	// 密码标志
	FuturePwdFlag tTThostFtdcFuturePwdFlagType
	// 密码
	FutureAccPwd tTThostFtdcFutureAccPwdType
	// 转账金额
	TradeAmt tTThostFtdcMoneyType
	// 客户手续费
	CustFee tTThostFtdcMoneyType
	// 币种：RMB-人民币 USD-美圆 HKD-港元
	CurrencyCode tTThostFtdcCurrencyCodeType
}

// 期货资金转银行请求响应
type tCThostFtdcTransferFutureToBankRspField struct {
	// 响应代码
	RetCode tTThostFtdcRetCodeType
	// 响应信息
	RetInfo tTThostFtdcRetInfoType
	// 资金账户
	FutureAccount tTThostFtdcAccountIDType
	// 转帐金额
	TradeAmt tTThostFtdcMoneyType
	// 应收客户手续费
	CustFee tTThostFtdcMoneyType
	// 币种
	CurrencyCode tTThostFtdcCurrencyCodeType
}

// 查询银行资金请求，TradeCode=204002
type tCThostFtdcTransferQryBankReqField struct {
	// 期货资金账户
	FutureAccount tTThostFtdcAccountIDType
	// 密码标志
	FuturePwdFlag tTThostFtdcFuturePwdFlagType
	// 密码
	FutureAccPwd tTThostFtdcFutureAccPwdType
	// 币种：RMB-人民币 USD-美圆 HKD-港元
	CurrencyCode tTThostFtdcCurrencyCodeType
}

// 查询银行资金请求响应
type tCThostFtdcTransferQryBankRspField struct {
	// 响应代码
	RetCode tTThostFtdcRetCodeType
	// 响应信息
	RetInfo tTThostFtdcRetInfoType
	// 资金账户
	FutureAccount tTThostFtdcAccountIDType
	// 银行余额
	TradeAmt tTThostFtdcMoneyType
	// 银行可用余额
	UseAmt tTThostFtdcMoneyType
	// 银行可取余额
	FetchAmt tTThostFtdcMoneyType
	// 币种
	CurrencyCode tTThostFtdcCurrencyCodeType
}

// 查询银行交易明细请求，TradeCode=204999
type tCThostFtdcTransferQryDetailReqField struct {
	// 期货资金账户
	FutureAccount tTThostFtdcAccountIDType
}

// 查询银行交易明细请求响应
type tCThostFtdcTransferQryDetailRspField struct {
	// 交易日期
	TradeDate tTThostFtdcDateType
	// 交易时间
	TradeTime tTThostFtdcTradeTimeType
	// 交易代码
	TradeCode tTThostFtdcTradeCodeType
	// 期货流水号
	FutureSerial tTThostFtdcTradeSerialNoType
	// 期货公司代码
	FutureID tTThostFtdcFutureIDType
	// 资金帐号
	FutureAccount tTThostFtdcFutureAccountType
	// 银行流水号
	BankSerial tTThostFtdcTradeSerialNoType
	// 银行代码
	BankID tTThostFtdcBankIDType
	// 银行分中心代码
	BankBrchID tTThostFtdcBankBrchIDType
	// 银行账号
	BankAccount tTThostFtdcBankAccountType
	// 证件号码
	CertCode tTThostFtdcCertCodeType
	// 货币代码
	CurrencyCode tTThostFtdcCurrencyCodeType
	// 发生金额
	TxAmount tTThostFtdcMoneyType
	// 有效标志
	Flag tTThostFtdcTransferValidFlagType
}

// 响应信息
type tCThostFtdcRspInfoField struct {
	// 错误代码
	ErrorID tTThostFtdcErrorIDType
	// 错误信息
	ErrorMsg tTThostFtdcErrorMsgType
}

// 交易所
type tCThostFtdcExchangeField struct {
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
	// 交易所名称
	ExchangeName tTThostFtdcExchangeNameType
	// 交易所属性
	ExchangeProperty tTThostFtdcExchangePropertyType
}

// 产品
type tCThostFtdcProductField struct {
	// 产品代码
	ProductID tTThostFtdcInstrumentIDType
	// 产品名称
	ProductName tTThostFtdcProductNameType
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
	// 产品类型
	ProductClass tTThostFtdcProductClassType
	// 合约数量乘数
	VolumeMultiple tTThostFtdcVolumeMultipleType
	// 最小变动价位
	PriceTick tTThostFtdcPriceType
	// 市价单最大下单量
	MaxMarketOrderVolume tTThostFtdcVolumeType
	// 市价单最小下单量
	MinMarketOrderVolume tTThostFtdcVolumeType
	// 限价单最大下单量
	MaxLimitOrderVolume tTThostFtdcVolumeType
	// 限价单最小下单量
	MinLimitOrderVolume tTThostFtdcVolumeType
	// 持仓类型
	PositionType tTThostFtdcPositionTypeType
	// 持仓日期类型
	PositionDateType tTThostFtdcPositionDateTypeType
	// 平仓处理类型
	CloseDealType tTThostFtdcCloseDealTypeType
	// 交易币种类型
	TradeCurrencyID tTThostFtdcCurrencyIDType
	// 质押资金可用范围
	MortgageFundUseRange tTThostFtdcMortgageFundUseRangeType
	// 交易所产品代码
	ExchangeProductID tTThostFtdcInstrumentIDType
	// 合约基础商品乘数
	UnderlyingMultiple tTThostFtdcUnderlyingMultipleType
}

// 合约
type tCThostFtdcInstrumentField struct {
	// 合约代码
	InstrumentID tTThostFtdcInstrumentIDType
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
	// 合约名称
	InstrumentName tTThostFtdcInstrumentNameType
	// 合约在交易所的代码
	ExchangeInstID tTThostFtdcExchangeInstIDType
	// 产品代码
	ProductID tTThostFtdcInstrumentIDType
	// 产品类型
	ProductClass tTThostFtdcProductClassType
	// 交割年份
	DeliveryYear tTThostFtdcYearType
	// 交割月
	DeliveryMonth tTThostFtdcMonthType
	// 市价单最大下单量
	MaxMarketOrderVolume tTThostFtdcVolumeType
	// 市价单最小下单量
	MinMarketOrderVolume tTThostFtdcVolumeType
	// 限价单最大下单量
	MaxLimitOrderVolume tTThostFtdcVolumeType
	// 限价单最小下单量
	MinLimitOrderVolume tTThostFtdcVolumeType
	// 合约数量乘数
	VolumeMultiple tTThostFtdcVolumeMultipleType
	// 最小变动价位
	PriceTick tTThostFtdcPriceType
	// 创建日
	CreateDate tTThostFtdcDateType
	// 上市日
	OpenDate tTThostFtdcDateType
	// 到期日
	ExpireDate tTThostFtdcDateType
	// 开始交割日
	StartDelivDate tTThostFtdcDateType
	// 结束交割日
	EndDelivDate tTThostFtdcDateType
	// 合约生命周期状态
	InstLifePhase tTThostFtdcInstLifePhaseType
	// 当前是否交易
	IsTrading tTThostFtdcBoolType
	// 持仓类型
	PositionType tTThostFtdcPositionTypeType
	// 持仓日期类型
	PositionDateType tTThostFtdcPositionDateTypeType
	// 多头保证金率
	LongMarginRatio tTThostFtdcRatioType
	// 空头保证金率
	ShortMarginRatio tTThostFtdcRatioType
	// 是否使用大额单边保证金算法
	MaxMarginSideAlgorithm tTThostFtdcMaxMarginSideAlgorithmType
	// 基础商品代码
	UnderlyingInstrID tTThostFtdcInstrumentIDType
	// 执行价
	StrikePrice tTThostFtdcPriceType
	// 期权类型
	OptionsType tTThostFtdcOptionsTypeType
	// 合约基础商品乘数
	UnderlyingMultiple tTThostFtdcUnderlyingMultipleType
	// 组合类型
	CombinationType tTThostFtdcCombinationTypeType
}

// 经纪公司
type tCThostFtdcBrokerField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 经纪公司简称
	BrokerAbbr tTThostFtdcBrokerAbbrType
	// 经纪公司名称
	BrokerName tTThostFtdcBrokerNameType
	// 是否活跃
	IsActive tTThostFtdcBoolType
}

// 交易所交易员
type tCThostFtdcTraderField struct {
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
	// 交易所交易员代码
	TraderID tTThostFtdcTraderIDType
	// 会员代码
	ParticipantID tTThostFtdcParticipantIDType
	// 密码
	Password tTThostFtdcPasswordType
	// 安装数量
	InstallCount tTThostFtdcInstallCountType
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
}

// 投资者
type tCThostFtdcInvestorField struct {
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者分组代码
	InvestorGroupID tTThostFtdcInvestorIDType
	// 投资者名称
	InvestorName tTThostFtdcPartyNameType
	// 证件类型
	IdentifiedCardType tTThostFtdcIdCardTypeType
	// 证件号码
	IdentifiedCardNo tTThostFtdcIdentifiedCardNoType
	// 是否活跃
	IsActive tTThostFtdcBoolType
	// 联系电话
	Telephone tTThostFtdcTelephoneType
	// 通讯地址
	Address tTThostFtdcAddressType
	// 开户日期
	OpenDate tTThostFtdcDateType
	// 手机
	Mobile tTThostFtdcMobileType
	// 手续费率模板代码
	CommModelID tTThostFtdcInvestorIDType
	// 保证金率模板代码
	MarginModelID tTThostFtdcInvestorIDType
}

// 交易编码
type tCThostFtdcTradingCodeField struct {
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
	// 客户代码
	ClientID tTThostFtdcClientIDType
	// 是否活跃
	IsActive tTThostFtdcBoolType
	// 交易编码类型
	ClientIDType tTThostFtdcClientIDTypeType
	// 营业部编号
	BranchID tTThostFtdcBranchIDType
	// 业务类型
	BizType tTThostFtdcBizTypeType
	// 投资单元代码
	InvestUnitID tTThostFtdcInvestUnitIDType
}

// 会员编码和经纪公司编码对照表
type tCThostFtdcPartBrokerField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
	// 会员代码
	ParticipantID tTThostFtdcParticipantIDType
	// 是否活跃
	IsActive tTThostFtdcBoolType
}

// 管理用户
type tCThostFtdcSuperUserField struct {
	// 用户代码
	UserID tTThostFtdcUserIDType
	// 用户名称
	UserName tTThostFtdcUserNameType
	// 密码
	Password tTThostFtdcPasswordType
	// 是否活跃
	IsActive tTThostFtdcBoolType
}

// 管理用户功能权限
type tCThostFtdcSuperUserFunctionField struct {
	// 用户代码
	UserID tTThostFtdcUserIDType
	// 功能代码
	FunctionCode tTThostFtdcFunctionCodeType
}

// 投资者组
type tCThostFtdcInvestorGroupField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者分组代码
	InvestorGroupID tTThostFtdcInvestorIDType
	// 投资者分组名称
	InvestorGroupName tTThostFtdcInvestorGroupNameType
}

// 资金账户
type tCThostFtdcTradingAccountField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者帐号
	AccountID tTThostFtdcAccountIDType
	// 上次质押金额
	PreMortgage tTThostFtdcMoneyType
	// 上次信用额度
	PreCredit tTThostFtdcMoneyType
	// 上次存款额
	PreDeposit tTThostFtdcMoneyType
	// 上次结算准备金
	PreBalance tTThostFtdcMoneyType
	// 上次占用的保证金
	PreMargin tTThostFtdcMoneyType
	// 利息基数
	InterestBase tTThostFtdcMoneyType
	// 利息收入
	Interest tTThostFtdcMoneyType
	// 入金金额
	Deposit tTThostFtdcMoneyType
	// 出金金额
	Withdraw tTThostFtdcMoneyType
	// 冻结的保证金
	FrozenMargin tTThostFtdcMoneyType
	// 冻结的资金
	FrozenCash tTThostFtdcMoneyType
	// 冻结的手续费
	FrozenCommission tTThostFtdcMoneyType
	// 当前保证金总额
	CurrMargin tTThostFtdcMoneyType
	// 资金差额
	CashIn tTThostFtdcMoneyType
	// 手续费
	Commission tTThostFtdcMoneyType
	// 平仓盈亏
	CloseProfit tTThostFtdcMoneyType
	// 持仓盈亏
	PositionProfit tTThostFtdcMoneyType
	// 期货结算准备金
	Balance tTThostFtdcMoneyType
	// 可用资金
	Available tTThostFtdcMoneyType
	// 可取资金
	WithdrawQuota tTThostFtdcMoneyType
	// 基本准备金
	Reserve tTThostFtdcMoneyType
	// 交易日
	TradingDay tTThostFtdcDateType
	// 结算编号
	SettlementID tTThostFtdcSettlementIDType
	// 信用额度
	Credit tTThostFtdcMoneyType
	// 质押金额
	Mortgage tTThostFtdcMoneyType
	// 交易所保证金
	ExchangeMargin tTThostFtdcMoneyType
	// 投资者交割保证金
	DeliveryMargin tTThostFtdcMoneyType
	// 交易所交割保证金
	ExchangeDeliveryMargin tTThostFtdcMoneyType
	// 保底期货结算准备金
	ReserveBalance tTThostFtdcMoneyType
	// 币种代码
	CurrencyID tTThostFtdcCurrencyIDType
	// 上次货币质入金额
	PreFundMortgageIn tTThostFtdcMoneyType
	// 上次货币质出金额
	PreFundMortgageOut tTThostFtdcMoneyType
	// 货币质入金额
	FundMortgageIn tTThostFtdcMoneyType
	// 货币质出金额
	FundMortgageOut tTThostFtdcMoneyType
	// 货币质押余额
	FundMortgageAvailable tTThostFtdcMoneyType
	// 可质押货币金额
	MortgageableFund tTThostFtdcMoneyType
	// 特殊产品占用保证金
	SpecProductMargin tTThostFtdcMoneyType
	// 特殊产品冻结保证金
	SpecProductFrozenMargin tTThostFtdcMoneyType
	// 特殊产品手续费
	SpecProductCommission tTThostFtdcMoneyType
	// 特殊产品冻结手续费
	SpecProductFrozenCommission tTThostFtdcMoneyType
	// 特殊产品持仓盈亏
	SpecProductPositionProfit tTThostFtdcMoneyType
	// 特殊产品平仓盈亏
	SpecProductCloseProfit tTThostFtdcMoneyType
	// 根据持仓盈亏算法计算的特殊产品持仓盈亏
	SpecProductPositionProfitByAlg tTThostFtdcMoneyType
	// 特殊产品交易所保证金
	SpecProductExchangeMargin tTThostFtdcMoneyType
	// 业务类型
	BizType tTThostFtdcBizTypeType
	// 延时换汇冻结金额
	FrozenSwap tTThostFtdcMoneyType
	// 剩余换汇额度
	RemainSwap tTThostFtdcMoneyType
}

// 投资者持仓
type tCThostFtdcInvestorPositionField struct {
	// 合约代码
	InstrumentID tTThostFtdcInstrumentIDType
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 持仓多空方向
	PosiDirection tTThostFtdcPosiDirectionType
	// 投机套保标志
	HedgeFlag tTThostFtdcHedgeFlagType
	// 持仓日期
	PositionDate tTThostFtdcPositionDateType
	// 上日持仓
	YdPosition tTThostFtdcVolumeType
	// 今日持仓
	Position tTThostFtdcVolumeType
	// 多头冻结
	LongFrozen tTThostFtdcVolumeType
	// 空头冻结
	ShortFrozen tTThostFtdcVolumeType
	// 开仓冻结金额
	LongFrozenAmount tTThostFtdcMoneyType
	// 开仓冻结金额
	ShortFrozenAmount tTThostFtdcMoneyType
	// 开仓量
	OpenVolume tTThostFtdcVolumeType
	// 平仓量
	CloseVolume tTThostFtdcVolumeType
	// 开仓金额
	OpenAmount tTThostFtdcMoneyType
	// 平仓金额
	CloseAmount tTThostFtdcMoneyType
	// 持仓成本
	PositionCost tTThostFtdcMoneyType
	// 上次占用的保证金
	PreMargin tTThostFtdcMoneyType
	// 占用的保证金
	UseMargin tTThostFtdcMoneyType
	// 冻结的保证金
	FrozenMargin tTThostFtdcMoneyType
	// 冻结的资金
	FrozenCash tTThostFtdcMoneyType
	// 冻结的手续费
	FrozenCommission tTThostFtdcMoneyType
	// 资金差额
	CashIn tTThostFtdcMoneyType
	// 手续费
	Commission tTThostFtdcMoneyType
	// 平仓盈亏
	CloseProfit tTThostFtdcMoneyType
	// 持仓盈亏
	PositionProfit tTThostFtdcMoneyType
	// 上次结算价
	PreSettlementPrice tTThostFtdcPriceType
	// 本次结算价
	SettlementPrice tTThostFtdcPriceType
	// 交易日
	TradingDay tTThostFtdcDateType
	// 结算编号
	SettlementID tTThostFtdcSettlementIDType
	// 开仓成本
	OpenCost tTThostFtdcMoneyType
	// 交易所保证金
	ExchangeMargin tTThostFtdcMoneyType
	// 组合成交形成的持仓
	CombPosition tTThostFtdcVolumeType
	// 组合多头冻结
	CombLongFrozen tTThostFtdcVolumeType
	// 组合空头冻结
	CombShortFrozen tTThostFtdcVolumeType
	// 逐日盯市平仓盈亏
	CloseProfitByDate tTThostFtdcMoneyType
	// 逐笔对冲平仓盈亏
	CloseProfitByTrade tTThostFtdcMoneyType
	// 今日持仓
	TodayPosition tTThostFtdcVolumeType
	// 保证金率
	MarginRateByMoney tTThostFtdcRatioType
	// 保证金率(按手数)
	MarginRateByVolume tTThostFtdcRatioType
	// 执行冻结
	StrikeFrozen tTThostFtdcVolumeType
	// 执行冻结金额
	StrikeFrozenAmount tTThostFtdcMoneyType
	// 放弃执行冻结
	AbandonFrozen tTThostFtdcVolumeType
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
	// 执行冻结的昨仓
	YdStrikeFrozen tTThostFtdcVolumeType
	// 投资单元代码
	InvestUnitID tTThostFtdcInvestUnitIDType
	// 大商所持仓成本差值，只有大商所使用
	PositionCostOffset tTThostFtdcMoneyType
}

// 合约保证金率
type tCThostFtdcInstrumentMarginRateField struct {
	// 合约代码
	InstrumentID tTThostFtdcInstrumentIDType
	// 投资者范围
	InvestorRange tTThostFtdcInvestorRangeType
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 投机套保标志
	HedgeFlag tTThostFtdcHedgeFlagType
	// 多头保证金率
	LongMarginRatioByMoney tTThostFtdcRatioType
	// 多头保证金费
	LongMarginRatioByVolume tTThostFtdcMoneyType
	// 空头保证金率
	ShortMarginRatioByMoney tTThostFtdcRatioType
	// 空头保证金费
	ShortMarginRatioByVolume tTThostFtdcMoneyType
	// 是否相对交易所收取
	IsRelative tTThostFtdcBoolType
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
	// 投资单元代码
	InvestUnitID tTThostFtdcInvestUnitIDType
}

// 合约手续费率
type tCThostFtdcInstrumentCommissionRateField struct {
	// 合约代码
	InstrumentID tTThostFtdcInstrumentIDType
	// 投资者范围
	InvestorRange tTThostFtdcInvestorRangeType
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 开仓手续费率
	OpenRatioByMoney tTThostFtdcRatioType
	// 开仓手续费
	OpenRatioByVolume tTThostFtdcRatioType
	// 平仓手续费率
	CloseRatioByMoney tTThostFtdcRatioType
	// 平仓手续费
	CloseRatioByVolume tTThostFtdcRatioType
	// 平今手续费率
	CloseTodayRatioByMoney tTThostFtdcRatioType
	// 平今手续费
	CloseTodayRatioByVolume tTThostFtdcRatioType
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
	// 业务类型
	BizType tTThostFtdcBizTypeType
	// 投资单元代码
	InvestUnitID tTThostFtdcInvestUnitIDType
}

// 深度行情
type tCThostFtdcDepthMarketDataField struct {
	// 交易日
	TradingDay tTThostFtdcDateType
	// 合约代码
	InstrumentID tTThostFtdcInstrumentIDType
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
	// 合约在交易所的代码
	ExchangeInstID tTThostFtdcExchangeInstIDType
	// 最新价
	LastPrice tTThostFtdcPriceType
	// 上次结算价
	PreSettlementPrice tTThostFtdcPriceType
	// 昨收盘
	PreClosePrice tTThostFtdcPriceType
	// 昨持仓量
	PreOpenInterest tTThostFtdcLargeVolumeType
	// 今开盘
	OpenPrice tTThostFtdcPriceType
	// 最高价
	HighestPrice tTThostFtdcPriceType
	// 最低价
	LowestPrice tTThostFtdcPriceType
	// 数量
	Volume tTThostFtdcVolumeType
	// 成交金额
	Turnover tTThostFtdcMoneyType
	// 持仓量
	OpenInterest tTThostFtdcLargeVolumeType
	// 今收盘
	ClosePrice tTThostFtdcPriceType
	// 本次结算价
	SettlementPrice tTThostFtdcPriceType
	// 涨停板价
	UpperLimitPrice tTThostFtdcPriceType
	// 跌停板价
	LowerLimitPrice tTThostFtdcPriceType
	// 昨虚实度
	PreDelta tTThostFtdcRatioType
	// 今虚实度
	CurrDelta tTThostFtdcRatioType
	// 最后修改时间
	UpdateTime tTThostFtdcTimeType
	// 最后修改毫秒
	UpdateMillisec tTThostFtdcMillisecType
	// 申买价一
	BidPrice1 tTThostFtdcPriceType
	// 申买量一
	BidVolume1 tTThostFtdcVolumeType
	// 申卖价一
	AskPrice1 tTThostFtdcPriceType
	// 申卖量一
	AskVolume1 tTThostFtdcVolumeType
	// 申买价二
	BidPrice2 tTThostFtdcPriceType
	// 申买量二
	BidVolume2 tTThostFtdcVolumeType
	// 申卖价二
	AskPrice2 tTThostFtdcPriceType
	// 申卖量二
	AskVolume2 tTThostFtdcVolumeType
	// 申买价三
	BidPrice3 tTThostFtdcPriceType
	// 申买量三
	BidVolume3 tTThostFtdcVolumeType
	// 申卖价三
	AskPrice3 tTThostFtdcPriceType
	// 申卖量三
	AskVolume3 tTThostFtdcVolumeType
	// 申买价四
	BidPrice4 tTThostFtdcPriceType
	// 申买量四
	BidVolume4 tTThostFtdcVolumeType
	// 申卖价四
	AskPrice4 tTThostFtdcPriceType
	// 申卖量四
	AskVolume4 tTThostFtdcVolumeType
	// 申买价五
	BidPrice5 tTThostFtdcPriceType
	// 申买量五
	BidVolume5 tTThostFtdcVolumeType
	// 申卖价五
	AskPrice5 tTThostFtdcPriceType
	// 申卖量五
	AskVolume5 tTThostFtdcVolumeType
	// 当日均价
	AveragePrice tTThostFtdcPriceType
	// 业务日期
	ActionDay tTThostFtdcDateType
}

// 投资者合约交易权限
type tCThostFtdcInstrumentTradingRightField struct {
	// 合约代码
	InstrumentID tTThostFtdcInstrumentIDType
	// 投资者范围
	InvestorRange tTThostFtdcInvestorRangeType
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 交易权限
	TradingRight tTThostFtdcTradingRightType
}

// 经纪公司用户
type tCThostFtdcBrokerUserField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 用户代码
	UserID tTThostFtdcUserIDType
	// 用户名称
	UserName tTThostFtdcUserNameType
	// 用户类型
	UserType tTThostFtdcUserTypeType
	// 是否活跃
	IsActive tTThostFtdcBoolType
	// 是否使用令牌
	IsUsingOTP tTThostFtdcBoolType
	// 是否强制终端认证
	IsAuthForce tTThostFtdcBoolType
}

// 经纪公司用户口令
type tCThostFtdcBrokerUserPasswordField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 用户代码
	UserID tTThostFtdcUserIDType
	// 密码
	Password tTThostFtdcPasswordType
	// 上次修改时间
	LastUpdateTime tTThostFtdcDateTimeType
	// 上次登陆时间
	LastLoginTime tTThostFtdcDateTimeType
	// 密码过期时间
	ExpireDate tTThostFtdcDateType
	// 弱密码过期时间
	WeakExpireDate tTThostFtdcDateType
}

// 经纪公司用户功能权限
type tCThostFtdcBrokerUserFunctionField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 用户代码
	UserID tTThostFtdcUserIDType
	// 经纪公司功能代码
	BrokerFunctionCode tTThostFtdcBrokerFunctionCodeType
}

// 交易所交易员报盘机
type tCThostFtdcTraderOfferField struct {
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
	// 交易所交易员代码
	TraderID tTThostFtdcTraderIDType
	// 会员代码
	ParticipantID tTThostFtdcParticipantIDType
	// 密码
	Password tTThostFtdcPasswordType
	// 安装编号
	InstallID tTThostFtdcInstallIDType
	// 本地报单编号
	OrderLocalID tTThostFtdcOrderLocalIDType
	// 交易所交易员连接状态
	TraderConnectStatus tTThostFtdcTraderConnectStatusType
	// 发出连接请求的日期
	ConnectRequestDate tTThostFtdcDateType
	// 发出连接请求的时间
	ConnectRequestTime tTThostFtdcTimeType
	// 上次报告日期
	LastReportDate tTThostFtdcDateType
	// 上次报告时间
	LastReportTime tTThostFtdcTimeType
	// 完成连接日期
	ConnectDate tTThostFtdcDateType
	// 完成连接时间
	ConnectTime tTThostFtdcTimeType
	// 启动日期
	StartDate tTThostFtdcDateType
	// 启动时间
	StartTime tTThostFtdcTimeType
	// 交易日
	TradingDay tTThostFtdcDateType
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 本席位最大成交编号
	MaxTradeID tTThostFtdcTradeIDType
	// 本席位最大报单备拷
	MaxOrderMessageReference tTThostFtdcReturnCodeType
}

// 投资者结算结果
type tCThostFtdcSettlementInfoField struct {
	// 交易日
	TradingDay tTThostFtdcDateType
	// 结算编号
	SettlementID tTThostFtdcSettlementIDType
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 序号
	SequenceNo tTThostFtdcSequenceNoType
	// 消息正文
	Content tTThostFtdcContentType
	// 投资者帐号
	AccountID tTThostFtdcAccountIDType
	// 币种代码
	CurrencyID tTThostFtdcCurrencyIDType
}

// 合约保证金率调整
type tCThostFtdcInstrumentMarginRateAdjustField struct {
	// 合约代码
	InstrumentID tTThostFtdcInstrumentIDType
	// 投资者范围
	InvestorRange tTThostFtdcInvestorRangeType
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 投机套保标志
	HedgeFlag tTThostFtdcHedgeFlagType
	// 多头保证金率
	LongMarginRatioByMoney tTThostFtdcRatioType
	// 多头保证金费
	LongMarginRatioByVolume tTThostFtdcMoneyType
	// 空头保证金率
	ShortMarginRatioByMoney tTThostFtdcRatioType
	// 空头保证金费
	ShortMarginRatioByVolume tTThostFtdcMoneyType
	// 是否相对交易所收取
	IsRelative tTThostFtdcBoolType
}

// 交易所保证金率
type tCThostFtdcExchangeMarginRateField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 合约代码
	InstrumentID tTThostFtdcInstrumentIDType
	// 投机套保标志
	HedgeFlag tTThostFtdcHedgeFlagType
	// 多头保证金率
	LongMarginRatioByMoney tTThostFtdcRatioType
	// 多头保证金费
	LongMarginRatioByVolume tTThostFtdcMoneyType
	// 空头保证金率
	ShortMarginRatioByMoney tTThostFtdcRatioType
	// 空头保证金费
	ShortMarginRatioByVolume tTThostFtdcMoneyType
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
}

// 交易所保证金率调整
type tCThostFtdcExchangeMarginRateAdjustField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 合约代码
	InstrumentID tTThostFtdcInstrumentIDType
	// 投机套保标志
	HedgeFlag tTThostFtdcHedgeFlagType
	// 跟随交易所投资者多头保证金率
	LongMarginRatioByMoney tTThostFtdcRatioType
	// 跟随交易所投资者多头保证金费
	LongMarginRatioByVolume tTThostFtdcMoneyType
	// 跟随交易所投资者空头保证金率
	ShortMarginRatioByMoney tTThostFtdcRatioType
	// 跟随交易所投资者空头保证金费
	ShortMarginRatioByVolume tTThostFtdcMoneyType
	// 交易所多头保证金率
	ExchLongMarginRatioByMoney tTThostFtdcRatioType
	// 交易所多头保证金费
	ExchLongMarginRatioByVolume tTThostFtdcMoneyType
	// 交易所空头保证金率
	ExchShortMarginRatioByMoney tTThostFtdcRatioType
	// 交易所空头保证金费
	ExchShortMarginRatioByVolume tTThostFtdcMoneyType
	// 不跟随交易所投资者多头保证金率
	NoLongMarginRatioByMoney tTThostFtdcRatioType
	// 不跟随交易所投资者多头保证金费
	NoLongMarginRatioByVolume tTThostFtdcMoneyType
	// 不跟随交易所投资者空头保证金率
	NoShortMarginRatioByMoney tTThostFtdcRatioType
	// 不跟随交易所投资者空头保证金费
	NoShortMarginRatioByVolume tTThostFtdcMoneyType
}

// 汇率
type tCThostFtdcExchangeRateField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 源币种
	FromCurrencyID tTThostFtdcCurrencyIDType
	// 源币种单位数量
	FromCurrencyUnit tTThostFtdcCurrencyUnitType
	// 目标币种
	ToCurrencyID tTThostFtdcCurrencyIDType
	// 汇率
	ExchangeRate tTThostFtdcExchangeRateType
}

// 结算引用
type tCThostFtdcSettlementRefField struct {
	// 交易日
	TradingDay tTThostFtdcDateType
	// 结算编号
	SettlementID tTThostFtdcSettlementIDType
}

// 当前时间
type tCThostFtdcCurrentTimeField struct {
	// 当前日期
	CurrDate tTThostFtdcDateType
	// 当前时间
	CurrTime tTThostFtdcTimeType
	// 当前时间（毫秒）
	CurrMillisec tTThostFtdcMillisecType
	// 业务日期
	ActionDay tTThostFtdcDateType
}

// 通讯阶段
type tCThostFtdcCommPhaseField struct {
	// 交易日
	TradingDay tTThostFtdcDateType
	// 通讯时段编号
	CommPhaseNo tTThostFtdcCommPhaseNoType
	// 系统编号
	SystemID tTThostFtdcSystemIDType
}

// 登录信息
type tCThostFtdcLoginInfoField struct {
	// 前置编号
	FrontID tTThostFtdcFrontIDType
	// 会话编号
	SessionID tTThostFtdcSessionIDType
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 用户代码
	UserID tTThostFtdcUserIDType
	// 登录日期
	LoginDate tTThostFtdcDateType
	// 登录时间
	LoginTime tTThostFtdcTimeType
	// IP地址
	IPAddress tTThostFtdcIPAddressType
	// 用户端产品信息
	UserProductInfo tTThostFtdcProductInfoType
	// 接口端产品信息
	InterfaceProductInfo tTThostFtdcProductInfoType
	// 协议信息
	ProtocolInfo tTThostFtdcProtocolInfoType
	// 系统名称
	SystemName tTThostFtdcSystemNameType
	// 密码,已弃用
	PasswordDeprecated tTThostFtdcPasswordType
	// 最大报单引用
	MaxOrderRef tTThostFtdcOrderRefType
	// 上期所时间
	SHFETime tTThostFtdcTimeType
	// 大商所时间
	DCETime tTThostFtdcTimeType
	// 郑商所时间
	CZCETime tTThostFtdcTimeType
	// 中金所时间
	FFEXTime tTThostFtdcTimeType
	// Mac地址
	MacAddress tTThostFtdcMacAddressType
	// 动态密码
	OneTimePassword tTThostFtdcPasswordType
	// 能源中心时间
	INETime tTThostFtdcTimeType
	// 查询时是否需要流控
	IsQryControl tTThostFtdcBoolType
	// 登录备注
	LoginRemark tTThostFtdcLoginRemarkType
	// 密码
	Password tTThostFtdcPasswordType
}

// 登录信息
type tCThostFtdcLogoutAllField struct {
	// 前置编号
	FrontID tTThostFtdcFrontIDType
	// 会话编号
	SessionID tTThostFtdcSessionIDType
	// 系统名称
	SystemName tTThostFtdcSystemNameType
}

// 前置状态
type tCThostFtdcFrontStatusField struct {
	// 前置编号
	FrontID tTThostFtdcFrontIDType
	// 上次报告日期
	LastReportDate tTThostFtdcDateType
	// 上次报告时间
	LastReportTime tTThostFtdcTimeType
	// 是否活跃
	IsActive tTThostFtdcBoolType
}

// 用户口令变更
type tCThostFtdcUserPasswordUpdateField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 用户代码
	UserID tTThostFtdcUserIDType
	// 原来的口令
	OldPassword tTThostFtdcPasswordType
	// 新的口令
	NewPassword tTThostFtdcPasswordType
}

// 输入报单
type tCThostFtdcInputOrderField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 合约代码
	InstrumentID tTThostFtdcInstrumentIDType
	// 报单引用
	OrderRef tTThostFtdcOrderRefType
	// 用户代码
	UserID tTThostFtdcUserIDType
	// 报单价格条件
	OrderPriceType tTThostFtdcOrderPriceTypeType
	// 买卖方向
	Direction tTThostFtdcDirectionType
	// 组合开平标志
	CombOffsetFlag tTThostFtdcCombOffsetFlagType
	// 组合投机套保标志
	CombHedgeFlag tTThostFtdcCombHedgeFlagType
	// 价格
	LimitPrice tTThostFtdcPriceType
	// 数量
	VolumeTotalOriginal tTThostFtdcVolumeType
	// 有效期类型
	TimeCondition tTThostFtdcTimeConditionType
	// GTD日期
	GTDDate tTThostFtdcDateType
	// 成交量类型
	VolumeCondition tTThostFtdcVolumeConditionType
	// 最小成交量
	MinVolume tTThostFtdcVolumeType
	// 触发条件
	ContingentCondition tTThostFtdcContingentConditionType
	// 止损价
	StopPrice tTThostFtdcPriceType
	// 强平原因
	ForceCloseReason tTThostFtdcForceCloseReasonType
	// 自动挂起标志
	IsAutoSuspend tTThostFtdcBoolType
	// 业务单元
	BusinessUnit tTThostFtdcBusinessUnitType
	// 请求编号
	RequestID tTThostFtdcRequestIDType
	// 用户强评标志
	UserForceClose tTThostFtdcBoolType
	// 互换单标志
	IsSwapOrder tTThostFtdcBoolType
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
	// 投资单元代码
	InvestUnitID tTThostFtdcInvestUnitIDType
	// 资金账号
	AccountID tTThostFtdcAccountIDType
	// 币种代码
	CurrencyID tTThostFtdcCurrencyIDType
	// 交易编码
	ClientID tTThostFtdcClientIDType
	// IP地址
	IPAddress tTThostFtdcIPAddressType
	// Mac地址
	MacAddress tTThostFtdcMacAddressType
}

// 报单
type tCThostFtdcOrderField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 合约代码
	InstrumentID tTThostFtdcInstrumentIDType
	// 报单引用
	OrderRef tTThostFtdcOrderRefType
	// 用户代码
	UserID tTThostFtdcUserIDType
	// 报单价格条件
	OrderPriceType tTThostFtdcOrderPriceTypeType
	// 买卖方向
	Direction tTThostFtdcDirectionType
	// 组合开平标志
	CombOffsetFlag tTThostFtdcCombOffsetFlagType
	// 组合投机套保标志
	CombHedgeFlag tTThostFtdcCombHedgeFlagType
	// 价格
	LimitPrice tTThostFtdcPriceType
	// 数量
	VolumeTotalOriginal tTThostFtdcVolumeType
	// 有效期类型
	TimeCondition tTThostFtdcTimeConditionType
	// GTD日期
	GTDDate tTThostFtdcDateType
	// 成交量类型
	VolumeCondition tTThostFtdcVolumeConditionType
	// 最小成交量
	MinVolume tTThostFtdcVolumeType
	// 触发条件
	ContingentCondition tTThostFtdcContingentConditionType
	// 止损价
	StopPrice tTThostFtdcPriceType
	// 强平原因
	ForceCloseReason tTThostFtdcForceCloseReasonType
	// 自动挂起标志
	IsAutoSuspend tTThostFtdcBoolType
	// 业务单元
	BusinessUnit tTThostFtdcBusinessUnitType
	// 请求编号
	RequestID tTThostFtdcRequestIDType
	// 本地报单编号
	OrderLocalID tTThostFtdcOrderLocalIDType
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
	// 会员代码
	ParticipantID tTThostFtdcParticipantIDType
	// 客户代码
	ClientID tTThostFtdcClientIDType
	// 合约在交易所的代码
	ExchangeInstID tTThostFtdcExchangeInstIDType
	// 交易所交易员代码
	TraderID tTThostFtdcTraderIDType
	// 安装编号
	InstallID tTThostFtdcInstallIDType
	// 报单提交状态
	OrderSubmitStatus tTThostFtdcOrderSubmitStatusType
	// 报单提示序号
	NotifySequence tTThostFtdcSequenceNoType
	// 交易日
	TradingDay tTThostFtdcDateType
	// 结算编号
	SettlementID tTThostFtdcSettlementIDType
	// 报单编号
	OrderSysID tTThostFtdcOrderSysIDType
	// 报单来源
	OrderSource tTThostFtdcOrderSourceType
	// 报单状态
	OrderStatus tTThostFtdcOrderStatusType
	// 报单类型
	OrderType tTThostFtdcOrderTypeType
	// 今成交数量
	VolumeTraded tTThostFtdcVolumeType
	// 剩余数量
	VolumeTotal tTThostFtdcVolumeType
	// 报单日期
	InsertDate tTThostFtdcDateType
	// 委托时间
	InsertTime tTThostFtdcTimeType
	// 激活时间
	ActiveTime tTThostFtdcTimeType
	// 挂起时间
	SuspendTime tTThostFtdcTimeType
	// 最后修改时间
	UpdateTime tTThostFtdcTimeType
	// 撤销时间
	CancelTime tTThostFtdcTimeType
	// 最后修改交易所交易员代码
	ActiveTraderID tTThostFtdcTraderIDType
	// 结算会员编号
	ClearingPartID tTThostFtdcParticipantIDType
	// 序号
	SequenceNo tTThostFtdcSequenceNoType
	// 前置编号
	FrontID tTThostFtdcFrontIDType
	// 会话编号
	SessionID tTThostFtdcSessionIDType
	// 用户端产品信息
	UserProductInfo tTThostFtdcProductInfoType
	// 状态信息
	StatusMsg tTThostFtdcErrorMsgType
	// 用户强评标志
	UserForceClose tTThostFtdcBoolType
	// 操作用户代码
	ActiveUserID tTThostFtdcUserIDType
	// 经纪公司报单编号
	BrokerOrderSeq tTThostFtdcSequenceNoType
	// 相关报单
	RelativeOrderSysID tTThostFtdcOrderSysIDType
	// 郑商所成交数量
	ZCETotalTradedVolume tTThostFtdcVolumeType
	// 互换单标志
	IsSwapOrder tTThostFtdcBoolType
	// 营业部编号
	BranchID tTThostFtdcBranchIDType
	// 投资单元代码
	InvestUnitID tTThostFtdcInvestUnitIDType
	// 资金账号
	AccountID tTThostFtdcAccountIDType
	// 币种代码
	CurrencyID tTThostFtdcCurrencyIDType
	// IP地址
	IPAddress tTThostFtdcIPAddressType
	// Mac地址
	MacAddress tTThostFtdcMacAddressType
}

// 交易所报单
type tCThostFtdcExchangeOrderField struct {
	// 报单价格条件
	OrderPriceType tTThostFtdcOrderPriceTypeType
	// 买卖方向
	Direction tTThostFtdcDirectionType
	// 组合开平标志
	CombOffsetFlag tTThostFtdcCombOffsetFlagType
	// 组合投机套保标志
	CombHedgeFlag tTThostFtdcCombHedgeFlagType
	// 价格
	LimitPrice tTThostFtdcPriceType
	// 数量
	VolumeTotalOriginal tTThostFtdcVolumeType
	// 有效期类型
	TimeCondition tTThostFtdcTimeConditionType
	// GTD日期
	GTDDate tTThostFtdcDateType
	// 成交量类型
	VolumeCondition tTThostFtdcVolumeConditionType
	// 最小成交量
	MinVolume tTThostFtdcVolumeType
	// 触发条件
	ContingentCondition tTThostFtdcContingentConditionType
	// 止损价
	StopPrice tTThostFtdcPriceType
	// 强平原因
	ForceCloseReason tTThostFtdcForceCloseReasonType
	// 自动挂起标志
	IsAutoSuspend tTThostFtdcBoolType
	// 业务单元
	BusinessUnit tTThostFtdcBusinessUnitType
	// 请求编号
	RequestID tTThostFtdcRequestIDType
	// 本地报单编号
	OrderLocalID tTThostFtdcOrderLocalIDType
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
	// 会员代码
	ParticipantID tTThostFtdcParticipantIDType
	// 客户代码
	ClientID tTThostFtdcClientIDType
	// 合约在交易所的代码
	ExchangeInstID tTThostFtdcExchangeInstIDType
	// 交易所交易员代码
	TraderID tTThostFtdcTraderIDType
	// 安装编号
	InstallID tTThostFtdcInstallIDType
	// 报单提交状态
	OrderSubmitStatus tTThostFtdcOrderSubmitStatusType
	// 报单提示序号
	NotifySequence tTThostFtdcSequenceNoType
	// 交易日
	TradingDay tTThostFtdcDateType
	// 结算编号
	SettlementID tTThostFtdcSettlementIDType
	// 报单编号
	OrderSysID tTThostFtdcOrderSysIDType
	// 报单来源
	OrderSource tTThostFtdcOrderSourceType
	// 报单状态
	OrderStatus tTThostFtdcOrderStatusType
	// 报单类型
	OrderType tTThostFtdcOrderTypeType
	// 今成交数量
	VolumeTraded tTThostFtdcVolumeType
	// 剩余数量
	VolumeTotal tTThostFtdcVolumeType
	// 报单日期
	InsertDate tTThostFtdcDateType
	// 委托时间
	InsertTime tTThostFtdcTimeType
	// 激活时间
	ActiveTime tTThostFtdcTimeType
	// 挂起时间
	SuspendTime tTThostFtdcTimeType
	// 最后修改时间
	UpdateTime tTThostFtdcTimeType
	// 撤销时间
	CancelTime tTThostFtdcTimeType
	// 最后修改交易所交易员代码
	ActiveTraderID tTThostFtdcTraderIDType
	// 结算会员编号
	ClearingPartID tTThostFtdcParticipantIDType
	// 序号
	SequenceNo tTThostFtdcSequenceNoType
	// 营业部编号
	BranchID tTThostFtdcBranchIDType
	// IP地址
	IPAddress tTThostFtdcIPAddressType
	// Mac地址
	MacAddress tTThostFtdcMacAddressType
}

// 交易所报单插入失败
type tCThostFtdcExchangeOrderInsertErrorField struct {
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
	// 会员代码
	ParticipantID tTThostFtdcParticipantIDType
	// 交易所交易员代码
	TraderID tTThostFtdcTraderIDType
	// 安装编号
	InstallID tTThostFtdcInstallIDType
	// 本地报单编号
	OrderLocalID tTThostFtdcOrderLocalIDType
	// 错误代码
	ErrorID tTThostFtdcErrorIDType
	// 错误信息
	ErrorMsg tTThostFtdcErrorMsgType
}

// 输入报单操作
type tCThostFtdcInputOrderActionField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 报单操作引用
	OrderActionRef tTThostFtdcOrderActionRefType
	// 报单引用
	OrderRef tTThostFtdcOrderRefType
	// 请求编号
	RequestID tTThostFtdcRequestIDType
	// 前置编号
	FrontID tTThostFtdcFrontIDType
	// 会话编号
	SessionID tTThostFtdcSessionIDType
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
	// 报单编号
	OrderSysID tTThostFtdcOrderSysIDType
	// 操作标志
	ActionFlag tTThostFtdcActionFlagType
	// 价格
	LimitPrice tTThostFtdcPriceType
	// 数量变化
	VolumeChange tTThostFtdcVolumeType
	// 用户代码
	UserID tTThostFtdcUserIDType
	// 合约代码
	InstrumentID tTThostFtdcInstrumentIDType
	// 投资单元代码
	InvestUnitID tTThostFtdcInvestUnitIDType
	// IP地址
	IPAddress tTThostFtdcIPAddressType
	// Mac地址
	MacAddress tTThostFtdcMacAddressType
}

// 报单操作
type tCThostFtdcOrderActionField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 报单操作引用
	OrderActionRef tTThostFtdcOrderActionRefType
	// 报单引用
	OrderRef tTThostFtdcOrderRefType
	// 请求编号
	RequestID tTThostFtdcRequestIDType
	// 前置编号
	FrontID tTThostFtdcFrontIDType
	// 会话编号
	SessionID tTThostFtdcSessionIDType
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
	// 报单编号
	OrderSysID tTThostFtdcOrderSysIDType
	// 操作标志
	ActionFlag tTThostFtdcActionFlagType
	// 价格
	LimitPrice tTThostFtdcPriceType
	// 数量变化
	VolumeChange tTThostFtdcVolumeType
	// 操作日期
	ActionDate tTThostFtdcDateType
	// 操作时间
	ActionTime tTThostFtdcTimeType
	// 交易所交易员代码
	TraderID tTThostFtdcTraderIDType
	// 安装编号
	InstallID tTThostFtdcInstallIDType
	// 本地报单编号
	OrderLocalID tTThostFtdcOrderLocalIDType
	// 操作本地编号
	ActionLocalID tTThostFtdcOrderLocalIDType
	// 会员代码
	ParticipantID tTThostFtdcParticipantIDType
	// 客户代码
	ClientID tTThostFtdcClientIDType
	// 业务单元
	BusinessUnit tTThostFtdcBusinessUnitType
	// 报单操作状态
	OrderActionStatus tTThostFtdcOrderActionStatusType
	// 用户代码
	UserID tTThostFtdcUserIDType
	// 状态信息
	StatusMsg tTThostFtdcErrorMsgType
	// 合约代码
	InstrumentID tTThostFtdcInstrumentIDType
	// 营业部编号
	BranchID tTThostFtdcBranchIDType
	// 投资单元代码
	InvestUnitID tTThostFtdcInvestUnitIDType
	// IP地址
	IPAddress tTThostFtdcIPAddressType
	// Mac地址
	MacAddress tTThostFtdcMacAddressType
}

// 交易所报单操作
type tCThostFtdcExchangeOrderActionField struct {
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
	// 报单编号
	OrderSysID tTThostFtdcOrderSysIDType
	// 操作标志
	ActionFlag tTThostFtdcActionFlagType
	// 价格
	LimitPrice tTThostFtdcPriceType
	// 数量变化
	VolumeChange tTThostFtdcVolumeType
	// 操作日期
	ActionDate tTThostFtdcDateType
	// 操作时间
	ActionTime tTThostFtdcTimeType
	// 交易所交易员代码
	TraderID tTThostFtdcTraderIDType
	// 安装编号
	InstallID tTThostFtdcInstallIDType
	// 本地报单编号
	OrderLocalID tTThostFtdcOrderLocalIDType
	// 操作本地编号
	ActionLocalID tTThostFtdcOrderLocalIDType
	// 会员代码
	ParticipantID tTThostFtdcParticipantIDType
	// 客户代码
	ClientID tTThostFtdcClientIDType
	// 业务单元
	BusinessUnit tTThostFtdcBusinessUnitType
	// 报单操作状态
	OrderActionStatus tTThostFtdcOrderActionStatusType
	// 用户代码
	UserID tTThostFtdcUserIDType
	// 营业部编号
	BranchID tTThostFtdcBranchIDType
	// IP地址
	IPAddress tTThostFtdcIPAddressType
	// Mac地址
	MacAddress tTThostFtdcMacAddressType
}

// 交易所报单操作失败
type tCThostFtdcExchangeOrderActionErrorField struct {
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
	// 报单编号
	OrderSysID tTThostFtdcOrderSysIDType
	// 交易所交易员代码
	TraderID tTThostFtdcTraderIDType
	// 安装编号
	InstallID tTThostFtdcInstallIDType
	// 本地报单编号
	OrderLocalID tTThostFtdcOrderLocalIDType
	// 操作本地编号
	ActionLocalID tTThostFtdcOrderLocalIDType
	// 错误代码
	ErrorID tTThostFtdcErrorIDType
	// 错误信息
	ErrorMsg tTThostFtdcErrorMsgType
}

// 交易所成交
type tCThostFtdcExchangeTradeField struct {
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
	// 成交编号
	TradeID tTThostFtdcTradeIDType
	// 买卖方向
	Direction tTThostFtdcDirectionType
	// 报单编号
	OrderSysID tTThostFtdcOrderSysIDType
	// 会员代码
	ParticipantID tTThostFtdcParticipantIDType
	// 客户代码
	ClientID tTThostFtdcClientIDType
	// 交易角色
	TradingRole tTThostFtdcTradingRoleType
	// 合约在交易所的代码
	ExchangeInstID tTThostFtdcExchangeInstIDType
	// 开平标志
	OffsetFlag tTThostFtdcOffsetFlagType
	// 投机套保标志
	HedgeFlag tTThostFtdcHedgeFlagType
	// 价格
	Price tTThostFtdcPriceType
	// 数量
	Volume tTThostFtdcVolumeType
	// 成交时期
	TradeDate tTThostFtdcDateType
	// 成交时间
	TradeTime tTThostFtdcTimeType
	// 成交类型
	TradeType tTThostFtdcTradeTypeType
	// 成交价来源
	PriceSource tTThostFtdcPriceSourceType
	// 交易所交易员代码
	TraderID tTThostFtdcTraderIDType
	// 本地报单编号
	OrderLocalID tTThostFtdcOrderLocalIDType
	// 结算会员编号
	ClearingPartID tTThostFtdcParticipantIDType
	// 业务单元
	BusinessUnit tTThostFtdcBusinessUnitType
	// 序号
	SequenceNo tTThostFtdcSequenceNoType
	// 成交来源
	TradeSource tTThostFtdcTradeSourceType
}

// 成交
type tCThostFtdcTradeField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 合约代码
	InstrumentID tTThostFtdcInstrumentIDType
	// 报单引用
	OrderRef tTThostFtdcOrderRefType
	// 用户代码
	UserID tTThostFtdcUserIDType
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
	// 成交编号
	TradeID tTThostFtdcTradeIDType
	// 买卖方向
	Direction tTThostFtdcDirectionType
	// 报单编号
	OrderSysID tTThostFtdcOrderSysIDType
	// 会员代码
	ParticipantID tTThostFtdcParticipantIDType
	// 客户代码
	ClientID tTThostFtdcClientIDType
	// 交易角色
	TradingRole tTThostFtdcTradingRoleType
	// 合约在交易所的代码
	ExchangeInstID tTThostFtdcExchangeInstIDType
	// 开平标志
	OffsetFlag tTThostFtdcOffsetFlagType
	// 投机套保标志
	HedgeFlag tTThostFtdcHedgeFlagType
	// 价格
	Price tTThostFtdcPriceType
	// 数量
	Volume tTThostFtdcVolumeType
	// 成交时期
	TradeDate tTThostFtdcDateType
	// 成交时间
	TradeTime tTThostFtdcTimeType
	// 成交类型
	TradeType tTThostFtdcTradeTypeType
	// 成交价来源
	PriceSource tTThostFtdcPriceSourceType
	// 交易所交易员代码
	TraderID tTThostFtdcTraderIDType
	// 本地报单编号
	OrderLocalID tTThostFtdcOrderLocalIDType
	// 结算会员编号
	ClearingPartID tTThostFtdcParticipantIDType
	// 业务单元
	BusinessUnit tTThostFtdcBusinessUnitType
	// 序号
	SequenceNo tTThostFtdcSequenceNoType
	// 交易日
	TradingDay tTThostFtdcDateType
	// 结算编号
	SettlementID tTThostFtdcSettlementIDType
	// 经纪公司报单编号
	BrokerOrderSeq tTThostFtdcSequenceNoType
	// 成交来源
	TradeSource tTThostFtdcTradeSourceType
	// 投资单元代码
	InvestUnitID tTThostFtdcInvestUnitIDType
}

// 用户会话
type tCThostFtdcUserSessionField struct {
	// 前置编号
	FrontID tTThostFtdcFrontIDType
	// 会话编号
	SessionID tTThostFtdcSessionIDType
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 用户代码
	UserID tTThostFtdcUserIDType
	// 登录日期
	LoginDate tTThostFtdcDateType
	// 登录时间
	LoginTime tTThostFtdcTimeType
	// IP地址
	IPAddress tTThostFtdcIPAddressType
	// 用户端产品信息
	UserProductInfo tTThostFtdcProductInfoType
	// 接口端产品信息
	InterfaceProductInfo tTThostFtdcProductInfoType
	// 协议信息
	ProtocolInfo tTThostFtdcProtocolInfoType
	// Mac地址
	MacAddress tTThostFtdcMacAddressType
	// 登录备注
	LoginRemark tTThostFtdcLoginRemarkType
}

// 查询最大报单数量
type tCThostFtdcQueryMaxOrderVolumeField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 合约代码
	InstrumentID tTThostFtdcInstrumentIDType
	// 买卖方向
	Direction tTThostFtdcDirectionType
	// 开平标志
	OffsetFlag tTThostFtdcOffsetFlagType
	// 投机套保标志
	HedgeFlag tTThostFtdcHedgeFlagType
	// 最大允许报单数量
	MaxVolume tTThostFtdcVolumeType
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
	// 投资单元代码
	InvestUnitID tTThostFtdcInvestUnitIDType
}

// 投资者结算结果确认信息
type tCThostFtdcSettlementInfoConfirmField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 确认日期
	ConfirmDate tTThostFtdcDateType
	// 确认时间
	ConfirmTime tTThostFtdcTimeType
	// 结算编号
	SettlementID tTThostFtdcSettlementIDType
	// 投资者帐号
	AccountID tTThostFtdcAccountIDType
	// 币种代码
	CurrencyID tTThostFtdcCurrencyIDType
}

// 出入金同步
type tCThostFtdcSyncDepositField struct {
	// 出入金流水号
	DepositSeqNo tTThostFtdcDepositSeqNoType
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 入金金额
	Deposit tTThostFtdcMoneyType
	// 是否强制进行
	IsForce tTThostFtdcBoolType
	// 币种代码
	CurrencyID tTThostFtdcCurrencyIDType
}

// 货币质押同步
type tCThostFtdcSyncFundMortgageField struct {
	// 货币质押流水号
	MortgageSeqNo tTThostFtdcDepositSeqNoType
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 源币种
	FromCurrencyID tTThostFtdcCurrencyIDType
	// 质押金额
	MortgageAmount tTThostFtdcMoneyType
	// 目标币种
	ToCurrencyID tTThostFtdcCurrencyIDType
}

// 经纪公司同步
type tCThostFtdcBrokerSyncField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
}

// 正在同步中的投资者
type tCThostFtdcSyncingInvestorField struct {
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者分组代码
	InvestorGroupID tTThostFtdcInvestorIDType
	// 投资者名称
	InvestorName tTThostFtdcPartyNameType
	// 证件类型
	IdentifiedCardType tTThostFtdcIdCardTypeType
	// 证件号码
	IdentifiedCardNo tTThostFtdcIdentifiedCardNoType
	// 是否活跃
	IsActive tTThostFtdcBoolType
	// 联系电话
	Telephone tTThostFtdcTelephoneType
	// 通讯地址
	Address tTThostFtdcAddressType
	// 开户日期
	OpenDate tTThostFtdcDateType
	// 手机
	Mobile tTThostFtdcMobileType
	// 手续费率模板代码
	CommModelID tTThostFtdcInvestorIDType
	// 保证金率模板代码
	MarginModelID tTThostFtdcInvestorIDType
}

// 正在同步中的交易代码
type tCThostFtdcSyncingTradingCodeField struct {
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
	// 客户代码
	ClientID tTThostFtdcClientIDType
	// 是否活跃
	IsActive tTThostFtdcBoolType
	// 交易编码类型
	ClientIDType tTThostFtdcClientIDTypeType
}

// 正在同步中的投资者分组
type tCThostFtdcSyncingInvestorGroupField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者分组代码
	InvestorGroupID tTThostFtdcInvestorIDType
	// 投资者分组名称
	InvestorGroupName tTThostFtdcInvestorGroupNameType
}

// 正在同步中的交易账号
type tCThostFtdcSyncingTradingAccountField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者帐号
	AccountID tTThostFtdcAccountIDType
	// 上次质押金额
	PreMortgage tTThostFtdcMoneyType
	// 上次信用额度
	PreCredit tTThostFtdcMoneyType
	// 上次存款额
	PreDeposit tTThostFtdcMoneyType
	// 上次结算准备金
	PreBalance tTThostFtdcMoneyType
	// 上次占用的保证金
	PreMargin tTThostFtdcMoneyType
	// 利息基数
	InterestBase tTThostFtdcMoneyType
	// 利息收入
	Interest tTThostFtdcMoneyType
	// 入金金额
	Deposit tTThostFtdcMoneyType
	// 出金金额
	Withdraw tTThostFtdcMoneyType
	// 冻结的保证金
	FrozenMargin tTThostFtdcMoneyType
	// 冻结的资金
	FrozenCash tTThostFtdcMoneyType
	// 冻结的手续费
	FrozenCommission tTThostFtdcMoneyType
	// 当前保证金总额
	CurrMargin tTThostFtdcMoneyType
	// 资金差额
	CashIn tTThostFtdcMoneyType
	// 手续费
	Commission tTThostFtdcMoneyType
	// 平仓盈亏
	CloseProfit tTThostFtdcMoneyType
	// 持仓盈亏
	PositionProfit tTThostFtdcMoneyType
	// 期货结算准备金
	Balance tTThostFtdcMoneyType
	// 可用资金
	Available tTThostFtdcMoneyType
	// 可取资金
	WithdrawQuota tTThostFtdcMoneyType
	// 基本准备金
	Reserve tTThostFtdcMoneyType
	// 交易日
	TradingDay tTThostFtdcDateType
	// 结算编号
	SettlementID tTThostFtdcSettlementIDType
	// 信用额度
	Credit tTThostFtdcMoneyType
	// 质押金额
	Mortgage tTThostFtdcMoneyType
	// 交易所保证金
	ExchangeMargin tTThostFtdcMoneyType
	// 投资者交割保证金
	DeliveryMargin tTThostFtdcMoneyType
	// 交易所交割保证金
	ExchangeDeliveryMargin tTThostFtdcMoneyType
	// 保底期货结算准备金
	ReserveBalance tTThostFtdcMoneyType
	// 币种代码
	CurrencyID tTThostFtdcCurrencyIDType
	// 上次货币质入金额
	PreFundMortgageIn tTThostFtdcMoneyType
	// 上次货币质出金额
	PreFundMortgageOut tTThostFtdcMoneyType
	// 货币质入金额
	FundMortgageIn tTThostFtdcMoneyType
	// 货币质出金额
	FundMortgageOut tTThostFtdcMoneyType
	// 货币质押余额
	FundMortgageAvailable tTThostFtdcMoneyType
	// 可质押货币金额
	MortgageableFund tTThostFtdcMoneyType
	// 特殊产品占用保证金
	SpecProductMargin tTThostFtdcMoneyType
	// 特殊产品冻结保证金
	SpecProductFrozenMargin tTThostFtdcMoneyType
	// 特殊产品手续费
	SpecProductCommission tTThostFtdcMoneyType
	// 特殊产品冻结手续费
	SpecProductFrozenCommission tTThostFtdcMoneyType
	// 特殊产品持仓盈亏
	SpecProductPositionProfit tTThostFtdcMoneyType
	// 特殊产品平仓盈亏
	SpecProductCloseProfit tTThostFtdcMoneyType
	// 根据持仓盈亏算法计算的特殊产品持仓盈亏
	SpecProductPositionProfitByAlg tTThostFtdcMoneyType
	// 特殊产品交易所保证金
	SpecProductExchangeMargin tTThostFtdcMoneyType
	// 延时换汇冻结金额
	FrozenSwap tTThostFtdcMoneyType
	// 剩余换汇额度
	RemainSwap tTThostFtdcMoneyType
}

// 正在同步中的投资者持仓
type tCThostFtdcSyncingInvestorPositionField struct {
	// 合约代码
	InstrumentID tTThostFtdcInstrumentIDType
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 持仓多空方向
	PosiDirection tTThostFtdcPosiDirectionType
	// 投机套保标志
	HedgeFlag tTThostFtdcHedgeFlagType
	// 持仓日期
	PositionDate tTThostFtdcPositionDateType
	// 上日持仓
	YdPosition tTThostFtdcVolumeType
	// 今日持仓
	Position tTThostFtdcVolumeType
	// 多头冻结
	LongFrozen tTThostFtdcVolumeType
	// 空头冻结
	ShortFrozen tTThostFtdcVolumeType
	// 开仓冻结金额
	LongFrozenAmount tTThostFtdcMoneyType
	// 开仓冻结金额
	ShortFrozenAmount tTThostFtdcMoneyType
	// 开仓量
	OpenVolume tTThostFtdcVolumeType
	// 平仓量
	CloseVolume tTThostFtdcVolumeType
	// 开仓金额
	OpenAmount tTThostFtdcMoneyType
	// 平仓金额
	CloseAmount tTThostFtdcMoneyType
	// 持仓成本
	PositionCost tTThostFtdcMoneyType
	// 上次占用的保证金
	PreMargin tTThostFtdcMoneyType
	// 占用的保证金
	UseMargin tTThostFtdcMoneyType
	// 冻结的保证金
	FrozenMargin tTThostFtdcMoneyType
	// 冻结的资金
	FrozenCash tTThostFtdcMoneyType
	// 冻结的手续费
	FrozenCommission tTThostFtdcMoneyType
	// 资金差额
	CashIn tTThostFtdcMoneyType
	// 手续费
	Commission tTThostFtdcMoneyType
	// 平仓盈亏
	CloseProfit tTThostFtdcMoneyType
	// 持仓盈亏
	PositionProfit tTThostFtdcMoneyType
	// 上次结算价
	PreSettlementPrice tTThostFtdcPriceType
	// 本次结算价
	SettlementPrice tTThostFtdcPriceType
	// 交易日
	TradingDay tTThostFtdcDateType
	// 结算编号
	SettlementID tTThostFtdcSettlementIDType
	// 开仓成本
	OpenCost tTThostFtdcMoneyType
	// 交易所保证金
	ExchangeMargin tTThostFtdcMoneyType
	// 组合成交形成的持仓
	CombPosition tTThostFtdcVolumeType
	// 组合多头冻结
	CombLongFrozen tTThostFtdcVolumeType
	// 组合空头冻结
	CombShortFrozen tTThostFtdcVolumeType
	// 逐日盯市平仓盈亏
	CloseProfitByDate tTThostFtdcMoneyType
	// 逐笔对冲平仓盈亏
	CloseProfitByTrade tTThostFtdcMoneyType
	// 今日持仓
	TodayPosition tTThostFtdcVolumeType
	// 保证金率
	MarginRateByMoney tTThostFtdcRatioType
	// 保证金率(按手数)
	MarginRateByVolume tTThostFtdcRatioType
	// 执行冻结
	StrikeFrozen tTThostFtdcVolumeType
	// 执行冻结金额
	StrikeFrozenAmount tTThostFtdcMoneyType
	// 放弃执行冻结
	AbandonFrozen tTThostFtdcVolumeType
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
	// 执行冻结的昨仓
	YdStrikeFrozen tTThostFtdcVolumeType
	// 投资单元代码
	InvestUnitID tTThostFtdcInvestUnitIDType
	// 大商所持仓成本差值，只有大商所使用
	PositionCostOffset tTThostFtdcMoneyType
}

// 正在同步中的合约保证金率
type tCThostFtdcSyncingInstrumentMarginRateField struct {
	// 合约代码
	InstrumentID tTThostFtdcInstrumentIDType
	// 投资者范围
	InvestorRange tTThostFtdcInvestorRangeType
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 投机套保标志
	HedgeFlag tTThostFtdcHedgeFlagType
	// 多头保证金率
	LongMarginRatioByMoney tTThostFtdcRatioType
	// 多头保证金费
	LongMarginRatioByVolume tTThostFtdcMoneyType
	// 空头保证金率
	ShortMarginRatioByMoney tTThostFtdcRatioType
	// 空头保证金费
	ShortMarginRatioByVolume tTThostFtdcMoneyType
	// 是否相对交易所收取
	IsRelative tTThostFtdcBoolType
}

// 正在同步中的合约手续费率
type tCThostFtdcSyncingInstrumentCommissionRateField struct {
	// 合约代码
	InstrumentID tTThostFtdcInstrumentIDType
	// 投资者范围
	InvestorRange tTThostFtdcInvestorRangeType
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 开仓手续费率
	OpenRatioByMoney tTThostFtdcRatioType
	// 开仓手续费
	OpenRatioByVolume tTThostFtdcRatioType
	// 平仓手续费率
	CloseRatioByMoney tTThostFtdcRatioType
	// 平仓手续费
	CloseRatioByVolume tTThostFtdcRatioType
	// 平今手续费率
	CloseTodayRatioByMoney tTThostFtdcRatioType
	// 平今手续费
	CloseTodayRatioByVolume tTThostFtdcRatioType
}

// 正在同步中的合约交易权限
type tCThostFtdcSyncingInstrumentTradingRightField struct {
	// 合约代码
	InstrumentID tTThostFtdcInstrumentIDType
	// 投资者范围
	InvestorRange tTThostFtdcInvestorRangeType
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 交易权限
	TradingRight tTThostFtdcTradingRightType
}

// 查询报单
type tCThostFtdcQryOrderField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 合约代码
	InstrumentID tTThostFtdcInstrumentIDType
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
	// 报单编号
	OrderSysID tTThostFtdcOrderSysIDType
	// 开始时间
	InsertTimeStart tTThostFtdcTimeType
	// 结束时间
	InsertTimeEnd tTThostFtdcTimeType
	// 投资单元代码
	InvestUnitID tTThostFtdcInvestUnitIDType
}

// 查询成交
type tCThostFtdcQryTradeField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 合约代码
	InstrumentID tTThostFtdcInstrumentIDType
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
	// 成交编号
	TradeID tTThostFtdcTradeIDType
	// 开始时间
	TradeTimeStart tTThostFtdcTimeType
	// 结束时间
	TradeTimeEnd tTThostFtdcTimeType
	// 投资单元代码
	InvestUnitID tTThostFtdcInvestUnitIDType
}

// 查询投资者持仓
type tCThostFtdcQryInvestorPositionField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 合约代码
	InstrumentID tTThostFtdcInstrumentIDType
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
	// 投资单元代码
	InvestUnitID tTThostFtdcInvestUnitIDType
}

// 查询资金账户
type tCThostFtdcQryTradingAccountField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 币种代码
	CurrencyID tTThostFtdcCurrencyIDType
	// 业务类型
	BizType tTThostFtdcBizTypeType
	// 投资者帐号
	AccountID tTThostFtdcAccountIDType
}

// 查询投资者
type tCThostFtdcQryInvestorField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
}

// 查询交易编码
type tCThostFtdcQryTradingCodeField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
	// 客户代码
	ClientID tTThostFtdcClientIDType
	// 交易编码类型
	ClientIDType tTThostFtdcClientIDTypeType
	// 投资单元代码
	InvestUnitID tTThostFtdcInvestUnitIDType
}

// 查询投资者组
type tCThostFtdcQryInvestorGroupField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
}

// 查询合约保证金率
type tCThostFtdcQryInstrumentMarginRateField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 合约代码
	InstrumentID tTThostFtdcInstrumentIDType
	// 投机套保标志
	HedgeFlag tTThostFtdcHedgeFlagType
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
	// 投资单元代码
	InvestUnitID tTThostFtdcInvestUnitIDType
}

// 查询手续费率
type tCThostFtdcQryInstrumentCommissionRateField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 合约代码
	InstrumentID tTThostFtdcInstrumentIDType
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
	// 投资单元代码
	InvestUnitID tTThostFtdcInvestUnitIDType
}

// 查询合约交易权限
type tCThostFtdcQryInstrumentTradingRightField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 合约代码
	InstrumentID tTThostFtdcInstrumentIDType
}

// 查询经纪公司
type tCThostFtdcQryBrokerField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
}

// 查询交易员
type tCThostFtdcQryTraderField struct {
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
	// 会员代码
	ParticipantID tTThostFtdcParticipantIDType
	// 交易所交易员代码
	TraderID tTThostFtdcTraderIDType
}

// 查询管理用户功能权限
type tCThostFtdcQrySuperUserFunctionField struct {
	// 用户代码
	UserID tTThostFtdcUserIDType
}

// 查询用户会话
type tCThostFtdcQryUserSessionField struct {
	// 前置编号
	FrontID tTThostFtdcFrontIDType
	// 会话编号
	SessionID tTThostFtdcSessionIDType
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 用户代码
	UserID tTThostFtdcUserIDType
}

// 查询经纪公司会员代码
type tCThostFtdcQryPartBrokerField struct {
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 会员代码
	ParticipantID tTThostFtdcParticipantIDType
}

// 查询前置状态
type tCThostFtdcQryFrontStatusField struct {
	// 前置编号
	FrontID tTThostFtdcFrontIDType
}

// 查询交易所报单
type tCThostFtdcQryExchangeOrderField struct {
	// 会员代码
	ParticipantID tTThostFtdcParticipantIDType
	// 客户代码
	ClientID tTThostFtdcClientIDType
	// 合约在交易所的代码
	ExchangeInstID tTThostFtdcExchangeInstIDType
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
	// 交易所交易员代码
	TraderID tTThostFtdcTraderIDType
}

// 查询报单操作
type tCThostFtdcQryOrderActionField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
}

// 查询交易所报单操作
type tCThostFtdcQryExchangeOrderActionField struct {
	// 会员代码
	ParticipantID tTThostFtdcParticipantIDType
	// 客户代码
	ClientID tTThostFtdcClientIDType
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
	// 交易所交易员代码
	TraderID tTThostFtdcTraderIDType
}

// 查询管理用户
type tCThostFtdcQrySuperUserField struct {
	// 用户代码
	UserID tTThostFtdcUserIDType
}

// 查询交易所
type tCThostFtdcQryExchangeField struct {
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
}

// 查询产品
type tCThostFtdcQryProductField struct {
	// 产品代码
	ProductID tTThostFtdcInstrumentIDType
	// 产品类型
	ProductClass tTThostFtdcProductClassType
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
}

// 查询合约
type tCThostFtdcQryInstrumentField struct {
	// 合约代码
	InstrumentID tTThostFtdcInstrumentIDType
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
	// 合约在交易所的代码
	ExchangeInstID tTThostFtdcExchangeInstIDType
	// 产品代码
	ProductID tTThostFtdcInstrumentIDType
}

// 查询行情
type tCThostFtdcQryDepthMarketDataField struct {
	// 合约代码
	InstrumentID tTThostFtdcInstrumentIDType
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
}

// 查询经纪公司用户
type tCThostFtdcQryBrokerUserField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 用户代码
	UserID tTThostFtdcUserIDType
}

// 查询经纪公司用户权限
type tCThostFtdcQryBrokerUserFunctionField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 用户代码
	UserID tTThostFtdcUserIDType
}

// 查询交易员报盘机
type tCThostFtdcQryTraderOfferField struct {
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
	// 会员代码
	ParticipantID tTThostFtdcParticipantIDType
	// 交易所交易员代码
	TraderID tTThostFtdcTraderIDType
}

// 查询出入金流水
type tCThostFtdcQrySyncDepositField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 出入金流水号
	DepositSeqNo tTThostFtdcDepositSeqNoType
}

// 查询投资者结算结果
type tCThostFtdcQrySettlementInfoField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 交易日
	TradingDay tTThostFtdcDateType
	// 投资者帐号
	AccountID tTThostFtdcAccountIDType
	// 币种代码
	CurrencyID tTThostFtdcCurrencyIDType
}

// 查询交易所保证金率
type tCThostFtdcQryExchangeMarginRateField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 合约代码
	InstrumentID tTThostFtdcInstrumentIDType
	// 投机套保标志
	HedgeFlag tTThostFtdcHedgeFlagType
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
}

// 查询交易所调整保证金率
type tCThostFtdcQryExchangeMarginRateAdjustField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 合约代码
	InstrumentID tTThostFtdcInstrumentIDType
	// 投机套保标志
	HedgeFlag tTThostFtdcHedgeFlagType
}

// 查询汇率
type tCThostFtdcQryExchangeRateField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 源币种
	FromCurrencyID tTThostFtdcCurrencyIDType
	// 目标币种
	ToCurrencyID tTThostFtdcCurrencyIDType
}

// 查询货币质押流水
type tCThostFtdcQrySyncFundMortgageField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 货币质押流水号
	MortgageSeqNo tTThostFtdcDepositSeqNoType
}

// 查询报单
type tCThostFtdcQryHisOrderField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 合约代码
	InstrumentID tTThostFtdcInstrumentIDType
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
	// 报单编号
	OrderSysID tTThostFtdcOrderSysIDType
	// 开始时间
	InsertTimeStart tTThostFtdcTimeType
	// 结束时间
	InsertTimeEnd tTThostFtdcTimeType
	// 交易日
	TradingDay tTThostFtdcDateType
	// 结算编号
	SettlementID tTThostFtdcSettlementIDType
}

// 当前期权合约最小保证金
type tCThostFtdcOptionInstrMiniMarginField struct {
	// 合约代码
	InstrumentID tTThostFtdcInstrumentIDType
	// 投资者范围
	InvestorRange tTThostFtdcInvestorRangeType
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 单位（手）期权合约最小保证金
	MinMargin tTThostFtdcMoneyType
	// 取值方式
	ValueMethod tTThostFtdcValueMethodType
	// 是否跟随交易所收取
	IsRelative tTThostFtdcBoolType
}

// 当前期权合约保证金调整系数
type tCThostFtdcOptionInstrMarginAdjustField struct {
	// 合约代码
	InstrumentID tTThostFtdcInstrumentIDType
	// 投资者范围
	InvestorRange tTThostFtdcInvestorRangeType
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 投机空头保证金调整系数
	SShortMarginRatioByMoney tTThostFtdcRatioType
	// 投机空头保证金调整系数
	SShortMarginRatioByVolume tTThostFtdcMoneyType
	// 保值空头保证金调整系数
	HShortMarginRatioByMoney tTThostFtdcRatioType
	// 保值空头保证金调整系数
	HShortMarginRatioByVolume tTThostFtdcMoneyType
	// 套利空头保证金调整系数
	AShortMarginRatioByMoney tTThostFtdcRatioType
	// 套利空头保证金调整系数
	AShortMarginRatioByVolume tTThostFtdcMoneyType
	// 是否跟随交易所收取
	IsRelative tTThostFtdcBoolType
	// 做市商空头保证金调整系数
	MShortMarginRatioByMoney tTThostFtdcRatioType
	// 做市商空头保证金调整系数
	MShortMarginRatioByVolume tTThostFtdcMoneyType
}

// 当前期权合约手续费的详细内容
type tCThostFtdcOptionInstrCommRateField struct {
	// 合约代码
	InstrumentID tTThostFtdcInstrumentIDType
	// 投资者范围
	InvestorRange tTThostFtdcInvestorRangeType
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 开仓手续费率
	OpenRatioByMoney tTThostFtdcRatioType
	// 开仓手续费
	OpenRatioByVolume tTThostFtdcRatioType
	// 平仓手续费率
	CloseRatioByMoney tTThostFtdcRatioType
	// 平仓手续费
	CloseRatioByVolume tTThostFtdcRatioType
	// 平今手续费率
	CloseTodayRatioByMoney tTThostFtdcRatioType
	// 平今手续费
	CloseTodayRatioByVolume tTThostFtdcRatioType
	// 执行手续费率
	StrikeRatioByMoney tTThostFtdcRatioType
	// 执行手续费
	StrikeRatioByVolume tTThostFtdcRatioType
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
	// 投资单元代码
	InvestUnitID tTThostFtdcInvestUnitIDType
}

// 期权交易成本
type tCThostFtdcOptionInstrTradeCostField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 合约代码
	InstrumentID tTThostFtdcInstrumentIDType
	// 投机套保标志
	HedgeFlag tTThostFtdcHedgeFlagType
	// 期权合约保证金不变部分
	FixedMargin tTThostFtdcMoneyType
	// 期权合约最小保证金
	MiniMargin tTThostFtdcMoneyType
	// 期权合约权利金
	Royalty tTThostFtdcMoneyType
	// 交易所期权合约保证金不变部分
	ExchFixedMargin tTThostFtdcMoneyType
	// 交易所期权合约最小保证金
	ExchMiniMargin tTThostFtdcMoneyType
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
	// 投资单元代码
	InvestUnitID tTThostFtdcInvestUnitIDType
}

// 期权交易成本查询
type tCThostFtdcQryOptionInstrTradeCostField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 合约代码
	InstrumentID tTThostFtdcInstrumentIDType
	// 投机套保标志
	HedgeFlag tTThostFtdcHedgeFlagType
	// 期权合约报价
	InputPrice tTThostFtdcPriceType
	// 标的价格,填0则用昨结算价
	UnderlyingPrice tTThostFtdcPriceType
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
	// 投资单元代码
	InvestUnitID tTThostFtdcInvestUnitIDType
}

// 期权手续费率查询
type tCThostFtdcQryOptionInstrCommRateField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 合约代码
	InstrumentID tTThostFtdcInstrumentIDType
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
	// 投资单元代码
	InvestUnitID tTThostFtdcInvestUnitIDType
}

// 股指现货指数
type tCThostFtdcIndexPriceField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 合约代码
	InstrumentID tTThostFtdcInstrumentIDType
	// 指数现货收盘价
	ClosePrice tTThostFtdcPriceType
}

// 输入的执行宣告
type tCThostFtdcInputExecOrderField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 合约代码
	InstrumentID tTThostFtdcInstrumentIDType
	// 执行宣告引用
	ExecOrderRef tTThostFtdcOrderRefType
	// 用户代码
	UserID tTThostFtdcUserIDType
	// 数量
	Volume tTThostFtdcVolumeType
	// 请求编号
	RequestID tTThostFtdcRequestIDType
	// 业务单元
	BusinessUnit tTThostFtdcBusinessUnitType
	// 开平标志
	OffsetFlag tTThostFtdcOffsetFlagType
	// 投机套保标志
	HedgeFlag tTThostFtdcHedgeFlagType
	// 执行类型
	ActionType tTThostFtdcActionTypeType
	// 保留头寸申请的持仓方向
	PosiDirection tTThostFtdcPosiDirectionType
	// 期权行权后是否保留期货头寸的标记,该字段已废弃
	ReservePositionFlag tTThostFtdcExecOrderPositionFlagType
	// 期权行权后生成的头寸是否自动平仓
	CloseFlag tTThostFtdcExecOrderCloseFlagType
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
	// 投资单元代码
	InvestUnitID tTThostFtdcInvestUnitIDType
	// 资金账号
	AccountID tTThostFtdcAccountIDType
	// 币种代码
	CurrencyID tTThostFtdcCurrencyIDType
	// 交易编码
	ClientID tTThostFtdcClientIDType
	// IP地址
	IPAddress tTThostFtdcIPAddressType
	// Mac地址
	MacAddress tTThostFtdcMacAddressType
}

// 输入执行宣告操作
type tCThostFtdcInputExecOrderActionField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 执行宣告操作引用
	ExecOrderActionRef tTThostFtdcOrderActionRefType
	// 执行宣告引用
	ExecOrderRef tTThostFtdcOrderRefType
	// 请求编号
	RequestID tTThostFtdcRequestIDType
	// 前置编号
	FrontID tTThostFtdcFrontIDType
	// 会话编号
	SessionID tTThostFtdcSessionIDType
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
	// 执行宣告操作编号
	ExecOrderSysID tTThostFtdcExecOrderSysIDType
	// 操作标志
	ActionFlag tTThostFtdcActionFlagType
	// 用户代码
	UserID tTThostFtdcUserIDType
	// 合约代码
	InstrumentID tTThostFtdcInstrumentIDType
	// 投资单元代码
	InvestUnitID tTThostFtdcInvestUnitIDType
	// IP地址
	IPAddress tTThostFtdcIPAddressType
	// Mac地址
	MacAddress tTThostFtdcMacAddressType
}

// 执行宣告
type tCThostFtdcExecOrderField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 合约代码
	InstrumentID tTThostFtdcInstrumentIDType
	// 执行宣告引用
	ExecOrderRef tTThostFtdcOrderRefType
	// 用户代码
	UserID tTThostFtdcUserIDType
	// 数量
	Volume tTThostFtdcVolumeType
	// 请求编号
	RequestID tTThostFtdcRequestIDType
	// 业务单元
	BusinessUnit tTThostFtdcBusinessUnitType
	// 开平标志
	OffsetFlag tTThostFtdcOffsetFlagType
	// 投机套保标志
	HedgeFlag tTThostFtdcHedgeFlagType
	// 执行类型
	ActionType tTThostFtdcActionTypeType
	// 保留头寸申请的持仓方向
	PosiDirection tTThostFtdcPosiDirectionType
	// 期权行权后是否保留期货头寸的标记,该字段已废弃
	ReservePositionFlag tTThostFtdcExecOrderPositionFlagType
	// 期权行权后生成的头寸是否自动平仓
	CloseFlag tTThostFtdcExecOrderCloseFlagType
	// 本地执行宣告编号
	ExecOrderLocalID tTThostFtdcOrderLocalIDType
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
	// 会员代码
	ParticipantID tTThostFtdcParticipantIDType
	// 客户代码
	ClientID tTThostFtdcClientIDType
	// 合约在交易所的代码
	ExchangeInstID tTThostFtdcExchangeInstIDType
	// 交易所交易员代码
	TraderID tTThostFtdcTraderIDType
	// 安装编号
	InstallID tTThostFtdcInstallIDType
	// 执行宣告提交状态
	OrderSubmitStatus tTThostFtdcOrderSubmitStatusType
	// 报单提示序号
	NotifySequence tTThostFtdcSequenceNoType
	// 交易日
	TradingDay tTThostFtdcDateType
	// 结算编号
	SettlementID tTThostFtdcSettlementIDType
	// 执行宣告编号
	ExecOrderSysID tTThostFtdcExecOrderSysIDType
	// 报单日期
	InsertDate tTThostFtdcDateType
	// 插入时间
	InsertTime tTThostFtdcTimeType
	// 撤销时间
	CancelTime tTThostFtdcTimeType
	// 执行结果
	ExecResult tTThostFtdcExecResultType
	// 结算会员编号
	ClearingPartID tTThostFtdcParticipantIDType
	// 序号
	SequenceNo tTThostFtdcSequenceNoType
	// 前置编号
	FrontID tTThostFtdcFrontIDType
	// 会话编号
	SessionID tTThostFtdcSessionIDType
	// 用户端产品信息
	UserProductInfo tTThostFtdcProductInfoType
	// 状态信息
	StatusMsg tTThostFtdcErrorMsgType
	// 操作用户代码
	ActiveUserID tTThostFtdcUserIDType
	// 经纪公司报单编号
	BrokerExecOrderSeq tTThostFtdcSequenceNoType
	// 营业部编号
	BranchID tTThostFtdcBranchIDType
	// 投资单元代码
	InvestUnitID tTThostFtdcInvestUnitIDType
	// 资金账号
	AccountID tTThostFtdcAccountIDType
	// 币种代码
	CurrencyID tTThostFtdcCurrencyIDType
	// IP地址
	IPAddress tTThostFtdcIPAddressType
	// Mac地址
	MacAddress tTThostFtdcMacAddressType
}

// 执行宣告操作
type tCThostFtdcExecOrderActionField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 执行宣告操作引用
	ExecOrderActionRef tTThostFtdcOrderActionRefType
	// 执行宣告引用
	ExecOrderRef tTThostFtdcOrderRefType
	// 请求编号
	RequestID tTThostFtdcRequestIDType
	// 前置编号
	FrontID tTThostFtdcFrontIDType
	// 会话编号
	SessionID tTThostFtdcSessionIDType
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
	// 执行宣告操作编号
	ExecOrderSysID tTThostFtdcExecOrderSysIDType
	// 操作标志
	ActionFlag tTThostFtdcActionFlagType
	// 操作日期
	ActionDate tTThostFtdcDateType
	// 操作时间
	ActionTime tTThostFtdcTimeType
	// 交易所交易员代码
	TraderID tTThostFtdcTraderIDType
	// 安装编号
	InstallID tTThostFtdcInstallIDType
	// 本地执行宣告编号
	ExecOrderLocalID tTThostFtdcOrderLocalIDType
	// 操作本地编号
	ActionLocalID tTThostFtdcOrderLocalIDType
	// 会员代码
	ParticipantID tTThostFtdcParticipantIDType
	// 客户代码
	ClientID tTThostFtdcClientIDType
	// 业务单元
	BusinessUnit tTThostFtdcBusinessUnitType
	// 报单操作状态
	OrderActionStatus tTThostFtdcOrderActionStatusType
	// 用户代码
	UserID tTThostFtdcUserIDType
	// 执行类型
	ActionType tTThostFtdcActionTypeType
	// 状态信息
	StatusMsg tTThostFtdcErrorMsgType
	// 合约代码
	InstrumentID tTThostFtdcInstrumentIDType
	// 营业部编号
	BranchID tTThostFtdcBranchIDType
	// 投资单元代码
	InvestUnitID tTThostFtdcInvestUnitIDType
	// IP地址
	IPAddress tTThostFtdcIPAddressType
	// Mac地址
	MacAddress tTThostFtdcMacAddressType
}

// 执行宣告查询
type tCThostFtdcQryExecOrderField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 合约代码
	InstrumentID tTThostFtdcInstrumentIDType
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
	// 执行宣告编号
	ExecOrderSysID tTThostFtdcExecOrderSysIDType
	// 开始时间
	InsertTimeStart tTThostFtdcTimeType
	// 结束时间
	InsertTimeEnd tTThostFtdcTimeType
}

// 交易所执行宣告信息
type tCThostFtdcExchangeExecOrderField struct {
	// 数量
	Volume tTThostFtdcVolumeType
	// 请求编号
	RequestID tTThostFtdcRequestIDType
	// 业务单元
	BusinessUnit tTThostFtdcBusinessUnitType
	// 开平标志
	OffsetFlag tTThostFtdcOffsetFlagType
	// 投机套保标志
	HedgeFlag tTThostFtdcHedgeFlagType
	// 执行类型
	ActionType tTThostFtdcActionTypeType
	// 保留头寸申请的持仓方向
	PosiDirection tTThostFtdcPosiDirectionType
	// 期权行权后是否保留期货头寸的标记,该字段已废弃
	ReservePositionFlag tTThostFtdcExecOrderPositionFlagType
	// 期权行权后生成的头寸是否自动平仓
	CloseFlag tTThostFtdcExecOrderCloseFlagType
	// 本地执行宣告编号
	ExecOrderLocalID tTThostFtdcOrderLocalIDType
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
	// 会员代码
	ParticipantID tTThostFtdcParticipantIDType
	// 客户代码
	ClientID tTThostFtdcClientIDType
	// 合约在交易所的代码
	ExchangeInstID tTThostFtdcExchangeInstIDType
	// 交易所交易员代码
	TraderID tTThostFtdcTraderIDType
	// 安装编号
	InstallID tTThostFtdcInstallIDType
	// 执行宣告提交状态
	OrderSubmitStatus tTThostFtdcOrderSubmitStatusType
	// 报单提示序号
	NotifySequence tTThostFtdcSequenceNoType
	// 交易日
	TradingDay tTThostFtdcDateType
	// 结算编号
	SettlementID tTThostFtdcSettlementIDType
	// 执行宣告编号
	ExecOrderSysID tTThostFtdcExecOrderSysIDType
	// 报单日期
	InsertDate tTThostFtdcDateType
	// 插入时间
	InsertTime tTThostFtdcTimeType
	// 撤销时间
	CancelTime tTThostFtdcTimeType
	// 执行结果
	ExecResult tTThostFtdcExecResultType
	// 结算会员编号
	ClearingPartID tTThostFtdcParticipantIDType
	// 序号
	SequenceNo tTThostFtdcSequenceNoType
	// 营业部编号
	BranchID tTThostFtdcBranchIDType
	// IP地址
	IPAddress tTThostFtdcIPAddressType
	// Mac地址
	MacAddress tTThostFtdcMacAddressType
}

// 交易所执行宣告查询
type tCThostFtdcQryExchangeExecOrderField struct {
	// 会员代码
	ParticipantID tTThostFtdcParticipantIDType
	// 客户代码
	ClientID tTThostFtdcClientIDType
	// 合约在交易所的代码
	ExchangeInstID tTThostFtdcExchangeInstIDType
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
	// 交易所交易员代码
	TraderID tTThostFtdcTraderIDType
}

// 执行宣告操作查询
type tCThostFtdcQryExecOrderActionField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
}

// 交易所执行宣告操作
type tCThostFtdcExchangeExecOrderActionField struct {
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
	// 执行宣告操作编号
	ExecOrderSysID tTThostFtdcExecOrderSysIDType
	// 操作标志
	ActionFlag tTThostFtdcActionFlagType
	// 操作日期
	ActionDate tTThostFtdcDateType
	// 操作时间
	ActionTime tTThostFtdcTimeType
	// 交易所交易员代码
	TraderID tTThostFtdcTraderIDType
	// 安装编号
	InstallID tTThostFtdcInstallIDType
	// 本地执行宣告编号
	ExecOrderLocalID tTThostFtdcOrderLocalIDType
	// 操作本地编号
	ActionLocalID tTThostFtdcOrderLocalIDType
	// 会员代码
	ParticipantID tTThostFtdcParticipantIDType
	// 客户代码
	ClientID tTThostFtdcClientIDType
	// 业务单元
	BusinessUnit tTThostFtdcBusinessUnitType
	// 报单操作状态
	OrderActionStatus tTThostFtdcOrderActionStatusType
	// 用户代码
	UserID tTThostFtdcUserIDType
	// 执行类型
	ActionType tTThostFtdcActionTypeType
	// 营业部编号
	BranchID tTThostFtdcBranchIDType
	// IP地址
	IPAddress tTThostFtdcIPAddressType
	// Mac地址
	MacAddress tTThostFtdcMacAddressType
	// 合约在交易所的代码
	ExchangeInstID tTThostFtdcExchangeInstIDType
	// 数量
	Volume tTThostFtdcVolumeType
}

// 交易所执行宣告操作查询
type tCThostFtdcQryExchangeExecOrderActionField struct {
	// 会员代码
	ParticipantID tTThostFtdcParticipantIDType
	// 客户代码
	ClientID tTThostFtdcClientIDType
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
	// 交易所交易员代码
	TraderID tTThostFtdcTraderIDType
}

// 错误执行宣告
type tCThostFtdcErrExecOrderField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 合约代码
	InstrumentID tTThostFtdcInstrumentIDType
	// 执行宣告引用
	ExecOrderRef tTThostFtdcOrderRefType
	// 用户代码
	UserID tTThostFtdcUserIDType
	// 数量
	Volume tTThostFtdcVolumeType
	// 请求编号
	RequestID tTThostFtdcRequestIDType
	// 业务单元
	BusinessUnit tTThostFtdcBusinessUnitType
	// 开平标志
	OffsetFlag tTThostFtdcOffsetFlagType
	// 投机套保标志
	HedgeFlag tTThostFtdcHedgeFlagType
	// 执行类型
	ActionType tTThostFtdcActionTypeType
	// 保留头寸申请的持仓方向
	PosiDirection tTThostFtdcPosiDirectionType
	// 期权行权后是否保留期货头寸的标记,该字段已废弃
	ReservePositionFlag tTThostFtdcExecOrderPositionFlagType
	// 期权行权后生成的头寸是否自动平仓
	CloseFlag tTThostFtdcExecOrderCloseFlagType
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
	// 投资单元代码
	InvestUnitID tTThostFtdcInvestUnitIDType
	// 资金账号
	AccountID tTThostFtdcAccountIDType
	// 币种代码
	CurrencyID tTThostFtdcCurrencyIDType
	// 交易编码
	ClientID tTThostFtdcClientIDType
	// IP地址
	IPAddress tTThostFtdcIPAddressType
	// Mac地址
	MacAddress tTThostFtdcMacAddressType
	// 错误代码
	ErrorID tTThostFtdcErrorIDType
	// 错误信息
	ErrorMsg tTThostFtdcErrorMsgType
}

// 查询错误执行宣告
type tCThostFtdcQryErrExecOrderField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
}

// 错误执行宣告操作
type tCThostFtdcErrExecOrderActionField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 执行宣告操作引用
	ExecOrderActionRef tTThostFtdcOrderActionRefType
	// 执行宣告引用
	ExecOrderRef tTThostFtdcOrderRefType
	// 请求编号
	RequestID tTThostFtdcRequestIDType
	// 前置编号
	FrontID tTThostFtdcFrontIDType
	// 会话编号
	SessionID tTThostFtdcSessionIDType
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
	// 执行宣告操作编号
	ExecOrderSysID tTThostFtdcExecOrderSysIDType
	// 操作标志
	ActionFlag tTThostFtdcActionFlagType
	// 用户代码
	UserID tTThostFtdcUserIDType
	// 合约代码
	InstrumentID tTThostFtdcInstrumentIDType
	// 投资单元代码
	InvestUnitID tTThostFtdcInvestUnitIDType
	// IP地址
	IPAddress tTThostFtdcIPAddressType
	// Mac地址
	MacAddress tTThostFtdcMacAddressType
	// 错误代码
	ErrorID tTThostFtdcErrorIDType
	// 错误信息
	ErrorMsg tTThostFtdcErrorMsgType
}

// 查询错误执行宣告操作
type tCThostFtdcQryErrExecOrderActionField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
}

// 投资者期权合约交易权限
type tCThostFtdcOptionInstrTradingRightField struct {
	// 合约代码
	InstrumentID tTThostFtdcInstrumentIDType
	// 投资者范围
	InvestorRange tTThostFtdcInvestorRangeType
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 买卖方向
	Direction tTThostFtdcDirectionType
	// 交易权限
	TradingRight tTThostFtdcTradingRightType
}

// 查询期权合约交易权限
type tCThostFtdcQryOptionInstrTradingRightField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 合约代码
	InstrumentID tTThostFtdcInstrumentIDType
	// 买卖方向
	Direction tTThostFtdcDirectionType
}

// 输入的询价
type tCThostFtdcInputForQuoteField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 合约代码
	InstrumentID tTThostFtdcInstrumentIDType
	// 询价引用
	ForQuoteRef tTThostFtdcOrderRefType
	// 用户代码
	UserID tTThostFtdcUserIDType
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
	// 投资单元代码
	InvestUnitID tTThostFtdcInvestUnitIDType
	// IP地址
	IPAddress tTThostFtdcIPAddressType
	// Mac地址
	MacAddress tTThostFtdcMacAddressType
}

// 询价
type tCThostFtdcForQuoteField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 合约代码
	InstrumentID tTThostFtdcInstrumentIDType
	// 询价引用
	ForQuoteRef tTThostFtdcOrderRefType
	// 用户代码
	UserID tTThostFtdcUserIDType
	// 本地询价编号
	ForQuoteLocalID tTThostFtdcOrderLocalIDType
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
	// 会员代码
	ParticipantID tTThostFtdcParticipantIDType
	// 客户代码
	ClientID tTThostFtdcClientIDType
	// 合约在交易所的代码
	ExchangeInstID tTThostFtdcExchangeInstIDType
	// 交易所交易员代码
	TraderID tTThostFtdcTraderIDType
	// 安装编号
	InstallID tTThostFtdcInstallIDType
	// 报单日期
	InsertDate tTThostFtdcDateType
	// 插入时间
	InsertTime tTThostFtdcTimeType
	// 询价状态
	ForQuoteStatus tTThostFtdcForQuoteStatusType
	// 前置编号
	FrontID tTThostFtdcFrontIDType
	// 会话编号
	SessionID tTThostFtdcSessionIDType
	// 状态信息
	StatusMsg tTThostFtdcErrorMsgType
	// 操作用户代码
	ActiveUserID tTThostFtdcUserIDType
	// 经纪公司询价编号
	BrokerForQutoSeq tTThostFtdcSequenceNoType
	// 投资单元代码
	InvestUnitID tTThostFtdcInvestUnitIDType
	// IP地址
	IPAddress tTThostFtdcIPAddressType
	// Mac地址
	MacAddress tTThostFtdcMacAddressType
}

// 询价查询
type tCThostFtdcQryForQuoteField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 合约代码
	InstrumentID tTThostFtdcInstrumentIDType
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
	// 开始时间
	InsertTimeStart tTThostFtdcTimeType
	// 结束时间
	InsertTimeEnd tTThostFtdcTimeType
	// 投资单元代码
	InvestUnitID tTThostFtdcInvestUnitIDType
}

// 交易所询价信息
type tCThostFtdcExchangeForQuoteField struct {
	// 本地询价编号
	ForQuoteLocalID tTThostFtdcOrderLocalIDType
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
	// 会员代码
	ParticipantID tTThostFtdcParticipantIDType
	// 客户代码
	ClientID tTThostFtdcClientIDType
	// 合约在交易所的代码
	ExchangeInstID tTThostFtdcExchangeInstIDType
	// 交易所交易员代码
	TraderID tTThostFtdcTraderIDType
	// 安装编号
	InstallID tTThostFtdcInstallIDType
	// 报单日期
	InsertDate tTThostFtdcDateType
	// 插入时间
	InsertTime tTThostFtdcTimeType
	// 询价状态
	ForQuoteStatus tTThostFtdcForQuoteStatusType
	// IP地址
	IPAddress tTThostFtdcIPAddressType
	// Mac地址
	MacAddress tTThostFtdcMacAddressType
}

// 交易所询价查询
type tCThostFtdcQryExchangeForQuoteField struct {
	// 会员代码
	ParticipantID tTThostFtdcParticipantIDType
	// 客户代码
	ClientID tTThostFtdcClientIDType
	// 合约在交易所的代码
	ExchangeInstID tTThostFtdcExchangeInstIDType
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
	// 交易所交易员代码
	TraderID tTThostFtdcTraderIDType
}

// 输入的报价
type tCThostFtdcInputQuoteField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 合约代码
	InstrumentID tTThostFtdcInstrumentIDType
	// 报价引用
	QuoteRef tTThostFtdcOrderRefType
	// 用户代码
	UserID tTThostFtdcUserIDType
	// 卖价格
	AskPrice tTThostFtdcPriceType
	// 买价格
	BidPrice tTThostFtdcPriceType
	// 卖数量
	AskVolume tTThostFtdcVolumeType
	// 买数量
	BidVolume tTThostFtdcVolumeType
	// 请求编号
	RequestID tTThostFtdcRequestIDType
	// 业务单元
	BusinessUnit tTThostFtdcBusinessUnitType
	// 卖开平标志
	AskOffsetFlag tTThostFtdcOffsetFlagType
	// 买开平标志
	BidOffsetFlag tTThostFtdcOffsetFlagType
	// 卖投机套保标志
	AskHedgeFlag tTThostFtdcHedgeFlagType
	// 买投机套保标志
	BidHedgeFlag tTThostFtdcHedgeFlagType
	// 衍生卖报单引用
	AskOrderRef tTThostFtdcOrderRefType
	// 衍生买报单引用
	BidOrderRef tTThostFtdcOrderRefType
	// 应价编号
	ForQuoteSysID tTThostFtdcOrderSysIDType
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
	// 投资单元代码
	InvestUnitID tTThostFtdcInvestUnitIDType
	// 交易编码
	ClientID tTThostFtdcClientIDType
	// IP地址
	IPAddress tTThostFtdcIPAddressType
	// Mac地址
	MacAddress tTThostFtdcMacAddressType
}

// 输入报价操作
type tCThostFtdcInputQuoteActionField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 报价操作引用
	QuoteActionRef tTThostFtdcOrderActionRefType
	// 报价引用
	QuoteRef tTThostFtdcOrderRefType
	// 请求编号
	RequestID tTThostFtdcRequestIDType
	// 前置编号
	FrontID tTThostFtdcFrontIDType
	// 会话编号
	SessionID tTThostFtdcSessionIDType
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
	// 报价操作编号
	QuoteSysID tTThostFtdcOrderSysIDType
	// 操作标志
	ActionFlag tTThostFtdcActionFlagType
	// 用户代码
	UserID tTThostFtdcUserIDType
	// 合约代码
	InstrumentID tTThostFtdcInstrumentIDType
	// 投资单元代码
	InvestUnitID tTThostFtdcInvestUnitIDType
	// 交易编码
	ClientID tTThostFtdcClientIDType
	// IP地址
	IPAddress tTThostFtdcIPAddressType
	// Mac地址
	MacAddress tTThostFtdcMacAddressType
}

// 报价
type tCThostFtdcQuoteField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 合约代码
	InstrumentID tTThostFtdcInstrumentIDType
	// 报价引用
	QuoteRef tTThostFtdcOrderRefType
	// 用户代码
	UserID tTThostFtdcUserIDType
	// 卖价格
	AskPrice tTThostFtdcPriceType
	// 买价格
	BidPrice tTThostFtdcPriceType
	// 卖数量
	AskVolume tTThostFtdcVolumeType
	// 买数量
	BidVolume tTThostFtdcVolumeType
	// 请求编号
	RequestID tTThostFtdcRequestIDType
	// 业务单元
	BusinessUnit tTThostFtdcBusinessUnitType
	// 卖开平标志
	AskOffsetFlag tTThostFtdcOffsetFlagType
	// 买开平标志
	BidOffsetFlag tTThostFtdcOffsetFlagType
	// 卖投机套保标志
	AskHedgeFlag tTThostFtdcHedgeFlagType
	// 买投机套保标志
	BidHedgeFlag tTThostFtdcHedgeFlagType
	// 本地报价编号
	QuoteLocalID tTThostFtdcOrderLocalIDType
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
	// 会员代码
	ParticipantID tTThostFtdcParticipantIDType
	// 客户代码
	ClientID tTThostFtdcClientIDType
	// 合约在交易所的代码
	ExchangeInstID tTThostFtdcExchangeInstIDType
	// 交易所交易员代码
	TraderID tTThostFtdcTraderIDType
	// 安装编号
	InstallID tTThostFtdcInstallIDType
	// 报价提示序号
	NotifySequence tTThostFtdcSequenceNoType
	// 报价提交状态
	OrderSubmitStatus tTThostFtdcOrderSubmitStatusType
	// 交易日
	TradingDay tTThostFtdcDateType
	// 结算编号
	SettlementID tTThostFtdcSettlementIDType
	// 报价编号
	QuoteSysID tTThostFtdcOrderSysIDType
	// 报单日期
	InsertDate tTThostFtdcDateType
	// 插入时间
	InsertTime tTThostFtdcTimeType
	// 撤销时间
	CancelTime tTThostFtdcTimeType
	// 报价状态
	QuoteStatus tTThostFtdcOrderStatusType
	// 结算会员编号
	ClearingPartID tTThostFtdcParticipantIDType
	// 序号
	SequenceNo tTThostFtdcSequenceNoType
	// 卖方报单编号
	AskOrderSysID tTThostFtdcOrderSysIDType
	// 买方报单编号
	BidOrderSysID tTThostFtdcOrderSysIDType
	// 前置编号
	FrontID tTThostFtdcFrontIDType
	// 会话编号
	SessionID tTThostFtdcSessionIDType
	// 用户端产品信息
	UserProductInfo tTThostFtdcProductInfoType
	// 状态信息
	StatusMsg tTThostFtdcErrorMsgType
	// 操作用户代码
	ActiveUserID tTThostFtdcUserIDType
	// 经纪公司报价编号
	BrokerQuoteSeq tTThostFtdcSequenceNoType
	// 衍生卖报单引用
	AskOrderRef tTThostFtdcOrderRefType
	// 衍生买报单引用
	BidOrderRef tTThostFtdcOrderRefType
	// 应价编号
	ForQuoteSysID tTThostFtdcOrderSysIDType
	// 营业部编号
	BranchID tTThostFtdcBranchIDType
	// 投资单元代码
	InvestUnitID tTThostFtdcInvestUnitIDType
	// 资金账号
	AccountID tTThostFtdcAccountIDType
	// 币种代码
	CurrencyID tTThostFtdcCurrencyIDType
	// IP地址
	IPAddress tTThostFtdcIPAddressType
	// Mac地址
	MacAddress tTThostFtdcMacAddressType
}

// 报价操作
type tCThostFtdcQuoteActionField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 报价操作引用
	QuoteActionRef tTThostFtdcOrderActionRefType
	// 报价引用
	QuoteRef tTThostFtdcOrderRefType
	// 请求编号
	RequestID tTThostFtdcRequestIDType
	// 前置编号
	FrontID tTThostFtdcFrontIDType
	// 会话编号
	SessionID tTThostFtdcSessionIDType
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
	// 报价操作编号
	QuoteSysID tTThostFtdcOrderSysIDType
	// 操作标志
	ActionFlag tTThostFtdcActionFlagType
	// 操作日期
	ActionDate tTThostFtdcDateType
	// 操作时间
	ActionTime tTThostFtdcTimeType
	// 交易所交易员代码
	TraderID tTThostFtdcTraderIDType
	// 安装编号
	InstallID tTThostFtdcInstallIDType
	// 本地报价编号
	QuoteLocalID tTThostFtdcOrderLocalIDType
	// 操作本地编号
	ActionLocalID tTThostFtdcOrderLocalIDType
	// 会员代码
	ParticipantID tTThostFtdcParticipantIDType
	// 客户代码
	ClientID tTThostFtdcClientIDType
	// 业务单元
	BusinessUnit tTThostFtdcBusinessUnitType
	// 报单操作状态
	OrderActionStatus tTThostFtdcOrderActionStatusType
	// 用户代码
	UserID tTThostFtdcUserIDType
	// 状态信息
	StatusMsg tTThostFtdcErrorMsgType
	// 合约代码
	InstrumentID tTThostFtdcInstrumentIDType
	// 营业部编号
	BranchID tTThostFtdcBranchIDType
	// 投资单元代码
	InvestUnitID tTThostFtdcInvestUnitIDType
	// IP地址
	IPAddress tTThostFtdcIPAddressType
	// Mac地址
	MacAddress tTThostFtdcMacAddressType
}

// 报价查询
type tCThostFtdcQryQuoteField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 合约代码
	InstrumentID tTThostFtdcInstrumentIDType
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
	// 报价编号
	QuoteSysID tTThostFtdcOrderSysIDType
	// 开始时间
	InsertTimeStart tTThostFtdcTimeType
	// 结束时间
	InsertTimeEnd tTThostFtdcTimeType
	// 投资单元代码
	InvestUnitID tTThostFtdcInvestUnitIDType
}

// 交易所报价信息
type tCThostFtdcExchangeQuoteField struct {
	// 卖价格
	AskPrice tTThostFtdcPriceType
	// 买价格
	BidPrice tTThostFtdcPriceType
	// 卖数量
	AskVolume tTThostFtdcVolumeType
	// 买数量
	BidVolume tTThostFtdcVolumeType
	// 请求编号
	RequestID tTThostFtdcRequestIDType
	// 业务单元
	BusinessUnit tTThostFtdcBusinessUnitType
	// 卖开平标志
	AskOffsetFlag tTThostFtdcOffsetFlagType
	// 买开平标志
	BidOffsetFlag tTThostFtdcOffsetFlagType
	// 卖投机套保标志
	AskHedgeFlag tTThostFtdcHedgeFlagType
	// 买投机套保标志
	BidHedgeFlag tTThostFtdcHedgeFlagType
	// 本地报价编号
	QuoteLocalID tTThostFtdcOrderLocalIDType
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
	// 会员代码
	ParticipantID tTThostFtdcParticipantIDType
	// 客户代码
	ClientID tTThostFtdcClientIDType
	// 合约在交易所的代码
	ExchangeInstID tTThostFtdcExchangeInstIDType
	// 交易所交易员代码
	TraderID tTThostFtdcTraderIDType
	// 安装编号
	InstallID tTThostFtdcInstallIDType
	// 报价提示序号
	NotifySequence tTThostFtdcSequenceNoType
	// 报价提交状态
	OrderSubmitStatus tTThostFtdcOrderSubmitStatusType
	// 交易日
	TradingDay tTThostFtdcDateType
	// 结算编号
	SettlementID tTThostFtdcSettlementIDType
	// 报价编号
	QuoteSysID tTThostFtdcOrderSysIDType
	// 报单日期
	InsertDate tTThostFtdcDateType
	// 插入时间
	InsertTime tTThostFtdcTimeType
	// 撤销时间
	CancelTime tTThostFtdcTimeType
	// 报价状态
	QuoteStatus tTThostFtdcOrderStatusType
	// 结算会员编号
	ClearingPartID tTThostFtdcParticipantIDType
	// 序号
	SequenceNo tTThostFtdcSequenceNoType
	// 卖方报单编号
	AskOrderSysID tTThostFtdcOrderSysIDType
	// 买方报单编号
	BidOrderSysID tTThostFtdcOrderSysIDType
	// 应价编号
	ForQuoteSysID tTThostFtdcOrderSysIDType
	// 营业部编号
	BranchID tTThostFtdcBranchIDType
	// IP地址
	IPAddress tTThostFtdcIPAddressType
	// Mac地址
	MacAddress tTThostFtdcMacAddressType
}

// 交易所报价查询
type tCThostFtdcQryExchangeQuoteField struct {
	// 会员代码
	ParticipantID tTThostFtdcParticipantIDType
	// 客户代码
	ClientID tTThostFtdcClientIDType
	// 合约在交易所的代码
	ExchangeInstID tTThostFtdcExchangeInstIDType
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
	// 交易所交易员代码
	TraderID tTThostFtdcTraderIDType
}

// 报价操作查询
type tCThostFtdcQryQuoteActionField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
}

// 交易所报价操作
type tCThostFtdcExchangeQuoteActionField struct {
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
	// 报价操作编号
	QuoteSysID tTThostFtdcOrderSysIDType
	// 操作标志
	ActionFlag tTThostFtdcActionFlagType
	// 操作日期
	ActionDate tTThostFtdcDateType
	// 操作时间
	ActionTime tTThostFtdcTimeType
	// 交易所交易员代码
	TraderID tTThostFtdcTraderIDType
	// 安装编号
	InstallID tTThostFtdcInstallIDType
	// 本地报价编号
	QuoteLocalID tTThostFtdcOrderLocalIDType
	// 操作本地编号
	ActionLocalID tTThostFtdcOrderLocalIDType
	// 会员代码
	ParticipantID tTThostFtdcParticipantIDType
	// 客户代码
	ClientID tTThostFtdcClientIDType
	// 业务单元
	BusinessUnit tTThostFtdcBusinessUnitType
	// 报单操作状态
	OrderActionStatus tTThostFtdcOrderActionStatusType
	// 用户代码
	UserID tTThostFtdcUserIDType
	// IP地址
	IPAddress tTThostFtdcIPAddressType
	// Mac地址
	MacAddress tTThostFtdcMacAddressType
}

// 交易所报价操作查询
type tCThostFtdcQryExchangeQuoteActionField struct {
	// 会员代码
	ParticipantID tTThostFtdcParticipantIDType
	// 客户代码
	ClientID tTThostFtdcClientIDType
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
	// 交易所交易员代码
	TraderID tTThostFtdcTraderIDType
}

// 期权合约delta值
type tCThostFtdcOptionInstrDeltaField struct {
	// 合约代码
	InstrumentID tTThostFtdcInstrumentIDType
	// 投资者范围
	InvestorRange tTThostFtdcInvestorRangeType
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// Delta值
	Delta tTThostFtdcRatioType
}

// 发给做市商的询价请求
type tCThostFtdcForQuoteRspField struct {
	// 交易日
	TradingDay tTThostFtdcDateType
	// 合约代码
	InstrumentID tTThostFtdcInstrumentIDType
	// 询价编号
	ForQuoteSysID tTThostFtdcOrderSysIDType
	// 询价时间
	ForQuoteTime tTThostFtdcTimeType
	// 业务日期
	ActionDay tTThostFtdcDateType
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
}

// 当前期权合约执行偏移值的详细内容
type tCThostFtdcStrikeOffsetField struct {
	// 合约代码
	InstrumentID tTThostFtdcInstrumentIDType
	// 投资者范围
	InvestorRange tTThostFtdcInvestorRangeType
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 执行偏移值
	Offset tTThostFtdcMoneyType
	// 执行偏移类型
	OffsetType tTThostFtdcStrikeOffsetTypeType
}

// 期权执行偏移值查询
type tCThostFtdcQryStrikeOffsetField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 合约代码
	InstrumentID tTThostFtdcInstrumentIDType
}

// 输入批量报单操作
type tCThostFtdcInputBatchOrderActionField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 报单操作引用
	OrderActionRef tTThostFtdcOrderActionRefType
	// 请求编号
	RequestID tTThostFtdcRequestIDType
	// 前置编号
	FrontID tTThostFtdcFrontIDType
	// 会话编号
	SessionID tTThostFtdcSessionIDType
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
	// 用户代码
	UserID tTThostFtdcUserIDType
	// 投资单元代码
	InvestUnitID tTThostFtdcInvestUnitIDType
	// IP地址
	IPAddress tTThostFtdcIPAddressType
	// Mac地址
	MacAddress tTThostFtdcMacAddressType
}

// 批量报单操作
type tCThostFtdcBatchOrderActionField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 报单操作引用
	OrderActionRef tTThostFtdcOrderActionRefType
	// 请求编号
	RequestID tTThostFtdcRequestIDType
	// 前置编号
	FrontID tTThostFtdcFrontIDType
	// 会话编号
	SessionID tTThostFtdcSessionIDType
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
	// 操作日期
	ActionDate tTThostFtdcDateType
	// 操作时间
	ActionTime tTThostFtdcTimeType
	// 交易所交易员代码
	TraderID tTThostFtdcTraderIDType
	// 安装编号
	InstallID tTThostFtdcInstallIDType
	// 操作本地编号
	ActionLocalID tTThostFtdcOrderLocalIDType
	// 会员代码
	ParticipantID tTThostFtdcParticipantIDType
	// 客户代码
	ClientID tTThostFtdcClientIDType
	// 业务单元
	BusinessUnit tTThostFtdcBusinessUnitType
	// 报单操作状态
	OrderActionStatus tTThostFtdcOrderActionStatusType
	// 用户代码
	UserID tTThostFtdcUserIDType
	// 状态信息
	StatusMsg tTThostFtdcErrorMsgType
	// 投资单元代码
	InvestUnitID tTThostFtdcInvestUnitIDType
	// IP地址
	IPAddress tTThostFtdcIPAddressType
	// Mac地址
	MacAddress tTThostFtdcMacAddressType
}

// 交易所批量报单操作
type tCThostFtdcExchangeBatchOrderActionField struct {
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
	// 操作日期
	ActionDate tTThostFtdcDateType
	// 操作时间
	ActionTime tTThostFtdcTimeType
	// 交易所交易员代码
	TraderID tTThostFtdcTraderIDType
	// 安装编号
	InstallID tTThostFtdcInstallIDType
	// 操作本地编号
	ActionLocalID tTThostFtdcOrderLocalIDType
	// 会员代码
	ParticipantID tTThostFtdcParticipantIDType
	// 客户代码
	ClientID tTThostFtdcClientIDType
	// 业务单元
	BusinessUnit tTThostFtdcBusinessUnitType
	// 报单操作状态
	OrderActionStatus tTThostFtdcOrderActionStatusType
	// 用户代码
	UserID tTThostFtdcUserIDType
	// IP地址
	IPAddress tTThostFtdcIPAddressType
	// Mac地址
	MacAddress tTThostFtdcMacAddressType
}

// 查询批量报单操作
type tCThostFtdcQryBatchOrderActionField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
}

// 组合合约安全系数
type tCThostFtdcCombInstrumentGuardField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 合约代码
	InstrumentID tTThostFtdcInstrumentIDType
	//
	GuarantRatio tTThostFtdcRatioType
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
}

// 组合合约安全系数查询
type tCThostFtdcQryCombInstrumentGuardField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 合约代码
	InstrumentID tTThostFtdcInstrumentIDType
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
}

// 输入的申请组合
type tCThostFtdcInputCombActionField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 合约代码
	InstrumentID tTThostFtdcInstrumentIDType
	// 组合引用
	CombActionRef tTThostFtdcOrderRefType
	// 用户代码
	UserID tTThostFtdcUserIDType
	// 买卖方向
	Direction tTThostFtdcDirectionType
	// 数量
	Volume tTThostFtdcVolumeType
	// 组合指令方向
	CombDirection tTThostFtdcCombDirectionType
	// 投机套保标志
	HedgeFlag tTThostFtdcHedgeFlagType
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
	// IP地址
	IPAddress tTThostFtdcIPAddressType
	// Mac地址
	MacAddress tTThostFtdcMacAddressType
	// 投资单元代码
	InvestUnitID tTThostFtdcInvestUnitIDType
}

// 申请组合
type tCThostFtdcCombActionField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 合约代码
	InstrumentID tTThostFtdcInstrumentIDType
	// 组合引用
	CombActionRef tTThostFtdcOrderRefType
	// 用户代码
	UserID tTThostFtdcUserIDType
	// 买卖方向
	Direction tTThostFtdcDirectionType
	// 数量
	Volume tTThostFtdcVolumeType
	// 组合指令方向
	CombDirection tTThostFtdcCombDirectionType
	// 投机套保标志
	HedgeFlag tTThostFtdcHedgeFlagType
	// 本地申请组合编号
	ActionLocalID tTThostFtdcOrderLocalIDType
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
	// 会员代码
	ParticipantID tTThostFtdcParticipantIDType
	// 客户代码
	ClientID tTThostFtdcClientIDType
	// 合约在交易所的代码
	ExchangeInstID tTThostFtdcExchangeInstIDType
	// 交易所交易员代码
	TraderID tTThostFtdcTraderIDType
	// 安装编号
	InstallID tTThostFtdcInstallIDType
	// 组合状态
	ActionStatus tTThostFtdcOrderActionStatusType
	// 报单提示序号
	NotifySequence tTThostFtdcSequenceNoType
	// 交易日
	TradingDay tTThostFtdcDateType
	// 结算编号
	SettlementID tTThostFtdcSettlementIDType
	// 序号
	SequenceNo tTThostFtdcSequenceNoType
	// 前置编号
	FrontID tTThostFtdcFrontIDType
	// 会话编号
	SessionID tTThostFtdcSessionIDType
	// 用户端产品信息
	UserProductInfo tTThostFtdcProductInfoType
	// 状态信息
	StatusMsg tTThostFtdcErrorMsgType
	// IP地址
	IPAddress tTThostFtdcIPAddressType
	// Mac地址
	MacAddress tTThostFtdcMacAddressType
	// 组合编号
	ComTradeID tTThostFtdcTradeIDType
	// 营业部编号
	BranchID tTThostFtdcBranchIDType
	// 投资单元代码
	InvestUnitID tTThostFtdcInvestUnitIDType
}

// 申请组合查询
type tCThostFtdcQryCombActionField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 合约代码
	InstrumentID tTThostFtdcInstrumentIDType
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
	// 投资单元代码
	InvestUnitID tTThostFtdcInvestUnitIDType
}

// 交易所申请组合信息
type tCThostFtdcExchangeCombActionField struct {
	// 买卖方向
	Direction tTThostFtdcDirectionType
	// 数量
	Volume tTThostFtdcVolumeType
	// 组合指令方向
	CombDirection tTThostFtdcCombDirectionType
	// 投机套保标志
	HedgeFlag tTThostFtdcHedgeFlagType
	// 本地申请组合编号
	ActionLocalID tTThostFtdcOrderLocalIDType
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
	// 会员代码
	ParticipantID tTThostFtdcParticipantIDType
	// 客户代码
	ClientID tTThostFtdcClientIDType
	// 合约在交易所的代码
	ExchangeInstID tTThostFtdcExchangeInstIDType
	// 交易所交易员代码
	TraderID tTThostFtdcTraderIDType
	// 安装编号
	InstallID tTThostFtdcInstallIDType
	// 组合状态
	ActionStatus tTThostFtdcOrderActionStatusType
	// 报单提示序号
	NotifySequence tTThostFtdcSequenceNoType
	// 交易日
	TradingDay tTThostFtdcDateType
	// 结算编号
	SettlementID tTThostFtdcSettlementIDType
	// 序号
	SequenceNo tTThostFtdcSequenceNoType
	// IP地址
	IPAddress tTThostFtdcIPAddressType
	// Mac地址
	MacAddress tTThostFtdcMacAddressType
	// 组合编号
	ComTradeID tTThostFtdcTradeIDType
	// 营业部编号
	BranchID tTThostFtdcBranchIDType
}

// 交易所申请组合查询
type tCThostFtdcQryExchangeCombActionField struct {
	// 会员代码
	ParticipantID tTThostFtdcParticipantIDType
	// 客户代码
	ClientID tTThostFtdcClientIDType
	// 合约在交易所的代码
	ExchangeInstID tTThostFtdcExchangeInstIDType
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
	// 交易所交易员代码
	TraderID tTThostFtdcTraderIDType
}

// 产品报价汇率
type tCThostFtdcProductExchRateField struct {
	// 产品代码
	ProductID tTThostFtdcInstrumentIDType
	// 报价币种类型
	QuoteCurrencyID tTThostFtdcCurrencyIDType
	// 汇率
	ExchangeRate tTThostFtdcExchangeRateType
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
}

// 产品报价汇率查询
type tCThostFtdcQryProductExchRateField struct {
	// 产品代码
	ProductID tTThostFtdcInstrumentIDType
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
}

// 查询询价价差参数
type tCThostFtdcQryForQuoteParamField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 合约代码
	InstrumentID tTThostFtdcInstrumentIDType
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
}

// 询价价差参数
type tCThostFtdcForQuoteParamField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 合约代码
	InstrumentID tTThostFtdcInstrumentIDType
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
	// 最新价
	LastPrice tTThostFtdcPriceType
	// 价差
	PriceInterval tTThostFtdcPriceType
}

// 当前做市商期权合约手续费的详细内容
type tCThostFtdcMMOptionInstrCommRateField struct {
	// 合约代码
	InstrumentID tTThostFtdcInstrumentIDType
	// 投资者范围
	InvestorRange tTThostFtdcInvestorRangeType
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 开仓手续费率
	OpenRatioByMoney tTThostFtdcRatioType
	// 开仓手续费
	OpenRatioByVolume tTThostFtdcRatioType
	// 平仓手续费率
	CloseRatioByMoney tTThostFtdcRatioType
	// 平仓手续费
	CloseRatioByVolume tTThostFtdcRatioType
	// 平今手续费率
	CloseTodayRatioByMoney tTThostFtdcRatioType
	// 平今手续费
	CloseTodayRatioByVolume tTThostFtdcRatioType
	// 执行手续费率
	StrikeRatioByMoney tTThostFtdcRatioType
	// 执行手续费
	StrikeRatioByVolume tTThostFtdcRatioType
}

// 做市商期权手续费率查询
type tCThostFtdcQryMMOptionInstrCommRateField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 合约代码
	InstrumentID tTThostFtdcInstrumentIDType
}

// 做市商合约手续费率
type tCThostFtdcMMInstrumentCommissionRateField struct {
	// 合约代码
	InstrumentID tTThostFtdcInstrumentIDType
	// 投资者范围
	InvestorRange tTThostFtdcInvestorRangeType
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 开仓手续费率
	OpenRatioByMoney tTThostFtdcRatioType
	// 开仓手续费
	OpenRatioByVolume tTThostFtdcRatioType
	// 平仓手续费率
	CloseRatioByMoney tTThostFtdcRatioType
	// 平仓手续费
	CloseRatioByVolume tTThostFtdcRatioType
	// 平今手续费率
	CloseTodayRatioByMoney tTThostFtdcRatioType
	// 平今手续费
	CloseTodayRatioByVolume tTThostFtdcRatioType
}

// 查询做市商合约手续费率
type tCThostFtdcQryMMInstrumentCommissionRateField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 合约代码
	InstrumentID tTThostFtdcInstrumentIDType
}

// 当前报单手续费的详细内容
type tCThostFtdcInstrumentOrderCommRateField struct {
	// 合约代码
	InstrumentID tTThostFtdcInstrumentIDType
	// 投资者范围
	InvestorRange tTThostFtdcInvestorRangeType
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 投机套保标志
	HedgeFlag tTThostFtdcHedgeFlagType
	// 报单手续费
	OrderCommByVolume tTThostFtdcRatioType
	// 撤单手续费
	OrderActionCommByVolume tTThostFtdcRatioType
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
	// 投资单元代码
	InvestUnitID tTThostFtdcInvestUnitIDType
}

// 报单手续费率查询
type tCThostFtdcQryInstrumentOrderCommRateField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 合约代码
	InstrumentID tTThostFtdcInstrumentIDType
}

// 交易参数
type tCThostFtdcTradeParamField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 参数代码
	TradeParamID tTThostFtdcTradeParamIDType
	// 参数代码值
	TradeParamValue tTThostFtdcSettlementParamValueType
	// 备注
	Memo tTThostFtdcMemoType
}

// 合约保证金率调整
type tCThostFtdcInstrumentMarginRateULField struct {
	// 合约代码
	InstrumentID tTThostFtdcInstrumentIDType
	// 投资者范围
	InvestorRange tTThostFtdcInvestorRangeType
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 投机套保标志
	HedgeFlag tTThostFtdcHedgeFlagType
	// 多头保证金率
	LongMarginRatioByMoney tTThostFtdcRatioType
	// 多头保证金费
	LongMarginRatioByVolume tTThostFtdcMoneyType
	// 空头保证金率
	ShortMarginRatioByMoney tTThostFtdcRatioType
	// 空头保证金费
	ShortMarginRatioByVolume tTThostFtdcMoneyType
}

// 期货持仓限制参数
type tCThostFtdcFutureLimitPosiParamField struct {
	// 投资者范围
	InvestorRange tTThostFtdcInvestorRangeType
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 产品代码
	ProductID tTThostFtdcInstrumentIDType
	// 当日投机开仓数量限制
	SpecOpenVolume tTThostFtdcVolumeType
	// 当日套利开仓数量限制
	ArbiOpenVolume tTThostFtdcVolumeType
	// 当日投机+套利开仓数量限制
	OpenVolume tTThostFtdcVolumeType
}

// 禁止登录IP
type tCThostFtdcLoginForbiddenIPField struct {
	// IP地址
	IPAddress tTThostFtdcIPAddressType
}

// IP列表
type tCThostFtdcIPListField struct {
	// IP地址
	IPAddress tTThostFtdcIPAddressType
	// 是否白名单
	IsWhite tTThostFtdcBoolType
}

// 输入的期权自对冲
type tCThostFtdcInputOptionSelfCloseField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 合约代码
	InstrumentID tTThostFtdcInstrumentIDType
	// 期权自对冲引用
	OptionSelfCloseRef tTThostFtdcOrderRefType
	// 用户代码
	UserID tTThostFtdcUserIDType
	// 数量
	Volume tTThostFtdcVolumeType
	// 请求编号
	RequestID tTThostFtdcRequestIDType
	// 业务单元
	BusinessUnit tTThostFtdcBusinessUnitType
	// 投机套保标志
	HedgeFlag tTThostFtdcHedgeFlagType
	// 期权行权的头寸是否自对冲
	OptSelfCloseFlag tTThostFtdcOptSelfCloseFlagType
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
	// 投资单元代码
	InvestUnitID tTThostFtdcInvestUnitIDType
	// 资金账号
	AccountID tTThostFtdcAccountIDType
	// 币种代码
	CurrencyID tTThostFtdcCurrencyIDType
	// 交易编码
	ClientID tTThostFtdcClientIDType
	// IP地址
	IPAddress tTThostFtdcIPAddressType
	// Mac地址
	MacAddress tTThostFtdcMacAddressType
}

// 输入期权自对冲操作
type tCThostFtdcInputOptionSelfCloseActionField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 期权自对冲操作引用
	OptionSelfCloseActionRef tTThostFtdcOrderActionRefType
	// 期权自对冲引用
	OptionSelfCloseRef tTThostFtdcOrderRefType
	// 请求编号
	RequestID tTThostFtdcRequestIDType
	// 前置编号
	FrontID tTThostFtdcFrontIDType
	// 会话编号
	SessionID tTThostFtdcSessionIDType
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
	// 期权自对冲操作编号
	OptionSelfCloseSysID tTThostFtdcOrderSysIDType
	// 操作标志
	ActionFlag tTThostFtdcActionFlagType
	// 用户代码
	UserID tTThostFtdcUserIDType
	// 合约代码
	InstrumentID tTThostFtdcInstrumentIDType
	// 投资单元代码
	InvestUnitID tTThostFtdcInvestUnitIDType
	// IP地址
	IPAddress tTThostFtdcIPAddressType
	// Mac地址
	MacAddress tTThostFtdcMacAddressType
}

// 期权自对冲
type tCThostFtdcOptionSelfCloseField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 合约代码
	InstrumentID tTThostFtdcInstrumentIDType
	// 期权自对冲引用
	OptionSelfCloseRef tTThostFtdcOrderRefType
	// 用户代码
	UserID tTThostFtdcUserIDType
	// 数量
	Volume tTThostFtdcVolumeType
	// 请求编号
	RequestID tTThostFtdcRequestIDType
	// 业务单元
	BusinessUnit tTThostFtdcBusinessUnitType
	// 投机套保标志
	HedgeFlag tTThostFtdcHedgeFlagType
	// 期权行权的头寸是否自对冲
	OptSelfCloseFlag tTThostFtdcOptSelfCloseFlagType
	// 本地期权自对冲编号
	OptionSelfCloseLocalID tTThostFtdcOrderLocalIDType
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
	// 会员代码
	ParticipantID tTThostFtdcParticipantIDType
	// 客户代码
	ClientID tTThostFtdcClientIDType
	// 合约在交易所的代码
	ExchangeInstID tTThostFtdcExchangeInstIDType
	// 交易所交易员代码
	TraderID tTThostFtdcTraderIDType
	// 安装编号
	InstallID tTThostFtdcInstallIDType
	// 期权自对冲提交状态
	OrderSubmitStatus tTThostFtdcOrderSubmitStatusType
	// 报单提示序号
	NotifySequence tTThostFtdcSequenceNoType
	// 交易日
	TradingDay tTThostFtdcDateType
	// 结算编号
	SettlementID tTThostFtdcSettlementIDType
	// 期权自对冲编号
	OptionSelfCloseSysID tTThostFtdcOrderSysIDType
	// 报单日期
	InsertDate tTThostFtdcDateType
	// 插入时间
	InsertTime tTThostFtdcTimeType
	// 撤销时间
	CancelTime tTThostFtdcTimeType
	// 自对冲结果
	ExecResult tTThostFtdcExecResultType
	// 结算会员编号
	ClearingPartID tTThostFtdcParticipantIDType
	// 序号
	SequenceNo tTThostFtdcSequenceNoType
	// 前置编号
	FrontID tTThostFtdcFrontIDType
	// 会话编号
	SessionID tTThostFtdcSessionIDType
	// 用户端产品信息
	UserProductInfo tTThostFtdcProductInfoType
	// 状态信息
	StatusMsg tTThostFtdcErrorMsgType
	// 操作用户代码
	ActiveUserID tTThostFtdcUserIDType
	// 经纪公司报单编号
	BrokerOptionSelfCloseSeq tTThostFtdcSequenceNoType
	// 营业部编号
	BranchID tTThostFtdcBranchIDType
	// 投资单元代码
	InvestUnitID tTThostFtdcInvestUnitIDType
	// 资金账号
	AccountID tTThostFtdcAccountIDType
	// 币种代码
	CurrencyID tTThostFtdcCurrencyIDType
	// IP地址
	IPAddress tTThostFtdcIPAddressType
	// Mac地址
	MacAddress tTThostFtdcMacAddressType
}

// 期权自对冲操作
type tCThostFtdcOptionSelfCloseActionField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 期权自对冲操作引用
	OptionSelfCloseActionRef tTThostFtdcOrderActionRefType
	// 期权自对冲引用
	OptionSelfCloseRef tTThostFtdcOrderRefType
	// 请求编号
	RequestID tTThostFtdcRequestIDType
	// 前置编号
	FrontID tTThostFtdcFrontIDType
	// 会话编号
	SessionID tTThostFtdcSessionIDType
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
	// 期权自对冲操作编号
	OptionSelfCloseSysID tTThostFtdcOrderSysIDType
	// 操作标志
	ActionFlag tTThostFtdcActionFlagType
	// 操作日期
	ActionDate tTThostFtdcDateType
	// 操作时间
	ActionTime tTThostFtdcTimeType
	// 交易所交易员代码
	TraderID tTThostFtdcTraderIDType
	// 安装编号
	InstallID tTThostFtdcInstallIDType
	// 本地期权自对冲编号
	OptionSelfCloseLocalID tTThostFtdcOrderLocalIDType
	// 操作本地编号
	ActionLocalID tTThostFtdcOrderLocalIDType
	// 会员代码
	ParticipantID tTThostFtdcParticipantIDType
	// 客户代码
	ClientID tTThostFtdcClientIDType
	// 业务单元
	BusinessUnit tTThostFtdcBusinessUnitType
	// 报单操作状态
	OrderActionStatus tTThostFtdcOrderActionStatusType
	// 用户代码
	UserID tTThostFtdcUserIDType
	// 状态信息
	StatusMsg tTThostFtdcErrorMsgType
	// 合约代码
	InstrumentID tTThostFtdcInstrumentIDType
	// 营业部编号
	BranchID tTThostFtdcBranchIDType
	// 投资单元代码
	InvestUnitID tTThostFtdcInvestUnitIDType
	// IP地址
	IPAddress tTThostFtdcIPAddressType
	// Mac地址
	MacAddress tTThostFtdcMacAddressType
}

// 期权自对冲查询
type tCThostFtdcQryOptionSelfCloseField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 合约代码
	InstrumentID tTThostFtdcInstrumentIDType
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
	// 期权自对冲编号
	OptionSelfCloseSysID tTThostFtdcOrderSysIDType
	// 开始时间
	InsertTimeStart tTThostFtdcTimeType
	// 结束时间
	InsertTimeEnd tTThostFtdcTimeType
}

// 交易所期权自对冲信息
type tCThostFtdcExchangeOptionSelfCloseField struct {
	// 数量
	Volume tTThostFtdcVolumeType
	// 请求编号
	RequestID tTThostFtdcRequestIDType
	// 业务单元
	BusinessUnit tTThostFtdcBusinessUnitType
	// 投机套保标志
	HedgeFlag tTThostFtdcHedgeFlagType
	// 期权行权的头寸是否自对冲
	OptSelfCloseFlag tTThostFtdcOptSelfCloseFlagType
	// 本地期权自对冲编号
	OptionSelfCloseLocalID tTThostFtdcOrderLocalIDType
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
	// 会员代码
	ParticipantID tTThostFtdcParticipantIDType
	// 客户代码
	ClientID tTThostFtdcClientIDType
	// 合约在交易所的代码
	ExchangeInstID tTThostFtdcExchangeInstIDType
	// 交易所交易员代码
	TraderID tTThostFtdcTraderIDType
	// 安装编号
	InstallID tTThostFtdcInstallIDType
	// 期权自对冲提交状态
	OrderSubmitStatus tTThostFtdcOrderSubmitStatusType
	// 报单提示序号
	NotifySequence tTThostFtdcSequenceNoType
	// 交易日
	TradingDay tTThostFtdcDateType
	// 结算编号
	SettlementID tTThostFtdcSettlementIDType
	// 期权自对冲编号
	OptionSelfCloseSysID tTThostFtdcOrderSysIDType
	// 报单日期
	InsertDate tTThostFtdcDateType
	// 插入时间
	InsertTime tTThostFtdcTimeType
	// 撤销时间
	CancelTime tTThostFtdcTimeType
	// 自对冲结果
	ExecResult tTThostFtdcExecResultType
	// 结算会员编号
	ClearingPartID tTThostFtdcParticipantIDType
	// 序号
	SequenceNo tTThostFtdcSequenceNoType
	// 营业部编号
	BranchID tTThostFtdcBranchIDType
	// IP地址
	IPAddress tTThostFtdcIPAddressType
	// Mac地址
	MacAddress tTThostFtdcMacAddressType
}

// 期权自对冲操作查询
type tCThostFtdcQryOptionSelfCloseActionField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
}

// 交易所期权自对冲操作
type tCThostFtdcExchangeOptionSelfCloseActionField struct {
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
	// 期权自对冲操作编号
	OptionSelfCloseSysID tTThostFtdcOrderSysIDType
	// 操作标志
	ActionFlag tTThostFtdcActionFlagType
	// 操作日期
	ActionDate tTThostFtdcDateType
	// 操作时间
	ActionTime tTThostFtdcTimeType
	// 交易所交易员代码
	TraderID tTThostFtdcTraderIDType
	// 安装编号
	InstallID tTThostFtdcInstallIDType
	// 本地期权自对冲编号
	OptionSelfCloseLocalID tTThostFtdcOrderLocalIDType
	// 操作本地编号
	ActionLocalID tTThostFtdcOrderLocalIDType
	// 会员代码
	ParticipantID tTThostFtdcParticipantIDType
	// 客户代码
	ClientID tTThostFtdcClientIDType
	// 业务单元
	BusinessUnit tTThostFtdcBusinessUnitType
	// 报单操作状态
	OrderActionStatus tTThostFtdcOrderActionStatusType
	// 用户代码
	UserID tTThostFtdcUserIDType
	// 营业部编号
	BranchID tTThostFtdcBranchIDType
	// IP地址
	IPAddress tTThostFtdcIPAddressType
	// Mac地址
	MacAddress tTThostFtdcMacAddressType
	// 合约在交易所的代码
	ExchangeInstID tTThostFtdcExchangeInstIDType
	// 期权行权的头寸是否自对冲
	OptSelfCloseFlag tTThostFtdcOptSelfCloseFlagType
}

// 延时换汇同步
type tCThostFtdcSyncDelaySwapField struct {
	// 换汇流水号
	DelaySwapSeqNo tTThostFtdcDepositSeqNoType
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 源币种
	FromCurrencyID tTThostFtdcCurrencyIDType
	// 源金额
	FromAmount tTThostFtdcMoneyType
	// 源换汇冻结金额(可用冻结)
	FromFrozenSwap tTThostFtdcMoneyType
	// 源剩余换汇额度(可提冻结)
	FromRemainSwap tTThostFtdcMoneyType
	// 目标币种
	ToCurrencyID tTThostFtdcCurrencyIDType
	// 目标金额
	ToAmount tTThostFtdcMoneyType
}

// 查询延时换汇同步
type tCThostFtdcQrySyncDelaySwapField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 延时换汇流水号
	DelaySwapSeqNo tTThostFtdcDepositSeqNoType
}

// 投资单元
type tCThostFtdcInvestUnitField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 投资单元代码
	InvestUnitID tTThostFtdcInvestUnitIDType
	// 投资者单元名称
	InvestorUnitName tTThostFtdcPartyNameType
	// 投资者分组代码
	InvestorGroupID tTThostFtdcInvestorIDType
	// 手续费率模板代码
	CommModelID tTThostFtdcInvestorIDType
	// 保证金率模板代码
	MarginModelID tTThostFtdcInvestorIDType
	// 资金账号
	AccountID tTThostFtdcAccountIDType
	// 币种代码
	CurrencyID tTThostFtdcCurrencyIDType
}

// 查询投资单元
type tCThostFtdcQryInvestUnitField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 投资单元代码
	InvestUnitID tTThostFtdcInvestUnitIDType
}

// 二级代理商资金校验模式
type tCThostFtdcSecAgentCheckModeField struct {
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 币种
	CurrencyID tTThostFtdcCurrencyIDType
	// 境外中介机构资金帐号
	BrokerSecAgentID tTThostFtdcAccountIDType
	// 是否需要校验自己的资金账户
	CheckSelfAccount tTThostFtdcBoolType
}

// 二级代理商信息
type tCThostFtdcSecAgentTradeInfoField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 境外中介机构资金帐号
	BrokerSecAgentID tTThostFtdcAccountIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 二级代理商姓名
	LongCustomerName tTThostFtdcLongIndividualNameType
}

// 市场行情
type tCThostFtdcMarketDataField struct {
	// 交易日
	TradingDay tTThostFtdcDateType
	// 合约代码
	InstrumentID tTThostFtdcInstrumentIDType
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
	// 合约在交易所的代码
	ExchangeInstID tTThostFtdcExchangeInstIDType
	// 最新价
	LastPrice tTThostFtdcPriceType
	// 上次结算价
	PreSettlementPrice tTThostFtdcPriceType
	// 昨收盘
	PreClosePrice tTThostFtdcPriceType
	// 昨持仓量
	PreOpenInterest tTThostFtdcLargeVolumeType
	// 今开盘
	OpenPrice tTThostFtdcPriceType
	// 最高价
	HighestPrice tTThostFtdcPriceType
	// 最低价
	LowestPrice tTThostFtdcPriceType
	// 数量
	Volume tTThostFtdcVolumeType
	// 成交金额
	Turnover tTThostFtdcMoneyType
	// 持仓量
	OpenInterest tTThostFtdcLargeVolumeType
	// 今收盘
	ClosePrice tTThostFtdcPriceType
	// 本次结算价
	SettlementPrice tTThostFtdcPriceType
	// 涨停板价
	UpperLimitPrice tTThostFtdcPriceType
	// 跌停板价
	LowerLimitPrice tTThostFtdcPriceType
	// 昨虚实度
	PreDelta tTThostFtdcRatioType
	// 今虚实度
	CurrDelta tTThostFtdcRatioType
	// 最后修改时间
	UpdateTime tTThostFtdcTimeType
	// 最后修改毫秒
	UpdateMillisec tTThostFtdcMillisecType
	// 业务日期
	ActionDay tTThostFtdcDateType
}

// 行情基础属性
type tCThostFtdcMarketDataBaseField struct {
	// 交易日
	TradingDay tTThostFtdcDateType
	// 上次结算价
	PreSettlementPrice tTThostFtdcPriceType
	// 昨收盘
	PreClosePrice tTThostFtdcPriceType
	// 昨持仓量
	PreOpenInterest tTThostFtdcLargeVolumeType
	// 昨虚实度
	PreDelta tTThostFtdcRatioType
}

// 行情静态属性
type tCThostFtdcMarketDataStaticField struct {
	// 今开盘
	OpenPrice tTThostFtdcPriceType
	// 最高价
	HighestPrice tTThostFtdcPriceType
	// 最低价
	LowestPrice tTThostFtdcPriceType
	// 今收盘
	ClosePrice tTThostFtdcPriceType
	// 涨停板价
	UpperLimitPrice tTThostFtdcPriceType
	// 跌停板价
	LowerLimitPrice tTThostFtdcPriceType
	// 本次结算价
	SettlementPrice tTThostFtdcPriceType
	// 今虚实度
	CurrDelta tTThostFtdcRatioType
}

// 行情最新成交属性
type tCThostFtdcMarketDataLastMatchField struct {
	// 最新价
	LastPrice tTThostFtdcPriceType
	// 数量
	Volume tTThostFtdcVolumeType
	// 成交金额
	Turnover tTThostFtdcMoneyType
	// 持仓量
	OpenInterest tTThostFtdcLargeVolumeType
}

// 行情最优价属性
type tCThostFtdcMarketDataBestPriceField struct {
	// 申买价一
	BidPrice1 tTThostFtdcPriceType
	// 申买量一
	BidVolume1 tTThostFtdcVolumeType
	// 申卖价一
	AskPrice1 tTThostFtdcPriceType
	// 申卖量一
	AskVolume1 tTThostFtdcVolumeType
}

// 行情申买二、三属性
type tCThostFtdcMarketDataBid23Field struct {
	// 申买价二
	BidPrice2 tTThostFtdcPriceType
	// 申买量二
	BidVolume2 tTThostFtdcVolumeType
	// 申买价三
	BidPrice3 tTThostFtdcPriceType
	// 申买量三
	BidVolume3 tTThostFtdcVolumeType
}

// 行情申卖二、三属性
type tCThostFtdcMarketDataAsk23Field struct {
	// 申卖价二
	AskPrice2 tTThostFtdcPriceType
	// 申卖量二
	AskVolume2 tTThostFtdcVolumeType
	// 申卖价三
	AskPrice3 tTThostFtdcPriceType
	// 申卖量三
	AskVolume3 tTThostFtdcVolumeType
}

// 行情申买四、五属性
type tCThostFtdcMarketDataBid45Field struct {
	// 申买价四
	BidPrice4 tTThostFtdcPriceType
	// 申买量四
	BidVolume4 tTThostFtdcVolumeType
	// 申买价五
	BidPrice5 tTThostFtdcPriceType
	// 申买量五
	BidVolume5 tTThostFtdcVolumeType
}

// 行情申卖四、五属性
type tCThostFtdcMarketDataAsk45Field struct {
	// 申卖价四
	AskPrice4 tTThostFtdcPriceType
	// 申卖量四
	AskVolume4 tTThostFtdcVolumeType
	// 申卖价五
	AskPrice5 tTThostFtdcPriceType
	// 申卖量五
	AskVolume5 tTThostFtdcVolumeType
}

// 行情更新时间属性
type tCThostFtdcMarketDataUpdateTimeField struct {
	// 合约代码
	InstrumentID tTThostFtdcInstrumentIDType
	// 最后修改时间
	UpdateTime tTThostFtdcTimeType
	// 最后修改毫秒
	UpdateMillisec tTThostFtdcMillisecType
	// 业务日期
	ActionDay tTThostFtdcDateType
}

// 行情交易所代码属性
type tCThostFtdcMarketDataExchangeField struct {
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
}

// 指定的合约
type tCThostFtdcSpecificInstrumentField struct {
	// 合约代码
	InstrumentID tTThostFtdcInstrumentIDType
}

// 合约状态
type tCThostFtdcInstrumentStatusField struct {
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
	// 合约在交易所的代码
	ExchangeInstID tTThostFtdcExchangeInstIDType
	// 结算组代码
	SettlementGroupID tTThostFtdcSettlementGroupIDType
	// 合约代码
	InstrumentID tTThostFtdcInstrumentIDType
	// 合约交易状态
	InstrumentStatus tTThostFtdcInstrumentStatusType
	// 交易阶段编号
	TradingSegmentSN tTThostFtdcTradingSegmentSNType
	// 进入本状态时间
	EnterTime tTThostFtdcTimeType
	// 进入本状态原因
	EnterReason tTThostFtdcInstStatusEnterReasonType
}

// 查询合约状态
type tCThostFtdcQryInstrumentStatusField struct {
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
	// 合约在交易所的代码
	ExchangeInstID tTThostFtdcExchangeInstIDType
}

// 投资者账户
type tCThostFtdcInvestorAccountField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 投资者帐号
	AccountID tTThostFtdcAccountIDType
	// 币种代码
	CurrencyID tTThostFtdcCurrencyIDType
}

// 浮动盈亏算法
type tCThostFtdcPositionProfitAlgorithmField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者帐号
	AccountID tTThostFtdcAccountIDType
	// 盈亏算法
	Algorithm tTThostFtdcAlgorithmType
	// 备注
	Memo tTThostFtdcMemoType
	// 币种代码
	CurrencyID tTThostFtdcCurrencyIDType
}

// 会员资金折扣
type tCThostFtdcDiscountField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者范围
	InvestorRange tTThostFtdcInvestorRangeType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 资金折扣比例
	Discount tTThostFtdcRatioType
}

// 查询转帐银行
type tCThostFtdcQryTransferBankField struct {
	// 银行代码
	BankID tTThostFtdcBankIDType
	// 银行分中心代码
	BankBrchID tTThostFtdcBankBrchIDType
}

// 转帐银行
type tCThostFtdcTransferBankField struct {
	// 银行代码
	BankID tTThostFtdcBankIDType
	// 银行分中心代码
	BankBrchID tTThostFtdcBankBrchIDType
	// 银行名称
	BankName tTThostFtdcBankNameType
	// 是否活跃
	IsActive tTThostFtdcBoolType
}

// 查询投资者持仓明细
type tCThostFtdcQryInvestorPositionDetailField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 合约代码
	InstrumentID tTThostFtdcInstrumentIDType
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
	// 投资单元代码
	InvestUnitID tTThostFtdcInvestUnitIDType
}

// 投资者持仓明细
type tCThostFtdcInvestorPositionDetailField struct {
	// 合约代码
	InstrumentID tTThostFtdcInstrumentIDType
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 投机套保标志
	HedgeFlag tTThostFtdcHedgeFlagType
	// 买卖
	Direction tTThostFtdcDirectionType
	// 开仓日期
	OpenDate tTThostFtdcDateType
	// 成交编号
	TradeID tTThostFtdcTradeIDType
	// 数量
	Volume tTThostFtdcVolumeType
	// 开仓价
	OpenPrice tTThostFtdcPriceType
	// 交易日
	TradingDay tTThostFtdcDateType
	// 结算编号
	SettlementID tTThostFtdcSettlementIDType
	// 成交类型
	TradeType tTThostFtdcTradeTypeType
	// 组合合约代码
	CombInstrumentID tTThostFtdcInstrumentIDType
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
	// 逐日盯市平仓盈亏
	CloseProfitByDate tTThostFtdcMoneyType
	// 逐笔对冲平仓盈亏
	CloseProfitByTrade tTThostFtdcMoneyType
	// 逐日盯市持仓盈亏
	PositionProfitByDate tTThostFtdcMoneyType
	// 逐笔对冲持仓盈亏
	PositionProfitByTrade tTThostFtdcMoneyType
	// 投资者保证金
	Margin tTThostFtdcMoneyType
	// 交易所保证金
	ExchMargin tTThostFtdcMoneyType
	// 保证金率
	MarginRateByMoney tTThostFtdcRatioType
	// 保证金率(按手数)
	MarginRateByVolume tTThostFtdcRatioType
	// 昨结算价
	LastSettlementPrice tTThostFtdcPriceType
	// 结算价
	SettlementPrice tTThostFtdcPriceType
	// 平仓量
	CloseVolume tTThostFtdcVolumeType
	// 平仓金额
	CloseAmount tTThostFtdcMoneyType
	// 按照时间顺序平仓的笔数,大商所专用
	TimeFirstVolume tTThostFtdcVolumeType
	// 投资单元代码
	InvestUnitID tTThostFtdcInvestUnitIDType
}

// 资金账户口令域
type tCThostFtdcTradingAccountPasswordField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者帐号
	AccountID tTThostFtdcAccountIDType
	// 密码
	Password tTThostFtdcPasswordType
	// 币种代码
	CurrencyID tTThostFtdcCurrencyIDType
}

// 交易所行情报盘机
type tCThostFtdcMDTraderOfferField struct {
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
	// 交易所交易员代码
	TraderID tTThostFtdcTraderIDType
	// 会员代码
	ParticipantID tTThostFtdcParticipantIDType
	// 密码
	Password tTThostFtdcPasswordType
	// 安装编号
	InstallID tTThostFtdcInstallIDType
	// 本地报单编号
	OrderLocalID tTThostFtdcOrderLocalIDType
	// 交易所交易员连接状态
	TraderConnectStatus tTThostFtdcTraderConnectStatusType
	// 发出连接请求的日期
	ConnectRequestDate tTThostFtdcDateType
	// 发出连接请求的时间
	ConnectRequestTime tTThostFtdcTimeType
	// 上次报告日期
	LastReportDate tTThostFtdcDateType
	// 上次报告时间
	LastReportTime tTThostFtdcTimeType
	// 完成连接日期
	ConnectDate tTThostFtdcDateType
	// 完成连接时间
	ConnectTime tTThostFtdcTimeType
	// 启动日期
	StartDate tTThostFtdcDateType
	// 启动时间
	StartTime tTThostFtdcTimeType
	// 交易日
	TradingDay tTThostFtdcDateType
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 本席位最大成交编号
	MaxTradeID tTThostFtdcTradeIDType
	// 本席位最大报单备拷
	MaxOrderMessageReference tTThostFtdcReturnCodeType
}

// 查询行情报盘机
type tCThostFtdcQryMDTraderOfferField struct {
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
	// 会员代码
	ParticipantID tTThostFtdcParticipantIDType
	// 交易所交易员代码
	TraderID tTThostFtdcTraderIDType
}

// 查询客户通知
type tCThostFtdcQryNoticeField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
}

// 客户通知
type tCThostFtdcNoticeField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 消息正文
	Content tTThostFtdcContentType
	// 经纪公司通知内容序列号
	SequenceLabel tTThostFtdcSequenceLabelType
}

// 用户权限
type tCThostFtdcUserRightField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 用户代码
	UserID tTThostFtdcUserIDType
	// 客户权限类型
	UserRightType tTThostFtdcUserRightTypeType
	// 是否禁止
	IsForbidden tTThostFtdcBoolType
}

// 查询结算信息确认域
type tCThostFtdcQrySettlementInfoConfirmField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 投资者帐号
	AccountID tTThostFtdcAccountIDType
	// 币种代码
	CurrencyID tTThostFtdcCurrencyIDType
}

// 装载结算信息
type tCThostFtdcLoadSettlementInfoField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
}

// 经纪公司可提资金算法表
type tCThostFtdcBrokerWithdrawAlgorithmField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 可提资金算法
	WithdrawAlgorithm tTThostFtdcAlgorithmType
	// 资金使用率
	UsingRatio tTThostFtdcRatioType
	// 可提是否包含平仓盈利
	IncludeCloseProfit tTThostFtdcIncludeCloseProfitType
	// 本日无仓且无成交客户是否受可提比例限制
	AllWithoutTrade tTThostFtdcAllWithoutTradeType
	// 可用是否包含平仓盈利
	AvailIncludeCloseProfit tTThostFtdcIncludeCloseProfitType
	// 是否启用用户事件
	IsBrokerUserEvent tTThostFtdcBoolType
	// 币种代码
	CurrencyID tTThostFtdcCurrencyIDType
	// 货币质押比率
	FundMortgageRatio tTThostFtdcRatioType
	// 权益算法
	BalanceAlgorithm tTThostFtdcBalanceAlgorithmType
}

// 资金账户口令变更域
type tCThostFtdcTradingAccountPasswordUpdateV1Field struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 原来的口令
	OldPassword tTThostFtdcPasswordType
	// 新的口令
	NewPassword tTThostFtdcPasswordType
}

// 资金账户口令变更域
type tCThostFtdcTradingAccountPasswordUpdateField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者帐号
	AccountID tTThostFtdcAccountIDType
	// 原来的口令
	OldPassword tTThostFtdcPasswordType
	// 新的口令
	NewPassword tTThostFtdcPasswordType
	// 币种代码
	CurrencyID tTThostFtdcCurrencyIDType
}

// 查询组合合约分腿
type tCThostFtdcQryCombinationLegField struct {
	// 组合合约代码
	CombInstrumentID tTThostFtdcInstrumentIDType
	// 单腿编号
	LegID tTThostFtdcLegIDType
	// 单腿合约代码
	LegInstrumentID tTThostFtdcInstrumentIDType
}

// 查询组合合约分腿
type tCThostFtdcQrySyncStatusField struct {
	// 交易日
	TradingDay tTThostFtdcDateType
}

// 组合交易合约的单腿
type tCThostFtdcCombinationLegField struct {
	// 组合合约代码
	CombInstrumentID tTThostFtdcInstrumentIDType
	// 单腿编号
	LegID tTThostFtdcLegIDType
	// 单腿合约代码
	LegInstrumentID tTThostFtdcInstrumentIDType
	// 买卖方向
	Direction tTThostFtdcDirectionType
	// 单腿乘数
	LegMultiple tTThostFtdcLegMultipleType
	// 派生层数
	ImplyLevel tTThostFtdcImplyLevelType
}

// 数据同步状态
type tCThostFtdcSyncStatusField struct {
	// 交易日
	TradingDay tTThostFtdcDateType
	// 数据同步状态
	DataSyncStatus tTThostFtdcDataSyncStatusType
}

// 查询联系人
type tCThostFtdcQryLinkManField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
}

// 联系人
type tCThostFtdcLinkManField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 联系人类型
	PersonType tTThostFtdcPersonTypeType
	// 证件类型
	IdentifiedCardType tTThostFtdcIdCardTypeType
	// 证件号码
	IdentifiedCardNo tTThostFtdcIdentifiedCardNoType
	// 名称
	PersonName tTThostFtdcPartyNameType
	// 联系电话
	Telephone tTThostFtdcTelephoneType
	// 通讯地址
	Address tTThostFtdcAddressType
	// 邮政编码
	ZipCode tTThostFtdcZipCodeType
	// 优先级
	Priority tTThostFtdcPriorityType
	// 开户邮政编码
	UOAZipCode tTThostFtdcUOAZipCodeType
	// 全称
	PersonFullName tTThostFtdcInvestorFullNameType
}

// 查询经纪公司用户事件
type tCThostFtdcQryBrokerUserEventField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 用户代码
	UserID tTThostFtdcUserIDType
	// 用户事件类型
	UserEventType tTThostFtdcUserEventTypeType
}

// 查询经纪公司用户事件
type tCThostFtdcBrokerUserEventField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 用户代码
	UserID tTThostFtdcUserIDType
	// 用户事件类型
	UserEventType tTThostFtdcUserEventTypeType
	// 用户事件序号
	EventSequenceNo tTThostFtdcSequenceNoType
	// 事件发生日期
	EventDate tTThostFtdcDateType
	// 事件发生时间
	EventTime tTThostFtdcTimeType
	// 用户事件信息
	UserEventInfo tTThostFtdcUserEventInfoType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 合约代码
	InstrumentID tTThostFtdcInstrumentIDType
}

// 查询签约银行请求
type tCThostFtdcQryContractBankField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 银行代码
	BankID tTThostFtdcBankIDType
	// 银行分中心代码
	BankBrchID tTThostFtdcBankBrchIDType
}

// 查询签约银行响应
type tCThostFtdcContractBankField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 银行代码
	BankID tTThostFtdcBankIDType
	// 银行分中心代码
	BankBrchID tTThostFtdcBankBrchIDType
	// 银行名称
	BankName tTThostFtdcBankNameType
}

// 投资者组合持仓明细
type tCThostFtdcInvestorPositionCombineDetailField struct {
	// 交易日
	TradingDay tTThostFtdcDateType
	// 开仓日期
	OpenDate tTThostFtdcDateType
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
	// 结算编号
	SettlementID tTThostFtdcSettlementIDType
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 组合编号
	ComTradeID tTThostFtdcTradeIDType
	// 撮合编号
	TradeID tTThostFtdcTradeIDType
	// 合约代码
	InstrumentID tTThostFtdcInstrumentIDType
	// 投机套保标志
	HedgeFlag tTThostFtdcHedgeFlagType
	// 买卖
	Direction tTThostFtdcDirectionType
	// 持仓量
	TotalAmt tTThostFtdcVolumeType
	// 投资者保证金
	Margin tTThostFtdcMoneyType
	// 交易所保证金
	ExchMargin tTThostFtdcMoneyType
	// 保证金率
	MarginRateByMoney tTThostFtdcRatioType
	// 保证金率(按手数)
	MarginRateByVolume tTThostFtdcRatioType
	// 单腿编号
	LegID tTThostFtdcLegIDType
	// 单腿乘数
	LegMultiple tTThostFtdcLegMultipleType
	// 组合持仓合约编码
	CombInstrumentID tTThostFtdcInstrumentIDType
	// 成交组号
	TradeGroupID tTThostFtdcTradeGroupIDType
	// 投资单元代码
	InvestUnitID tTThostFtdcInvestUnitIDType
}

// 预埋单
type tCThostFtdcParkedOrderField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 合约代码
	InstrumentID tTThostFtdcInstrumentIDType
	// 报单引用
	OrderRef tTThostFtdcOrderRefType
	// 用户代码
	UserID tTThostFtdcUserIDType
	// 报单价格条件
	OrderPriceType tTThostFtdcOrderPriceTypeType
	// 买卖方向
	Direction tTThostFtdcDirectionType
	// 组合开平标志
	CombOffsetFlag tTThostFtdcCombOffsetFlagType
	// 组合投机套保标志
	CombHedgeFlag tTThostFtdcCombHedgeFlagType
	// 价格
	LimitPrice tTThostFtdcPriceType
	// 数量
	VolumeTotalOriginal tTThostFtdcVolumeType
	// 有效期类型
	TimeCondition tTThostFtdcTimeConditionType
	// GTD日期
	GTDDate tTThostFtdcDateType
	// 成交量类型
	VolumeCondition tTThostFtdcVolumeConditionType
	// 最小成交量
	MinVolume tTThostFtdcVolumeType
	// 触发条件
	ContingentCondition tTThostFtdcContingentConditionType
	// 止损价
	StopPrice tTThostFtdcPriceType
	// 强平原因
	ForceCloseReason tTThostFtdcForceCloseReasonType
	// 自动挂起标志
	IsAutoSuspend tTThostFtdcBoolType
	// 业务单元
	BusinessUnit tTThostFtdcBusinessUnitType
	// 请求编号
	RequestID tTThostFtdcRequestIDType
	// 用户强评标志
	UserForceClose tTThostFtdcBoolType
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
	// 预埋报单编号
	ParkedOrderID tTThostFtdcParkedOrderIDType
	// 用户类型
	UserType tTThostFtdcUserTypeType
	// 预埋单状态
	Status tTThostFtdcParkedOrderStatusType
	// 错误代码
	ErrorID tTThostFtdcErrorIDType
	// 错误信息
	ErrorMsg tTThostFtdcErrorMsgType
	// 互换单标志
	IsSwapOrder tTThostFtdcBoolType
	// 资金账号
	AccountID tTThostFtdcAccountIDType
	// 币种代码
	CurrencyID tTThostFtdcCurrencyIDType
	// 交易编码
	ClientID tTThostFtdcClientIDType
	// 投资单元代码
	InvestUnitID tTThostFtdcInvestUnitIDType
	// IP地址
	IPAddress tTThostFtdcIPAddressType
	// Mac地址
	MacAddress tTThostFtdcMacAddressType
}

// 输入预埋单操作
type tCThostFtdcParkedOrderActionField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 报单操作引用
	OrderActionRef tTThostFtdcOrderActionRefType
	// 报单引用
	OrderRef tTThostFtdcOrderRefType
	// 请求编号
	RequestID tTThostFtdcRequestIDType
	// 前置编号
	FrontID tTThostFtdcFrontIDType
	// 会话编号
	SessionID tTThostFtdcSessionIDType
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
	// 报单编号
	OrderSysID tTThostFtdcOrderSysIDType
	// 操作标志
	ActionFlag tTThostFtdcActionFlagType
	// 价格
	LimitPrice tTThostFtdcPriceType
	// 数量变化
	VolumeChange tTThostFtdcVolumeType
	// 用户代码
	UserID tTThostFtdcUserIDType
	// 合约代码
	InstrumentID tTThostFtdcInstrumentIDType
	// 预埋撤单单编号
	ParkedOrderActionID tTThostFtdcParkedOrderActionIDType
	// 用户类型
	UserType tTThostFtdcUserTypeType
	// 预埋撤单状态
	Status tTThostFtdcParkedOrderStatusType
	// 错误代码
	ErrorID tTThostFtdcErrorIDType
	// 错误信息
	ErrorMsg tTThostFtdcErrorMsgType
	// 投资单元代码
	InvestUnitID tTThostFtdcInvestUnitIDType
	// IP地址
	IPAddress tTThostFtdcIPAddressType
	// Mac地址
	MacAddress tTThostFtdcMacAddressType
}

// 查询预埋单
type tCThostFtdcQryParkedOrderField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 合约代码
	InstrumentID tTThostFtdcInstrumentIDType
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
	// 投资单元代码
	InvestUnitID tTThostFtdcInvestUnitIDType
}

// 查询预埋撤单
type tCThostFtdcQryParkedOrderActionField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 合约代码
	InstrumentID tTThostFtdcInstrumentIDType
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
	// 投资单元代码
	InvestUnitID tTThostFtdcInvestUnitIDType
}

// 删除预埋单
type tCThostFtdcRemoveParkedOrderField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 预埋报单编号
	ParkedOrderID tTThostFtdcParkedOrderIDType
	// 投资单元代码
	InvestUnitID tTThostFtdcInvestUnitIDType
}

// 删除预埋撤单
type tCThostFtdcRemoveParkedOrderActionField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 预埋撤单编号
	ParkedOrderActionID tTThostFtdcParkedOrderActionIDType
	// 投资单元代码
	InvestUnitID tTThostFtdcInvestUnitIDType
}

// 经纪公司可提资金算法表
type tCThostFtdcInvestorWithdrawAlgorithmField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者范围
	InvestorRange tTThostFtdcInvestorRangeType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 可提资金比例
	UsingRatio tTThostFtdcRatioType
	// 币种代码
	CurrencyID tTThostFtdcCurrencyIDType
	// 货币质押比率
	FundMortgageRatio tTThostFtdcRatioType
}

// 查询组合持仓明细
type tCThostFtdcQryInvestorPositionCombineDetailField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 组合持仓合约编码
	CombInstrumentID tTThostFtdcInstrumentIDType
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
	// 投资单元代码
	InvestUnitID tTThostFtdcInvestUnitIDType
}

// 成交均价
type tCThostFtdcMarketDataAveragePriceField struct {
	// 当日均价
	AveragePrice tTThostFtdcPriceType
}

// 校验投资者密码
type tCThostFtdcVerifyInvestorPasswordField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 密码
	Password tTThostFtdcPasswordType
}

// 用户IP
type tCThostFtdcUserIPField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 用户代码
	UserID tTThostFtdcUserIDType
	// IP地址
	IPAddress tTThostFtdcIPAddressType
	// IP地址掩码
	IPMask tTThostFtdcIPAddressType
	// Mac地址
	MacAddress tTThostFtdcMacAddressType
}

// 用户事件通知信息
type tCThostFtdcTradingNoticeInfoField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 发送时间
	SendTime tTThostFtdcTimeType
	// 消息正文
	FieldContent tTThostFtdcContentType
	// 序列系列号
	SequenceSeries tTThostFtdcSequenceSeriesType
	// 序列号
	SequenceNo tTThostFtdcSequenceNoType
	// 投资单元代码
	InvestUnitID tTThostFtdcInvestUnitIDType
}

// 用户事件通知
type tCThostFtdcTradingNoticeField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者范围
	InvestorRange tTThostFtdcInvestorRangeType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 序列系列号
	SequenceSeries tTThostFtdcSequenceSeriesType
	// 用户代码
	UserID tTThostFtdcUserIDType
	// 发送时间
	SendTime tTThostFtdcTimeType
	// 序列号
	SequenceNo tTThostFtdcSequenceNoType
	// 消息正文
	FieldContent tTThostFtdcContentType
	// 投资单元代码
	InvestUnitID tTThostFtdcInvestUnitIDType
}

// 查询交易事件通知
type tCThostFtdcQryTradingNoticeField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 投资单元代码
	InvestUnitID tTThostFtdcInvestUnitIDType
}

// 查询错误报单
type tCThostFtdcQryErrOrderField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
}

// 错误报单
type tCThostFtdcErrOrderField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 合约代码
	InstrumentID tTThostFtdcInstrumentIDType
	// 报单引用
	OrderRef tTThostFtdcOrderRefType
	// 用户代码
	UserID tTThostFtdcUserIDType
	// 报单价格条件
	OrderPriceType tTThostFtdcOrderPriceTypeType
	// 买卖方向
	Direction tTThostFtdcDirectionType
	// 组合开平标志
	CombOffsetFlag tTThostFtdcCombOffsetFlagType
	// 组合投机套保标志
	CombHedgeFlag tTThostFtdcCombHedgeFlagType
	// 价格
	LimitPrice tTThostFtdcPriceType
	// 数量
	VolumeTotalOriginal tTThostFtdcVolumeType
	// 有效期类型
	TimeCondition tTThostFtdcTimeConditionType
	// GTD日期
	GTDDate tTThostFtdcDateType
	// 成交量类型
	VolumeCondition tTThostFtdcVolumeConditionType
	// 最小成交量
	MinVolume tTThostFtdcVolumeType
	// 触发条件
	ContingentCondition tTThostFtdcContingentConditionType
	// 止损价
	StopPrice tTThostFtdcPriceType
	// 强平原因
	ForceCloseReason tTThostFtdcForceCloseReasonType
	// 自动挂起标志
	IsAutoSuspend tTThostFtdcBoolType
	// 业务单元
	BusinessUnit tTThostFtdcBusinessUnitType
	// 请求编号
	RequestID tTThostFtdcRequestIDType
	// 用户强评标志
	UserForceClose tTThostFtdcBoolType
	// 错误代码
	ErrorID tTThostFtdcErrorIDType
	// 错误信息
	ErrorMsg tTThostFtdcErrorMsgType
	// 互换单标志
	IsSwapOrder tTThostFtdcBoolType
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
	// 投资单元代码
	InvestUnitID tTThostFtdcInvestUnitIDType
	// 资金账号
	AccountID tTThostFtdcAccountIDType
	// 币种代码
	CurrencyID tTThostFtdcCurrencyIDType
	// 交易编码
	ClientID tTThostFtdcClientIDType
	// IP地址
	IPAddress tTThostFtdcIPAddressType
	// Mac地址
	MacAddress tTThostFtdcMacAddressType
}

// 查询错误报单操作
type tCThostFtdcErrorConditionalOrderField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 合约代码
	InstrumentID tTThostFtdcInstrumentIDType
	// 报单引用
	OrderRef tTThostFtdcOrderRefType
	// 用户代码
	UserID tTThostFtdcUserIDType
	// 报单价格条件
	OrderPriceType tTThostFtdcOrderPriceTypeType
	// 买卖方向
	Direction tTThostFtdcDirectionType
	// 组合开平标志
	CombOffsetFlag tTThostFtdcCombOffsetFlagType
	// 组合投机套保标志
	CombHedgeFlag tTThostFtdcCombHedgeFlagType
	// 价格
	LimitPrice tTThostFtdcPriceType
	// 数量
	VolumeTotalOriginal tTThostFtdcVolumeType
	// 有效期类型
	TimeCondition tTThostFtdcTimeConditionType
	// GTD日期
	GTDDate tTThostFtdcDateType
	// 成交量类型
	VolumeCondition tTThostFtdcVolumeConditionType
	// 最小成交量
	MinVolume tTThostFtdcVolumeType
	// 触发条件
	ContingentCondition tTThostFtdcContingentConditionType
	// 止损价
	StopPrice tTThostFtdcPriceType
	// 强平原因
	ForceCloseReason tTThostFtdcForceCloseReasonType
	// 自动挂起标志
	IsAutoSuspend tTThostFtdcBoolType
	// 业务单元
	BusinessUnit tTThostFtdcBusinessUnitType
	// 请求编号
	RequestID tTThostFtdcRequestIDType
	// 本地报单编号
	OrderLocalID tTThostFtdcOrderLocalIDType
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
	// 会员代码
	ParticipantID tTThostFtdcParticipantIDType
	// 客户代码
	ClientID tTThostFtdcClientIDType
	// 合约在交易所的代码
	ExchangeInstID tTThostFtdcExchangeInstIDType
	// 交易所交易员代码
	TraderID tTThostFtdcTraderIDType
	// 安装编号
	InstallID tTThostFtdcInstallIDType
	// 报单提交状态
	OrderSubmitStatus tTThostFtdcOrderSubmitStatusType
	// 报单提示序号
	NotifySequence tTThostFtdcSequenceNoType
	// 交易日
	TradingDay tTThostFtdcDateType
	// 结算编号
	SettlementID tTThostFtdcSettlementIDType
	// 报单编号
	OrderSysID tTThostFtdcOrderSysIDType
	// 报单来源
	OrderSource tTThostFtdcOrderSourceType
	// 报单状态
	OrderStatus tTThostFtdcOrderStatusType
	// 报单类型
	OrderType tTThostFtdcOrderTypeType
	// 今成交数量
	VolumeTraded tTThostFtdcVolumeType
	// 剩余数量
	VolumeTotal tTThostFtdcVolumeType
	// 报单日期
	InsertDate tTThostFtdcDateType
	// 委托时间
	InsertTime tTThostFtdcTimeType
	// 激活时间
	ActiveTime tTThostFtdcTimeType
	// 挂起时间
	SuspendTime tTThostFtdcTimeType
	// 最后修改时间
	UpdateTime tTThostFtdcTimeType
	// 撤销时间
	CancelTime tTThostFtdcTimeType
	// 最后修改交易所交易员代码
	ActiveTraderID tTThostFtdcTraderIDType
	// 结算会员编号
	ClearingPartID tTThostFtdcParticipantIDType
	// 序号
	SequenceNo tTThostFtdcSequenceNoType
	// 前置编号
	FrontID tTThostFtdcFrontIDType
	// 会话编号
	SessionID tTThostFtdcSessionIDType
	// 用户端产品信息
	UserProductInfo tTThostFtdcProductInfoType
	// 状态信息
	StatusMsg tTThostFtdcErrorMsgType
	// 用户强评标志
	UserForceClose tTThostFtdcBoolType
	// 操作用户代码
	ActiveUserID tTThostFtdcUserIDType
	// 经纪公司报单编号
	BrokerOrderSeq tTThostFtdcSequenceNoType
	// 相关报单
	RelativeOrderSysID tTThostFtdcOrderSysIDType
	// 郑商所成交数量
	ZCETotalTradedVolume tTThostFtdcVolumeType
	// 错误代码
	ErrorID tTThostFtdcErrorIDType
	// 错误信息
	ErrorMsg tTThostFtdcErrorMsgType
	// 互换单标志
	IsSwapOrder tTThostFtdcBoolType
	// 营业部编号
	BranchID tTThostFtdcBranchIDType
	// 投资单元代码
	InvestUnitID tTThostFtdcInvestUnitIDType
	// 资金账号
	AccountID tTThostFtdcAccountIDType
	// 币种代码
	CurrencyID tTThostFtdcCurrencyIDType
	// IP地址
	IPAddress tTThostFtdcIPAddressType
	// Mac地址
	MacAddress tTThostFtdcMacAddressType
}

// 查询错误报单操作
type tCThostFtdcQryErrOrderActionField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
}

// 错误报单操作
type tCThostFtdcErrOrderActionField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 报单操作引用
	OrderActionRef tTThostFtdcOrderActionRefType
	// 报单引用
	OrderRef tTThostFtdcOrderRefType
	// 请求编号
	RequestID tTThostFtdcRequestIDType
	// 前置编号
	FrontID tTThostFtdcFrontIDType
	// 会话编号
	SessionID tTThostFtdcSessionIDType
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
	// 报单编号
	OrderSysID tTThostFtdcOrderSysIDType
	// 操作标志
	ActionFlag tTThostFtdcActionFlagType
	// 价格
	LimitPrice tTThostFtdcPriceType
	// 数量变化
	VolumeChange tTThostFtdcVolumeType
	// 操作日期
	ActionDate tTThostFtdcDateType
	// 操作时间
	ActionTime tTThostFtdcTimeType
	// 交易所交易员代码
	TraderID tTThostFtdcTraderIDType
	// 安装编号
	InstallID tTThostFtdcInstallIDType
	// 本地报单编号
	OrderLocalID tTThostFtdcOrderLocalIDType
	// 操作本地编号
	ActionLocalID tTThostFtdcOrderLocalIDType
	// 会员代码
	ParticipantID tTThostFtdcParticipantIDType
	// 客户代码
	ClientID tTThostFtdcClientIDType
	// 业务单元
	BusinessUnit tTThostFtdcBusinessUnitType
	// 报单操作状态
	OrderActionStatus tTThostFtdcOrderActionStatusType
	// 用户代码
	UserID tTThostFtdcUserIDType
	// 状态信息
	StatusMsg tTThostFtdcErrorMsgType
	// 合约代码
	InstrumentID tTThostFtdcInstrumentIDType
	// 营业部编号
	BranchID tTThostFtdcBranchIDType
	// 投资单元代码
	InvestUnitID tTThostFtdcInvestUnitIDType
	// IP地址
	IPAddress tTThostFtdcIPAddressType
	// Mac地址
	MacAddress tTThostFtdcMacAddressType
	// 错误代码
	ErrorID tTThostFtdcErrorIDType
	// 错误信息
	ErrorMsg tTThostFtdcErrorMsgType
}

// 查询交易所状态
type tCThostFtdcQryExchangeSequenceField struct {
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
}

// 交易所状态
type tCThostFtdcExchangeSequenceField struct {
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
	// 序号
	SequenceNo tTThostFtdcSequenceNoType
	// 合约交易状态
	MarketStatus tTThostFtdcInstrumentStatusType
}

// 根据价格查询最大报单数量
type tCThostFtdcQueryMaxOrderVolumeWithPriceField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 合约代码
	InstrumentID tTThostFtdcInstrumentIDType
	// 买卖方向
	Direction tTThostFtdcDirectionType
	// 开平标志
	OffsetFlag tTThostFtdcOffsetFlagType
	// 投机套保标志
	HedgeFlag tTThostFtdcHedgeFlagType
	// 最大允许报单数量
	MaxVolume tTThostFtdcVolumeType
	// 报单价格
	Price tTThostFtdcPriceType
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
	// 投资单元代码
	InvestUnitID tTThostFtdcInvestUnitIDType
}

// 查询经纪公司交易参数
type tCThostFtdcQryBrokerTradingParamsField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 币种代码
	CurrencyID tTThostFtdcCurrencyIDType
	// 投资者帐号
	AccountID tTThostFtdcAccountIDType
}

// 经纪公司交易参数
type tCThostFtdcBrokerTradingParamsField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 保证金价格类型
	MarginPriceType tTThostFtdcMarginPriceTypeType
	// 盈亏算法
	Algorithm tTThostFtdcAlgorithmType
	// 可用是否包含平仓盈利
	AvailIncludeCloseProfit tTThostFtdcIncludeCloseProfitType
	// 币种代码
	CurrencyID tTThostFtdcCurrencyIDType
	// 期权权利金价格类型
	OptionRoyaltyPriceType tTThostFtdcOptionRoyaltyPriceTypeType
	// 投资者帐号
	AccountID tTThostFtdcAccountIDType
}

// 查询经纪公司交易算法
type tCThostFtdcQryBrokerTradingAlgosField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
	// 合约代码
	InstrumentID tTThostFtdcInstrumentIDType
}

// 经纪公司交易算法
type tCThostFtdcBrokerTradingAlgosField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
	// 合约代码
	InstrumentID tTThostFtdcInstrumentIDType
	// 持仓处理算法编号
	HandlePositionAlgoID tTThostFtdcHandlePositionAlgoIDType
	// 寻找保证金率算法编号
	FindMarginRateAlgoID tTThostFtdcFindMarginRateAlgoIDType
	// 资金处理算法编号
	HandleTradingAccountAlgoID tTThostFtdcHandleTradingAccountAlgoIDType
}

// 查询经纪公司资金
type tCThostFtdcQueryBrokerDepositField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
}

// 经纪公司资金
type tCThostFtdcBrokerDepositField struct {
	// 交易日期
	TradingDay tTThostFtdcTradeDateType
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 会员代码
	ParticipantID tTThostFtdcParticipantIDType
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
	// 上次结算准备金
	PreBalance tTThostFtdcMoneyType
	// 当前保证金总额
	CurrMargin tTThostFtdcMoneyType
	// 平仓盈亏
	CloseProfit tTThostFtdcMoneyType
	// 期货结算准备金
	Balance tTThostFtdcMoneyType
	// 入金金额
	Deposit tTThostFtdcMoneyType
	// 出金金额
	Withdraw tTThostFtdcMoneyType
	// 可提资金
	Available tTThostFtdcMoneyType
	// 基本准备金
	Reserve tTThostFtdcMoneyType
	// 冻结的保证金
	FrozenMargin tTThostFtdcMoneyType
}

// 查询保证金监管系统经纪公司密钥
type tCThostFtdcQryCFMMCBrokerKeyField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
}

// 保证金监管系统经纪公司密钥
type tCThostFtdcCFMMCBrokerKeyField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 经纪公司统一编码
	ParticipantID tTThostFtdcParticipantIDType
	// 密钥生成日期
	CreateDate tTThostFtdcDateType
	// 密钥生成时间
	CreateTime tTThostFtdcTimeType
	// 密钥编号
	KeyID tTThostFtdcSequenceNoType
	// 动态密钥
	CurrentKey tTThostFtdcCFMMCKeyType
	// 动态密钥类型
	KeyKind tTThostFtdcCFMMCKeyKindType
}

// 保证金监管系统经纪公司资金账户密钥
type tCThostFtdcCFMMCTradingAccountKeyField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 经纪公司统一编码
	ParticipantID tTThostFtdcParticipantIDType
	// 投资者帐号
	AccountID tTThostFtdcAccountIDType
	// 密钥编号
	KeyID tTThostFtdcSequenceNoType
	// 动态密钥
	CurrentKey tTThostFtdcCFMMCKeyType
}

// 请求查询保证金监管系统经纪公司资金账户密钥
type tCThostFtdcQryCFMMCTradingAccountKeyField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
}

// 用户动态令牌参数
type tCThostFtdcBrokerUserOTPParamField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 用户代码
	UserID tTThostFtdcUserIDType
	// 动态令牌提供商
	OTPVendorsID tTThostFtdcOTPVendorsIDType
	// 动态令牌序列号
	SerialNumber tTThostFtdcSerialNumberType
	// 令牌密钥
	AuthKey tTThostFtdcAuthKeyType
	// 漂移值
	LastDrift tTThostFtdcLastDriftType
	// 成功值
	LastSuccess tTThostFtdcLastSuccessType
	// 动态令牌类型
	OTPType tTThostFtdcOTPTypeType
}

// 手工同步用户动态令牌
type tCThostFtdcManualSyncBrokerUserOTPField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 用户代码
	UserID tTThostFtdcUserIDType
	// 动态令牌类型
	OTPType tTThostFtdcOTPTypeType
	// 第一个动态密码
	FirstOTP tTThostFtdcPasswordType
	// 第二个动态密码
	SecondOTP tTThostFtdcPasswordType
}

// 投资者手续费率模板
type tCThostFtdcCommRateModelField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 手续费率模板代码
	CommModelID tTThostFtdcInvestorIDType
	// 模板名称
	CommModelName tTThostFtdcCommModelNameType
}

// 请求查询投资者手续费率模板
type tCThostFtdcQryCommRateModelField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 手续费率模板代码
	CommModelID tTThostFtdcInvestorIDType
}

// 投资者保证金率模板
type tCThostFtdcMarginModelField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 保证金率模板代码
	MarginModelID tTThostFtdcInvestorIDType
	// 模板名称
	MarginModelName tTThostFtdcCommModelNameType
}

// 请求查询投资者保证金率模板
type tCThostFtdcQryMarginModelField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 保证金率模板代码
	MarginModelID tTThostFtdcInvestorIDType
}

// 仓单折抵信息
type tCThostFtdcEWarrantOffsetField struct {
	// 交易日期
	TradingDay tTThostFtdcTradeDateType
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
	// 合约代码
	InstrumentID tTThostFtdcInstrumentIDType
	// 买卖方向
	Direction tTThostFtdcDirectionType
	// 投机套保标志
	HedgeFlag tTThostFtdcHedgeFlagType
	// 数量
	Volume tTThostFtdcVolumeType
	// 投资单元代码
	InvestUnitID tTThostFtdcInvestUnitIDType
}

// 查询仓单折抵信息
type tCThostFtdcQryEWarrantOffsetField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
	// 合约代码
	InstrumentID tTThostFtdcInstrumentIDType
	// 投资单元代码
	InvestUnitID tTThostFtdcInvestUnitIDType
}

// 查询投资者品种/跨品种保证金
type tCThostFtdcQryInvestorProductGroupMarginField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 品种/跨品种标示
	ProductGroupID tTThostFtdcInstrumentIDType
	// 投机套保标志
	HedgeFlag tTThostFtdcHedgeFlagType
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
	// 投资单元代码
	InvestUnitID tTThostFtdcInvestUnitIDType
}

// 投资者品种/跨品种保证金
type tCThostFtdcInvestorProductGroupMarginField struct {
	// 品种/跨品种标示
	ProductGroupID tTThostFtdcInstrumentIDType
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 交易日
	TradingDay tTThostFtdcDateType
	// 结算编号
	SettlementID tTThostFtdcSettlementIDType
	// 冻结的保证金
	FrozenMargin tTThostFtdcMoneyType
	// 多头冻结的保证金
	LongFrozenMargin tTThostFtdcMoneyType
	// 空头冻结的保证金
	ShortFrozenMargin tTThostFtdcMoneyType
	// 占用的保证金
	UseMargin tTThostFtdcMoneyType
	// 多头保证金
	LongUseMargin tTThostFtdcMoneyType
	// 空头保证金
	ShortUseMargin tTThostFtdcMoneyType
	// 交易所保证金
	ExchMargin tTThostFtdcMoneyType
	// 交易所多头保证金
	LongExchMargin tTThostFtdcMoneyType
	// 交易所空头保证金
	ShortExchMargin tTThostFtdcMoneyType
	// 平仓盈亏
	CloseProfit tTThostFtdcMoneyType
	// 冻结的手续费
	FrozenCommission tTThostFtdcMoneyType
	// 手续费
	Commission tTThostFtdcMoneyType
	// 冻结的资金
	FrozenCash tTThostFtdcMoneyType
	// 资金差额
	CashIn tTThostFtdcMoneyType
	// 持仓盈亏
	PositionProfit tTThostFtdcMoneyType
	// 折抵总金额
	OffsetAmount tTThostFtdcMoneyType
	// 多头折抵总金额
	LongOffsetAmount tTThostFtdcMoneyType
	// 空头折抵总金额
	ShortOffsetAmount tTThostFtdcMoneyType
	// 交易所折抵总金额
	ExchOffsetAmount tTThostFtdcMoneyType
	// 交易所多头折抵总金额
	LongExchOffsetAmount tTThostFtdcMoneyType
	// 交易所空头折抵总金额
	ShortExchOffsetAmount tTThostFtdcMoneyType
	// 投机套保标志
	HedgeFlag tTThostFtdcHedgeFlagType
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
	// 投资单元代码
	InvestUnitID tTThostFtdcInvestUnitIDType
}

// 查询监控中心用户令牌
type tCThostFtdcQueryCFMMCTradingAccountTokenField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 投资单元代码
	InvestUnitID tTThostFtdcInvestUnitIDType
}

// 监控中心用户令牌
type tCThostFtdcCFMMCTradingAccountTokenField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 经纪公司统一编码
	ParticipantID tTThostFtdcParticipantIDType
	// 投资者帐号
	AccountID tTThostFtdcAccountIDType
	// 密钥编号
	KeyID tTThostFtdcSequenceNoType
	// 动态令牌
	Token tTThostFtdcCFMMCTokenType
}

// 查询产品组
type tCThostFtdcQryProductGroupField struct {
	// 产品代码
	ProductID tTThostFtdcInstrumentIDType
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
}

// 投资者品种/跨品种保证金产品组
type tCThostFtdcProductGroupField struct {
	// 产品代码
	ProductID tTThostFtdcInstrumentIDType
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
	// 产品组代码
	ProductGroupID tTThostFtdcInstrumentIDType
}

// 交易所公告
type tCThostFtdcBulletinField struct {
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
	// 交易日
	TradingDay tTThostFtdcDateType
	// 公告编号
	BulletinID tTThostFtdcBulletinIDType
	// 序列号
	SequenceNo tTThostFtdcSequenceNoType
	// 公告类型
	NewsType tTThostFtdcNewsTypeType
	// 紧急程度
	NewsUrgency tTThostFtdcNewsUrgencyType
	// 发送时间
	SendTime tTThostFtdcTimeType
	// 消息摘要
	Abstract tTThostFtdcAbstractType
	// 消息来源
	ComeFrom tTThostFtdcComeFromType
	// 消息正文
	Content tTThostFtdcContentType
	// WEB地址
	URLLink tTThostFtdcURLLinkType
	// 市场代码
	MarketID tTThostFtdcMarketIDType
}

// 查询交易所公告
type tCThostFtdcQryBulletinField struct {
	// 交易所代码
	ExchangeID tTThostFtdcExchangeIDType
	// 公告编号
	BulletinID tTThostFtdcBulletinIDType
	// 序列号
	SequenceNo tTThostFtdcSequenceNoType
	// 公告类型
	NewsType tTThostFtdcNewsTypeType
	// 紧急程度
	NewsUrgency tTThostFtdcNewsUrgencyType
}

// 转帐开户请求
type tCThostFtdcReqOpenAccountField struct {
	// 业务功能码
	TradeCode tTThostFtdcTradeCodeType
	// 银行代码
	BankID tTThostFtdcBankIDType
	// 银行分支机构代码
	BankBranchID tTThostFtdcBankBrchIDType
	// 期商代码
	BrokerID tTThostFtdcBrokerIDType
	// 期商分支机构代码
	BrokerBranchID tTThostFtdcFutureBranchIDType
	// 交易日期
	TradeDate tTThostFtdcTradeDateType
	// 交易时间
	TradeTime tTThostFtdcTradeTimeType
	// 银行流水号
	BankSerial tTThostFtdcBankSerialType
	// 交易系统日期
	TradingDay tTThostFtdcTradeDateType
	// 银期平台消息流水号
	PlateSerial tTThostFtdcSerialType
	// 最后分片标志
	LastFragment tTThostFtdcLastFragmentType
	// 会话号
	SessionID tTThostFtdcSessionIDType
	// 客户姓名
	CustomerName tTThostFtdcIndividualNameType
	// 证件类型
	IdCardType tTThostFtdcIdCardTypeType
	// 证件号码
	IdentifiedCardNo tTThostFtdcIdentifiedCardNoType
	// 性别
	Gender tTThostFtdcGenderType
	// 国家代码
	CountryCode tTThostFtdcCountryCodeType
	// 客户类型
	CustType tTThostFtdcCustTypeType
	// 地址
	Address tTThostFtdcAddressType
	// 邮编
	ZipCode tTThostFtdcZipCodeType
	// 电话号码
	Telephone tTThostFtdcTelephoneType
	// 手机
	MobilePhone tTThostFtdcMobilePhoneType
	// 传真
	Fax tTThostFtdcFaxType
	// 电子邮件
	EMail tTThostFtdcEMailType
	// 资金账户状态
	MoneyAccountStatus tTThostFtdcMoneyAccountStatusType
	// 银行帐号
	BankAccount tTThostFtdcBankAccountType
	// 银行密码
	BankPassWord tTThostFtdcPasswordType
	// 投资者帐号
	AccountID tTThostFtdcAccountIDType
	// 期货密码
	Password tTThostFtdcPasswordType
	// 安装编号
	InstallID tTThostFtdcInstallIDType
	// 验证客户证件号码标志
	VerifyCertNoFlag tTThostFtdcYesNoIndicatorType
	// 币种代码
	CurrencyID tTThostFtdcCurrencyIDType
	// 汇钞标志
	CashExchangeCode tTThostFtdcCashExchangeCodeType
	// 摘要
	Digest tTThostFtdcDigestType
	// 银行帐号类型
	BankAccType tTThostFtdcBankAccTypeType
	// 渠道标志
	DeviceID tTThostFtdcDeviceIDType
	// 期货单位帐号类型
	BankSecuAccType tTThostFtdcBankAccTypeType
	// 期货公司银行编码
	BrokerIDByBank tTThostFtdcBankCodingForFutureType
	// 期货单位帐号
	BankSecuAcc tTThostFtdcBankAccountType
	// 银行密码标志
	BankPwdFlag tTThostFtdcPwdFlagType
	// 期货资金密码核对标志
	SecuPwdFlag tTThostFtdcPwdFlagType
	// 交易柜员
	OperNo tTThostFtdcOperNoType
	// 交易ID
	TID tTThostFtdcTIDType
	// 用户标识
	UserID tTThostFtdcUserIDType
	// 长客户姓名
	LongCustomerName tTThostFtdcLongIndividualNameType
}

// 转帐销户请求
type tCThostFtdcReqCancelAccountField struct {
	// 业务功能码
	TradeCode tTThostFtdcTradeCodeType
	// 银行代码
	BankID tTThostFtdcBankIDType
	// 银行分支机构代码
	BankBranchID tTThostFtdcBankBrchIDType
	// 期商代码
	BrokerID tTThostFtdcBrokerIDType
	// 期商分支机构代码
	BrokerBranchID tTThostFtdcFutureBranchIDType
	// 交易日期
	TradeDate tTThostFtdcTradeDateType
	// 交易时间
	TradeTime tTThostFtdcTradeTimeType
	// 银行流水号
	BankSerial tTThostFtdcBankSerialType
	// 交易系统日期
	TradingDay tTThostFtdcTradeDateType
	// 银期平台消息流水号
	PlateSerial tTThostFtdcSerialType
	// 最后分片标志
	LastFragment tTThostFtdcLastFragmentType
	// 会话号
	SessionID tTThostFtdcSessionIDType
	// 客户姓名
	CustomerName tTThostFtdcIndividualNameType
	// 证件类型
	IdCardType tTThostFtdcIdCardTypeType
	// 证件号码
	IdentifiedCardNo tTThostFtdcIdentifiedCardNoType
	// 性别
	Gender tTThostFtdcGenderType
	// 国家代码
	CountryCode tTThostFtdcCountryCodeType
	// 客户类型
	CustType tTThostFtdcCustTypeType
	// 地址
	Address tTThostFtdcAddressType
	// 邮编
	ZipCode tTThostFtdcZipCodeType
	// 电话号码
	Telephone tTThostFtdcTelephoneType
	// 手机
	MobilePhone tTThostFtdcMobilePhoneType
	// 传真
	Fax tTThostFtdcFaxType
	// 电子邮件
	EMail tTThostFtdcEMailType
	// 资金账户状态
	MoneyAccountStatus tTThostFtdcMoneyAccountStatusType
	// 银行帐号
	BankAccount tTThostFtdcBankAccountType
	// 银行密码
	BankPassWord tTThostFtdcPasswordType
	// 投资者帐号
	AccountID tTThostFtdcAccountIDType
	// 期货密码
	Password tTThostFtdcPasswordType
	// 安装编号
	InstallID tTThostFtdcInstallIDType
	// 验证客户证件号码标志
	VerifyCertNoFlag tTThostFtdcYesNoIndicatorType
	// 币种代码
	CurrencyID tTThostFtdcCurrencyIDType
	// 汇钞标志
	CashExchangeCode tTThostFtdcCashExchangeCodeType
	// 摘要
	Digest tTThostFtdcDigestType
	// 银行帐号类型
	BankAccType tTThostFtdcBankAccTypeType
	// 渠道标志
	DeviceID tTThostFtdcDeviceIDType
	// 期货单位帐号类型
	BankSecuAccType tTThostFtdcBankAccTypeType
	// 期货公司银行编码
	BrokerIDByBank tTThostFtdcBankCodingForFutureType
	// 期货单位帐号
	BankSecuAcc tTThostFtdcBankAccountType
	// 银行密码标志
	BankPwdFlag tTThostFtdcPwdFlagType
	// 期货资金密码核对标志
	SecuPwdFlag tTThostFtdcPwdFlagType
	// 交易柜员
	OperNo tTThostFtdcOperNoType
	// 交易ID
	TID tTThostFtdcTIDType
	// 用户标识
	UserID tTThostFtdcUserIDType
	// 长客户姓名
	LongCustomerName tTThostFtdcLongIndividualNameType
}

// 变更银行账户请求
type tCThostFtdcReqChangeAccountField struct {
	// 业务功能码
	TradeCode tTThostFtdcTradeCodeType
	// 银行代码
	BankID tTThostFtdcBankIDType
	// 银行分支机构代码
	BankBranchID tTThostFtdcBankBrchIDType
	// 期商代码
	BrokerID tTThostFtdcBrokerIDType
	// 期商分支机构代码
	BrokerBranchID tTThostFtdcFutureBranchIDType
	// 交易日期
	TradeDate tTThostFtdcTradeDateType
	// 交易时间
	TradeTime tTThostFtdcTradeTimeType
	// 银行流水号
	BankSerial tTThostFtdcBankSerialType
	// 交易系统日期
	TradingDay tTThostFtdcTradeDateType
	// 银期平台消息流水号
	PlateSerial tTThostFtdcSerialType
	// 最后分片标志
	LastFragment tTThostFtdcLastFragmentType
	// 会话号
	SessionID tTThostFtdcSessionIDType
	// 客户姓名
	CustomerName tTThostFtdcIndividualNameType
	// 证件类型
	IdCardType tTThostFtdcIdCardTypeType
	// 证件号码
	IdentifiedCardNo tTThostFtdcIdentifiedCardNoType
	// 性别
	Gender tTThostFtdcGenderType
	// 国家代码
	CountryCode tTThostFtdcCountryCodeType
	// 客户类型
	CustType tTThostFtdcCustTypeType
	// 地址
	Address tTThostFtdcAddressType
	// 邮编
	ZipCode tTThostFtdcZipCodeType
	// 电话号码
	Telephone tTThostFtdcTelephoneType
	// 手机
	MobilePhone tTThostFtdcMobilePhoneType
	// 传真
	Fax tTThostFtdcFaxType
	// 电子邮件
	EMail tTThostFtdcEMailType
	// 资金账户状态
	MoneyAccountStatus tTThostFtdcMoneyAccountStatusType
	// 银行帐号
	BankAccount tTThostFtdcBankAccountType
	// 银行密码
	BankPassWord tTThostFtdcPasswordType
	// 新银行帐号
	NewBankAccount tTThostFtdcBankAccountType
	// 新银行密码
	NewBankPassWord tTThostFtdcPasswordType
	// 投资者帐号
	AccountID tTThostFtdcAccountIDType
	// 期货密码
	Password tTThostFtdcPasswordType
	// 银行帐号类型
	BankAccType tTThostFtdcBankAccTypeType
	// 安装编号
	InstallID tTThostFtdcInstallIDType
	// 验证客户证件号码标志
	VerifyCertNoFlag tTThostFtdcYesNoIndicatorType
	// 币种代码
	CurrencyID tTThostFtdcCurrencyIDType
	// 期货公司银行编码
	BrokerIDByBank tTThostFtdcBankCodingForFutureType
	// 银行密码标志
	BankPwdFlag tTThostFtdcPwdFlagType
	// 期货资金密码核对标志
	SecuPwdFlag tTThostFtdcPwdFlagType
	// 交易ID
	TID tTThostFtdcTIDType
	// 摘要
	Digest tTThostFtdcDigestType
	// 长客户姓名
	LongCustomerName tTThostFtdcLongIndividualNameType
}

// 转账请求
type tCThostFtdcReqTransferField struct {
	// 业务功能码
	TradeCode tTThostFtdcTradeCodeType
	// 银行代码
	BankID tTThostFtdcBankIDType
	// 银行分支机构代码
	BankBranchID tTThostFtdcBankBrchIDType
	// 期商代码
	BrokerID tTThostFtdcBrokerIDType
	// 期商分支机构代码
	BrokerBranchID tTThostFtdcFutureBranchIDType
	// 交易日期
	TradeDate tTThostFtdcTradeDateType
	// 交易时间
	TradeTime tTThostFtdcTradeTimeType
	// 银行流水号
	BankSerial tTThostFtdcBankSerialType
	// 交易系统日期
	TradingDay tTThostFtdcTradeDateType
	// 银期平台消息流水号
	PlateSerial tTThostFtdcSerialType
	// 最后分片标志
	LastFragment tTThostFtdcLastFragmentType
	// 会话号
	SessionID tTThostFtdcSessionIDType
	// 客户姓名
	CustomerName tTThostFtdcIndividualNameType
	// 证件类型
	IdCardType tTThostFtdcIdCardTypeType
	// 证件号码
	IdentifiedCardNo tTThostFtdcIdentifiedCardNoType
	// 客户类型
	CustType tTThostFtdcCustTypeType
	// 银行帐号
	BankAccount tTThostFtdcBankAccountType
	// 银行密码
	BankPassWord tTThostFtdcPasswordType
	// 投资者帐号
	AccountID tTThostFtdcAccountIDType
	// 期货密码
	Password tTThostFtdcPasswordType
	// 安装编号
	InstallID tTThostFtdcInstallIDType
	// 期货公司流水号
	FutureSerial tTThostFtdcFutureSerialType
	// 用户标识
	UserID tTThostFtdcUserIDType
	// 验证客户证件号码标志
	VerifyCertNoFlag tTThostFtdcYesNoIndicatorType
	// 币种代码
	CurrencyID tTThostFtdcCurrencyIDType
	// 转帐金额
	TradeAmount tTThostFtdcTradeAmountType
	// 期货可取金额
	FutureFetchAmount tTThostFtdcTradeAmountType
	// 费用支付标志
	FeePayFlag tTThostFtdcFeePayFlagType
	// 应收客户费用
	CustFee tTThostFtdcCustFeeType
	// 应收期货公司费用
	BrokerFee tTThostFtdcFutureFeeType
	// 发送方给接收方的消息
	Message tTThostFtdcAddInfoType
	// 摘要
	Digest tTThostFtdcDigestType
	// 银行帐号类型
	BankAccType tTThostFtdcBankAccTypeType
	// 渠道标志
	DeviceID tTThostFtdcDeviceIDType
	// 期货单位帐号类型
	BankSecuAccType tTThostFtdcBankAccTypeType
	// 期货公司银行编码
	BrokerIDByBank tTThostFtdcBankCodingForFutureType
	// 期货单位帐号
	BankSecuAcc tTThostFtdcBankAccountType
	// 银行密码标志
	BankPwdFlag tTThostFtdcPwdFlagType
	// 期货资金密码核对标志
	SecuPwdFlag tTThostFtdcPwdFlagType
	// 交易柜员
	OperNo tTThostFtdcOperNoType
	// 请求编号
	RequestID tTThostFtdcRequestIDType
	// 交易ID
	TID tTThostFtdcTIDType
	// 转账交易状态
	TransferStatus tTThostFtdcTransferStatusType
	// 长客户姓名
	LongCustomerName tTThostFtdcLongIndividualNameType
}

// 银行发起银行资金转期货响应
type tCThostFtdcRspTransferField struct {
	// 业务功能码
	TradeCode tTThostFtdcTradeCodeType
	// 银行代码
	BankID tTThostFtdcBankIDType
	// 银行分支机构代码
	BankBranchID tTThostFtdcBankBrchIDType
	// 期商代码
	BrokerID tTThostFtdcBrokerIDType
	// 期商分支机构代码
	BrokerBranchID tTThostFtdcFutureBranchIDType
	// 交易日期
	TradeDate tTThostFtdcTradeDateType
	// 交易时间
	TradeTime tTThostFtdcTradeTimeType
	// 银行流水号
	BankSerial tTThostFtdcBankSerialType
	// 交易系统日期
	TradingDay tTThostFtdcTradeDateType
	// 银期平台消息流水号
	PlateSerial tTThostFtdcSerialType
	// 最后分片标志
	LastFragment tTThostFtdcLastFragmentType
	// 会话号
	SessionID tTThostFtdcSessionIDType
	// 客户姓名
	CustomerName tTThostFtdcIndividualNameType
	// 证件类型
	IdCardType tTThostFtdcIdCardTypeType
	// 证件号码
	IdentifiedCardNo tTThostFtdcIdentifiedCardNoType
	// 客户类型
	CustType tTThostFtdcCustTypeType
	// 银行帐号
	BankAccount tTThostFtdcBankAccountType
	// 银行密码
	BankPassWord tTThostFtdcPasswordType
	// 投资者帐号
	AccountID tTThostFtdcAccountIDType
	// 期货密码
	Password tTThostFtdcPasswordType
	// 安装编号
	InstallID tTThostFtdcInstallIDType
	// 期货公司流水号
	FutureSerial tTThostFtdcFutureSerialType
	// 用户标识
	UserID tTThostFtdcUserIDType
	// 验证客户证件号码标志
	VerifyCertNoFlag tTThostFtdcYesNoIndicatorType
	// 币种代码
	CurrencyID tTThostFtdcCurrencyIDType
	// 转帐金额
	TradeAmount tTThostFtdcTradeAmountType
	// 期货可取金额
	FutureFetchAmount tTThostFtdcTradeAmountType
	// 费用支付标志
	FeePayFlag tTThostFtdcFeePayFlagType
	// 应收客户费用
	CustFee tTThostFtdcCustFeeType
	// 应收期货公司费用
	BrokerFee tTThostFtdcFutureFeeType
	// 发送方给接收方的消息
	Message tTThostFtdcAddInfoType
	// 摘要
	Digest tTThostFtdcDigestType
	// 银行帐号类型
	BankAccType tTThostFtdcBankAccTypeType
	// 渠道标志
	DeviceID tTThostFtdcDeviceIDType
	// 期货单位帐号类型
	BankSecuAccType tTThostFtdcBankAccTypeType
	// 期货公司银行编码
	BrokerIDByBank tTThostFtdcBankCodingForFutureType
	// 期货单位帐号
	BankSecuAcc tTThostFtdcBankAccountType
	// 银行密码标志
	BankPwdFlag tTThostFtdcPwdFlagType
	// 期货资金密码核对标志
	SecuPwdFlag tTThostFtdcPwdFlagType
	// 交易柜员
	OperNo tTThostFtdcOperNoType
	// 请求编号
	RequestID tTThostFtdcRequestIDType
	// 交易ID
	TID tTThostFtdcTIDType
	// 转账交易状态
	TransferStatus tTThostFtdcTransferStatusType
	// 错误代码
	ErrorID tTThostFtdcErrorIDType
	// 错误信息
	ErrorMsg tTThostFtdcErrorMsgType
	// 长客户姓名
	LongCustomerName tTThostFtdcLongIndividualNameType
}

// 冲正请求
type tCThostFtdcReqRepealField struct {
	// 冲正时间间隔
	RepealTimeInterval tTThostFtdcRepealTimeIntervalType
	// 已经冲正次数
	RepealedTimes tTThostFtdcRepealedTimesType
	// 银行冲正标志
	BankRepealFlag tTThostFtdcBankRepealFlagType
	// 期商冲正标志
	BrokerRepealFlag tTThostFtdcBrokerRepealFlagType
	// 被冲正平台流水号
	PlateRepealSerial tTThostFtdcPlateSerialType
	// 被冲正银行流水号
	BankRepealSerial tTThostFtdcBankSerialType
	// 被冲正期货流水号
	FutureRepealSerial tTThostFtdcFutureSerialType
	// 业务功能码
	TradeCode tTThostFtdcTradeCodeType
	// 银行代码
	BankID tTThostFtdcBankIDType
	// 银行分支机构代码
	BankBranchID tTThostFtdcBankBrchIDType
	// 期商代码
	BrokerID tTThostFtdcBrokerIDType
	// 期商分支机构代码
	BrokerBranchID tTThostFtdcFutureBranchIDType
	// 交易日期
	TradeDate tTThostFtdcTradeDateType
	// 交易时间
	TradeTime tTThostFtdcTradeTimeType
	// 银行流水号
	BankSerial tTThostFtdcBankSerialType
	// 交易系统日期
	TradingDay tTThostFtdcTradeDateType
	// 银期平台消息流水号
	PlateSerial tTThostFtdcSerialType
	// 最后分片标志
	LastFragment tTThostFtdcLastFragmentType
	// 会话号
	SessionID tTThostFtdcSessionIDType
	// 客户姓名
	CustomerName tTThostFtdcIndividualNameType
	// 证件类型
	IdCardType tTThostFtdcIdCardTypeType
	// 证件号码
	IdentifiedCardNo tTThostFtdcIdentifiedCardNoType
	// 客户类型
	CustType tTThostFtdcCustTypeType
	// 银行帐号
	BankAccount tTThostFtdcBankAccountType
	// 银行密码
	BankPassWord tTThostFtdcPasswordType
	// 投资者帐号
	AccountID tTThostFtdcAccountIDType
	// 期货密码
	Password tTThostFtdcPasswordType
	// 安装编号
	InstallID tTThostFtdcInstallIDType
	// 期货公司流水号
	FutureSerial tTThostFtdcFutureSerialType
	// 用户标识
	UserID tTThostFtdcUserIDType
	// 验证客户证件号码标志
	VerifyCertNoFlag tTThostFtdcYesNoIndicatorType
	// 币种代码
	CurrencyID tTThostFtdcCurrencyIDType
	// 转帐金额
	TradeAmount tTThostFtdcTradeAmountType
	// 期货可取金额
	FutureFetchAmount tTThostFtdcTradeAmountType
	// 费用支付标志
	FeePayFlag tTThostFtdcFeePayFlagType
	// 应收客户费用
	CustFee tTThostFtdcCustFeeType
	// 应收期货公司费用
	BrokerFee tTThostFtdcFutureFeeType
	// 发送方给接收方的消息
	Message tTThostFtdcAddInfoType
	// 摘要
	Digest tTThostFtdcDigestType
	// 银行帐号类型
	BankAccType tTThostFtdcBankAccTypeType
	// 渠道标志
	DeviceID tTThostFtdcDeviceIDType
	// 期货单位帐号类型
	BankSecuAccType tTThostFtdcBankAccTypeType
	// 期货公司银行编码
	BrokerIDByBank tTThostFtdcBankCodingForFutureType
	// 期货单位帐号
	BankSecuAcc tTThostFtdcBankAccountType
	// 银行密码标志
	BankPwdFlag tTThostFtdcPwdFlagType
	// 期货资金密码核对标志
	SecuPwdFlag tTThostFtdcPwdFlagType
	// 交易柜员
	OperNo tTThostFtdcOperNoType
	// 请求编号
	RequestID tTThostFtdcRequestIDType
	// 交易ID
	TID tTThostFtdcTIDType
	// 转账交易状态
	TransferStatus tTThostFtdcTransferStatusType
	// 长客户姓名
	LongCustomerName tTThostFtdcLongIndividualNameType
}

// 冲正响应
type tCThostFtdcRspRepealField struct {
	// 冲正时间间隔
	RepealTimeInterval tTThostFtdcRepealTimeIntervalType
	// 已经冲正次数
	RepealedTimes tTThostFtdcRepealedTimesType
	// 银行冲正标志
	BankRepealFlag tTThostFtdcBankRepealFlagType
	// 期商冲正标志
	BrokerRepealFlag tTThostFtdcBrokerRepealFlagType
	// 被冲正平台流水号
	PlateRepealSerial tTThostFtdcPlateSerialType
	// 被冲正银行流水号
	BankRepealSerial tTThostFtdcBankSerialType
	// 被冲正期货流水号
	FutureRepealSerial tTThostFtdcFutureSerialType
	// 业务功能码
	TradeCode tTThostFtdcTradeCodeType
	// 银行代码
	BankID tTThostFtdcBankIDType
	// 银行分支机构代码
	BankBranchID tTThostFtdcBankBrchIDType
	// 期商代码
	BrokerID tTThostFtdcBrokerIDType
	// 期商分支机构代码
	BrokerBranchID tTThostFtdcFutureBranchIDType
	// 交易日期
	TradeDate tTThostFtdcTradeDateType
	// 交易时间
	TradeTime tTThostFtdcTradeTimeType
	// 银行流水号
	BankSerial tTThostFtdcBankSerialType
	// 交易系统日期
	TradingDay tTThostFtdcTradeDateType
	// 银期平台消息流水号
	PlateSerial tTThostFtdcSerialType
	// 最后分片标志
	LastFragment tTThostFtdcLastFragmentType
	// 会话号
	SessionID tTThostFtdcSessionIDType
	// 客户姓名
	CustomerName tTThostFtdcIndividualNameType
	// 证件类型
	IdCardType tTThostFtdcIdCardTypeType
	// 证件号码
	IdentifiedCardNo tTThostFtdcIdentifiedCardNoType
	// 客户类型
	CustType tTThostFtdcCustTypeType
	// 银行帐号
	BankAccount tTThostFtdcBankAccountType
	// 银行密码
	BankPassWord tTThostFtdcPasswordType
	// 投资者帐号
	AccountID tTThostFtdcAccountIDType
	// 期货密码
	Password tTThostFtdcPasswordType
	// 安装编号
	InstallID tTThostFtdcInstallIDType
	// 期货公司流水号
	FutureSerial tTThostFtdcFutureSerialType
	// 用户标识
	UserID tTThostFtdcUserIDType
	// 验证客户证件号码标志
	VerifyCertNoFlag tTThostFtdcYesNoIndicatorType
	// 币种代码
	CurrencyID tTThostFtdcCurrencyIDType
	// 转帐金额
	TradeAmount tTThostFtdcTradeAmountType
	// 期货可取金额
	FutureFetchAmount tTThostFtdcTradeAmountType
	// 费用支付标志
	FeePayFlag tTThostFtdcFeePayFlagType
	// 应收客户费用
	CustFee tTThostFtdcCustFeeType
	// 应收期货公司费用
	BrokerFee tTThostFtdcFutureFeeType
	// 发送方给接收方的消息
	Message tTThostFtdcAddInfoType
	// 摘要
	Digest tTThostFtdcDigestType
	// 银行帐号类型
	BankAccType tTThostFtdcBankAccTypeType
	// 渠道标志
	DeviceID tTThostFtdcDeviceIDType
	// 期货单位帐号类型
	BankSecuAccType tTThostFtdcBankAccTypeType
	// 期货公司银行编码
	BrokerIDByBank tTThostFtdcBankCodingForFutureType
	// 期货单位帐号
	BankSecuAcc tTThostFtdcBankAccountType
	// 银行密码标志
	BankPwdFlag tTThostFtdcPwdFlagType
	// 期货资金密码核对标志
	SecuPwdFlag tTThostFtdcPwdFlagType
	// 交易柜员
	OperNo tTThostFtdcOperNoType
	// 请求编号
	RequestID tTThostFtdcRequestIDType
	// 交易ID
	TID tTThostFtdcTIDType
	// 转账交易状态
	TransferStatus tTThostFtdcTransferStatusType
	// 错误代码
	ErrorID tTThostFtdcErrorIDType
	// 错误信息
	ErrorMsg tTThostFtdcErrorMsgType
	// 长客户姓名
	LongCustomerName tTThostFtdcLongIndividualNameType
}

// 查询账户信息请求
type tCThostFtdcReqQueryAccountField struct {
	// 业务功能码
	TradeCode tTThostFtdcTradeCodeType
	// 银行代码
	BankID tTThostFtdcBankIDType
	// 银行分支机构代码
	BankBranchID tTThostFtdcBankBrchIDType
	// 期商代码
	BrokerID tTThostFtdcBrokerIDType
	// 期商分支机构代码
	BrokerBranchID tTThostFtdcFutureBranchIDType
	// 交易日期
	TradeDate tTThostFtdcTradeDateType
	// 交易时间
	TradeTime tTThostFtdcTradeTimeType
	// 银行流水号
	BankSerial tTThostFtdcBankSerialType
	// 交易系统日期
	TradingDay tTThostFtdcTradeDateType
	// 银期平台消息流水号
	PlateSerial tTThostFtdcSerialType
	// 最后分片标志
	LastFragment tTThostFtdcLastFragmentType
	// 会话号
	SessionID tTThostFtdcSessionIDType
	// 客户姓名
	CustomerName tTThostFtdcIndividualNameType
	// 证件类型
	IdCardType tTThostFtdcIdCardTypeType
	// 证件号码
	IdentifiedCardNo tTThostFtdcIdentifiedCardNoType
	// 客户类型
	CustType tTThostFtdcCustTypeType
	// 银行帐号
	BankAccount tTThostFtdcBankAccountType
	// 银行密码
	BankPassWord tTThostFtdcPasswordType
	// 投资者帐号
	AccountID tTThostFtdcAccountIDType
	// 期货密码
	Password tTThostFtdcPasswordType
	// 期货公司流水号
	FutureSerial tTThostFtdcFutureSerialType
	// 安装编号
	InstallID tTThostFtdcInstallIDType
	// 用户标识
	UserID tTThostFtdcUserIDType
	// 验证客户证件号码标志
	VerifyCertNoFlag tTThostFtdcYesNoIndicatorType
	// 币种代码
	CurrencyID tTThostFtdcCurrencyIDType
	// 摘要
	Digest tTThostFtdcDigestType
	// 银行帐号类型
	BankAccType tTThostFtdcBankAccTypeType
	// 渠道标志
	DeviceID tTThostFtdcDeviceIDType
	// 期货单位帐号类型
	BankSecuAccType tTThostFtdcBankAccTypeType
	// 期货公司银行编码
	BrokerIDByBank tTThostFtdcBankCodingForFutureType
	// 期货单位帐号
	BankSecuAcc tTThostFtdcBankAccountType
	// 银行密码标志
	BankPwdFlag tTThostFtdcPwdFlagType
	// 期货资金密码核对标志
	SecuPwdFlag tTThostFtdcPwdFlagType
	// 交易柜员
	OperNo tTThostFtdcOperNoType
	// 请求编号
	RequestID tTThostFtdcRequestIDType
	// 交易ID
	TID tTThostFtdcTIDType
	// 长客户姓名
	LongCustomerName tTThostFtdcLongIndividualNameType
}

// 查询账户信息响应
type tCThostFtdcRspQueryAccountField struct {
	// 业务功能码
	TradeCode tTThostFtdcTradeCodeType
	// 银行代码
	BankID tTThostFtdcBankIDType
	// 银行分支机构代码
	BankBranchID tTThostFtdcBankBrchIDType
	// 期商代码
	BrokerID tTThostFtdcBrokerIDType
	// 期商分支机构代码
	BrokerBranchID tTThostFtdcFutureBranchIDType
	// 交易日期
	TradeDate tTThostFtdcTradeDateType
	// 交易时间
	TradeTime tTThostFtdcTradeTimeType
	// 银行流水号
	BankSerial tTThostFtdcBankSerialType
	// 交易系统日期
	TradingDay tTThostFtdcTradeDateType
	// 银期平台消息流水号
	PlateSerial tTThostFtdcSerialType
	// 最后分片标志
	LastFragment tTThostFtdcLastFragmentType
	// 会话号
	SessionID tTThostFtdcSessionIDType
	// 客户姓名
	CustomerName tTThostFtdcIndividualNameType
	// 证件类型
	IdCardType tTThostFtdcIdCardTypeType
	// 证件号码
	IdentifiedCardNo tTThostFtdcIdentifiedCardNoType
	// 客户类型
	CustType tTThostFtdcCustTypeType
	// 银行帐号
	BankAccount tTThostFtdcBankAccountType
	// 银行密码
	BankPassWord tTThostFtdcPasswordType
	// 投资者帐号
	AccountID tTThostFtdcAccountIDType
	// 期货密码
	Password tTThostFtdcPasswordType
	// 期货公司流水号
	FutureSerial tTThostFtdcFutureSerialType
	// 安装编号
	InstallID tTThostFtdcInstallIDType
	// 用户标识
	UserID tTThostFtdcUserIDType
	// 验证客户证件号码标志
	VerifyCertNoFlag tTThostFtdcYesNoIndicatorType
	// 币种代码
	CurrencyID tTThostFtdcCurrencyIDType
	// 摘要
	Digest tTThostFtdcDigestType
	// 银行帐号类型
	BankAccType tTThostFtdcBankAccTypeType
	// 渠道标志
	DeviceID tTThostFtdcDeviceIDType
	// 期货单位帐号类型
	BankSecuAccType tTThostFtdcBankAccTypeType
	// 期货公司银行编码
	BrokerIDByBank tTThostFtdcBankCodingForFutureType
	// 期货单位帐号
	BankSecuAcc tTThostFtdcBankAccountType
	// 银行密码标志
	BankPwdFlag tTThostFtdcPwdFlagType
	// 期货资金密码核对标志
	SecuPwdFlag tTThostFtdcPwdFlagType
	// 交易柜员
	OperNo tTThostFtdcOperNoType
	// 请求编号
	RequestID tTThostFtdcRequestIDType
	// 交易ID
	TID tTThostFtdcTIDType
	// 银行可用金额
	BankUseAmount tTThostFtdcTradeAmountType
	// 银行可取金额
	BankFetchAmount tTThostFtdcTradeAmountType
	// 长客户姓名
	LongCustomerName tTThostFtdcLongIndividualNameType
}

// 期商签到签退
type tCThostFtdcFutureSignIOField struct {
	// 业务功能码
	TradeCode tTThostFtdcTradeCodeType
	// 银行代码
	BankID tTThostFtdcBankIDType
	// 银行分支机构代码
	BankBranchID tTThostFtdcBankBrchIDType
	// 期商代码
	BrokerID tTThostFtdcBrokerIDType
	// 期商分支机构代码
	BrokerBranchID tTThostFtdcFutureBranchIDType
	// 交易日期
	TradeDate tTThostFtdcTradeDateType
	// 交易时间
	TradeTime tTThostFtdcTradeTimeType
	// 银行流水号
	BankSerial tTThostFtdcBankSerialType
	// 交易系统日期
	TradingDay tTThostFtdcTradeDateType
	// 银期平台消息流水号
	PlateSerial tTThostFtdcSerialType
	// 最后分片标志
	LastFragment tTThostFtdcLastFragmentType
	// 会话号
	SessionID tTThostFtdcSessionIDType
	// 安装编号
	InstallID tTThostFtdcInstallIDType
	// 用户标识
	UserID tTThostFtdcUserIDType
	// 摘要
	Digest tTThostFtdcDigestType
	// 币种代码
	CurrencyID tTThostFtdcCurrencyIDType
	// 渠道标志
	DeviceID tTThostFtdcDeviceIDType
	// 期货公司银行编码
	BrokerIDByBank tTThostFtdcBankCodingForFutureType
	// 交易柜员
	OperNo tTThostFtdcOperNoType
	// 请求编号
	RequestID tTThostFtdcRequestIDType
	// 交易ID
	TID tTThostFtdcTIDType
}

// 期商签到响应
type tCThostFtdcRspFutureSignInField struct {
	// 业务功能码
	TradeCode tTThostFtdcTradeCodeType
	// 银行代码
	BankID tTThostFtdcBankIDType
	// 银行分支机构代码
	BankBranchID tTThostFtdcBankBrchIDType
	// 期商代码
	BrokerID tTThostFtdcBrokerIDType
	// 期商分支机构代码
	BrokerBranchID tTThostFtdcFutureBranchIDType
	// 交易日期
	TradeDate tTThostFtdcTradeDateType
	// 交易时间
	TradeTime tTThostFtdcTradeTimeType
	// 银行流水号
	BankSerial tTThostFtdcBankSerialType
	// 交易系统日期
	TradingDay tTThostFtdcTradeDateType
	// 银期平台消息流水号
	PlateSerial tTThostFtdcSerialType
	// 最后分片标志
	LastFragment tTThostFtdcLastFragmentType
	// 会话号
	SessionID tTThostFtdcSessionIDType
	// 安装编号
	InstallID tTThostFtdcInstallIDType
	// 用户标识
	UserID tTThostFtdcUserIDType
	// 摘要
	Digest tTThostFtdcDigestType
	// 币种代码
	CurrencyID tTThostFtdcCurrencyIDType
	// 渠道标志
	DeviceID tTThostFtdcDeviceIDType
	// 期货公司银行编码
	BrokerIDByBank tTThostFtdcBankCodingForFutureType
	// 交易柜员
	OperNo tTThostFtdcOperNoType
	// 请求编号
	RequestID tTThostFtdcRequestIDType
	// 交易ID
	TID tTThostFtdcTIDType
	// 错误代码
	ErrorID tTThostFtdcErrorIDType
	// 错误信息
	ErrorMsg tTThostFtdcErrorMsgType
	// PIN密钥
	PinKey tTThostFtdcPasswordKeyType
	// MAC密钥
	MacKey tTThostFtdcPasswordKeyType
}

// 期商签退请求
type tCThostFtdcReqFutureSignOutField struct {
	// 业务功能码
	TradeCode tTThostFtdcTradeCodeType
	// 银行代码
	BankID tTThostFtdcBankIDType
	// 银行分支机构代码
	BankBranchID tTThostFtdcBankBrchIDType
	// 期商代码
	BrokerID tTThostFtdcBrokerIDType
	// 期商分支机构代码
	BrokerBranchID tTThostFtdcFutureBranchIDType
	// 交易日期
	TradeDate tTThostFtdcTradeDateType
	// 交易时间
	TradeTime tTThostFtdcTradeTimeType
	// 银行流水号
	BankSerial tTThostFtdcBankSerialType
	// 交易系统日期
	TradingDay tTThostFtdcTradeDateType
	// 银期平台消息流水号
	PlateSerial tTThostFtdcSerialType
	// 最后分片标志
	LastFragment tTThostFtdcLastFragmentType
	// 会话号
	SessionID tTThostFtdcSessionIDType
	// 安装编号
	InstallID tTThostFtdcInstallIDType
	// 用户标识
	UserID tTThostFtdcUserIDType
	// 摘要
	Digest tTThostFtdcDigestType
	// 币种代码
	CurrencyID tTThostFtdcCurrencyIDType
	// 渠道标志
	DeviceID tTThostFtdcDeviceIDType
	// 期货公司银行编码
	BrokerIDByBank tTThostFtdcBankCodingForFutureType
	// 交易柜员
	OperNo tTThostFtdcOperNoType
	// 请求编号
	RequestID tTThostFtdcRequestIDType
	// 交易ID
	TID tTThostFtdcTIDType
}

// 期商签退响应
type tCThostFtdcRspFutureSignOutField struct {
	// 业务功能码
	TradeCode tTThostFtdcTradeCodeType
	// 银行代码
	BankID tTThostFtdcBankIDType
	// 银行分支机构代码
	BankBranchID tTThostFtdcBankBrchIDType
	// 期商代码
	BrokerID tTThostFtdcBrokerIDType
	// 期商分支机构代码
	BrokerBranchID tTThostFtdcFutureBranchIDType
	// 交易日期
	TradeDate tTThostFtdcTradeDateType
	// 交易时间
	TradeTime tTThostFtdcTradeTimeType
	// 银行流水号
	BankSerial tTThostFtdcBankSerialType
	// 交易系统日期
	TradingDay tTThostFtdcTradeDateType
	// 银期平台消息流水号
	PlateSerial tTThostFtdcSerialType
	// 最后分片标志
	LastFragment tTThostFtdcLastFragmentType
	// 会话号
	SessionID tTThostFtdcSessionIDType
	// 安装编号
	InstallID tTThostFtdcInstallIDType
	// 用户标识
	UserID tTThostFtdcUserIDType
	// 摘要
	Digest tTThostFtdcDigestType
	// 币种代码
	CurrencyID tTThostFtdcCurrencyIDType
	// 渠道标志
	DeviceID tTThostFtdcDeviceIDType
	// 期货公司银行编码
	BrokerIDByBank tTThostFtdcBankCodingForFutureType
	// 交易柜员
	OperNo tTThostFtdcOperNoType
	// 请求编号
	RequestID tTThostFtdcRequestIDType
	// 交易ID
	TID tTThostFtdcTIDType
	// 错误代码
	ErrorID tTThostFtdcErrorIDType
	// 错误信息
	ErrorMsg tTThostFtdcErrorMsgType
}

// 查询指定流水号的交易结果请求
type tCThostFtdcReqQueryTradeResultBySerialField struct {
	// 业务功能码
	TradeCode tTThostFtdcTradeCodeType
	// 银行代码
	BankID tTThostFtdcBankIDType
	// 银行分支机构代码
	BankBranchID tTThostFtdcBankBrchIDType
	// 期商代码
	BrokerID tTThostFtdcBrokerIDType
	// 期商分支机构代码
	BrokerBranchID tTThostFtdcFutureBranchIDType
	// 交易日期
	TradeDate tTThostFtdcTradeDateType
	// 交易时间
	TradeTime tTThostFtdcTradeTimeType
	// 银行流水号
	BankSerial tTThostFtdcBankSerialType
	// 交易系统日期
	TradingDay tTThostFtdcTradeDateType
	// 银期平台消息流水号
	PlateSerial tTThostFtdcSerialType
	// 最后分片标志
	LastFragment tTThostFtdcLastFragmentType
	// 会话号
	SessionID tTThostFtdcSessionIDType
	// 流水号
	Reference tTThostFtdcSerialType
	// 本流水号发布者的机构类型
	RefrenceIssureType tTThostFtdcInstitutionTypeType
	// 本流水号发布者机构编码
	RefrenceIssure tTThostFtdcOrganCodeType
	// 客户姓名
	CustomerName tTThostFtdcIndividualNameType
	// 证件类型
	IdCardType tTThostFtdcIdCardTypeType
	// 证件号码
	IdentifiedCardNo tTThostFtdcIdentifiedCardNoType
	// 客户类型
	CustType tTThostFtdcCustTypeType
	// 银行帐号
	BankAccount tTThostFtdcBankAccountType
	// 银行密码
	BankPassWord tTThostFtdcPasswordType
	// 投资者帐号
	AccountID tTThostFtdcAccountIDType
	// 期货密码
	Password tTThostFtdcPasswordType
	// 币种代码
	CurrencyID tTThostFtdcCurrencyIDType
	// 转帐金额
	TradeAmount tTThostFtdcTradeAmountType
	// 摘要
	Digest tTThostFtdcDigestType
	// 长客户姓名
	LongCustomerName tTThostFtdcLongIndividualNameType
}

// 查询指定流水号的交易结果响应
type tCThostFtdcRspQueryTradeResultBySerialField struct {
	// 业务功能码
	TradeCode tTThostFtdcTradeCodeType
	// 银行代码
	BankID tTThostFtdcBankIDType
	// 银行分支机构代码
	BankBranchID tTThostFtdcBankBrchIDType
	// 期商代码
	BrokerID tTThostFtdcBrokerIDType
	// 期商分支机构代码
	BrokerBranchID tTThostFtdcFutureBranchIDType
	// 交易日期
	TradeDate tTThostFtdcTradeDateType
	// 交易时间
	TradeTime tTThostFtdcTradeTimeType
	// 银行流水号
	BankSerial tTThostFtdcBankSerialType
	// 交易系统日期
	TradingDay tTThostFtdcTradeDateType
	// 银期平台消息流水号
	PlateSerial tTThostFtdcSerialType
	// 最后分片标志
	LastFragment tTThostFtdcLastFragmentType
	// 会话号
	SessionID tTThostFtdcSessionIDType
	// 错误代码
	ErrorID tTThostFtdcErrorIDType
	// 错误信息
	ErrorMsg tTThostFtdcErrorMsgType
	// 流水号
	Reference tTThostFtdcSerialType
	// 本流水号发布者的机构类型
	RefrenceIssureType tTThostFtdcInstitutionTypeType
	// 本流水号发布者机构编码
	RefrenceIssure tTThostFtdcOrganCodeType
	// 原始返回代码
	OriginReturnCode tTThostFtdcReturnCodeType
	// 原始返回码描述
	OriginDescrInfoForReturnCode tTThostFtdcDescrInfoForReturnCodeType
	// 银行帐号
	BankAccount tTThostFtdcBankAccountType
	// 银行密码
	BankPassWord tTThostFtdcPasswordType
	// 投资者帐号
	AccountID tTThostFtdcAccountIDType
	// 期货密码
	Password tTThostFtdcPasswordType
	// 币种代码
	CurrencyID tTThostFtdcCurrencyIDType
	// 转帐金额
	TradeAmount tTThostFtdcTradeAmountType
	// 摘要
	Digest tTThostFtdcDigestType
}

// 日终文件就绪请求
type tCThostFtdcReqDayEndFileReadyField struct {
	// 业务功能码
	TradeCode tTThostFtdcTradeCodeType
	// 银行代码
	BankID tTThostFtdcBankIDType
	// 银行分支机构代码
	BankBranchID tTThostFtdcBankBrchIDType
	// 期商代码
	BrokerID tTThostFtdcBrokerIDType
	// 期商分支机构代码
	BrokerBranchID tTThostFtdcFutureBranchIDType
	// 交易日期
	TradeDate tTThostFtdcTradeDateType
	// 交易时间
	TradeTime tTThostFtdcTradeTimeType
	// 银行流水号
	BankSerial tTThostFtdcBankSerialType
	// 交易系统日期
	TradingDay tTThostFtdcTradeDateType
	// 银期平台消息流水号
	PlateSerial tTThostFtdcSerialType
	// 最后分片标志
	LastFragment tTThostFtdcLastFragmentType
	// 会话号
	SessionID tTThostFtdcSessionIDType
	// 文件业务功能
	FileBusinessCode tTThostFtdcFileBusinessCodeType
	// 摘要
	Digest tTThostFtdcDigestType
}

// 返回结果
type tCThostFtdcReturnResultField struct {
	// 返回代码
	ReturnCode tTThostFtdcReturnCodeType
	// 返回码描述
	DescrInfoForReturnCode tTThostFtdcDescrInfoForReturnCodeType
}

// 验证期货资金密码
type tCThostFtdcVerifyFuturePasswordField struct {
	// 业务功能码
	TradeCode tTThostFtdcTradeCodeType
	// 银行代码
	BankID tTThostFtdcBankIDType
	// 银行分支机构代码
	BankBranchID tTThostFtdcBankBrchIDType
	// 期商代码
	BrokerID tTThostFtdcBrokerIDType
	// 期商分支机构代码
	BrokerBranchID tTThostFtdcFutureBranchIDType
	// 交易日期
	TradeDate tTThostFtdcTradeDateType
	// 交易时间
	TradeTime tTThostFtdcTradeTimeType
	// 银行流水号
	BankSerial tTThostFtdcBankSerialType
	// 交易系统日期
	TradingDay tTThostFtdcTradeDateType
	// 银期平台消息流水号
	PlateSerial tTThostFtdcSerialType
	// 最后分片标志
	LastFragment tTThostFtdcLastFragmentType
	// 会话号
	SessionID tTThostFtdcSessionIDType
	// 投资者帐号
	AccountID tTThostFtdcAccountIDType
	// 期货密码
	Password tTThostFtdcPasswordType
	// 银行帐号
	BankAccount tTThostFtdcBankAccountType
	// 银行密码
	BankPassWord tTThostFtdcPasswordType
	// 安装编号
	InstallID tTThostFtdcInstallIDType
	// 交易ID
	TID tTThostFtdcTIDType
	// 币种代码
	CurrencyID tTThostFtdcCurrencyIDType
}

// 验证客户信息
type tCThostFtdcVerifyCustInfoField struct {
	// 客户姓名
	CustomerName tTThostFtdcIndividualNameType
	// 证件类型
	IdCardType tTThostFtdcIdCardTypeType
	// 证件号码
	IdentifiedCardNo tTThostFtdcIdentifiedCardNoType
	// 客户类型
	CustType tTThostFtdcCustTypeType
	// 长客户姓名
	LongCustomerName tTThostFtdcLongIndividualNameType
}

// 验证期货资金密码和客户信息
type tCThostFtdcVerifyFuturePasswordAndCustInfoField struct {
	// 客户姓名
	CustomerName tTThostFtdcIndividualNameType
	// 证件类型
	IdCardType tTThostFtdcIdCardTypeType
	// 证件号码
	IdentifiedCardNo tTThostFtdcIdentifiedCardNoType
	// 客户类型
	CustType tTThostFtdcCustTypeType
	// 投资者帐号
	AccountID tTThostFtdcAccountIDType
	// 期货密码
	Password tTThostFtdcPasswordType
	// 币种代码
	CurrencyID tTThostFtdcCurrencyIDType
	// 长客户姓名
	LongCustomerName tTThostFtdcLongIndividualNameType
}

// 验证期货资金密码和客户信息
type tCThostFtdcDepositResultInformField struct {
	// 出入金流水号，该流水号为银期报盘返回的流水号
	DepositSeqNo tTThostFtdcDepositSeqNoType
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 入金金额
	Deposit tTThostFtdcMoneyType
	// 请求编号
	RequestID tTThostFtdcRequestIDType
	// 返回代码
	ReturnCode tTThostFtdcReturnCodeType
	// 返回码描述
	DescrInfoForReturnCode tTThostFtdcDescrInfoForReturnCodeType
}

// 交易核心向银期报盘发出密钥同步请求
type tCThostFtdcReqSyncKeyField struct {
	// 业务功能码
	TradeCode tTThostFtdcTradeCodeType
	// 银行代码
	BankID tTThostFtdcBankIDType
	// 银行分支机构代码
	BankBranchID tTThostFtdcBankBrchIDType
	// 期商代码
	BrokerID tTThostFtdcBrokerIDType
	// 期商分支机构代码
	BrokerBranchID tTThostFtdcFutureBranchIDType
	// 交易日期
	TradeDate tTThostFtdcTradeDateType
	// 交易时间
	TradeTime tTThostFtdcTradeTimeType
	// 银行流水号
	BankSerial tTThostFtdcBankSerialType
	// 交易系统日期
	TradingDay tTThostFtdcTradeDateType
	// 银期平台消息流水号
	PlateSerial tTThostFtdcSerialType
	// 最后分片标志
	LastFragment tTThostFtdcLastFragmentType
	// 会话号
	SessionID tTThostFtdcSessionIDType
	// 安装编号
	InstallID tTThostFtdcInstallIDType
	// 用户标识
	UserID tTThostFtdcUserIDType
	// 交易核心给银期报盘的消息
	Message tTThostFtdcAddInfoType
	// 渠道标志
	DeviceID tTThostFtdcDeviceIDType
	// 期货公司银行编码
	BrokerIDByBank tTThostFtdcBankCodingForFutureType
	// 交易柜员
	OperNo tTThostFtdcOperNoType
	// 请求编号
	RequestID tTThostFtdcRequestIDType
	// 交易ID
	TID tTThostFtdcTIDType
}

// 交易核心向银期报盘发出密钥同步响应
type tCThostFtdcRspSyncKeyField struct {
	// 业务功能码
	TradeCode tTThostFtdcTradeCodeType
	// 银行代码
	BankID tTThostFtdcBankIDType
	// 银行分支机构代码
	BankBranchID tTThostFtdcBankBrchIDType
	// 期商代码
	BrokerID tTThostFtdcBrokerIDType
	// 期商分支机构代码
	BrokerBranchID tTThostFtdcFutureBranchIDType
	// 交易日期
	TradeDate tTThostFtdcTradeDateType
	// 交易时间
	TradeTime tTThostFtdcTradeTimeType
	// 银行流水号
	BankSerial tTThostFtdcBankSerialType
	// 交易系统日期
	TradingDay tTThostFtdcTradeDateType
	// 银期平台消息流水号
	PlateSerial tTThostFtdcSerialType
	// 最后分片标志
	LastFragment tTThostFtdcLastFragmentType
	// 会话号
	SessionID tTThostFtdcSessionIDType
	// 安装编号
	InstallID tTThostFtdcInstallIDType
	// 用户标识
	UserID tTThostFtdcUserIDType
	// 交易核心给银期报盘的消息
	Message tTThostFtdcAddInfoType
	// 渠道标志
	DeviceID tTThostFtdcDeviceIDType
	// 期货公司银行编码
	BrokerIDByBank tTThostFtdcBankCodingForFutureType
	// 交易柜员
	OperNo tTThostFtdcOperNoType
	// 请求编号
	RequestID tTThostFtdcRequestIDType
	// 交易ID
	TID tTThostFtdcTIDType
	// 错误代码
	ErrorID tTThostFtdcErrorIDType
	// 错误信息
	ErrorMsg tTThostFtdcErrorMsgType
}

// 查询账户信息通知
type tCThostFtdcNotifyQueryAccountField struct {
	// 业务功能码
	TradeCode tTThostFtdcTradeCodeType
	// 银行代码
	BankID tTThostFtdcBankIDType
	// 银行分支机构代码
	BankBranchID tTThostFtdcBankBrchIDType
	// 期商代码
	BrokerID tTThostFtdcBrokerIDType
	// 期商分支机构代码
	BrokerBranchID tTThostFtdcFutureBranchIDType
	// 交易日期
	TradeDate tTThostFtdcTradeDateType
	// 交易时间
	TradeTime tTThostFtdcTradeTimeType
	// 银行流水号
	BankSerial tTThostFtdcBankSerialType
	// 交易系统日期
	TradingDay tTThostFtdcTradeDateType
	// 银期平台消息流水号
	PlateSerial tTThostFtdcSerialType
	// 最后分片标志
	LastFragment tTThostFtdcLastFragmentType
	// 会话号
	SessionID tTThostFtdcSessionIDType
	// 客户姓名
	CustomerName tTThostFtdcIndividualNameType
	// 证件类型
	IdCardType tTThostFtdcIdCardTypeType
	// 证件号码
	IdentifiedCardNo tTThostFtdcIdentifiedCardNoType
	// 客户类型
	CustType tTThostFtdcCustTypeType
	// 银行帐号
	BankAccount tTThostFtdcBankAccountType
	// 银行密码
	BankPassWord tTThostFtdcPasswordType
	// 投资者帐号
	AccountID tTThostFtdcAccountIDType
	// 期货密码
	Password tTThostFtdcPasswordType
	// 期货公司流水号
	FutureSerial tTThostFtdcFutureSerialType
	// 安装编号
	InstallID tTThostFtdcInstallIDType
	// 用户标识
	UserID tTThostFtdcUserIDType
	// 验证客户证件号码标志
	VerifyCertNoFlag tTThostFtdcYesNoIndicatorType
	// 币种代码
	CurrencyID tTThostFtdcCurrencyIDType
	// 摘要
	Digest tTThostFtdcDigestType
	// 银行帐号类型
	BankAccType tTThostFtdcBankAccTypeType
	// 渠道标志
	DeviceID tTThostFtdcDeviceIDType
	// 期货单位帐号类型
	BankSecuAccType tTThostFtdcBankAccTypeType
	// 期货公司银行编码
	BrokerIDByBank tTThostFtdcBankCodingForFutureType
	// 期货单位帐号
	BankSecuAcc tTThostFtdcBankAccountType
	// 银行密码标志
	BankPwdFlag tTThostFtdcPwdFlagType
	// 期货资金密码核对标志
	SecuPwdFlag tTThostFtdcPwdFlagType
	// 交易柜员
	OperNo tTThostFtdcOperNoType
	// 请求编号
	RequestID tTThostFtdcRequestIDType
	// 交易ID
	TID tTThostFtdcTIDType
	// 银行可用金额
	BankUseAmount tTThostFtdcTradeAmountType
	// 银行可取金额
	BankFetchAmount tTThostFtdcTradeAmountType
	// 错误代码
	ErrorID tTThostFtdcErrorIDType
	// 错误信息
	ErrorMsg tTThostFtdcErrorMsgType
	// 长客户姓名
	LongCustomerName tTThostFtdcLongIndividualNameType
}

// 银期转账交易流水表
type tCThostFtdcTransferSerialField struct {
	// 平台流水号
	PlateSerial tTThostFtdcPlateSerialType
	// 交易发起方日期
	TradeDate tTThostFtdcTradeDateType
	// 交易日期
	TradingDay tTThostFtdcDateType
	// 交易时间
	TradeTime tTThostFtdcTradeTimeType
	// 交易代码
	TradeCode tTThostFtdcTradeCodeType
	// 会话编号
	SessionID tTThostFtdcSessionIDType
	// 银行编码
	BankID tTThostFtdcBankIDType
	// 银行分支机构编码
	BankBranchID tTThostFtdcBankBrchIDType
	// 银行帐号类型
	BankAccType tTThostFtdcBankAccTypeType
	// 银行帐号
	BankAccount tTThostFtdcBankAccountType
	// 银行流水号
	BankSerial tTThostFtdcBankSerialType
	// 期货公司编码
	BrokerID tTThostFtdcBrokerIDType
	// 期商分支机构代码
	BrokerBranchID tTThostFtdcFutureBranchIDType
	// 期货公司帐号类型
	FutureAccType tTThostFtdcFutureAccTypeType
	// 投资者帐号
	AccountID tTThostFtdcAccountIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
	// 期货公司流水号
	FutureSerial tTThostFtdcFutureSerialType
	// 证件类型
	IdCardType tTThostFtdcIdCardTypeType
	// 证件号码
	IdentifiedCardNo tTThostFtdcIdentifiedCardNoType
	// 币种代码
	CurrencyID tTThostFtdcCurrencyIDType
	// 交易金额
	TradeAmount tTThostFtdcTradeAmountType
	// 应收客户费用
	CustFee tTThostFtdcCustFeeType
	// 应收期货公司费用
	BrokerFee tTThostFtdcFutureFeeType
	// 有效标志
	AvailabilityFlag tTThostFtdcAvailabilityFlagType
	// 操作员
	OperatorCode tTThostFtdcOperatorCodeType
	// 新银行帐号
	BankNewAccount tTThostFtdcBankAccountType
	// 错误代码
	ErrorID tTThostFtdcErrorIDType
	// 错误信息
	ErrorMsg tTThostFtdcErrorMsgType
}

// 请求查询转帐流水
type tCThostFtdcQryTransferSerialField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者帐号
	AccountID tTThostFtdcAccountIDType
	// 银行编码
	BankID tTThostFtdcBankIDType
	// 币种代码
	CurrencyID tTThostFtdcCurrencyIDType
}

// 期商签到通知
type tCThostFtdcNotifyFutureSignInField struct {
	// 业务功能码
	TradeCode tTThostFtdcTradeCodeType
	// 银行代码
	BankID tTThostFtdcBankIDType
	// 银行分支机构代码
	BankBranchID tTThostFtdcBankBrchIDType
	// 期商代码
	BrokerID tTThostFtdcBrokerIDType
	// 期商分支机构代码
	BrokerBranchID tTThostFtdcFutureBranchIDType
	// 交易日期
	TradeDate tTThostFtdcTradeDateType
	// 交易时间
	TradeTime tTThostFtdcTradeTimeType
	// 银行流水号
	BankSerial tTThostFtdcBankSerialType
	// 交易系统日期
	TradingDay tTThostFtdcTradeDateType
	// 银期平台消息流水号
	PlateSerial tTThostFtdcSerialType
	// 最后分片标志
	LastFragment tTThostFtdcLastFragmentType
	// 会话号
	SessionID tTThostFtdcSessionIDType
	// 安装编号
	InstallID tTThostFtdcInstallIDType
	// 用户标识
	UserID tTThostFtdcUserIDType
	// 摘要
	Digest tTThostFtdcDigestType
	// 币种代码
	CurrencyID tTThostFtdcCurrencyIDType
	// 渠道标志
	DeviceID tTThostFtdcDeviceIDType
	// 期货公司银行编码
	BrokerIDByBank tTThostFtdcBankCodingForFutureType
	// 交易柜员
	OperNo tTThostFtdcOperNoType
	// 请求编号
	RequestID tTThostFtdcRequestIDType
	// 交易ID
	TID tTThostFtdcTIDType
	// 错误代码
	ErrorID tTThostFtdcErrorIDType
	// 错误信息
	ErrorMsg tTThostFtdcErrorMsgType
	// PIN密钥
	PinKey tTThostFtdcPasswordKeyType
	// MAC密钥
	MacKey tTThostFtdcPasswordKeyType
}

// 期商签退通知
type tCThostFtdcNotifyFutureSignOutField struct {
	// 业务功能码
	TradeCode tTThostFtdcTradeCodeType
	// 银行代码
	BankID tTThostFtdcBankIDType
	// 银行分支机构代码
	BankBranchID tTThostFtdcBankBrchIDType
	// 期商代码
	BrokerID tTThostFtdcBrokerIDType
	// 期商分支机构代码
	BrokerBranchID tTThostFtdcFutureBranchIDType
	// 交易日期
	TradeDate tTThostFtdcTradeDateType
	// 交易时间
	TradeTime tTThostFtdcTradeTimeType
	// 银行流水号
	BankSerial tTThostFtdcBankSerialType
	// 交易系统日期
	TradingDay tTThostFtdcTradeDateType
	// 银期平台消息流水号
	PlateSerial tTThostFtdcSerialType
	// 最后分片标志
	LastFragment tTThostFtdcLastFragmentType
	// 会话号
	SessionID tTThostFtdcSessionIDType
	// 安装编号
	InstallID tTThostFtdcInstallIDType
	// 用户标识
	UserID tTThostFtdcUserIDType
	// 摘要
	Digest tTThostFtdcDigestType
	// 币种代码
	CurrencyID tTThostFtdcCurrencyIDType
	// 渠道标志
	DeviceID tTThostFtdcDeviceIDType
	// 期货公司银行编码
	BrokerIDByBank tTThostFtdcBankCodingForFutureType
	// 交易柜员
	OperNo tTThostFtdcOperNoType
	// 请求编号
	RequestID tTThostFtdcRequestIDType
	// 交易ID
	TID tTThostFtdcTIDType
	// 错误代码
	ErrorID tTThostFtdcErrorIDType
	// 错误信息
	ErrorMsg tTThostFtdcErrorMsgType
}

// 交易核心向银期报盘发出密钥同步处理结果的通知
type tCThostFtdcNotifySyncKeyField struct {
	// 业务功能码
	TradeCode tTThostFtdcTradeCodeType
	// 银行代码
	BankID tTThostFtdcBankIDType
	// 银行分支机构代码
	BankBranchID tTThostFtdcBankBrchIDType
	// 期商代码
	BrokerID tTThostFtdcBrokerIDType
	// 期商分支机构代码
	BrokerBranchID tTThostFtdcFutureBranchIDType
	// 交易日期
	TradeDate tTThostFtdcTradeDateType
	// 交易时间
	TradeTime tTThostFtdcTradeTimeType
	// 银行流水号
	BankSerial tTThostFtdcBankSerialType
	// 交易系统日期
	TradingDay tTThostFtdcTradeDateType
	// 银期平台消息流水号
	PlateSerial tTThostFtdcSerialType
	// 最后分片标志
	LastFragment tTThostFtdcLastFragmentType
	// 会话号
	SessionID tTThostFtdcSessionIDType
	// 安装编号
	InstallID tTThostFtdcInstallIDType
	// 用户标识
	UserID tTThostFtdcUserIDType
	// 交易核心给银期报盘的消息
	Message tTThostFtdcAddInfoType
	// 渠道标志
	DeviceID tTThostFtdcDeviceIDType
	// 期货公司银行编码
	BrokerIDByBank tTThostFtdcBankCodingForFutureType
	// 交易柜员
	OperNo tTThostFtdcOperNoType
	// 请求编号
	RequestID tTThostFtdcRequestIDType
	// 交易ID
	TID tTThostFtdcTIDType
	// 错误代码
	ErrorID tTThostFtdcErrorIDType
	// 错误信息
	ErrorMsg tTThostFtdcErrorMsgType
}

// 请求查询银期签约关系
type tCThostFtdcQryAccountregisterField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者帐号
	AccountID tTThostFtdcAccountIDType
	// 银行编码
	BankID tTThostFtdcBankIDType
	// 银行分支机构编码
	BankBranchID tTThostFtdcBankBrchIDType
	// 币种代码
	CurrencyID tTThostFtdcCurrencyIDType
}

// 客户开销户信息表
type tCThostFtdcAccountregisterField struct {
	// 交易日期
	TradeDay tTThostFtdcTradeDateType
	// 银行编码
	BankID tTThostFtdcBankIDType
	// 银行分支机构编码
	BankBranchID tTThostFtdcBankBrchIDType
	// 银行帐号
	BankAccount tTThostFtdcBankAccountType
	// 期货公司编码
	BrokerID tTThostFtdcBrokerIDType
	// 期货公司分支机构编码
	BrokerBranchID tTThostFtdcFutureBranchIDType
	// 投资者帐号
	AccountID tTThostFtdcAccountIDType
	// 证件类型
	IdCardType tTThostFtdcIdCardTypeType
	// 证件号码
	IdentifiedCardNo tTThostFtdcIdentifiedCardNoType
	// 客户姓名
	CustomerName tTThostFtdcIndividualNameType
	// 币种代码
	CurrencyID tTThostFtdcCurrencyIDType
	// 开销户类别
	OpenOrDestroy tTThostFtdcOpenOrDestroyType
	// 签约日期
	RegDate tTThostFtdcTradeDateType
	// 解约日期
	OutDate tTThostFtdcTradeDateType
	// 交易ID
	TID tTThostFtdcTIDType
	// 客户类型
	CustType tTThostFtdcCustTypeType
	// 银行帐号类型
	BankAccType tTThostFtdcBankAccTypeType
	// 长客户姓名
	LongCustomerName tTThostFtdcLongIndividualNameType
}

// 银期开户信息
type tCThostFtdcOpenAccountField struct {
	// 业务功能码
	TradeCode tTThostFtdcTradeCodeType
	// 银行代码
	BankID tTThostFtdcBankIDType
	// 银行分支机构代码
	BankBranchID tTThostFtdcBankBrchIDType
	// 期商代码
	BrokerID tTThostFtdcBrokerIDType
	// 期商分支机构代码
	BrokerBranchID tTThostFtdcFutureBranchIDType
	// 交易日期
	TradeDate tTThostFtdcTradeDateType
	// 交易时间
	TradeTime tTThostFtdcTradeTimeType
	// 银行流水号
	BankSerial tTThostFtdcBankSerialType
	// 交易系统日期
	TradingDay tTThostFtdcTradeDateType
	// 银期平台消息流水号
	PlateSerial tTThostFtdcSerialType
	// 最后分片标志
	LastFragment tTThostFtdcLastFragmentType
	// 会话号
	SessionID tTThostFtdcSessionIDType
	// 客户姓名
	CustomerName tTThostFtdcIndividualNameType
	// 证件类型
	IdCardType tTThostFtdcIdCardTypeType
	// 证件号码
	IdentifiedCardNo tTThostFtdcIdentifiedCardNoType
	// 性别
	Gender tTThostFtdcGenderType
	// 国家代码
	CountryCode tTThostFtdcCountryCodeType
	// 客户类型
	CustType tTThostFtdcCustTypeType
	// 地址
	Address tTThostFtdcAddressType
	// 邮编
	ZipCode tTThostFtdcZipCodeType
	// 电话号码
	Telephone tTThostFtdcTelephoneType
	// 手机
	MobilePhone tTThostFtdcMobilePhoneType
	// 传真
	Fax tTThostFtdcFaxType
	// 电子邮件
	EMail tTThostFtdcEMailType
	// 资金账户状态
	MoneyAccountStatus tTThostFtdcMoneyAccountStatusType
	// 银行帐号
	BankAccount tTThostFtdcBankAccountType
	// 银行密码
	BankPassWord tTThostFtdcPasswordType
	// 投资者帐号
	AccountID tTThostFtdcAccountIDType
	// 期货密码
	Password tTThostFtdcPasswordType
	// 安装编号
	InstallID tTThostFtdcInstallIDType
	// 验证客户证件号码标志
	VerifyCertNoFlag tTThostFtdcYesNoIndicatorType
	// 币种代码
	CurrencyID tTThostFtdcCurrencyIDType
	// 汇钞标志
	CashExchangeCode tTThostFtdcCashExchangeCodeType
	// 摘要
	Digest tTThostFtdcDigestType
	// 银行帐号类型
	BankAccType tTThostFtdcBankAccTypeType
	// 渠道标志
	DeviceID tTThostFtdcDeviceIDType
	// 期货单位帐号类型
	BankSecuAccType tTThostFtdcBankAccTypeType
	// 期货公司银行编码
	BrokerIDByBank tTThostFtdcBankCodingForFutureType
	// 期货单位帐号
	BankSecuAcc tTThostFtdcBankAccountType
	// 银行密码标志
	BankPwdFlag tTThostFtdcPwdFlagType
	// 期货资金密码核对标志
	SecuPwdFlag tTThostFtdcPwdFlagType
	// 交易柜员
	OperNo tTThostFtdcOperNoType
	// 交易ID
	TID tTThostFtdcTIDType
	// 用户标识
	UserID tTThostFtdcUserIDType
	// 错误代码
	ErrorID tTThostFtdcErrorIDType
	// 错误信息
	ErrorMsg tTThostFtdcErrorMsgType
	// 长客户姓名
	LongCustomerName tTThostFtdcLongIndividualNameType
}

// 银期销户信息
type tCThostFtdcCancelAccountField struct {
	// 业务功能码
	TradeCode tTThostFtdcTradeCodeType
	// 银行代码
	BankID tTThostFtdcBankIDType
	// 银行分支机构代码
	BankBranchID tTThostFtdcBankBrchIDType
	// 期商代码
	BrokerID tTThostFtdcBrokerIDType
	// 期商分支机构代码
	BrokerBranchID tTThostFtdcFutureBranchIDType
	// 交易日期
	TradeDate tTThostFtdcTradeDateType
	// 交易时间
	TradeTime tTThostFtdcTradeTimeType
	// 银行流水号
	BankSerial tTThostFtdcBankSerialType
	// 交易系统日期
	TradingDay tTThostFtdcTradeDateType
	// 银期平台消息流水号
	PlateSerial tTThostFtdcSerialType
	// 最后分片标志
	LastFragment tTThostFtdcLastFragmentType
	// 会话号
	SessionID tTThostFtdcSessionIDType
	// 客户姓名
	CustomerName tTThostFtdcIndividualNameType
	// 证件类型
	IdCardType tTThostFtdcIdCardTypeType
	// 证件号码
	IdentifiedCardNo tTThostFtdcIdentifiedCardNoType
	// 性别
	Gender tTThostFtdcGenderType
	// 国家代码
	CountryCode tTThostFtdcCountryCodeType
	// 客户类型
	CustType tTThostFtdcCustTypeType
	// 地址
	Address tTThostFtdcAddressType
	// 邮编
	ZipCode tTThostFtdcZipCodeType
	// 电话号码
	Telephone tTThostFtdcTelephoneType
	// 手机
	MobilePhone tTThostFtdcMobilePhoneType
	// 传真
	Fax tTThostFtdcFaxType
	// 电子邮件
	EMail tTThostFtdcEMailType
	// 资金账户状态
	MoneyAccountStatus tTThostFtdcMoneyAccountStatusType
	// 银行帐号
	BankAccount tTThostFtdcBankAccountType
	// 银行密码
	BankPassWord tTThostFtdcPasswordType
	// 投资者帐号
	AccountID tTThostFtdcAccountIDType
	// 期货密码
	Password tTThostFtdcPasswordType
	// 安装编号
	InstallID tTThostFtdcInstallIDType
	// 验证客户证件号码标志
	VerifyCertNoFlag tTThostFtdcYesNoIndicatorType
	// 币种代码
	CurrencyID tTThostFtdcCurrencyIDType
	// 汇钞标志
	CashExchangeCode tTThostFtdcCashExchangeCodeType
	// 摘要
	Digest tTThostFtdcDigestType
	// 银行帐号类型
	BankAccType tTThostFtdcBankAccTypeType
	// 渠道标志
	DeviceID tTThostFtdcDeviceIDType
	// 期货单位帐号类型
	BankSecuAccType tTThostFtdcBankAccTypeType
	// 期货公司银行编码
	BrokerIDByBank tTThostFtdcBankCodingForFutureType
	// 期货单位帐号
	BankSecuAcc tTThostFtdcBankAccountType
	// 银行密码标志
	BankPwdFlag tTThostFtdcPwdFlagType
	// 期货资金密码核对标志
	SecuPwdFlag tTThostFtdcPwdFlagType
	// 交易柜员
	OperNo tTThostFtdcOperNoType
	// 交易ID
	TID tTThostFtdcTIDType
	// 用户标识
	UserID tTThostFtdcUserIDType
	// 错误代码
	ErrorID tTThostFtdcErrorIDType
	// 错误信息
	ErrorMsg tTThostFtdcErrorMsgType
	// 长客户姓名
	LongCustomerName tTThostFtdcLongIndividualNameType
}

// 银期变更银行账号信息
type tCThostFtdcChangeAccountField struct {
	// 业务功能码
	TradeCode tTThostFtdcTradeCodeType
	// 银行代码
	BankID tTThostFtdcBankIDType
	// 银行分支机构代码
	BankBranchID tTThostFtdcBankBrchIDType
	// 期商代码
	BrokerID tTThostFtdcBrokerIDType
	// 期商分支机构代码
	BrokerBranchID tTThostFtdcFutureBranchIDType
	// 交易日期
	TradeDate tTThostFtdcTradeDateType
	// 交易时间
	TradeTime tTThostFtdcTradeTimeType
	// 银行流水号
	BankSerial tTThostFtdcBankSerialType
	// 交易系统日期
	TradingDay tTThostFtdcTradeDateType
	// 银期平台消息流水号
	PlateSerial tTThostFtdcSerialType
	// 最后分片标志
	LastFragment tTThostFtdcLastFragmentType
	// 会话号
	SessionID tTThostFtdcSessionIDType
	// 客户姓名
	CustomerName tTThostFtdcIndividualNameType
	// 证件类型
	IdCardType tTThostFtdcIdCardTypeType
	// 证件号码
	IdentifiedCardNo tTThostFtdcIdentifiedCardNoType
	// 性别
	Gender tTThostFtdcGenderType
	// 国家代码
	CountryCode tTThostFtdcCountryCodeType
	// 客户类型
	CustType tTThostFtdcCustTypeType
	// 地址
	Address tTThostFtdcAddressType
	// 邮编
	ZipCode tTThostFtdcZipCodeType
	// 电话号码
	Telephone tTThostFtdcTelephoneType
	// 手机
	MobilePhone tTThostFtdcMobilePhoneType
	// 传真
	Fax tTThostFtdcFaxType
	// 电子邮件
	EMail tTThostFtdcEMailType
	// 资金账户状态
	MoneyAccountStatus tTThostFtdcMoneyAccountStatusType
	// 银行帐号
	BankAccount tTThostFtdcBankAccountType
	// 银行密码
	BankPassWord tTThostFtdcPasswordType
	// 新银行帐号
	NewBankAccount tTThostFtdcBankAccountType
	// 新银行密码
	NewBankPassWord tTThostFtdcPasswordType
	// 投资者帐号
	AccountID tTThostFtdcAccountIDType
	// 期货密码
	Password tTThostFtdcPasswordType
	// 银行帐号类型
	BankAccType tTThostFtdcBankAccTypeType
	// 安装编号
	InstallID tTThostFtdcInstallIDType
	// 验证客户证件号码标志
	VerifyCertNoFlag tTThostFtdcYesNoIndicatorType
	// 币种代码
	CurrencyID tTThostFtdcCurrencyIDType
	// 期货公司银行编码
	BrokerIDByBank tTThostFtdcBankCodingForFutureType
	// 银行密码标志
	BankPwdFlag tTThostFtdcPwdFlagType
	// 期货资金密码核对标志
	SecuPwdFlag tTThostFtdcPwdFlagType
	// 交易ID
	TID tTThostFtdcTIDType
	// 摘要
	Digest tTThostFtdcDigestType
	// 错误代码
	ErrorID tTThostFtdcErrorIDType
	// 错误信息
	ErrorMsg tTThostFtdcErrorMsgType
	// 长客户姓名
	LongCustomerName tTThostFtdcLongIndividualNameType
}

// 二级代理操作员银期权限
type tCThostFtdcSecAgentACIDMapField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 用户代码
	UserID tTThostFtdcUserIDType
	// 资金账户
	AccountID tTThostFtdcAccountIDType
	// 币种
	CurrencyID tTThostFtdcCurrencyIDType
	// 境外中介机构资金帐号
	BrokerSecAgentID tTThostFtdcAccountIDType
}

// 二级代理操作员银期权限查询
type tCThostFtdcQrySecAgentACIDMapField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 用户代码
	UserID tTThostFtdcUserIDType
	// 资金账户
	AccountID tTThostFtdcAccountIDType
	// 币种
	CurrencyID tTThostFtdcCurrencyIDType
}

// 灾备中心交易权限
type tCThostFtdcUserRightsAssignField struct {
	// 应用单元代码
	BrokerID tTThostFtdcBrokerIDType
	// 用户代码
	UserID tTThostFtdcUserIDType
	// 交易中心代码
	DRIdentityID tTThostFtdcDRIdentityIDType
}

// 经济公司是否有在本标示的交易权限
type tCThostFtdcBrokerUserRightAssignField struct {
	// 应用单元代码
	BrokerID tTThostFtdcBrokerIDType
	// 交易中心代码
	DRIdentityID tTThostFtdcDRIdentityIDType
	// 能否交易
	Tradeable tTThostFtdcBoolType
}

// 灾备交易转换报文
type tCThostFtdcDRTransferField struct {
	// 原交易中心代码
	OrigDRIdentityID tTThostFtdcDRIdentityIDType
	// 目标交易中心代码
	DestDRIdentityID tTThostFtdcDRIdentityIDType
	// 原应用单元代码
	OrigBrokerID tTThostFtdcBrokerIDType
	// 目标易用单元代码
	DestBrokerID tTThostFtdcBrokerIDType
}

// Fens用户信息
type tCThostFtdcFensUserInfoField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 用户代码
	UserID tTThostFtdcUserIDType
	// 登录模式
	LoginMode tTThostFtdcLoginModeType
}

// 当前银期所属交易中心
type tCThostFtdcCurrTransferIdentityField struct {
	// 交易中心代码
	IdentityID tTThostFtdcDRIdentityIDType
}

// 禁止登录用户
type tCThostFtdcLoginForbiddenUserField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 用户代码
	UserID tTThostFtdcUserIDType
	// IP地址
	IPAddress tTThostFtdcIPAddressType
}

// 查询禁止登录用户
type tCThostFtdcQryLoginForbiddenUserField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 用户代码
	UserID tTThostFtdcUserIDType
}

// UDP组播组信息
type tCThostFtdcMulticastGroupInfoField struct {
	// 组播组IP地址
	GroupIP tTThostFtdcIPAddressType
	// 组播组IP端口
	GroupPort tTThostFtdcIPPortType
	// 源地址
	SourceIP tTThostFtdcIPAddressType
}

// 资金账户基本准备金
type tCThostFtdcTradingAccountReserveField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者帐号
	AccountID tTThostFtdcAccountIDType
	// 基本准备金
	Reserve tTThostFtdcMoneyType
	// 币种代码
	CurrencyID tTThostFtdcCurrencyIDType
}

// 查询禁止登录IP
type tCThostFtdcQryLoginForbiddenIPField struct {
	// IP地址
	IPAddress tTThostFtdcIPAddressType
}

// 查询IP列表
type tCThostFtdcQryIPListField struct {
	// IP地址
	IPAddress tTThostFtdcIPAddressType
}

// 查询用户下单权限分配表
type tCThostFtdcQryUserRightsAssignField struct {
	// 应用单元代码
	BrokerID tTThostFtdcBrokerIDType
	// 用户代码
	UserID tTThostFtdcUserIDType
}

// 银期预约开户确认请求
type tCThostFtdcReserveOpenAccountConfirmField struct {
	// 业务功能码
	TradeCode tTThostFtdcTradeCodeType
	// 银行代码
	BankID tTThostFtdcBankIDType
	// 银行分支机构代码
	BankBranchID tTThostFtdcBankBrchIDType
	// 期商代码
	BrokerID tTThostFtdcBrokerIDType
	// 期商分支机构代码
	BrokerBranchID tTThostFtdcFutureBranchIDType
	// 交易日期
	TradeDate tTThostFtdcTradeDateType
	// 交易时间
	TradeTime tTThostFtdcTradeTimeType
	// 银行流水号
	BankSerial tTThostFtdcBankSerialType
	// 交易系统日期
	TradingDay tTThostFtdcTradeDateType
	// 银期平台消息流水号
	PlateSerial tTThostFtdcSerialType
	// 最后分片标志
	LastFragment tTThostFtdcLastFragmentType
	// 会话号
	SessionID tTThostFtdcSessionIDType
	// 客户姓名
	CustomerName tTThostFtdcLongIndividualNameType
	// 证件类型
	IdCardType tTThostFtdcIdCardTypeType
	// 证件号码
	IdentifiedCardNo tTThostFtdcIdentifiedCardNoType
	// 性别
	Gender tTThostFtdcGenderType
	// 国家代码
	CountryCode tTThostFtdcCountryCodeType
	// 客户类型
	CustType tTThostFtdcCustTypeType
	// 地址
	Address tTThostFtdcAddressType
	// 邮编
	ZipCode tTThostFtdcZipCodeType
	// 电话号码
	Telephone tTThostFtdcTelephoneType
	// 手机
	MobilePhone tTThostFtdcMobilePhoneType
	// 传真
	Fax tTThostFtdcFaxType
	// 电子邮件
	EMail tTThostFtdcEMailType
	// 资金账户状态
	MoneyAccountStatus tTThostFtdcMoneyAccountStatusType
	// 银行帐号
	BankAccount tTThostFtdcBankAccountType
	// 银行密码
	BankPassWord tTThostFtdcPasswordType
	// 安装编号
	InstallID tTThostFtdcInstallIDType
	// 验证客户证件号码标志
	VerifyCertNoFlag tTThostFtdcYesNoIndicatorType
	// 币种代码
	CurrencyID tTThostFtdcCurrencyIDType
	// 摘要
	Digest tTThostFtdcDigestType
	// 银行帐号类型
	BankAccType tTThostFtdcBankAccTypeType
	// 期货公司银行编码
	BrokerIDByBank tTThostFtdcBankCodingForFutureType
	// 交易ID
	TID tTThostFtdcTIDType
	// 投资者帐号
	AccountID tTThostFtdcAccountIDType
	// 期货密码
	Password tTThostFtdcPasswordType
	// 预约开户银行流水号
	BankReserveOpenSeq tTThostFtdcBankSerialType
	// 预约开户日期
	BookDate tTThostFtdcTradeDateType
	// 预约开户验证密码
	BookPsw tTThostFtdcPasswordType
	// 错误代码
	ErrorID tTThostFtdcErrorIDType
	// 错误信息
	ErrorMsg tTThostFtdcErrorMsgType
}

// 银期预约开户
type tCThostFtdcReserveOpenAccountField struct {
	// 业务功能码
	TradeCode tTThostFtdcTradeCodeType
	// 银行代码
	BankID tTThostFtdcBankIDType
	// 银行分支机构代码
	BankBranchID tTThostFtdcBankBrchIDType
	// 期商代码
	BrokerID tTThostFtdcBrokerIDType
	// 期商分支机构代码
	BrokerBranchID tTThostFtdcFutureBranchIDType
	// 交易日期
	TradeDate tTThostFtdcTradeDateType
	// 交易时间
	TradeTime tTThostFtdcTradeTimeType
	// 银行流水号
	BankSerial tTThostFtdcBankSerialType
	// 交易系统日期
	TradingDay tTThostFtdcTradeDateType
	// 银期平台消息流水号
	PlateSerial tTThostFtdcSerialType
	// 最后分片标志
	LastFragment tTThostFtdcLastFragmentType
	// 会话号
	SessionID tTThostFtdcSessionIDType
	// 客户姓名
	CustomerName tTThostFtdcLongIndividualNameType
	// 证件类型
	IdCardType tTThostFtdcIdCardTypeType
	// 证件号码
	IdentifiedCardNo tTThostFtdcIdentifiedCardNoType
	// 性别
	Gender tTThostFtdcGenderType
	// 国家代码
	CountryCode tTThostFtdcCountryCodeType
	// 客户类型
	CustType tTThostFtdcCustTypeType
	// 地址
	Address tTThostFtdcAddressType
	// 邮编
	ZipCode tTThostFtdcZipCodeType
	// 电话号码
	Telephone tTThostFtdcTelephoneType
	// 手机
	MobilePhone tTThostFtdcMobilePhoneType
	// 传真
	Fax tTThostFtdcFaxType
	// 电子邮件
	EMail tTThostFtdcEMailType
	// 资金账户状态
	MoneyAccountStatus tTThostFtdcMoneyAccountStatusType
	// 银行帐号
	BankAccount tTThostFtdcBankAccountType
	// 银行密码
	BankPassWord tTThostFtdcPasswordType
	// 安装编号
	InstallID tTThostFtdcInstallIDType
	// 验证客户证件号码标志
	VerifyCertNoFlag tTThostFtdcYesNoIndicatorType
	// 币种代码
	CurrencyID tTThostFtdcCurrencyIDType
	// 摘要
	Digest tTThostFtdcDigestType
	// 银行帐号类型
	BankAccType tTThostFtdcBankAccTypeType
	// 期货公司银行编码
	BrokerIDByBank tTThostFtdcBankCodingForFutureType
	// 交易ID
	TID tTThostFtdcTIDType
	// 预约开户状态
	ReserveOpenAccStas tTThostFtdcReserveOpenAccStasType
	// 错误代码
	ErrorID tTThostFtdcErrorIDType
	// 错误信息
	ErrorMsg tTThostFtdcErrorMsgType
}

// 银行账户属性
type tCThostFtdcAccountPropertyField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者帐号
	AccountID tTThostFtdcAccountIDType
	// 银行统一标识类型
	BankID tTThostFtdcBankIDType
	// 银行账户
	BankAccount tTThostFtdcBankAccountType
	// 银行账户的开户人名称
	OpenName tTThostFtdcInvestorFullNameType
	// 银行账户的开户行
	OpenBank tTThostFtdcOpenBankType
	// 是否活跃
	IsActive tTThostFtdcBoolType
	// 账户来源
	AccountSourceType tTThostFtdcAccountSourceTypeType
	// 开户日期
	OpenDate tTThostFtdcDateType
	// 注销日期
	CancelDate tTThostFtdcDateType
	// 录入员代码
	OperatorID tTThostFtdcOperatorIDType
	// 录入日期
	OperateDate tTThostFtdcDateType
	// 录入时间
	OperateTime tTThostFtdcTimeType
	// 币种代码
	CurrencyID tTThostFtdcCurrencyIDType
}

// 查询当前交易中心
type tCThostFtdcQryCurrDRIdentityField struct {
	// 交易中心代码
	DRIdentityID tTThostFtdcDRIdentityIDType
}

// 当前交易中心
type tCThostFtdcCurrDRIdentityField struct {
	// 交易中心代码
	DRIdentityID tTThostFtdcDRIdentityIDType
}

// 查询二级代理商资金校验模式
type tCThostFtdcQrySecAgentCheckModeField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
}

// 查询二级代理商信息
type tCThostFtdcQrySecAgentTradeInfoField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 境外中介机构资金帐号
	BrokerSecAgentID tTThostFtdcAccountIDType
}

// 用户系统信息
type tCThostFtdcUserSystemInfoField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 用户代码
	UserID tTThostFtdcUserIDType
	// 用户端系统内部信息长度
	ClientSystemInfoLen tTThostFtdcSystemInfoLenType
	// 用户端系统内部信息
	ClientSystemInfo tTThostFtdcClientSystemInfoType
	// 用户公网IP
	ClientPublicIP tTThostFtdcIPAddressType
	// 终端IP端口
	ClientIPPort tTThostFtdcIPPortType
	// 登录成功时间
	ClientLoginTime tTThostFtdcTimeType
	// App代码
	ClientAppID tTThostFtdcAppIDType
}

// 用户发出获取安全安全登陆方法请求
type tCThostFtdcReqUserAuthMethodField struct {
	// 交易日
	TradingDay tTThostFtdcDateType
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 用户代码
	UserID tTThostFtdcUserIDType
}

// 用户发出获取安全安全登陆方法回复
type tCThostFtdcRspUserAuthMethodField struct {
	// 当前可以用的认证模式
	UsableAuthMethod tTThostFtdcCurrentAuthMethodType
}

// 用户发出获取安全安全登陆方法请求
type tCThostFtdcReqGenUserCaptchaField struct {
	// 交易日
	TradingDay tTThostFtdcDateType
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 用户代码
	UserID tTThostFtdcUserIDType
}

// 生成的图片验证码信息
type tCThostFtdcRspGenUserCaptchaField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 用户代码
	UserID tTThostFtdcUserIDType
	// 图片信息长度
	CaptchaInfoLen tTThostFtdcCaptchaInfoLenType
	// 图片信息
	CaptchaInfo tTThostFtdcCaptchaInfoType
}

// 用户发出获取安全安全登陆方法请求
type tCThostFtdcReqGenUserTextField struct {
	// 交易日
	TradingDay tTThostFtdcDateType
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 用户代码
	UserID tTThostFtdcUserIDType
}

// 短信验证码生成的回复
type tCThostFtdcRspGenUserTextField struct {
	// 短信验证码序号
	UserTextSeq tTThostFtdcUserTextSeqType
}

// 用户发出带图形验证码的登录请求请求
type tCThostFtdcReqUserLoginWithCaptchaField struct {
	// 交易日
	TradingDay tTThostFtdcDateType
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 用户代码
	UserID tTThostFtdcUserIDType
	// 密码
	Password tTThostFtdcPasswordType
	// 用户端产品信息
	UserProductInfo tTThostFtdcProductInfoType
	// 接口端产品信息
	InterfaceProductInfo tTThostFtdcProductInfoType
	// 协议信息
	ProtocolInfo tTThostFtdcProtocolInfoType
	// Mac地址
	MacAddress tTThostFtdcMacAddressType
	// 终端IP地址
	ClientIPAddress tTThostFtdcIPAddressType
	// 登录备注
	LoginRemark tTThostFtdcLoginRemarkType
	// 图形验证码的文字内容
	Captcha tTThostFtdcPasswordType
	// 终端IP端口
	ClientIPPort tTThostFtdcIPPortType
}

// 用户发出带短信验证码的登录请求请求
type tCThostFtdcReqUserLoginWithTextField struct {
	// 交易日
	TradingDay tTThostFtdcDateType
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 用户代码
	UserID tTThostFtdcUserIDType
	// 密码
	Password tTThostFtdcPasswordType
	// 用户端产品信息
	UserProductInfo tTThostFtdcProductInfoType
	// 接口端产品信息
	InterfaceProductInfo tTThostFtdcProductInfoType
	// 协议信息
	ProtocolInfo tTThostFtdcProtocolInfoType
	// Mac地址
	MacAddress tTThostFtdcMacAddressType
	// 终端IP地址
	ClientIPAddress tTThostFtdcIPAddressType
	// 登录备注
	LoginRemark tTThostFtdcLoginRemarkType
	// 短信验证码文字内容
	Text tTThostFtdcPasswordType
	// 终端IP端口
	ClientIPPort tTThostFtdcIPPortType
}

// 用户发出带动态验证码的登录请求请求
type tCThostFtdcReqUserLoginWithOTPField struct {
	// 交易日
	TradingDay tTThostFtdcDateType
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 用户代码
	UserID tTThostFtdcUserIDType
	// 密码
	Password tTThostFtdcPasswordType
	// 用户端产品信息
	UserProductInfo tTThostFtdcProductInfoType
	// 接口端产品信息
	InterfaceProductInfo tTThostFtdcProductInfoType
	// 协议信息
	ProtocolInfo tTThostFtdcProtocolInfoType
	// Mac地址
	MacAddress tTThostFtdcMacAddressType
	// 终端IP地址
	ClientIPAddress tTThostFtdcIPAddressType
	// 登录备注
	LoginRemark tTThostFtdcLoginRemarkType
	// OTP密码
	OTPPassword tTThostFtdcPasswordType
	// 终端IP端口
	ClientIPPort tTThostFtdcIPPortType
}

// api握手请求
type tCThostFtdcReqApiHandshakeField struct {
	// api与front通信密钥版本号
	CryptoKeyVersion tTThostFtdcCryptoKeyVersionType
}

// front发给api的握手回复
type tCThostFtdcRspApiHandshakeField struct {
	// 握手回复数据长度
	FrontHandshakeDataLen tTThostFtdcHandshakeDataLenType
	// 握手回复数据
	FrontHandshakeData tTThostFtdcHandshakeDataType
	// API认证是否开启
	IsApiAuthEnabled tTThostFtdcBoolType
}

// api给front的验证key的请求
type tCThostFtdcReqVerifyApiKeyField struct {
	// 握手回复数据长度
	ApiHandshakeDataLen tTThostFtdcHandshakeDataLenType
	// 握手回复数据
	ApiHandshakeData tTThostFtdcHandshakeDataType
}

// 操作员组织架构关系
type tCThostFtdcDepartmentUserField struct {
	// 经纪公司代码
	BrokerID tTThostFtdcBrokerIDType
	// 用户代码
	UserID tTThostFtdcUserIDType
	// 投资者范围
	InvestorRange tTThostFtdcDepartmentRangeType
	// 投资者代码
	InvestorID tTThostFtdcInvestorIDType
}

// 查询频率，每秒查询比数
type tCThostFtdcQueryFreqField struct {
	// 查询频率
	QueryFreq tTThostFtdcQueryFreqType
}
