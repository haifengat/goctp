package goctp

import (
	"fmt"
	"testing"
)

func TestQuotePro(t *testing.T) {
	q := NewQuotePro()
	info, rsp := q.Start(LoginConfig{
		// Front: "tcp://180.168.146.187:10212",
		Front: "tcp://180.168.146.187:10131",
	})
	if rsp.ErrorID != 0 {
		fmt.Println(rsp.ErrorMsg.String())
	} else {
		fmt.Printf("%+v\n", info)
	}
}
