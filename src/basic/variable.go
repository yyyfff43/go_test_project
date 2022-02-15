// Copyright 2021 The Go Authors YangFan. All rights reserved.

/*
go语言工程项目测试，包含一些基本的变量方法调用和处理
*/
package basic

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

// RangeStringByByte 方法将输入的字符串已字节形式遍历
// 将传入的字符串打印，无返回值
func RangeStringByByte(str string) {
	for _, b := range []byte(str) {
		fmt.Printf("%X", b)
	}
	fmt.Println()
}

// RangeStringByByte 方法将输入的字符串已字节形式遍历
// 将传入的字符串打印，无返回值
func RangeStringByRune(str string) {
	for i, ch := range str { // ch is a rune
		fmt.Printf("(%d %X) ", i, ch)
	}
	fmt.Println()
}

//
//  RangeEachUtf8Ch
//  @Description: utf8编码的中文字符串遍历
//  @param str 中文字符串
//
func RangeEachUtf8Ch(str string) {
	bytes := []byte(str)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("*%c* ", ch)
	}
	fmt.Println()
}

//
//  DealForumNum
//  @Description://圈子的关注数和帖子数，超过9999的以单位为万，四舍五入保留1位小数展示，例10199,显示为1.0万
//  @param num
//  @return string
//
func DealForumNum(num int64) string {
	resStr := ""
	if num > 10000 {
		var tmpStr = strconv.FormatInt(num, 10)
		var tmpByte = []byte(tmpStr)
		var leftNumByte = tmpByte[len(tmpByte)-4 : len(tmpByte)] //取最后四位数
		var leftNumStr = "0." + string(leftNumByte)
		var leftFloat, _ = strconv.ParseFloat(leftNumStr, 64)
		var headNum = string(tmpByte[:len(tmpByte)-4]) //取最后四位以外的前边数字
		var headNumFloat, _ = strconv.ParseFloat(headNum, 64)
		var newFloat = headNumFloat + leftFloat            //加上变为新的浮点数
		resStr = strconv.FormatFloat(newFloat, 'f', 1, 64) //四舍五入，保留1位小数
		resStr = resStr + "万"
	}
	return resStr
}
