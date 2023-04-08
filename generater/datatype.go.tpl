package goctp

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
	[[- range .Define ]][[/*const 值*/]]
const [[ .Var ]] [[ $typedef.Name ]]  = [[ if eq (len .Value) 1 ]]'[[ .Value ]]'[[ else ]]"[[ .Value ]]"[[ end ]] // [[ .Comment ]]
	[[ end ]]
[[- if gt (len .Define) 0 ]][[/*有 const 定义*/]]	
	[[- $idx0 := index .Define 0 ]][[/* 取值的长度, 为1:char类型, >1 string类型(银转类型)*/]]
	[[- if eq (len $idx0.Value) 1 ]]
var mp[[ .Name ]] = map[[ `[` ]][[ .Name ]][[ `]` ]]string{[[ range $i, $v := .Define ]][[ if gt $i 0 ]], [[ end ]]'[[ .Value ]]': "[[ .Var ]]"[[ end ]]}
	[[ else ]]
var mp[[ .Name ]] = map[[ `[` ]][[ .Name ]][[ `]` ]]string{[[ range $i, $v := .Define ]][[ if gt $i 0 ]], [[ end ]]"[[ .Value ]]": "[[ .Var ]]"[[ end ]]}
	[[ end ]]
func (e [[ .Name ]]) String() string {
	if s, ok := mp[[ .Name ]][e];ok{
		return s[strings.LastIndex(s, "_")+1:]
	}
	return "值错误"
}
[[ end -]]
	[[- if and (eq .Type "char") (gt .Length 1)]][[/*const xxx byte 类型的 String*/]]
func (s [[ .Name ]]) String() string {
	return toGBK(s[:])
}
	[[- else if eq .Type "double"]][[/*double 类型 String()*/]]
func (f [[ .Name ]]) String() string {
	return strings.TrimRight(strings.TrimRight(fmt.Sprintf("%.6f", f), "0"), ".")
}
	[[ end ]]

[[ end ]]