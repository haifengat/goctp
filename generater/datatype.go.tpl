package def

import (
	"bytes"
	"fmt"
	"strings"

	"golang.org/x/text/encoding/simplifiedchinese"
)

type THOST_TE_RESUME_TYPE int32

const (
	THOST_TERT_RESTART THOST_TE_RESUME_TYPE = 0
	THOST_TERT_RESUME  THOST_TE_RESUME_TYPE = 1
	THOST_TERT_QUICK   THOST_TE_RESUME_TYPE = 2
)

func toGBK(bs []byte) string {
	msg, _ := simplifiedchinese.GB18030.NewDecoder().Bytes(bytes.TrimRight(bs, "\x00"))
	return string(msg)
}

[[ range $index, $typedef := .]]// [[ .Comment ]]
type [[ .Name ]] [[ toGo .Type .Length ]]
	[[- range .Define ]]
const [[ .Var ]] [[ $typedef.Name ]]  = [[ if eq (len .Value) 1 ]]'[[ .Value ]]'[[ else ]]"[[ .Value ]]"[[ end ]] // [[ .Comment ]]
	[[ end ]]
	[[- if and (eq .Type "char") (gt .Length 1) ]]
func (s [[ .Name ]]) String() string {
	return toGBK(s[:])
}
	[[- else if eq .Type "double" ]]
func (f [[ .Name ]]) String() string {
	return strings.TrimRight(strings.TrimRight(fmt.Sprintf("%.6f", f), "0"), ".")
}
	[[ end ]]

[[ end ]]