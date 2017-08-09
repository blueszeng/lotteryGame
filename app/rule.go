package app

import (
	"errors"
	"fmt"
	"strconv"
)

func init() {

}

func sizeScensResult(group string) (value string, error error) {

	number1, error := strconv.Atoi(group[len(group)-1:])
	if error != nil {
		return "", errors.New("字符串转换成整数失败")
	}
	number2, error := strconv.Atoi(group[len(group)-2 : len(group)-1])
	if error != nil {
		return "", errors.New("字符串转换成整数失败")
	}
	fmt.Println(number1, number2)
	// 是不是合
	if number1 == number2 {
		return "合", nil
	}

	// 是不是大
	if number1 >= 5 && number1 <= 9 {
		return "大", nil
	}
	if number1 >= 0 && number1 <= 4 {
		return "小", nil
	}
	return
}
