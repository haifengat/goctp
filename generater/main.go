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

var srcPath = "../CTPv6.7.2_20230913/"

func main() {
	genQuoteC()
	genTradeC()
	genDataType()
	genStruct()
	genQuoteGo()
	genTradeGo()
	fmt.Println("运行 sh make_lib.sh 重新编译 .so")
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

func genTradeGo() {
	fn, on := getFuncs("ThostFtdcTraderApi.h")
	tmpl("trade.go.tpl", map[string]any{"Fn": fn, "On": on}, "../", template.FuncMap{
		"APIpath": func() string {
			return path.Base(srcPath) // 返回 CTPv6.6.9_20220922
		},
		"toCGo": func(typ string) string {
			if strings.HasPrefix(typ, "CThostFtdc") {
				if typ == "CThostFtdcTraderSpi" {
					return "void"
				}
				return "struct " + typ
			}
			switch typ {
			case "bool":
				typ = "_Bool"
			case "int", "char":
			case "THOST_TE_RESUME_TYPE":
				typ = "int"
			default:
				fmt.Println("需处理:", typ)
			}
			return typ
		},
		"exToCGo": func(typ string) string {
			if strings.HasPrefix(typ, "CThostFtdc") {
				if typ == "CThostFtdcTraderSpi" {
					return "void"
				}
				return "C.struct_" + typ
			}
			switch typ {
			case "int":
				typ = "C.int"
			case "char":
				typ = "C.char"
			case "bool":
				typ = "C._Bool"
			default:
				fmt.Println("exToCGo 未处理", typ)
			}
			return typ
		},
		"toGoType": func(typ, name string) string {
			if typ == "CThostFtdcTraderSpi" { // spi
				typ = "unsafe.Pointer"
			} else if strings.HasPrefix(typ, "CThostFtdc") { // struct
				typ = "*" + typ
			} else {
				switch typ {
				case "char":
					typ = "string"
				case "int", "bool": // 类型名一样
				case "THOST_TE_RESUME_TYPE":
					typ = "THOST_TE_RESUME_TYPE"
				default:
					fmt.Println("toGoType 未处理", typ)
				}
			}
			return typ
		},
		"fnVar": func(typ, name string) string {
			if typ == "CThostFtdcTraderSpi" { // spi

			} else if strings.HasPrefix(typ, "CThostFtdc") { // struct
				name = fmt.Sprintf("(*C.struct_%s)(unsafe.Pointer(%s))", typ, name)
			} else {
				switch typ {
				case "char":
					name = fmt.Sprintf("C.CString(%s)", name)
				case "int", "THOST_TE_RESUME_TYPE":
					name = fmt.Sprintf("C.int(%s)", name)
				default:
					fmt.Println("fnVar 未处理", name)
				}
			}
			return name
		},
		"onVar": func(typ, name string) string {
			if strings.HasPrefix(typ, "CThostFtdc") { // struct
				name = fmt.Sprintf("(*%s)(unsafe.Pointer(%s))", typ, name)
			} else {
				switch typ {
				case "char":
					name = fmt.Sprintf("C.GString(%s)", name)
				case "int", "bool":
					name = fmt.Sprintf("%s(%s)", typ, name)
				default:
					fmt.Println("onVar 未处理", name)
				}
			}
			return name
		},
	})
}

func genTradeC() {
	fn, on := getFuncs("ThostFtdcTraderApi.h")
	tmpl("trade.h.tpl", map[string]any{"Fn": fn, "On": on}, "../c/", template.FuncMap{"APIpath": func() string {
		return path.Base(srcPath) // 返回 CTPv6.6.9_20220922
	}})
	tmpl("trade.cpp.tpl", map[string]any{"Fn": fn, "On": on}, "../c/", template.FuncMap{"APIpath": func() string {
		return path.Base(srcPath) // 返回 CTPv6.6.9_20220922
	}})
}

func genQuoteGo() {
	fn, on := getFuncs("ThostFtdcMdApi.h")
	tmpl("quote.go.tpl", map[string]any{"Fn": fn, "On": on}, "../", template.FuncMap{
		"APIpath": func() string {
			return path.Base(srcPath) // 返回 CTPv6.6.9_20220922
		},
		"toCGo": func(typ string) string {
			if strings.HasPrefix(typ, "CThostFtdc") {
				if typ == "CThostFtdcMdSpi" {
					return "void"
				}
				return "struct " + typ
			}
			switch typ {
			case "bool":
				typ = "_Bool"
			case "int", "char":

			default:
				fmt.Println("需处理:", typ)
			}
			return typ
		},
		"exToCGo": func(typ string) string {
			if strings.HasPrefix(typ, "CThostFtdc") {
				if typ == "CThostFtdcMdSpi" {
					return "void"
				}
				return "C.struct_" + typ
			}
			switch typ {
			case "int":
				typ = "C.int"
			case "char":
				typ = "C.char"
			case "bool":
				typ = "C._Bool"
			default:
				fmt.Println("未处理", typ)
			}
			return typ
		},
		"toGoType": func(typ, name string) string {
			if typ == "CThostFtdcMdSpi" { // spi
				typ = "unsafe.Pointer"
			} else if strings.HasPrefix(typ, "CThostFtdc") { // struct
				typ = "*" + typ
			} else if name == "ppInstrumentID" { // char 数组
				typ = "[]string"
			} else {
				switch typ {
				case "char":
					typ = "string"
				case "int", "bool": // 类型名一样
				default:
					fmt.Println("未处理", typ)
				}
			}
			return typ
		},
		"fnVar": func(typ, name string) string {
			if typ == "CThostFtdcMdSpi" { // spi

			} else if strings.HasPrefix(typ, "CThostFtdc") { // struct
				name = fmt.Sprintf("(*C.struct_%s)(unsafe.Pointer(%s))", typ, name)
			} else {
				switch typ {
				case "char":
					name = fmt.Sprintf("C.CString(%s)", name)
				case "int":
					name = fmt.Sprintf("C.int(%s)", name)
				default:
					fmt.Println("未处理", name)
				}
			}
			return name
		},
		"onVar": func(typ, name string) string {
			if strings.HasPrefix(typ, "CThostFtdc") { // struct
				name = fmt.Sprintf("(*%s)(unsafe.Pointer(%s))", typ, name)
			} else {
				switch typ {
				case "char":
					name = fmt.Sprintf("C.GString(%s)", name)
				case "int", "bool":
					name = fmt.Sprintf("%s(%s)", typ, name)
				default:
					fmt.Println("未处理", name)
				}
			}
			return name
		},
	})
}

func genQuoteC() {
	fn, on := getFuncs("ThostFtdcMdApi.h")
	tmpl("quote.h.tpl", map[string]any{"Fn": fn, "On": on}, "../c/", template.FuncMap{"APIpath": func() string {
		return path.Base(srcPath) // 返回 CTPv6.6.9_20220922
	}})
	tmpl("quote.cpp.tpl", map[string]any{"Fn": fn, "On": on}, "../c/", template.FuncMap{"APIpath": func() string {
		return path.Base(srcPath) // 返回 CTPv6.6.9_20220922
	}})
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
			if strings.HasPrefix(v[3], "reserve") {
				v[3] = strings.ReplaceAll(v[3], "reserve", "Reserve") // 避免 unused(U 1000)警告
			}
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

	tmpl("struct.go.tpl", datas, "../", nil)
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

	tmpl("./datatype.go.tpl", datas, "../", template.FuncMap{
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

	f, _ := os.Create("../datatype.go")
	f.WriteString(typeChars)
	f.WriteString(strings.Join(typeChar, "\n"))
}
