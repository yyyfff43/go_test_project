package err_test

import (
	"fmt"
)

/*
defer 在调用函数退出时执行，栈结构，先入后出，最后的defer先执行，以此类推
在程序中尽量少出现panic，panic触发会按照堆栈的先入后出的原则调用之前的defer，如果defer中有recover函数，则执行
recover函数中的错误处理，如果recover没有处理这个错误，则再次从recover里抛出panic
 */

func TryRecover() {
	defer func() {//使用匿名函数调用，defer之后就执行这个函数
		r := recover()
		if r == nil {
			fmt.Println("没有错误，请尝试取消下边的注释方法来测试 ")
			return
		}
		if err, ok := r.(error); ok {
			fmt.Println("Error occurred:", err)
		} else {
			panic(fmt.Sprintf(
				"I don't know what to do: %v", r))
		}
	}()//（）表示匿名函数定义完后立即执行

	// Uncomment each block to see different panic
	// scenarios.
	// Normal error
	//panic(errors.New("this is an error"))

	// Division by zero
	//b := 0
	//a := 5 / b
	//fmt.Println(a)

	// Causes re-panic
	//panic(123)
}
