package goctp

type BaseQuote interface {
	ReqConnect(addr string)
	ReqLogin(investor, pwd, broker string)
	Release()
	ReqSubscript(instrument string)

	RegOnFrontConnected(on OnFrontConnectedType)
	RegOnRspUserLogin(on OnRspUserLoginType)
	RegOnTick(on OnTickType)
}
