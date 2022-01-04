// Copyright 2021 The Go Authors Yangfan. All rights reserved.

/*
go语言工程项目测试，函数式变成学习
*/
package basic

//一般函数式变成实现整数值的累加
func Adder() func(int) int {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}

type iAdder func(int) (int, iAdder)

func Adder2(base int) iAdder {
	return func(v int) (int, iAdder) {
		return base + v, Adder2(base + v)
	}
}