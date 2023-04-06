package goctp

import (
	"fmt"
	"testing"
)

func TestTradePro(t *testing.T) {
	trd := NewTradePro()
	logInfo, rsp := trd.Start(LoginConfig{
		Front:    "tcp://180.168.146.187:10130",
		Broker:   "9999",
		UserID:   "008107",
		Password: "1",
		AppID:    "simnow_client_test",
		AuthCode: "0000000000000000",
	})
	fmt.Printf("%+v\n", logInfo)
	fmt.Printf("%+v\n", rsp)
	select {}
}
