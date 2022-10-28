package util

import (
	"regexp"
	"strconv"
	"strings"
)

var Reg = newReg()

type reg struct {
}

func newReg() *reg {
	return &reg{}
}

// RegexpReplace 正则表达式
func (ur *reg) RegexpReplace(reg, src, temp string) string {
	var result []byte
	pattern := regexp.MustCompile(reg)
	for _, subMatch := range pattern.FindAllStringSubmatchIndex(src, -1) {
		result = pattern.ExpandString(result, temp, src, subMatch)
	}
	return string(result)
}

// RegexpPhoneFormat 手机验证
// Phone format validation.
// 1. China Mobile:
//    134, 135, 136, 137, 138, 139, 150, 151, 152, 157, 158, 159, 182, 183, 184, 187, 188,
//    178(4G), 147(Net)；
//    172
//
// 2. China Unicom:
//    130, 131, 132, 155, 156, 185, 186 ,176(4G), 145(Net), 175
//
// 3. China Telecom:
//    133, 153, 180, 181, 189, 177(4G)
//
// 4. Satelite:
//    1349
//
// 5. Virtual:
//    170, 173
//
// 6. 2018:
//    16x, 19x
func (ur *reg) RegexpPhoneFormat(phone string) bool {
	regular := `^13[\d]{9}$|^14[5,7]{1}\d{8}$|^15[^4]{1}\d{8}$|^16[\d]{9}$|^17[0,2,3,5,6,7,8]{1}\d{8}$|^18[\d]{9}$|^19[\d]{9}$`

	reg := regexp.MustCompile(regular)
	return reg.MatchString(phone)
}

// RegexpPhoneLooseFormat
// Loose mobile phone number verification(宽松的手机号验证)
// As long as the 11 digit numbers beginning with
// 13, 14, 15, 16, 17, 18, 19 can pass the verification (只要满足 13、14、15、16、17、18、19开头的11位数字都可以通过验证)
func (ur *reg) RegexpPhoneLooseFormat(phone string) bool {
	regular := `^1(3|4|5|6|7|8|9)\d{9}$`

	reg := regexp.MustCompile(regular)
	return reg.MatchString(phone)
}

// RegexpTelephoneFormat 固话校验
func (ur *reg) RegexpTelephoneFormat(telephone string) bool {
	regular := `^((\d{3,4})|\d{3,4}-)?\d{7,8}$`

	reg := regexp.MustCompile(regular)
	return reg.MatchString(telephone)
}

// RegexpPostalCodeFormat 邮政编码
func (ur *reg) RegexpPostalCodeFormat(postalCode string) bool {
	regular := `^\d{6}$`

	reg := regexp.MustCompile(regular)
	return reg.MatchString(postalCode)
}

// 身份证ID校验
//
// xxxxxx yyyy MM dd 375 0  十八位
// xxxxxx   yy MM dd  75 0  十五位
//
// 地区：     [1-9]\d{5}
// 年的前两位：(18|19|([23]\d))  1800-2399
// 年的后两位：\d{2}
// 月份：     ((0[1-9])|(10|11|12))
// 天数：     (([0-2][1-9])|10|20|30|31) 闰年不能禁止29+
//
// 三位顺序码：\d{3}
// 两位顺序码：\d{2}
// 校验码：   [0-9Xx]
//
// 十八位：^[1-9]\d{5}(18|19|([23]\d))\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\d{3}[0-9Xx]$
// 十五位：^[1-9]\d{5}\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\d{3}$
//
// 总：
// (^[1-9]\d{5}(18|19|([23]\d))\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\d{3}[0-9Xx]$)|(^[1-9]\d{5}\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\d{3}$)
func (ur *reg) checkResidentId(id string) bool {
	id = strings.ToUpper(strings.TrimSpace(id))
	if len(id) != 18 {
		return false
	}
	var (
		weightFactor = []int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
		checkCode    = []byte{'1', '0', 'X', '9', '8', '7', '6', '5', '4', '3', '2'}
		last         = id[17]
		num          = 0
	)
	for i := 0; i < 17; i++ {
		tmp, err := strconv.Atoi(string(id[i]))
		if err != nil {
			return false
		}
		num = num + tmp*weightFactor[i]
	}
	if checkCode[num%11] != last {
		return false
	}
	regular := `(^[1-9]\d{5}(18|19|([23]\d))\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\d{3}[0-9Xx]$)|(^[1-9]\d{5}\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\d{3}$)`
	reg := regexp.MustCompile(regular)
	return reg.MatchString(id)
}

