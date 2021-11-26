package goctp

import (
	"strings"

	"golang.org/x/text/encoding/simplifiedchinese"
)

func Bytes2String(t []byte) string {
	msg, _ := simplifiedchinese.GB18030.NewDecoder().Bytes(t)
	return strings.Trim(string(msg), "\u0000")
}
