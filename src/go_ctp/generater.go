package go_ctp

import (
	"bytes"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"text/template"
)

func GenerateMd() {
	bsFile, err := ioutil.ReadFile("src/ctp_20180109_x64/ThostFtdcMdApi.h")
	checkErr(err)

	outFile := "src/go_ctp/ctp_quote.go"
	_, err = os.Stat(outFile)
	if err == nil || os.IsExist(err) { // 文件存在
		os.Remove(outFile)
	}
	f, err := os.OpenFile(outFile, os.O_RDWR|os.O_CREATE, os.ModePerm)
	checkErr(err)
	defer f.Close()
	strHead := `package go_ctp

import (
	"fmt"
	"os"
	"syscall"
	"unsafe"
)

type CThostFtdcMdSpi uintptr
{{.cbTypeDefine}}
type Quote struct {
	h        *syscall.DLL
	api, spi uintptr
	nRequestID      int
{{.funcVar}}
{{.cbVar}}
}

func (p *Quote) Quote(){
	// 创建 log目录
	_, err := os.Stat("log")
	if err != nil {
		os.Mkdir("log", os.ModePerm)
	}

	p.h = syscall.MustLoadDLL("ctp_quote.dll")
	//defer h.Release() // 函数结束后会释放导致后续函数执行失败

	p.api, _, _ = p.h.MustFindProc("CreateApi").Call()
	p.spi, _, _ = p.h.MustFindProc("CreateSpi").Call()
	p.h.MustFindProc("RegisterSpi").Call(p.api, p.spi)
{{.funDefine}}
{{.cbSet}}
}
`

	var funcVar string   // 函数在var中的声明 uintptr
	var funDefine string // 主调函数定义 funcRegisterFront := h.MustFindProc("RegisterFront")
	var funBody string   // 函数主体

	var cbSet string        // 在init函数中进行回调函数定义 h.MustFindProc("SetOnFrontConnected").Call(spi, syscall.NewCallback(OnConnect))
	var cbBody string       // 回调函数主体
	var cbTypeDefine string // 回调函数类型定义 type xxx func()
	var cbVar string        //回调函数变量

	// 汉字处理
	bsFile, _ = simplifiedchinese.GB18030.NewDecoder().Bytes(bsFile)
	//fmt.Println(string(bsFile))
	re := regexp.MustCompile(`\t///(.*)\r\n[^v]*virtual\s*(\w*)\s*(\w*)\(([^\)]*)\)`) // 分解函数定义:注释,返回类型,函数名,参数字段两部分
	funs := re.FindAllStringSubmatch(string(bsFile), -1)
	for _, fun := range funs {
		funComment, _, funName, funFields := fun[1], fun[2], fun[3], fun[4]
		// 回调函数
		if strings.Index(funName, "On") == 0 {
			cbTypeDefine += fmt.Sprintf("type %sType func(", funName) // 回调函数类型定义
			cbBody += fmt.Sprintf("// %s\nfunc (p *Quote) %s(", funComment, funName)
			funContent := fmt.Sprintf("\tif p.cb%s != nil{\n\t\tp.cb%s(", funName, funName)
			if funFields != "" { // 参数可能为空
				re = regexp.MustCompile(`(\w*)\s*([*]?\w*)[,]?\s*`) //参数分解:类型,名称
				fields := re.FindAllStringSubmatch(funFields, -1)
				for _, field := range fields {
					fieldType, fieldName := field[1], field[2]
					cbTypeDefine += fieldType + ", "

					if strings.Index(fieldName, "*") >= 0 {
						cbBody += fmt.Sprintf("%s uintptr, ", fieldName[1:]) // 函数参数
						funContent += fmt.Sprintf("*(*%s)(unsafe.Pointer(%s)), ", fieldType, fieldName[1:])
						//funContent += fmt.Sprintf("\tp%s := (*%s)(unsafe.Pointer(%s))\n", fieldType, fieldType, fieldName[1:]) //pCThostFtdcRspInfoField := (*CThostFtdcRspInfoField)(unsafe.Pointer(info))
						//funContent += fmt.Sprintf("\tfmt.Println(p%s)\n", fieldType)
					} else {
						cbBody += fmt.Sprintf("%s %s,", fieldName, fieldType)
						funContent += fieldName + ", "
					}
				}
			}
			cbTypeDefine = strings.TrimRight(cbTypeDefine, ", ") + ")\n"
			cbBody = strings.TrimRight(cbBody, ", ") // 去掉尾部,
			cbBody += fmt.Sprintf(") uintptr {\n%s)\n\t}\n\treturn 0\n}\n\n", strings.TrimRight(funContent, ", "))
			cbVar += fmt.Sprintf("\tcb%s %sType\n", funName, funName)
			cbSet += fmt.Sprintf("\tp.h.MustFindProc(\"Set%s\").Call(p.spi, syscall.NewCallback(p.%s))\n", funName, funName)
		} else { // 主调函数
			funContent, funParams, callParams := "", "", ""
			// 函数定义 funcRegisterFront := h.MustFindProc("RegisterFront")
			funDefine += fmt.Sprintf("\tp.func%s = p.h.MustFindProc(\"%s\")\n", funName, funName)
			funcVar += fmt.Sprintf("\tfunc%s *syscall.Proc\n", funName)

			if funFields != "" { // 参数可能为空
				re = regexp.MustCompile(`(\w*)\s*([*]?\s*\w*[\[]?[\]]?)[,]?\s*`) //参数分解:类型,名称
				fields := re.FindAllStringSubmatch(funFields, -1)
				for _, field := range fields {
					isPtrVar := strings.Index(field[2], "*") == 0
					fieldType, fieldName := field[1], strings.TrimRight(strings.TrimLeft(field[2], "*"), "[]")

					// 处理char *ppInstrumentID[] ==> []string
					if strings.Contains(field[2], "[]") {
						funParams += fmt.Sprintf("%s []string, nCount int", fieldName)
						callParams += fmt.Sprintf(", uintptr(unsafe.Pointer(&%s)), uintptr(nCount)", fieldName)
						break
					} else if isPtrVar { // 带*的变量
						// string -> uintptr(unsafe.Pointer()) -> uintptr(unsafe.Pointer(bs))
						if fieldType == "char" {
							funParams += fmt.Sprintf("%s string, ", fieldName)
							funContent += fmt.Sprintf("\n\tbs, _ := syscall.BytePtrFromString(%s)", fieldName)
							callParams += ", uintptr(unsafe.Pointer(bs))"
						} else { // struct
							funParams += fmt.Sprintf("%s %s, ", strings.TrimLeft(fieldName, " "), fieldType)
							callParams += fmt.Sprintf(", uintptr(unsafe.Pointer(&%s))", fieldName)
						}
					} else {
						if fieldName == "nRequestID" {
							funContent += "\n\tp.nRequestID++"
							callParams += fmt.Sprintf(", uintptr(p.%s)", fieldName)
						} else {
							funParams += fmt.Sprintf("%s %s, ", strings.TrimLeft(fieldName, " "), fieldType)
							callParams += fmt.Sprintf(", uintptr(%s)", fieldName)
						}
					}
				}
			}
			funParams = strings.TrimRight(funParams, ", ")
			data := map[string]interface{}{
				"funComment": funComment,
				"funName":    funName,
				"funParams":  funParams,
				"funContent": funContent,
				"callParams": callParams,
			}

			temp := `
// {{.funComment}}
func (p *Quote) {{.funName}}({{.funParams}}){ {{.funContent}}
	p.func{{.funName}}.Call(p.api{{.callParams}})
}
`
			t := template.Must(template.New("fun").Parse(temp))
			buf := &bytes.Buffer{}
			if err := t.Execute(buf, data); err != nil {
				panic(err)
			}
			funBody += buf.String()
		}
	}

	// 变量替换
	data := map[string]interface{}{
		"funcVar":      funcVar,
		"funDefine":    funDefine,
		"cbSet":        cbSet,
		"cbVar":        cbVar,
		"cbTypeDefine": cbTypeDefine,
	}
	t := template.Must(template.New("api").Parse(strHead))
	buf := &bytes.Buffer{}
	if err := t.Execute(buf, data); err != nil {
		panic(err)
	}
	_, _ = f.WriteString(buf.String())
	_, _ = f.WriteString("\n")
	_, _ = f.WriteString(cbBody)
	_, _ = f.WriteString(funBody)
}

