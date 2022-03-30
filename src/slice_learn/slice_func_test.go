/*
* @File : slice_func_test
* @Describe :
* @Author: yangfan@zongheng.com
* @Date : 2022/3/30 11:28
* @Software: GoLand
 */

package slice_learn

import (
	"fmt"
	"testing"
)

//测试删除切片中的元素
func TestDelElement(t *testing.T) {
	var testArr = []int{1, 3, 5, 7, 9}
	res := DelElement(3, testArr)
	fmt.Println(res)
	for k, v := range res {
		fmt.Printf("k:%v zhi:%v\n", k, v)
	}
	fmt.Println(testArr) //[1 5 7 9 0]这里将老数组删除掉后在末尾的位置坐了补0值处理，防止程序再次调用这个已删除元素时造成内存泄露
}

//测试剪切切片中一段元素
func TestCutElements(t *testing.T) {
	var testArr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	res := CutElements(2, 5, testArr)
	fmt.Println(res)
	for k, v := range res {
		fmt.Printf("k:%v zhi:%v\n", k, v)
	}
	fmt.Println(testArr)
}

//过滤出条件指定的切片元素
func TestFilterElements(t *testing.T) {
	var testArr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var testConArr = []int{5, 6, 7}
	res := FilterElements(testArr, testConArr)
	fmt.Println(res)
	fmt.Println(testArr)
	fmt.Println(testConArr)
}
