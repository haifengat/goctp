package goctp

import (
	"fmt"
	"strings"

	"golang.org/x/text/encoding/simplifiedchinese"
)

func ToGBK(bs []byte) string {
	msg, _ := simplifiedchinese.GB18030.NewDecoder().Bytes(bs)
	return string(msg)
}

func FormatFloat(f float64, precision int) string {
	strFmt := fmt.Sprintf("%%.%df", precision)
	return strings.TrimRight(strings.TrimRight(fmt.Sprintf(strFmt, f), "0"), ".")
}