func GenerateStruct() {
	bsFile, err := ioutil.ReadFile("src/ctp_20180109_x64/ThostFtdcUserApiStruct.h")
	checkErr(err)

	outFile := "src/go_ctp/ctp_struct.go"
	_, err = os.Stat(outFile)
	if err == nil || os.IsExist(err) { // 文件存在
		os.Remove(outFile)
	}

	f, err := os.OpenFile(outFile, os.O_RDWR|os.O_CREATE, os.ModePerm)

	checkErr(err)
	defer f.Close()
	_, _ = f.WriteString("package go_ctp\n\n")

	// 汉字处理
	bsFile, _ = simplifiedchinese.GB18030.NewDecoder().Bytes(bsFile)
	//fmt.Println(string(bsFile))
	re := regexp.MustCompile(`///(\S*)\s*struct\s*(\w*)\s*\{([^\}]*)}`) // 分成struct的注释,名称,字段两部分
	structs := re.FindAllStringSubmatch(string(bsFile), -1)
	for _, strc := range structs {
		//strcName := strc[1]
		strStruc := fmt.Sprintf("// %s\ntype %s struct{\n", strc[1], strc[2])
		re = regexp.MustCompile(`///([^\r\n]*)\s*(\w*)\s*([^;]*);`) // 所有字段再分解成各个单独字段: 注释(可能含空格),类型,名称
		fields := re.FindAllStringSubmatch(strc[3], -1)
		for _, field := range fields {
			strStruc += fmt.Sprintf("\t// %s\n\t%s %s\n", field[1], field[3], field[2])
			//fmt.Println(field)
		}
		strStruc += "}\n\n"
		_, _ = f.WriteString(strStruc)
		//fmt.Println(strStruc)
	}
}

