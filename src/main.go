package main

import (
	"fmt"
	"go_test_project/src/basic"
)

func main() {
	fmt.Println("项目文件入口")

	var str string = "First of all，这是一个关于go语言的项目工程的测试，I try to test it"
	basic.RangeStringByByte(str)
}
