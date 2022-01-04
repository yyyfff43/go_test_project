/*
go语言工程项目测试入口main函数
*/
package main

import (
	"fmt"
	"go_test_project/src/basic"
	"unicode/utf8"
)


func main() {
	fmt.Println("项目文件入口")

	var str string = "First of all，这是一个关于go语言的项目工程的测试，I try to test it"
	fmt.Println("byte格式遍历每一个字符")
	basic.RangeStringByByte(str)
    fmt.Println("rune格式遍历每一个字符")
	basic.RangeStringByRune(str)

	//统计字符串包括中文在内一共多少个字，每一个英文或者特殊符号都算
	fmt.Println("Rune count:",utf8.RuneCountInString(str))

	fmt.Println("遍历UTF8字符串中的每一个字符")
	basic.RangeEachUtf8Ch(str)

	for i, ch := range []rune(str) {
		fmt.Printf("(%d %c) ", i, ch)
	}
	fmt.Println()
	var strShow []rune = []rune(str)
	fmt.Printf("%c",strShow[1])

}