// RegexpQQFormat 腾讯qq校验
func (ur *reg) RegexpQQFormat(qq string) bool {
	regular := `^[1-9][0-9]{4,}$`

	reg := regexp.MustCompile(regular)
	return reg.MatchString(qq)
}

// RegexpPassportFormat 校验护照
// Universal passport format rule:
// Starting with letter, containing only numbers or underscores, length between 6 and 18.
func (ur *reg) RegexpPassportFormat(qq string) bool {
	regular := `^[a-zA-Z]{1}\w{5,17}$`

	reg := regexp.MustCompile(regular)
	return reg.MatchString(qq)
}

// RegexpPassWordLowFormat 弱密码(任意可见字符，长度在6~18之间)
func (ur *reg) RegexpPassWordLowFormat(p string) bool {
	regular := `^[\w\S]{6,18}$`
	reg := regexp.MustCompile(regular)
	return reg.MatchString(p)
}

// RegexpPassWordMidFormat 中等强度密码(在弱密码的基础上，必须包含大小写字母和数字)
func (ur *reg) RegexpPassWordMidFormat(p string) bool {
	regular1 := `^[a-zA-Z]{1}\w{5,17}$`
	regular2 := `[a-z]+`
	regular3 := `[A-Z]+`
	regular4 := `\d+`
	return regexp.MustCompile(regular1).MatchString(p) && regexp.MustCompile(regular2).MatchString(p) && regexp.MustCompile(regular3).MatchString(p) && regexp.MustCompile(regular4).MatchString(p)
}

// RegexpPassWordHighFormat 强等强度密码(在弱密码的基础上，必须包含大小写字母、数字和特殊字符)
func (ur *reg) RegexpPassWordHighFormat(p string) bool {
	regular1 := `^[\w\S]{6,18}$`
	regular2 := `[a-z]+`
	regular3 := `[A-Z]+`
	regular4 := `\d+`
	regular5 := `[^a-zA-Z0-9]+`
	return regexp.MustCompile(regular1).MatchString(p) && regexp.MustCompile(regular2).MatchString(p) && regexp.MustCompile(regular3).MatchString(p) && regexp.MustCompile(regular4).MatchString(p) && regexp.MustCompile(regular5).MatchString(p)
}

func (ur *reg) RegexpEmailFormat(p string) bool {
	regular1 := `^[a-zA-Z0-9_\-\.]+@[a-zA-Z0-9_\-]+(\.[a-zA-Z0-9_\-]+)+$`
	return regexp.MustCompile(regular1).MatchString(p)
}

func (ur *reg) RegexpUrlFormat(p string) bool {
	regular1 := `(https?|ftp|file)://[-A-Za-z0-9+&@#/%?=~_|!:,.;]+[-A-Za-z0-9+&@#/%=~_|]`
	return regexp.MustCompile(regular1).MatchString(p)
}

func (ur *reg) RegexpDomainFormat(p string) bool {
	regular1 := `^([0-9a-zA-Z][0-9a-zA-Z\-]{0,62}\.)+([a-zA-Z]{0,62})$`
	return regexp.MustCompile(regular1).MatchString(p)
}

func (ur *reg) RegexpMacFormat(p string) bool {
	regular1 := `^([0-9A-Fa-f]{2}[\-:]){5}[0-9A-Fa-f]{2}$`
	return regexp.MustCompile(regular1).MatchString(p)
}
