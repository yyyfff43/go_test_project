// Copyright 2021 The Go Authors. All rights reserved.

/*
go语言工程项目测试，包含一些基本的变量方法调用和处理
*/
package basic

import "fmt"

// RangeStringByByte 方法将输入的字符串已字节形式遍历
// 将传入的字符串打印，无返回值
func RangeStringByByte(str string) {
	for _, b := range []byte(str) {
		fmt.Printf("%X", b)
	}

}
