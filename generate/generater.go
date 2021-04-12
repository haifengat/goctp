package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"strings"
	"text/template"
	"unicode"

	"golang.org/x/text/encoding/simplifiedchinese"
)

// 接口源目录
var (
	srcPath     = "./v6.5.1_20200908/"
	outPath     = "./ctpdefine"
	packageName = "gitee.com/haifengat/goctp"
)

func main() {
	fmt.Println("run generater.go in parent dir of it.")
	generateDataType() // 生成后需和动处理出入金的几个定义
	generateStruct()
	generateCtp("trade")
	generateCtp("quote")
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func templateMap(templateString string, mapContent map[string]string) string {
	t := template.Must(template.New("fun").Parse(templateString))
	buf := &bytes.Buffer{}
	checkErr(t.Execute(buf, mapContent))
	return buf.String()
}

func generateCtp(tradeOrQuote string) {
	var (
		bsFile []byte
		err    error
		// Trade or Quote
		title = strings.Title(tradeOrQuote)
		// q or t
		firstChar = string(tradeOrQuote[0])
		// 输出文件名(带相对路径)
		outFileName = path.Join("./win", fmt.Sprintf("ctp_%s.go", tradeOrQuote))
		// 函数主体
		funBody string
		// 在init函数中进行回调函数定义 h.MustFindProc("SetOnFrontConnected").Call(spi, syscall.NewCallback(OnConnect))
		cbSet string
		// 回调函数类型定义 type xxx func()
		cbTypeDefine string
		// 底层接口的响应类名 ThostFtdcMdSpi ThostFtdcTraderSpi
		cbName string
	)

	// 创建输出文件
	_, err = os.Stat(outFileName)
	if err == nil || os.IsExist(err) { // 文件存在=>删除
		_ = os.Remove(outFileName)
	}
	f, err := os.OpenFile(outFileName, os.O_RDWR|os.O_CREATE, os.ModePerm)
	checkErr(err)
	defer func() { _ = f.Close() }()

	// 读取文件
	if tradeOrQuote == "trade" {
		bsFile, err = ioutil.ReadFile(path.Join(srcPath, "ThostFtdcTraderApi.h"))
		cbName = "CThostFtdcTraderSpi"
	} else {
		bsFile, err = ioutil.ReadFile(path.Join(srcPath, "ThostFtdcMdApi.h"))
		cbName = "CThostFtdcMdSpi"
	}
	checkErr(err)
	// 汉字处理
	bsFile, _ = simplifiedchinese.GB18030.NewDecoder().Bytes(bsFile)
	/*
		///登录请求响应
		virtual void OnRspUserLogin(CThostFtdcRspUserLoginField *pRspUserLogin, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {};
	*/
	re := regexp.MustCompile(`\t///(.*)\r\n[^v]*virtual\s+(\w+)\s+(\w+)\(([^)]*)\)`) // 分解函数定义:注释,返回类型,函数名,参数字段四部分
	funs := re.FindAllStringSubmatch(string(bsFile), -1)
	for _, fun := range funs {
		funComment, _, funName, funFields := fun[1], fun[2], fun[3], fun[4]
		// 回调函数
		if strings.Index(funName, "On") == 0 {
			// type OnRspUserLoginType func(*CThostFtdcRspUserLoginField, *CThostFtdcRspInfoField, int, bool)
			cbTypeDefine += fmt.Sprintf("type %sOn%sType func(", firstChar, funName[2:]) // 回调函数类型定义
			if funFields != "" {                                                         // 参数可能为空
				re = regexp.MustCompile(`(\w+)\s+([*]?\w+)[,]?\s*`) //参数分解:类型,名称
				fields := re.FindAllStringSubmatch(funFields, -1)
				for _, field := range fields {
					fieldType := []rune(field[1])
					if unicode.IsUpper(fieldType[0]) { // 首字母大写:自定义类型
						cbTypeDefine += "*ctp." // 类型前加t不会被外部看到
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

			if funFields != "" { // 参数可能为空
				re = regexp.MustCompile(`(\w+)\s+([*]?)(\w+)([[]?)`) //参数分解:类型,名称
				fields := re.FindAllStringSubmatch(funFields, -1)
				for _, field := range fields {
					isPtrVar := len(field[2]) > 0 // 带*的字段传指针
					fieldType, fieldName, isArray := field[1], field[3], len(field[4]) > 0

					// 处理char *ppInstrumentID[] ==> []byte
					if isArray {
						funParams += fmt.Sprintf("%s [1][]byte, nCount int", fieldName)
						callParams += fmt.Sprintf(", uintptr(unsafe.Pointer(&%s)), uintptr(1)", fieldName)
						break
					} else if isPtrVar { // 带*的变量
						// string -> uintptr(unsafe.Pointer()) -> uintptr(unsafe.Pointer(bs))
						if fieldType == "char" { // char* => syscall.BytePtrFromString(%s)
							funParams += fmt.Sprintf("%s string, ", fieldName)
							funContent += fmt.Sprintf("\n\tbs, _ := syscall.BytePtrFromString(%s)", fieldName)
							callParams += ", uintptr(unsafe.Pointer(bs))"
						} else { // struct => uintptr(unsafe.Pointer(&%s))
							funParams += fmt.Sprintf("%s ctp.%s, ", fieldName, fieldType)
							callParams += fmt.Sprintf(", uintptr(unsafe.Pointer(&%s))", fieldName)
						}
					} else if fieldName == "nRequestID" {
						funContent += fmt.Sprintf("\n\t%s.nRequestID++", firstChar)
						callParams += fmt.Sprintf(", uintptr(%s.%s)", firstChar, fieldName)
					} else {
						funParams += fmt.Sprintf("%s ctp.%s, ", fieldName, fieldType)
						callParams += fmt.Sprintf(", uintptr(%s)", fieldName)
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
	_, _, _ = {{.firstChar}}.h.MustFindProc("{{.funName}}").Call({{.firstChar}}.api{{.callParams}})
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

	temp := `package win

import (
	ctp "{{.packageName}}/ctpdefine"
	"os"
	"path/filepath"
	"runtime"
	"syscall"
	"unsafe"
)

{{.cbTypeDefine}}
type {{.name}} struct {
	h        *syscall.DLL
	api, spi  uintptr
	nRequestID      int
}

func ({{.firstChar}} *{{.name}}) loadDll() {
	// 执行目录下创建 log目录
	_, err := os.Stat("log")
	if err != nil {
		os.Mkdir("log", os.ModePerm)
	}
	workPath, _ := os.Getwd()
	_, curFile, _, ok := runtime.Caller(1)
	if !ok {
		panic("取当前文件路径失败")
	}
	dllPath := filepath.Dir(curFile)
	_ = os.Chdir(dllPath)
	{{.firstChar}}.h = syscall.MustLoadDLL("ctp_{{.name}}.dll")
	
	// 还原到之前的工作目录
	os.Chdir(workPath)
	//defer h.Release() // 函数结束后会释放导致后续函数执行失败
}

// 接口
func new{{.title}}() *{{.name}} {
	{{.firstChar}} := new({{.name}})

	{{.firstChar}}.loadDll()
	{{.firstChar}}.api, _, _ = {{.firstChar}}.h.MustFindProc("CreateApi").Call()
	{{.firstChar}}.spi, _, _ = {{.firstChar}}.h.MustFindProc("CreateSpi").Call()
	_, _, _ = {{.firstChar}}.h.MustFindProc("RegisterSpi").Call({{.firstChar}}.api, uintptr(unsafe.Pointer({{.firstChar}}.spi)))
	return {{.firstChar}}
}

{{.cbSet}}
`
	// 变量替换
	data := map[string]string{
		"packageName":  packageName,
		"cbSet":        cbSet,
		"cbTypeDefine": cbTypeDefine,
		"name":         tradeOrQuote,
		"firstChar":    firstChar,
		"title":        title,
		"cbName":       cbName,
	}
	buf := templateMap(temp, data)
	_, _ = f.WriteString(buf)
	_, _ = f.WriteString("\n")
	_, _ = f.WriteString(funBody)
}

func generateStruct() {
	bsFile, err := ioutil.ReadFile(path.Join(srcPath, "ThostFtdcUserApiStruct.h"))
	checkErr(err)

	outFile := path.Join(outPath, "ctp_struct.go")
	_, err = os.Stat(outFile)
	if err == nil || os.IsExist(err) { // 文件存在
		checkErr(os.Remove(outFile))
	}

	f, err := os.OpenFile(outFile, os.O_RDWR|os.O_CREATE, os.ModePerm)
	checkErr(err)
	defer func() { _ = f.Close() }()

	_, err = f.WriteString("package ctpdefine\n\n")
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
		re = regexp.MustCompile(`///([^\r\n]*)\s*(\w+)\s+([^;]+);`) // 所有字段再分解成各个单独字段: 注释(可能含空格),类型,名称
		strStruc += re.ReplaceAllString(strc[3], "\t// $1\n\t$3 $2\n")
		//fields := re.FindAllStringSubmatch(strc[3], -1)
		//for _, field := range fields {
		//	strStruc += fmt.Sprintf("\t// %s\n\t%s %s\n", field[1], field[3], field[2])
		//}
		strStruc += "}\n\n"
		_, _ = f.WriteString(strStruc)
		//fmt.Println(strStruc)
	}
}

func generateDataType() {
	bsFile, err := ioutil.ReadFile(path.Join(srcPath, "ThostFtdcUserApiDataType.h"))
	checkErr(err)
	// 汉字处理
	bsFile, _ = simplifiedchinese.GB18030.NewDecoder().Bytes(bsFile)

	transType := make(map[string]string)
	// typedef char TThostFtdcTraderIDType[21]; ==> type TThostFtdcTraderIDType [21]byte
	transType[`/+.+是一个(.*)[\r]?\n/*[\r]?\ntypedef\s+char\s+(\w+)\[(\d+)\]\s*;`] = "// $1\ntype $2 [$3]byte\n"
	transType[`/+.+是一个(.*)[\r]?\n/*[\r]?\ntypedef\s+int\s+(\w+)\s*;`] = "// $1\ntype $2 int32\n" // int与go int32对应
	transType[`/+.+是一个(.*)[\r]?\n/*[\r]?\ntypedef\s+double\s+(\w+)\s*;`] = "// $1\ntype $2 float64\n"
	transType[`/+.+是一个(.*)[\r]?\n/*[\r]?\ntypedef\s+short\s+(\w+)\s*;`] = "// $1\ntype $2 int16\n"

	outFile := path.Join(outPath, "ctp_datatype.go")
	_, err = os.Stat(outFile)
	if err == nil || os.IsExist(err) { // 文件存在
		err = os.Remove(outFile)
		checkErr(err)
	}

	f, err := os.OpenFile(outFile, os.O_RDWR|os.O_CREATE, os.ModePerm)
	checkErr(err)
	defer func() { _ = f.Close() }()

	_, _ = f.WriteString(fmt.Sprintf(`package ctpfefine

type THOST_TE_RESUME_TYPE int32

const (
	THOST_TERT_RESTART THOST_TE_RESUME_TYPE = 0
	THOST_TERT_RESUME  THOST_TE_RESUME_TYPE = 1
	THOST_TERT_QUICK   THOST_TE_RESUME_TYPE = 2
)
`))
	txt := string(bsFile)
	for match, repl := range transType {
		re := regexp.MustCompile(match)
		for _, matchString := range re.FindAllString(txt, -1) {
			goType := re.ReplaceAllString(matchString, repl)
			_, _ = f.WriteString(goType)
		}
	}
	/*/////////////////////////////////////////////////////////////////////////
	///TFtdcBatchStatusType是一个处理状态类型
	/////////////////////////////////////////////////////////////////////////
	///未上传
	#define THOST_FTDC_BS_NoUpload '1'
	///已上传
	#define THOST_FTDC_BS_Uploaded '2'
	///审核失败
	#define THOST_FTDC_BS_Failed '3'

	typedef char TThostFtdcBatchStatusType;*/
	re := regexp.MustCompile(`/+.+是一个(.*)[\r]?\n/*[\r]?\n([^;]+)typedef\s+char\s+(\w+)\s*;`)
	for _, typeDef := range re.FindAllStringSubmatch(txt, -1) {
		if strings.Contains(typeDef[0], "#define") {
			comment, defines, name := typeDef[1], typeDef[2], typeDef[3]
			_, _ = f.WriteString(fmt.Sprintf("// %s\ntype %s byte\n", comment, name))
			// transType[] = "// $1\nconst $2 = '$3'\n"
			reSub := regexp.MustCompile(`/+(.*)[\r]?\n#define\s+(\w+)\s+'(.+)'`) // \w改为.因为有'#'的情况
			_, _ = f.WriteString(reSub.ReplaceAllString(defines, "// $1\nconst $2 = '$3'\n"))
		}
	}
	// 有定义没有具体的值
	_, _ = f.WriteString("// 紧急程度类型\ntype TThostFtdcNewsUrgencyType byte\n")
}
