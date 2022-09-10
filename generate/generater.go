package main

import (
	"bytes"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
	"text/template"

	"golang.org/x/text/encoding/simplifiedchinese"
)

// 接口源目录
var (
	srcPath     = "../v6.6.8_20220712/"
	outPath     = "ctpdefine"
	packageName = "gitee.com/haifengat/goctp"
)

func main() {
	// genC()
	// genGo()
	// genGoStruct()
	genGoDataType()
	fmt.Println("finish.")
}

type fieldStruct struct {
	FieldType, FieldName, Comment string
}
type tplStruct struct {
	Comment, FuncTypeName, FuncName string
	FuncFields                      []fieldStruct
}

func tmpl(tplFileName string, content interface{}, funcMap template.FuncMap) {
	_, curFile, _, _ := runtime.Caller(1)
	tplPath := filepath.Dir(curFile) // 模板文件与执行文件在同一目录

	fm := make(template.FuncMap, 0)
	fm["trimStar"] = func(str string) string {
		return strings.TrimPrefix(str, "*")
	}
	if funcMap != nil {
		for k, v := range funcMap {
			fm[k] = v
		}
	}

	t := template.New(path.Base(tplFileName)).Delims("[[", "]]").Funcs(fm)
	t, err := t.ParseFiles(path.Join(tplPath, tplFileName))
	if err != nil {
		panic(err)
	}
	var buf = bytes.Buffer{}
	err = t.Execute(&buf, content) // ***
	if err != nil {
		panic(err)
	}
	// 写入 .h
	err = os.WriteFile(strings.TrimSuffix(tplFileName, filepath.Ext(tplFileName)), buf.Bytes(), os.ModePerm)
	if err != nil {
		panic(err)
	}
}

// generate 取接口中的函数,转换为对应的接口
func generate(tplExeFunc func(title string, on []*tplStruct, fn []*tplStruct)) {
	for _, hFileName := range []string{srcPath + "ThostFtdcMdApi.h", srcPath + "/ThostFtdcTraderApi.h"} {
		var title string
		if strings.Contains(hFileName, "ThostFtdcMdApi") {
			title = "quote"
		} else {
			title = "trade"
		}
		bsFile, err := os.ReadFile(hFileName)
		if err != nil {
			panic(err)
		}
		// 汉字处理
		bsFile, _ = simplifiedchinese.GB18030.NewDecoder().Bytes(bsFile)
		/*
			///登录请求响应
			virtual void OnRspUserLogin(CThostFtdcRspUserLoginField *pRspUserLogin, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {};
		*/
		strFile := string(bsFile)
		strFile = strings.ReplaceAll(strFile, "\r\n", "\n")                            // 换行符用 \n 避免 win和 lnx执行时不一致
		re := regexp.MustCompile(`\t///(.*)\n[^v]*virtual\s+(\w+)\s+(\w+)\(([^)]*)\)`) // 分解函数定义:注释,返回类型,函数名,参数字段四部分
		funs := re.FindAllStringSubmatch(strFile, -1)
		tplsOn := make([]*tplStruct, 0)
		tplsFn := make([]*tplStruct, 0)
		for _, fun := range funs {
			funComment, _, funName, funParams := fun[1], fun[2], fun[3], fun[4]
			re := regexp.MustCompile(`(\w+)\s+([*]?[ ]?\w+)[,]?\s*`) //参数分解:类型,名称
			fields := re.FindAllStringSubmatch(funParams, -1)
			funFields := make([]fieldStruct, 0)
			for _, field := range fields {
				funFields = append(funFields, fieldStruct{FieldType: field[1], FieldName: field[2]})
			}
			if strings.HasPrefix(funName, "On") { // On 响应函数
				tplsOn = append(tplsOn, &tplStruct{
					Comment:      funComment,
					FuncTypeName: funName[2:] + "Type",
					FuncName:     funName[2:],
					FuncFields:   funFields,
				})
			} else {
				tplsFn = append(tplsFn, &tplStruct{
					Comment:      funComment,
					FuncTypeName: funName + "Type",
					FuncName:     funName,
					FuncFields:   funFields,
				})
			}
		}
		tplExeFunc(title, tplsOn, tplsFn)
	}
}

