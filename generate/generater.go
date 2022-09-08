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
	srcPath     = "./v6.6.8_20220712/"
	outPath     = "./ctpdefine"
	packageName = "gitee.com/haifengat/goctp"
)

func main() {
	// generateDataType() // 生成后需和动处理出入金的几个定义
	// generateStruct()
	genC()
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
		return errors.Wrap(err, "写入 .h")
	}
	return nil
}

// genC 生成 c 的 .h .cpp
func genC() {
	for _, hFileName := range []string{"../" + srcPath + "/ThostFtdcMdApi.h", "../" + srcPath + "/ThostFtdcTraderApi.h"} {
		var title string
		if strings.Contains(hFileName, "ThostFtdcMdApi") {
			title = "quote"
		} else {
			title = "trade"
		}
		bsFile, _ := os.ReadFile(hFileName)
		// 汉字处理
		bsFile, _ = simplifiedchinese.GB18030.NewDecoder().Bytes(bsFile)
		/*
			///登录请求响应
			virtual void OnRspUserLogin(CThostFtdcRspUserLoginField *pRspUserLogin, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {};
		*/
		re := regexp.MustCompile(`\t///(.*)\r\n[^v]*virtual\s+(\w+)\s+(\w+)\(([^)]*)\)`) // 分解函数定义:注释,返回类型,函数名,参数字段四部分
		funs := re.FindAllStringSubmatch(string(bsFile), -1)
		tplsOn := make([]tplStruct, 0)
		tplsFn := make([]tplStruct, 0)
		for _, fun := range funs {
			funComment, _, funName, funParams := fun[1], fun[2], fun[3], fun[4]
			re := regexp.MustCompile(`(\w+)\s+([*]?[ ]?\w+)[,]?\s*`) //参数分解:类型,名称
			fields := re.FindAllStringSubmatch(funParams, -1)
			funFields := make([]fieldStruct, 0)
			for _, field := range fields {
				funFields = append(funFields, fieldStruct{FieldType: field[1], FieldName: field[2]})
			}
			if strings.HasPrefix(funName, "On") { // On 响应函数
				tplsOn = append(tplsOn, tplStruct{
					Comment:      funComment,
					FuncTypeName: funName[2:] + "Type",
					FuncName:     funName[2:],
					FuncFields:   funFields,
				})
			} else {
				tplsFn = append(tplsFn, tplStruct{
					Comment:      funComment,
					FuncTypeName: funName + "Type",
					FuncName:     funName,
					FuncFields:   funFields,
				})
			}
		}
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
	}
}
