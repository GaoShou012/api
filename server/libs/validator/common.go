package libs_validator

import (
	"errors"
	"fmt"
	"regexp"
)

/*
	6-32位
	数字和英文字母组成组成
*/
func Username(str string) error {
	strLen := len(str)
	if strLen < 6 || strLen > 32 {
		return errors.New("用户名长度错误，必须是6-32位")
	}

	exp := "^[a-zA-Z0-9]*$"

	regCom, err := regexp.Compile(exp)
	if err != nil {
		info := fmt.Sprintf("regexp=%v,err=%v\n", exp, err)
		return errors.New(info)
	}

	// 校验str
	ok := regCom.MatchString(str)
	if !ok {
		return errors.New("账号格式错误，必须是数字和英文字母组成")
	}

	return nil
}

/*
	6-32位
	数字和英文字母组成组成
*/
func Password(str string) error {
	strLen := len(str)
	if strLen < 6 || strLen > 32 {
		return errors.New("密码长度错误，必须是6-32位")
	}

	exp := "^[a-zA-Z0-9]*$"

	regCom, err := regexp.Compile(exp)
	if err != nil {
		info := fmt.Sprintf("regexp=%v,err=%v\n", exp, err)
		return errors.New(info)
	}

	// 校验str
	ok := regCom.MatchString(str)
	if !ok {
		return errors.New("密码格式错误，必须是数字和英文字母组成")
	}

	return nil
}

/*
	校验是否中国手机号码
*/
func Phone(str string) error {
	strLen := len(str)
	if strLen != 11 {
		return errors.New("手机号吗长度错误，必须是11位")
	}

	exp := "^1[0-9]*$"

	regCom, err := regexp.Compile(exp)
	if err != nil {
		info := fmt.Sprintf("regexp=%v,err=%v\n", exp, err)
		return errors.New(info)
	}

	// 校验str
	ok := regCom.MatchString(str)
	if !ok {
		return errors.New("手机号吗格式错误，必须是数字组成")
	}

	return nil
}

/*
	校验邮箱
*/
func Email(str string) (bool, error) {
	exp := "^[A-Za-z\\d]+([-_.][A-Za-z\\d]+)*@([A-Za-z\\d]+[-.])+[A-Za-z\\d]{2,4}$"
	return regexp.MatchString(exp, str)
}

/*
	判断是否纯数字，字符串
*/
func IsDigitalString(str string, minLen int, maxLen int) bool {
	if checkLength(str, minLen, maxLen) != true {
		return false
	}
	for _, r := range str {
		if '0' <= r && r <= '9' == false {
			return false
		}
	}
	return true
}

/*
	判断是否纯英文字母，字符串
*/
func IsLetterString(str string, minLen int, maxLen int) bool {
	if checkLength(str, minLen, maxLen) != true {
		return false
	}

	exp := "^[a-zA-Z]*$"
	regCom, err := regexp.Compile(exp)
	if err != nil {
		info := fmt.Sprintf("regexp=%v,err=%v\n", exp, err)
		panic(info)
	}
	return regCom.MatchString(str)
}
