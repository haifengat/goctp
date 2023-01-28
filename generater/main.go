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
)

var srcPath = "../CTPv6.6.8_20220712"

func main() {
	genDataType()
	os.Exit(0)
	var input string
	for {
		fmt.Println("1. data type")
		fmt.Scanln(&input)
		switch input {
		case "1":
			dataType()
		default:
			fmt.Println("输入有误")
		}
	}
}

type DataType struct {
	TypeName string
	TypeType string
	Comment  string
	Consts   []struct {
		Name    string
		Value   string
		Comment string
	}
}

func genDataType() {
	bs, _ := os.ReadFile(path.Join(srcPath, "ThostFtdcUserApiDataType.h"))
	str := string(bs)

	datas := make([]DataType, 0)

	re := regexp.MustCompile(`/+.+是一个(.+)\n[/\n]*(///[^;]+)?typedef\s+(\w+)\s+(\w+)(?:\[(\d+)])?`)
	charMatches := re.FindAllStringSubmatch(str, -1)
	for _, m := range charMatches {
		data := DataType{
			TypeName: m[4],
			Comment:  m[1],
		}
		switch m[3] {
		case "char":
			if len(m[5]) == 0 {
				// typedef char TThostFtdcIdCardTypeType;
				data.TypeType = "byte"
			} else {
				// typedef char TThostFtdcIPAddressType[33];
				data.TypeType = fmt.Sprintf("[%s]byte", m[5])
			}
		case "short":
			// typedef short TThostFtdcSequenceSeriesType;
			data.TypeType = "int16"
		case "double":
			// typedef double TThostFtdcUnderlyingMultipleType;
			data.TypeType = "float64"
		case "int":
			// typedef int TThostFtdcPriorityType;
			data.TypeType = "int32"
		default:
			data.TypeType = m[3]
		}

		reSub := regexp.MustCompile(`///\s*(.*)\n#define\s+(\w+)\s+'(.+)'`)
		consts := reSub.FindAllStringSubmatch(m[2], -1)
		for _, c := range consts {
			data.Consts = append(data.Consts, struct {
				Name    string
				Value   string
				Comment string
			}{
				// #define THOST_FTDC_ICT_EID '0'
				Name:    fmt.Sprintf("%s %s", c[2], m[4]), // .tpl中[[$.TypeName]]报错
				Value:   c[3],
				Comment: c[1],
			})
			if len(c[3]) > 1 { // 出入金相关的值
				data.TypeType = "string"
			}
		}
		datas = append(datas, data)
	}

	tmpl("./datatype.go.tpl", datas, "../go/def", nil)
}

func tmpl(tplFileName string, content interface{}, outPath string, funcMap template.FuncMap) {
	_, curFile, _, _ := runtime.Caller(1)
	tplPath := filepath.Dir(curFile) // 模板文件与执行文件在同一目录

	t := template.New(path.Base(tplFileName)).Delims("[[", "]]")
	if funcMap != nil {
		t.Funcs(funcMap)
	}

	t, err := t.ParseFiles(path.Join(tplPath, tplFileName))
	if err != nil {
		panic(err)
	}

	var buf = bytes.Buffer{}
	err = t.Execute(&buf, content)
	if err != nil {
		panic(err)
	}

	// 写入文件
	err = os.WriteFile(path.Join(outPath, strings.TrimSuffix(tplFileName, filepath.Ext(tplFileName))), buf.Bytes(), os.ModePerm)
	if err != nil {
		panic(err)
	}
}

func dataType() {
	bs, _ := os.ReadFile(path.Join(srcPath, "ThostFtdcUserApiDataType.h"))
	str := string(bs)

	re := regexp.MustCompile(`typedef\s+char\s+(\w+)\[(\d+)]`)
	matches := re.FindAllString(str, -1) // re.FindAllStringSubmatch(str, -1)
	typeChars := re.ReplaceAllString(strings.Join(matches, "\n"), `type $1 [$2]byte`)
	fmt.Println(typeChars)

	re = regexp.MustCompile(`/+.+是一个(.+)\n[/\n]*(///[^;]+)?typedef\s+(\w+)\s+(\w+)`)
	typeChar := make([]string, 0)
	submatches := re.FindAllStringSubmatch(str, -1)
	for _, sub := range submatches {
		// sub[0] 替换为 类型定义
		typeChar = append(typeChar, re.ReplaceAllString(sub[0], "// $1\ntype $4 byte\n"))

		// 类型的值
		reSub := regexp.MustCompile(`///\s*(.*)\n#define\s+(\w+)\s+'(.+)'`)
		typeChar = append(typeChar, reSub.ReplaceAllString(sub[2], fmt.Sprintf(`const $2 %s = '$3' // $1`, sub[4])))
	}

	f, _ := os.Create("../go/def/datatype.go")
	f.WriteString(typeChars)
	f.WriteString(strings.Join(typeChar, "\n"))
}