// genGo quote.go trade.go
func genGo() {
	generate(func(title string, tplsOn, tplsFn []*tplStruct) {
		// for _, v := range [][]*tplStruct{tplsOn, tplsFn} {
		for _, fn := range tplsFn { // 主调函数
			// 增加 void* api 首个参数
			tmp := []fieldStruct{fieldStruct{FieldType: "void*", FieldName: "api"}}
			tmp = append(tmp, fn.FuncFields...)
			fn.FuncFields = tmp
		}
		// }
		mpCpp := make(map[string]interface{})
		mpCpp["On"] = tplsOn
		mpCpp["Fn"] = tplsFn
		fm := make(template.FuncMap)
		fm["struct_Type"] = func(structType string) string {
			if structType == "CThostFtdcMdSpi" {
				return "void"
			}
			if structType == "CThostFtdcTraderSpi" {
				return "void"
			}
			if strings.HasPrefix(structType, "CThostFtdc") { // struct
				return "struct " + structType // struct CThostFtdcRspUserLoginField *pRspUserLogin
			}
			if structType == "bool" {
				return "_Bool"
			}
			if structType == "THOST_TE_RESUME_TYPE" {
				return "int"
			}
			return structType
		}
		fm["C_struct"] = func(structType string) string {
			if strings.HasPrefix(structType, "CThostFtdc") { // struct
				return "*C.struct_" + structType // field *C.struct_CThostFtdcRspUserLoginField
			}
			if structType == "int" {
				return "C.int"
			}
			if structType == "bool" {
				return "C._Bool"
			}
			return structType
		}
		fm["ctp_type"] = func(structType string) string {
			if strings.HasPrefix(structType, "CThostFtdc") { // struct
				return fmt.Sprintf("*ctp.%s", structType) // *ctp.CThostFtdcUserLogoutField
			}
			return structType
		}
		fm["ctp_param"] = func(structType, field string) string {
			if strings.HasPrefix(structType, "CThostFtdc") { // struct
				return fmt.Sprintf("(*ctp.%s)(unsafe.Pointer(%s))", structType, strings.TrimPrefix(field, "*")) // (*ctp.CThostFtdcRspUserLoginField)(unsafe.Pointer(field))
			}
			if structType == "int" {
				return "int(" + field + ")"
			}
			if structType == "bool" {
				return "bool(" + field + ")"
			}
			return field
		}
		tmpl(title+"_lnx.go.tpl", mpCpp, fm)
	})
}

// genC 生成 quote.h quote.cpp trade.h trade.cpp
func genC() {
	generate(func(title string, tplsOn, tplsFn []*tplStruct) {
		tmpl(title+".h.tpl", tplsOn, nil)
		mpCpp := make(map[string]interface{})
		mpCpp["On"] = tplsOn
		mpCpp["Fn"] = tplsFn
		tmpl(title+".cpp.tpl", mpCpp, nil)
	})
}

// 生成数据类型与 struct
func genGoStruct() {
	bsFile, err := os.ReadFile(srcPath + "ThostFtdcUserApiStruct.h")
	if err != nil {
		panic(err)
	}
	// 汉字处理
	bsFile, _ = simplifiedchinese.GB18030.NewDecoder().Bytes(bsFile)

	re := regexp.MustCompile(`///(\S*)\s*struct\s*(\w*)\s*{([^}]*)}`) // 分成struct的注释,名称,字段两部分
	structs := re.FindAllStringSubmatch(string(bsFile), -1)
	strs := make([]*tplStruct, 0)
	for _, strc := range structs {
		ts := &tplStruct{
			Comment:      strc[1],
			FuncTypeName: strc[2],
		}
		re = regexp.MustCompile(`///([^\r\n]*)\s*(\w+)\s+([^;]+);`) // 所有字段再分解成各个单独字段: 注释(可能含空格),类型,名称
		fields := re.FindAllStringSubmatch(strc[3], -1)
		for _, v := range fields {
			ts.FuncFields = append(ts.FuncFields, fieldStruct{
				Comment:   v[1],
				FieldType: v[2],
				FieldName: v[3],
			})
		}
		strs = append(strs, ts)
	}
	tmpl("struct.go.tpl", strs, nil)
}

