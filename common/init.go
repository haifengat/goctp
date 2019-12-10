package common

import (
	"bytes"
	"errors"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"reflect"
	"strings"
	"text/template"
)

// 公共-连接
type OnFrontConnectedType func()

// 公共-登录
type OnRspUserLoginType func(loginField *RspUserLoginField, info *RspInfoField)

// 行情
type OnTickType func(tick *TickField)

// 交易-委托响应
type OnRtnOrderType func(field *OrderField)

// 交易-错误委托
type OnRtnErrOrderType func(field *OrderField, info *RspInfoField)

// 交易-成交响应
type OnRtnTradeType func(field *TradeField)

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func TemplateMap(templateString string, mapContent map[string]string) string {
	t := template.Must(template.New("fun").Parse(templateString))
	buf := &bytes.Buffer{}
	checkErr(t.Execute(buf, mapContent))
	return buf.String()
}

func Bytes2String(t []byte) string {
	msg, _ := simplifiedchinese.GB18030.NewDecoder().Bytes(t)
	return strings.Trim(string(msg), "\u0000")
}

//func GB18030ToUtf8(s []byte) ([]byte, error) {
//	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GB18030.NewDecoder())
//	d, e := ioutil.ReadAll(reader)
//	if e != nil {
//		return nil, e
//	}
//	return d, nil
//}
//
//func Utf8ToGbk(s []byte) ([]byte, error) {
//	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewEncoder())
//	d, e := ioutil.ReadAll(reader)
//	if e != nil {
//		return nil, e
//	}
//	return d, nil
//}

// 判断obj是否在target中，target支持的类型 arrary, slice, map
func Contain(obj interface{}, target interface{}) (bool, error) {
	targetValue := reflect.ValueOf(target)
	switch reflect.TypeOf(target).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			if targetValue.Index(i).Interface() == obj {
				return true, nil
			}
		}
	case reflect.Map:
		if targetValue.MapIndex(reflect.ValueOf(obj)).IsValid() {
			return true, nil
		}
	}

	return false, errors.New("not in array")
}