func GenerateDataType() {
	bsFile, err := ioutil.ReadFile("src/ctp_20180109_x64/ThostFtdcUserApiDataType.h")
	checkErr(err)
	// 汉字处理
	bsFile, _ = simplifiedchinese.GB18030.NewDecoder().Bytes(bsFile)
	// 分行
	lines := strings.Split(string(bsFile), "\n")

	var dataType []string
	var constDef []string
	transType := make(map[string]string)
	transType[`typedef\s*char\s*(\w+)\[(\d+)\].*\s*$`] = "type $1 [$2]byte" // typedef char TThostFtdcTraderIDType[21]; ==> type TThostFtdcTraderIDType [21]byte
	transType[`typedef\s*char\s*(\w+);.*\s*$`] = "type $1 byte"             //typedef char TThostFtdcIdCardTypeType; ==> type TThostFtdcIdCardTypeType byte
	transType[`typedef\s*int\s*(\w+);.*\s*$`] = "type $1 int"
	transType[`typedef\s*double\s*(\w+);.*\s*$`] = "type $1 float64"
	transType[`typedef\s*short\s*(\w+);.*\s*$`] = "type $1 int16"
	//transType[`typedef\s*\b(int|double|short)\b\s*(\w+);.*\s*$`] = "type $2 $1" // typedef int TThostFtdcIPPortType; ==> type TThostFtdcIPPortType int / double /short
	transType[`#define\s*(\w+)\s*'(\w+)'\s*$`] = "const |enumtype|_$1 = '$2'" //#define THOST_FTDC_ICT_AccountsPermits 'J' ==> const THOST_FTDC_ICT_AccountsPermits = 'J'

	for i, line := range lines {
		var res []string
		for k, v := range transType {
			// 根据规则判断此等是否要转换
			re := regexp.MustCompile(k)
			res = re.FindAllString(line, -1)
			if res != nil {
				isConst := v == "const |enumtype|_$1 = '$2'" // 常量匹配
				// 取注释信息
				commentReg := `^.*是一个(.*)\s*$`
				if isConst {
					commentReg = `^///([^/\^]*)\s*$` // 常量注释信息排除全是/的行和///^\d的特殊行
				}

				for j := i - 1; j > 0; j-- {
					re := regexp.MustCompile(commentReg)
					res = re.FindAllString(lines[j], -1)
					if res != nil {
						comment := re.ReplaceAllString(lines[j], "// $1")
						if isConst { // 常量注释信息
							constDef = append(constDef, comment)
						} else {
							exists, _ := Contain(comment, dataType)
							if !exists {
								dataType = append(dataType, comment)
							}
						}
						break
					}
				}

				// 转换后的结果保存
				val := re.ReplaceAllString(line, v)
				// 处理 enum 的值
				if isConst {
					// THOST_FTDC_ICT_AccountsPermits ==> AccountsPermits 只保留最后一个元素,避免再加上类型前缀后过长
					valName := strings.Split(val, " ")[1]
					valNameSlice := strings.Split(valName, "_")
					val = strings.Replace(val, valName, valNameSlice[0]+"_"+valNameSlice[len(valNameSlice)-1], 1)
					// 处理特殊情况: '200001'
					if len(strings.Split(val, "'")[1]) > 1 {
						val = strings.Replace(val, "'", "\"", -1)
					}
					constDef = append(constDef, val)
					commentReg = `^///([^/]*)\s*$` // 常量注释信息
				} else {
					exists, _ := Contain(val, dataType)
					if !exists {
						dataType = append(dataType, val)
						// 处理 enum
						if v == "type $1 byte" {
							tName := re.ReplaceAllString(line, "$1")
							for _, cst := range constDef {
								dataType = append(dataType, strings.Replace(cst, "|enumtype|", tName, 1)) // 给enum的值加上前缀
							}
							constDef = make([]string, 0) //清除enum定义的值
						}
					}
				}
				break //跳出规则判断 下一行
			}
		}
	}

	outFile := "src/go_ctp/datatype.go"
	_, err = os.Stat(outFile)
	if err == nil || os.IsExist(err) { // 文件存在
		os.Remove(outFile)
	}

	f, err := os.OpenFile(outFile, os.O_RDWR|os.O_CREATE, os.ModePerm)
	checkErr(err)
	defer f.Close()
	_, _ = f.WriteString("package go_ctp\n\n")
	for _, str := range dataType {
		//fmt.Println(str)
		// 写文件
		if strings.Index(str, "//") == 0 {
			_, _ = f.WriteString("\n") // 注释前加一行
		}
		_, _ = f.WriteString(str + "\n")
	}
}
