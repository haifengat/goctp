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

	"github.com/pkg/errors"
	"golang.org/x/text/encoding/simplifiedchinese"
)

// 接口源目录
var (
	srcPath     = "v6.6.8_20220712"
	outPath     = "ctpdefine"
	packageName = "gitee.com/haifengat/goctp"
)

func main() {
	// generateDataType() // 生成后需和动处理出入金的几个定义
	// generateStruct()
	// genC()
	genGo()
	fmt.Println("finish.")
}

type fieldStruct struct {
	FieldType, FieldName string
}
type tplStruct struct {
	Comment, FuncTypeName, FuncName string
	FuncFields                      []fieldStruct
}

func tmpl(tplFileName string, content interface{}) error {
	_, curFile, _, _ := runtime.Caller(1)
	tplPath := filepath.Dir(curFile) // 模板文件与执行文件在同一目录

	fm := make(template.FuncMap, 0)
	fm["trimStar"] = func(str string) string {
		return strings.TrimPrefix(str, "*")
	}
	////////// golang /////////////////
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

	t := template.New(path.Base(tplFileName)).Delims("[[", "]]").Funcs(fm)
	t, err := t.ParseFiles(path.Join(tplPath, tplFileName))
	if err != nil {
		return errors.Wrap(err, "new template")
	}
	var buf = bytes.Buffer{}
	err = t.Execute(&buf, content) // ***
	if err != nil {
		return errors.Wrap(err, "execute template")
	}
	// 写入 .h
	err = os.WriteFile(strings.TrimSuffix(tplFileName, filepath.Ext(tplFileName)), buf.Bytes(), os.ModePerm)
	if err != nil {
		return errors.Wrap(err, tplFileName)
	}
	return nil
}

// generate 取接口中的函数,转换为对应的接口
func generate(tplExeFunc func(title string, on []*tplStruct, fn []*tplStruct)) {
	apiPath := path.Join("..", srcPath)
	for _, hFileName := range []string{apiPath + "/ThostFtdcMdApi.h", apiPath + "/ThostFtdcTraderApi.h"} {
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
		err := tmpl(title+"_lnx.go.tpl", mpCpp)
		if err != nil {
			fmt.Println(err)
		}
	})
}

// genC 生成 quote.h quote.cpp trade.h trade.cpp
func genC() {
	generate(func(title string, tplsOn, tplsFn []*tplStruct) {
		err := tmpl(title+".h.tpl", tplsOn)
		if err != nil {
			fmt.Println(err)
		}
		mpCpp := make(map[string]interface{})
		mpCpp["On"] = tplsOn
		mpCpp["Fn"] = tplsFn
		err = tmpl(title+".cpp.tpl", mpCpp)
		if err != nil {
			fmt.Println(err)
		}
	})
}
