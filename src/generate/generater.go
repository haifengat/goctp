package generate

import (
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"unicode"
)

func GenerateMd() {
	bsFile, err := ioutil.ReadFile("src/ctp_20180109_x64/ThostFtdcMdApi.h")
	checkErr(err)

	outFile := "src/go_ctp/ctp_quote.go"
	_, err = os.Stat(outFile)
	if err == nil || os.IsExist(err) { // 文件存在
		_ = os.Remove(outFile)
	}
	f, err := os.OpenFile(outFile, os.O_RDWR|os.O_CREATE, os.ModePerm)
	checkErr(err)
	defer func() { _ = f.Close() }()

	var funcVar string   // 函数在var中的声明 uintptr
	var funDefine string // 主调函数定义 funcRegisterFront := h.MustFindProc("RegisterFront")
	var funBody string   // 函数主体

	var cbSet string        // 在init函数中进行回调函数定义 h.MustFindProc("SetOnFrontConnected").Call(spi, syscall.NewCallback(OnConnect))
	var cbBody string       // 回调函数主体
	var cbTypeDefine string // 回调函数类型定义 type xxx func()
	var cbVar string        //回调函数变量

	// 汉字处理
	bsFile, _ = simplifiedchinese.GB18030.NewDecoder().Bytes(bsFile)
	/*
		///登录请求响应
		virtual void OnRspUserLogin(CThostFtdcRspUserLoginField *pRspUserLogin, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {};
	*/
	re := regexp.MustCompile(`\t///(.*)\r\n[^v]*virtual\s*(\w*)\s*(\w*)\(([^)]*)\)`) // 分解函数定义:注释,返回类型,函数名,参数字段四部分
	funs := re.FindAllStringSubmatch(string(bsFile), -1)
	for _, fun := range funs {
		funComment, _, funName, funFields := fun[1], fun[2], fun[3], fun[4]
		// 回调函数
		if strings.Index(funName, "On") == 0 {
			// type OnRspUserLoginType func(*CThostFtdcRspUserLoginField, *CThostFtdcRspInfoField, int, bool)
			cbTypeDefine += fmt.Sprintf("type on%sType func(", funName[2:]) // 回调函数类型定义
			if funFields != "" {                                            // 参数可能为空
				re = regexp.MustCompile(`(\w*)\s*([*]?\w*)[,]?\s*`) //参数分解:类型,名称
				fields := re.FindAllStringSubmatch(funFields, -1)
				for _, field := range fields {
					fieldType := []rune(field[1])
					if unicode.IsUpper(fieldType[0]) {
						cbTypeDefine += "*"
					}
					cbTypeDefine += string(fieldType) + ", "
				}
			}
			cbTypeDefine = strings.TrimRight(cbTypeDefine, ", ") + ") uintptr\n"

			/*
			   // 登录请求响应
			   func (p *Quote) RegOnRspUserLogin(on OnRspUserLoginType) {
			   	_, _, err := p.h.MustFindProc("SetOnRspUserLogin").Call(p.spi, syscall.NewCallback(on))
			   	checkErr(err)
			   }
			*/
			temp := `
// {{.funComment}}
func (q *quote) reg{{.OnFunName}}(on on{{.funName}}Type){
	_, _, _ = q.h.MustFindProc("Set{{.OnFunName}}").Call(q.spi, syscall.NewCallback(on))
}
`
			data := map[string]string{
				"funComment": funComment,
				"funName":    funName[2:],
				"OnFunName":  funName,
			}
			cbSet += templateMap(temp, data)
		} else {
			// 主调函数
			funContent, funParams, callParams := "", "", ""
			// 函数定义 p.funcReqUserLogin = p.h.MustFindProc("ReqUserLogin")
			funDefine += fmt.Sprintf("\tq.func%s = q.h.MustFindProc(\"%s\")\n", funName, funName)
			// 函数变量(放在结构体声明{}中) funcReqUserLogin  *syscall.Proc
			funcVar += fmt.Sprintf("\tfunc%s *syscall.Proc\n", funName)

			if funFields != "" { // 参数可能为空
				re = regexp.MustCompile(`(\w*)\s*([*]?\s*\w*[\[]?[]]?)[,]?\s*`) //参数分解:类型,名称
				fields := re.FindAllStringSubmatch(funFields, -1)
				for _, field := range fields {
					isPtrVar := strings.Index(field[2], "*") == 0
					fieldType, fieldName := field[1], strings.TrimRight(strings.TrimLeft(field[2], "*"), "[]")

					// 处理char *ppInstrumentID[] ==> []byte
					if strings.Contains(field[2], "[]") {
						funParams += fmt.Sprintf("%s [1][]byte, nCount int", fieldName)
						callParams += fmt.Sprintf(", uintptr(unsafe.Pointer(&%s)), uintptr(1)", fieldName)
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
							funContent += "\n\tq.nRequestID++"
							callParams += fmt.Sprintf(", uintptr(q.%s)", fieldName)
						} else {
							funParams += fmt.Sprintf("%s %s, ", strings.TrimLeft(fieldName, " "), fieldType)
							callParams += fmt.Sprintf(", uintptr(%s)", fieldName)
						}
					}
				}
			}
			funParams = strings.TrimRight(funParams, ", ")

			/*// 用户登录请求
			  func (p *Quote) ReqUserLogin(pReqUserLoginField CThostFtdcReqUserLoginField) {
			  	p.nRequestID++
			  	p.funcReqUserLogin.Call(p.api, uintptr(unsafe.Pointer(&pReqUserLoginField)), uintptr(p.nRequestID))
			  }*/
			temp := `
// {{.funComment}}
func (q *quote) {{.funName}}({{.funParams}}){ {{.funContent}}
	_, _, _ = q.func{{.funName}}.Call(q.api{{.callParams}})
}
`
			data := map[string]string{
				"funComment": funComment,
				"funName":    funName,
				"funParams":  funParams,
				"funContent": funContent,
				"callParams": callParams,
			}
			funBody += templateMap(temp, data)
		}
	}

	temp := `package go_ctp

import (
	"os"
	"path"
	"path/filepath"
	"runtime"
	"syscall"
	"unsafe"
)

type CThostFtdcMdSpi uintptr
{{.cbTypeDefine}}
type quote struct {
	h        *syscall.DLL
	api, spi  uintptr
	nRequestID      int
	funcCreateApi,funcCreateSpi *syscall.Proc
{{.funcVar}}
{{.cbVar}}
}

func (q *quote) loadDll() {
	// 执行目录下创建 log目录
	_, err := os.Stat("log")
	if err != nil {
		checkErr(os.Mkdir("log", os.ModePerm))
	}
	workPath, _ := os.Getwd()
	_, curFile, _, ok := runtime.Caller(1)
	if !ok {
		panic("取当前文件路径失败")
	}
	dllPath := filepath.Dir(curFile)
	checkErr(os.Chdir(path.Join(dllPath, "lib64")))
	q.h = syscall.MustLoadDLL("ctp_quote.dll")
	// 还原到之前的工作目录
	checkErr(os.Chdir(workPath))
	//defer h.Release() // 函数结束后会释放导致后续函数执行失败
}

// 行情接口
func newQuote() *quote {
	q := new(quote)

	q.loadDll()
{{.funDefine}}
	q.api, _, _ = q.h.MustFindProc("CreateApi").Call()
	q.spi, _, _ = q.h.MustFindProc("CreateSpi").Call()
	_, _, _ = q.funcRegisterSpi.Call(q.api, uintptr(unsafe.Pointer(q.spi)))
	return q
}

{{.cbSet}}
`
	// 变量替换
	data := map[string]string{
		"funcVar":      funcVar,
		"funDefine":    funDefine,
		"cbSet":        cbSet,
		"cbVar":        cbVar,
		"cbTypeDefine": cbTypeDefine,
	}
	buf := templateMap(temp, data)
	_, _ = f.WriteString(buf)
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
		checkErr(os.Remove(outFile))
	}

	f, err := os.OpenFile(outFile, os.O_RDWR|os.O_CREATE, os.ModePerm)
	checkErr(err)
	defer func() { _ = f.Close() }()

	_, err = f.WriteString("package go_ctp\n\n")
	checkErr(err)

	// 汉字处理
	bsFile, err = simplifiedchinese.GB18030.NewDecoder().Bytes(bsFile)
	checkErr(err)
	/*
		///信息分发
		struct CThostFtdcDisseminationField
		{
			///序列系列号
			TThostFtdcSequenceSeriesType	SequenceSeries;
			///序列号
			TThostFtdcSequenceNoType	SequenceNo;
		};
	*/
	re := regexp.MustCompile(`///(\S*)\s*struct\s*(\w*)\s*{([^}]*)}`) // 分成struct的注释,名称,字段两部分
	structs := re.FindAllStringSubmatch(string(bsFile), -1)
	for _, strc := range structs {
		/*
			// 信息分发
			type CThostFtdcDisseminationField struct {
				// 序列系列号
				SequenceSeries TThostFtdcSequenceSeriesType
				// 序列号
				SequenceNo TThostFtdcSequenceNoType
			}
		*/
		//strcName := strc[1]
		strStruc := fmt.Sprintf("// %s\ntype %s struct{\n", strc[1], strc[2])
		re = regexp.MustCompile(`///([^\r\n]*)\s*(\w*)\s*([^;]*);`) // 所有字段再分解成各个单独字段: 注释(可能含空格),类型,名称
		fields := re.FindAllStringSubmatch(strc[3], -1)
		for _, field := range fields {
			strStruc += fmt.Sprintf("\t// %s\n\t%s %s\n", field[1], field[3], field[2])
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

	outFile := "src/go_ctp/ctp_datatype.go"
	_, err = os.Stat(outFile)
	if err == nil || os.IsExist(err) { // 文件存在
		err = os.Remove(outFile)
		checkErr(err)
	}

	f, err := os.OpenFile(outFile, os.O_RDWR|os.O_CREATE, os.ModePerm)
	checkErr(err)
	defer func() { _ = f.Close() }()
	_, _ = f.WriteString("package generate\n\n")
	for _, str := range dataType {
		//fmt.Println(str)
		// 写文件
		if strings.Index(str, "//") == 0 {
			_, _ = f.WriteString("\n") // 注释前加一行
		}
		_, _ = f.WriteString(str + "\n")
	}
}
