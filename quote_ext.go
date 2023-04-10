/*
封装成 golang 更易使用的接口
*/
package goctp

type QuoteExt struct {
	*Quote
	id int
}

func NewQuoteExt() *QuoteExt {
	qExt := &QuoteExt{}
	qExt.Quote = NewQuote()
	return qExt
}

func (q *QuoteExt) getReqID() int {
	q.id++
	return q.id
}

func (q *QuoteExt) ReqUserLogin(broker, user, pwd string) {
	f := CThostFtdcReqUserLoginField{}
	copy(f.BrokerID[:], []byte(broker))
	copy(f.UserID[:], []byte(user))
	copy(f.Password[:], []byte(pwd))
	q.Quote.ReqUserLogin(&f, q.getReqID())
}

func (q *QuoteExt) ReqSubscript(instruments ...string) {
	q.SubscribeMarketData(instruments, len(instruments))
}

func (q *QuoteExt) Release() {
	q.RegisterSpi(nil)
	q.Quote.Release()
	q.Quote = nil
}
