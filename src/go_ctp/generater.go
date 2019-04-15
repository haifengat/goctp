package go_ctp

import (
	"golang.org/x/text/encoding/simplifiedchinese"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

func Run() {
	bsFile, err := ioutil.ReadFile("src/ctp_20180109_x64/ThostFtdcUserApiDataType.h")
	checkErr(err)
	// 汉字处理
	bsFile, _ = simplifiedchinese.GB18030.NewDecoder().Bytes(bsFile)
	// 分行
	lines := strings.Split(string(bsFile), "\n")

	var dataType []string
	var constDef []string
	transType := make(map[string]string)
	transType[`typedef\s*char\s*(\w+)\[(\d+)\].*\s*$`] = "type $1 [$2]byte"       // typedef char TThostFtdcTraderIDType[21]; ==> type TThostFtdcTraderIDType [21]byte
	transType[`typedef\s*char\s*(\w+);.*\s*$`] = "type $1 byte"                   //typedef char TThostFtdcIdCardTypeType; ==> type TThostFtdcIdCardTypeType byte
	transType[`typedef\s*int\s*(\w+);.*\s*$`] = "type $1 int"
	transType[`typedef\s*double\s*(\w+);.*\s*$`] = "type $1 float64"
	transType[`typedef\s*short\s*(\w+);.*\s*$`] = "type $1 int16"
	//transType[`typedef\s*\b(int|double|short)\b\s*(\w+);.*\s*$`] = "type $2 $1" // typedef int TThostFtdcIPPortType; ==> type TThostFtdcIPPortType int / double /short
	transType[`#define\s*(\w+)\s*'(\w+)'\s*$`] = "const |enumtype|_$1 = '$2'"     //#define THOST_FTDC_ICT_AccountsPermits 'J' ==> const THOST_FTDC_ICT_AccountsPermits = 'J'

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
	f, err := os.OpenFile("src/go_ctp/datatype.go", os.O_RDWR|os.O_CREATE, os.ModePerm)
	checkErr(err)
	defer f.Close()
	f.WriteString("package go_ctp\n")
	for _, str := range dataType {
		//fmt.Println(str)
		// 写文件
		f.WriteString(str + "\n")
	}
}
