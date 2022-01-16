/*
go语言工程项目测试入口main函数
*/
package main

import (
	"fmt"
	"go_test_project/src/basic"
	"go_test_project/src/err_learn"
	"go_test_project/src/protobuf"
	"os"
	"unicode/utf8"

	"github.com/golang/protobuf/proto"
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
	fmt.Printf("下标为1的元素值为：%c",strShow[1])

	//函数式变成测试，1累计加到10
    fLearn := basic.Adder()
    for i:=0; i<=10; i++{
			var s int
			s = fLearn(i)
			fmt.Printf("0 + 1 + ... + %d = %d\n",
				i, s)
	}

	fmt.Println("第二种实现方式：")
	fLearnAdder := basic.Adder2(0)
	for i := 0; i < 10; i++ {
		var s int
		s, fLearnAdder = fLearnAdder(i)
		fmt.Printf("0 + 1 + ... + %d = %d\n",
			i, s)
	}

	fmt.Println("")
	fmt.Println("测试go的err处理:")
	//defer,panic,recover的用法
	err_learn.TryRecover()

	//rpc proto练习
	fmt.Println("")
	fmt.Println("protobuf数据交互协议练习")
	msgTest := &protobuf.Person{
		Name: proto.String("张三"),
		Age:  proto.Int(18),
		From: proto.String("China"),
	}

	//序列化
	msgDataEncoding, err := proto.Marshal(msgTest)
	if err != nil {
		panic(err.Error())
		return
	}

	fmt.Printf("打包后的二进制数据：%b",msgDataEncoding)

	//反序列化：
	msgEntity := protobuf.Person{}
	err = proto.Unmarshal(msgDataEncoding, &msgEntity)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
		return
	}
	fmt.Println("")
	fmt.Printf("姓名：%s\n\n", msgEntity.GetName())
	fmt.Printf("年龄：%d\n\n", msgEntity.GetAge())
	fmt.Printf("国籍：%s\n\n", msgEntity.GetFrom())
}
