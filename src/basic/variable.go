// Copyright 2021 The Go Authors Yangfan. All rights reserved.

/*
go语言工程项目测试，包含一些基本的变量方法调用和处理
*/
package basic

import (
	"fmt"
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