func genGoDataType() {
	bsFile, err := os.ReadFile(srcPath + "ThostFtdcUserApiDataType.h")
	if err != nil {
		panic(err)
	}
	// 汉字处理
	bsFile, _ = simplifiedchinese.GB18030.NewDecoder().Bytes(bsFile)
	strFile := string(bsFile)
	strFile = strings.ReplaceAll(strings.ReplaceAll(strFile, "\r\n", "\n"), "\n\t", "\n")
	// 类型声明-enum
	re := regexp.MustCompile(`/+.+是一个(.*)\n/*\n([^;]+)typedef\s+(\w+)\s+(\w+)\s*;`) // 注释,defines,类型,名称
	types := re.FindAllStringSubmatch(strFile, -1)

	tss := make([]*tplStruct, 0)
	for _, v := range types {
		ts := &tplStruct{
			Comment:      v[1],
			FuncTypeName: v[3], // 基础类型
			FuncName:     v[4],
		}
		reSub := regexp.MustCompile(`/+(.*)\n#define\s+(\w+)\s+'(.+)'`) // 注释,名称,值 \w改为.因为有'#'的情况
		definnes := reSub.FindAllStringSubmatch(v[2], -1)
		for _, v := range definnes {
			ts.FuncFields = append(ts.FuncFields, fieldStruct{
				Comment:   v[1],
				FieldType: v[2],
				FieldName: v[3],
			})
		}
		tss = append(tss, ts)
	}
	/*/////////////////////////////////////////////////////////////////////////
	///TFtdcExchangePropertyType是一个交易所属性类型
	/////////////////////////////////////////////////////////////////////////
	///正常
	#define THOST_FTDC_EXP_Normal '0'
	///根据成交生成报单
	#define THOST_FTDC_EXP_GenOrderByTrade '1'

	typedef char TThostFtdcExchangePropertyType;
	=> // 交易所属性类型
	type TThostFtdcExchangePropertyType byte
	const THOST_FTDC_EXP_Normal = '0' // 正常
	const THOST_FTDC_EXP_GenOrderByTrade = '1' // 根据成交生成报单*/
	re = regexp.MustCompile(`/+.+是一个(.*)\n/*\ntypedef\s+(\w+)\s+(.+)\s*;`) // 注释,类型,自定义类型
	types = re.FindAllStringSubmatch(strFile, -1)
	for _, v := range types {
		ts := &tplStruct{
			Comment:      v[1],
			FuncTypeName: v[2], // 基础类型
			FuncName:     v[3],
		}
		// typedef char TThostFtdcTraderIDType[21]; -> type TThostFtdcTraderIDType [21]byte
		if strings.Contains(ts.FuncName, "[") {
			ts.FuncTypeName = "[" + strings.Split(ts.FuncName, "[")[1] + "byte"
			ts.FuncName = strings.Split(ts.FuncName, "[")[0]
		}
		tss = append(tss, ts)
	}

	fm := make(template.FuncMap)
	fm["baseType"] = func(preType string) string {
		if preType == "int" { // typedef int TThostFtdcIPPortType; -> type TThostFtdcIPPortType int32
			return "int32"
		}
		if preType == "double" { // typedef double TThostFtdcPriceType; -> type TThostFtdcPriceType float64
			return "float64"
		}
		if preType == "short" { // typedef short TThostFtdcSequenceSeriesType; -> type TThostFtdcSequenceSeriesType int16
			return "int16"
		}
		if preType == "char" {
			return "byte"
		}
		return preType
	}
	tmpl("datatype.go.tpl", tss, fm)
}
