package generate

import (
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"strings"
	"unicode"
)

// 接口源目录
var srcPath = "src/ctp_20190220_se_x64/"

func GenerateCtp(tradeOrQuote string) {
	var (
		bsFile []byte
		err    error
		// Trade or Quote
		title string
		// q or t
		firstChar string
		// 输出文件名(带相对路径)
		outFileName string
		// 函数在var中的声明 funcRegisterFront *syscall.Proc
		funcVar string
		// 主调函数定义 funcRegisterFront := h.MustFindProc("RegisterFront")
		funDefine string
		// 函数主体
		funBody string
		// 在init函数中进行回调函数定义 h.MustFindProc("SetOnFrontConnected").Call(spi, syscall.NewCallback(OnConnect))
		cbSet string
		// 回调函数主体
		cbBody string
		// 回调函数类型定义 type xxx func()
		cbTypeDefine string
		//回调函数变量 onFrontConnected OnFrontConnectedType
		cbVar string
		// 底层接口的响应类名 ThostFtdcMdSpi ThostFtdcTraderSpi
		cbName string
	)
	title = strings.Title(tradeOrQuote)
	firstChar = string(tradeOrQuote[0])
	if tradeOrQuote == "trade" {
		bsFile, err = ioutil.ReadFile(path.Join(srcPath, "ThostFtdcTraderApi.h"))
		cbName = "CThostFtdcTraderSpi"
	} else {
		bsFile, err = ioutil.ReadFile(path.Join(srcPath, "ThostFtdcMdApi.h"))
		cbName = "CThostFtdcMdSpi"
	}
	checkErr(err)

	outFileName = fmt.Sprintf("src/go_ctp/ctp_%s.go", tradeOrQuote)
	checkErr(err)
	_, err = os.Stat(outFileName)
	if err == nil || os.IsExist(err) { // 文件存在=>删除
		_ = os.Remove(outFileName)
	}
	f, err := os.OpenFile(outFileName, os.O_RDWR|os.O_CREATE, os.ModePerm)
	checkErr(err)
	defer func() { _ = f.Close() }()

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
			cbTypeDefine += fmt.Sprintf("type %sOn%sType func(", firstChar, funName[2:]) // 回调函数类型定义
			if funFields != "" {                                                         // 参数可能为空
				re = regexp.MustCompile(`(\w*)\s*([*]?\w*)[,]?\s*`) //参数分解:类型,名称
				fields := re.FindAllStringSubmatch(funFields, -1)
				for _, field := range fields {
					fieldType := []rune(field[1])
					if unicode.IsUpper(fieldType[0]) {
						cbTypeDefine += "*t"
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
func ({{.firstChar}} *{{.name}}) reg{{.OnFunName}}(on {{.firstChar}}On{{.funName}}Type){
	_, _, _ = {{.firstChar}}.h.MustFindProc("Set{{.OnFunName}}").Call({{.firstChar}}.spi, syscall.NewCallback(on))
}
`
			data := map[string]string{
				"funComment": funComment,
				"funName":    funName[2:],
				"OnFunName":  funName,
				"firstChar":  firstChar,
				"name":       tradeOrQuote,
			}
			cbSet += templateMap(temp, data)
		} else {
			if funName == "RegisterSpi" {
				continue
			}
			// 主调函数
			funContent, funParams, callParams := "", "", ""
			// 函数定义 p.funcReqUserLogin = p.h.MustFindProc("ReqUserLogin")
			funDefine += fmt.Sprintf("\t%s.func%s = %s.h.MustFindProc(\"%s\")\n", firstChar, funName, firstChar, funName)
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
							funParams += fmt.Sprintf("%s t%s, ", strings.TrimLeft(fieldName, " "), fieldType)
							callParams += fmt.Sprintf(", uintptr(unsafe.Pointer(&%s))", fieldName)
						}
					} else {
						if fieldName == "nRequestID" {
							funContent += fmt.Sprintf("\n\t%s.nRequestID++", firstChar)
							callParams += fmt.Sprintf(", uintptr(%s.%s)", firstChar, fieldName)
						} else {
							funParams += fmt.Sprintf("%s t%s, ", strings.TrimLeft(fieldName, " "), fieldType)
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
func ({{.firstChar}} *{{.name}}) {{.funName}}({{.funParams}}){ {{.funContent}}
	_, _, _ = {{.firstChar}}.func{{.funName}}.Call({{.firstChar}}.api{{.callParams}})
}
`
			data := map[string]string{
				"funComment": funComment,
				"funName":    funName,
				"funParams":  funParams,
				"funContent": funContent,
				"callParams": callParams,
				"firstChar":  firstChar,
				"name":       tradeOrQuote,
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

type {{.cbName}} uintptr
{{.cbTypeDefine}}
type {{.name}} struct {
	h        *syscall.DLL
	api, spi  uintptr
	nRequestID      int
	funcCreateApi,funcCreateSpi *syscall.Proc
{{.funcVar}}
{{.cbVar}}
}

func ({{.firstChar}} *{{.name}}) loadDll() {
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
	{{.firstChar}}.h = syscall.MustLoadDLL("ctp_{{.name}}.dll")
	// 还原到之前的工作目录
	checkErr(os.Chdir(workPath))
	//defer h.Release() // 函数结束后会释放导致后续函数执行失败
}

// 接口
func new{{.title}}() *{{.name}} {
	{{.firstChar}} := new({{.name}})

	{{.firstChar}}.loadDll()
{{.funDefine}}
	{{.firstChar}}.api, _, _ = {{.firstChar}}.h.MustFindProc("CreateApi").Call()
	{{.firstChar}}.spi, _, _ = {{.firstChar}}.h.MustFindProc("CreateSpi").Call()
	_, _, _ = {{.firstChar}}.h.MustFindProc("RegisterSpi").Call({{.firstChar}}.api, uintptr(unsafe.Pointer({{.firstChar}}.spi)))
	return {{.firstChar}}
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
		"name":         tradeOrQuote,
		"firstChar":    firstChar,
		"title":        title,
		"cbName":       cbName,
	}
	buf := templateMap(temp, data)
	_, _ = f.WriteString(buf)
	_, _ = f.WriteString("\n")
	_, _ = f.WriteString(cbBody)
	_, _ = f.WriteString(funBody)
}

func GenerateStruct() {
	bsFile, err := ioutil.ReadFile(path.Join(srcPath, "ThostFtdcUserApiStruct.h"))
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
		strStruc := fmt.Sprintf("// %s\ntype t%s struct{\n", strc[1], strc[2])
		re = regexp.MustCompile(`///([^\r\n]*)\s*(\w*)\s*([^;]*);`) // 所有字段再分解成各个单独字段: 注释(可能含空格),类型,名称
		fields := re.FindAllStringSubmatch(strc[3], -1)
		for _, field := range fields {
			strStruc += fmt.Sprintf("\t// %s\n\t%s t%s\n", field[1], field[3], field[2])
		}
		strStruc += "}\n\n"
		_, _ = f.WriteString(strStruc)
		//fmt.Println(strStruc)
	}
}

func GenerateDataType() {
	bsFile, err := ioutil.ReadFile(path.Join(srcPath, "ThostFtdcUserApiDataType.h"))
	checkErr(err)
	// 汉字处理
	bsFile, _ = simplifiedchinese.GB18030.NewDecoder().Bytes(bsFile)
	// 分行
	lines := strings.Split(string(bsFile), "\n")

	var dataType []string
	var constDef []string
	transType := make(map[string]string)
	transType[`typedef\s*char\s*(\w+)\[(\d+)\].*\s*$`] = "type t$1 [$2]byte" // typedef char TThostFtdcTraderIDType[21]; ==> type TThostFtdcTraderIDType [21]byte
	transType[`typedef\s*char\s*(\w+);.*\s*$`] = "type t$1 byte"             //typedef char TThostFtdcIdCardTypeType; ==> type TThostFtdcIdCardTypeType byte
	transType[`typedef\s*int\s*(\w+);.*\s*$`] = "type t$1 int32"             // int与go int32对应
	transType[`typedef\s*double\s*(\w+);.*\s*$`] = "type t$1 float64"
	transType[`typedef\s*short\s*(\w+);.*\s*$`] = "type t$1 int16"
	//transType[`typedef\s*\b(int|double|short)\b\s*(\w+);.*\s*$`] = "type $2 $1" // typedef int TThostFtdcIPPortType; ==> type TThostFtdcIPPortType int / double /short
	transType[`#define\s*(\w+)\s*'(\w+)'\s*$`] = "const |enumtype|_t$1 = '$2'" //#define THOST_FTDC_ICT_AccountsPermits 'J' ==> const THOST_FTDC_ICT_AccountsPermits = 'J'

	for i, line := range lines {
		var res []string
		for k, v := range transType {
			// 根据规则判断此等是否要转换
			re := regexp.MustCompile(k)
			res = re.FindAllString(line, -1)
			if res != nil {
				isConst := v == "const |enumtype|_t$1 = '$2'" // 常量匹配
				// 取注释信息
				commentReg := `^.*是一个(.*)\s*$`
				if isConst {
					commentReg = `^///([^/\^]*)\s*$` // 常量注释信息排除全是/的行和///^\d的特殊行
				}

				// 取注释信息
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
				//val = fmt.Sprintf("%s%s", strings.ToLower(string(val[0])), val[1:])  // 首字母小写,不让外部看到 THOST_FTDC_ICT_AccountsPermits=>tHOST_FTDC_ICT_AccountsPermits
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
						if v == "type t$1 byte" {
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
	_, _ = f.WriteString(`package go_ctp

type tTHOST_TE_RESUME_TYPE int8

const (
	THOST_TERT_RESTART tTHOST_TE_RESUME_TYPE = 0
	THOST_TERT_RESUME  tTHOST_TE_RESUME_TYPE = 1
	THOST_TERT_QUICK   tTHOST_TE_RESUME_TYPE = 2
)
`)
	for _, str := range dataType {
		//fmt.Println(str)
		// 写文件
		if strings.Index(str, "//") == 0 {
			_, _ = f.WriteString("\n") // 注释前加一行
		}
		_, _ = f.WriteString(str + "\n")
	}
}
