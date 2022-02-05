/*
* @File : learnDefer_test
* @Describe : defer单元测试文件
* @Author: Jerry Yang
* @Date : 2022/2/6 0:07
* @Software: GoLand
 */

package basic

import (
	"fmt"
	"testing"
)

func TestAll(t *testing.T) {
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())

	fmt.Println("defer注册要延迟执行的函数时该函数所有的参数都需要确定其值:")
	x := 1
	y := 2
	defer calc("AA", x, calc("A", x, y))
	x = 10
	defer calc("BB", x, calc("B", x, y))
	y = 20
}
