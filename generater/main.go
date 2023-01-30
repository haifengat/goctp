package main

import (
	"bytes"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"text/template"
)

var srcPath = "../CTPv6.6.8_20220712"

func main() {
	genQuote()
	// genStruct()
	// genDataType()
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

type Func struct {
	RtnType string
	Name    string
	Comment string
	Params  []struct {
		Type    string
		Var     string
		HasStar bool
	}
}

func genQuote() {
	fn, on := getFuncs("ThostFtdcMdApi.h")
	tmpl("quote.h.cpl", map[string]any{"Fn": fn, "On": on}, "../c/", nil)
	tmpl("quote.cpp.cpl", map[string]any{"Fn": fn, "On": on}, "../c/", nil)
}

func getFuncs(hFileName string) (fn []Func, on []Func) {
	bs, _ := os.ReadFile(path.Join(srcPath, hFileName))
	str := string(bs)
	fn = make([]Func, 0)
	on = make([]Func, 0)
	/*
		///用户登录请求
		virtual int ReqUserLogin(CThostFtdcReqUserLoginField *pReqUserLoginField, int nRequestID) = 0;
	*/
	re := regexp.MustCompile(`///(.*)\n[^v]*virtual\s+(\w+)\s+(\w+)\(([^)]*)\)`)
	for _, m := range re.FindAllStringSubmatch(str, -1) {
		f := Func{
			RtnType: m[2],
			Name:    m[3],
			Comment: m[1],
		}
		re = regexp.MustCompile(`(\w+)\s+([*])?\s?(\w+)`)
		for _, v := range re.FindAllStringSubmatch(m[4], -1) { // 函数参数
			f.Params = append(f.Params, struct {
				Type    string
				Var     string
				HasStar bool
			}{
				Type:    v[1],
				Var:     v[3],
				HasStar: len(v[2]) > 0,
			})
		}
		if strings.HasPrefix(f.Name, "On") { // 响应函数
			on = append(on, f)
		} else {
			fn = append(fn, f)
		}
	}
	fmt.Printf("%+v\n", on)
	return
}

func genStruct() {
	type Struct struct {
		Name    string
		Comment string
		Fields  []struct {
			Name    string
			Type    string
			Comment string
		}
	}

	bs, _ := os.ReadFile(path.Join(srcPath, "ThostFtdcUserApiStruct.h"))
	str := string(bs)
	datas := make([]Struct, 0)

	re := regexp.MustCompile(`///\s*(\S*)\s*struct\s+(\w+)\s*{([^}]*)}`)
	for _, m := range re.FindAllStringSubmatch(str, -1) {
		stru := Struct{
			Name:    m[2],
			Comment: m[1],
		}
		re = regexp.MustCompile(`///(\S*)\s*(\w+)\s+([^;]+)`)
		for _, v := range re.FindAllStringSubmatch(m[3], -1) {
			stru.Fields = append(stru.Fields, struct {
				Name    string
				Type    string
				Comment string
			}{
				Name:    v[3],
				Type:    v[2],
				Comment: v[1],
			})
		}
		datas = append(datas, stru)
	}

	tmpl("struct.go.tpl", datas, "../go/def", nil)
}

func genDataType() {
	type Typedef struct {
		Name    string
		Type    string
		Length  int
		Comment string
		Define  []struct {
			Var     string
			Value   string
			Comment string
		}
	}

	bs, _ := os.ReadFile(path.Join(srcPath, "ThostFtdcUserApiDataType.h"))
	str := string(bs)

	datas := make([]Typedef, 0)

	re := regexp.MustCompile(`/+.+是一个(.+)\n[/\n]*(///[^;]+)?typedef\s+(\w+)\s+(\w+)(?:\[(\d+)])?`)
	charMatches := re.FindAllStringSubmatch(str, -1)
	for _, m := range charMatches {
		data := Typedef{
			Name:    m[4],
			Type:    m[3],
			Comment: m[1],
		}
		data.Length, _ = strconv.Atoi(m[5])

		reSub := regexp.MustCompile(`///\s*(.*)\n#define\s+(\w+)\s+'(.+)'`)
		consts := reSub.FindAllStringSubmatch(m[2], -1)
		for _, c := range consts {
			data.Define = append(data.Define, struct {
				Var     string
				Value   string
				Comment string
			}{
				// #define THOST_FTDC_ICT_EID '0'
				Var:     c[2],
				Value:   c[3],
				Comment: c[1],
			})
			if len(c[3]) > 1 { // 出入金相关的值 '102001'
				data.Type = "string"
			}
		}
		datas = append(datas, data)
	}

	tmpl("./datatype.go.tpl", datas, "../go/def", template.FuncMap{
		"toGo": func(t string, l int) string {
			var goType string = t
			switch t {
			case "char":
				if l == 0 {
					// typedef char TThostFtdcIdCardTypeType;
					goType = "byte"
				} else {
					// typedef char TThostFtdcIPAddressType[33];
					goType = fmt.Sprintf("[%d]byte", l)
				}
			case "short":
				// typedef short TThostFtdcSequenceSeriesType;
				goType = "int16"
			case "double":
				// typedef double TThostFtdcUnderlyingMultipleType;
				goType = "float64"
			case "int":
				// typedef int TThostFtdcPriorityType;
				goType = "int32"
			case "string":
			default:
				fmt.Println("未处理类型: ", t)
			}
			return goType
		},
	})
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
