/*
* @File : slice_func
* @Describe :
* @Author: yangfan@zongheng.com
* @Date : 2022/3/30 11:27
* @Software: GoLand
 */

package slice_learn

import (
	"reflect"
	"sort"
)

//
//  RemoveDuplicatesInt64
//  @Description: 排序并去掉切片中重复的元素，类型是int64的数组
//  @param arr
//  @return []int64
//
func RemoveDuplicatesInt64(in []int64) []int64 {
	sort.Slice(in, func(i, j int) bool { return in[i] < in[j] })
	j := 0
	for i := 1; i < len(in); i++ {
		if in[j] == in[i] {
			continue
		}
		j++
		// 需要保存原始数据时
		// in[i], in[j] = in[j], in[i]
		// 只需要保存需要的数据时
		in[j] = in[i]
	}
	result := in[:j+1]
	return result
}

//
//  InArray
//  @Description: 查找字符是否在数组中
//  @param obj
//  @param target
//  @return bool
//
func InArray(obj interface{}, target interface{}) bool {
	targetValue := reflect.ValueOf(target)
	switch reflect.TypeOf(target).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			if targetValue.Index(i).Interface() == obj {
				return true
			}
		}
	case reflect.Map:
		if targetValue.MapIndex(reflect.ValueOf(obj)).IsValid() {
			return true
		}
	}

	return false
}

//
//  RemoveByAnotherSlice
//  @Description: 去掉目标数组中包含的相同元素
//  @param targetArr
//  @param conditionArr
//  @return []int64
//
func RemoveByAnotherSlice(targetArr, conditionArr []int64) []int64 {
	// fmt.Println(conditionArr)
	// fmt.Println(targetArr)
	var newResArr = make([]int64, 0)
	for _, v := range targetArr {
		if InArray(v, conditionArr) {
			//TODO 如果包含了相同元素，这里什么也不做

		} else {
			newResArr = append(newResArr, v)
		}
	}
	return newResArr
}

//
//  FilterElements
//  @Description: 过滤掉给定条件切片中的元素后，整理一个保留元素的切片返回
//  @param targetArr
//  @param conditionArr
//  @return []int64
//
func FilterElements(targetArr, conditionArr []int) []int {
	//复制一份目标数组防止函数外改变此切片的值，如果不执行这一项一路传递targetArr，函数外此执行完后targetArr会改变内布的元素，即下边[:n]会跑到切片前边去
	tmpSlice := make([]int, len(targetArr))
	copy(tmpSlice, targetArr)

	n := 0
	for _, x := range tmpSlice {
		if InArray(x, conditionArr) {
			tmpSlice[n] = x // 保留该元素
			n++
		}
	}
	tmpSlice = tmpSlice[:n] // 截取切片中需保留的元素
	return tmpSlice

	//这里假设过滤的条件已封装为keep函数，使用for range遍历切片a的所有元素逐一调用keep函数进行过滤。
	//n := 0
	//for _, x := range a {
	//	if keep(x) {
	//		a[n] = x  // 保留该元素
	//		n++
	//	}
	//}
	//a = a[:n]  // 截取切片中需保留的元素
}

//
//  DelElement
//  @Description: 删除数组中的指定元素
//  @param obj
//  @param target
//  @return []int 删除后的数组
//
func DelElement(obj int, target []int) []int {
	for i, tLen := 0, len(target); i < tLen; i++ {
		if obj == target[i] {
			//append法移动下标实现
			//			target = append(target[:i], target[i+1:]...)
			//copy法移动下标实现
			//			target = target[:i+copy(target[i:], target[i+1:])]

			/*
				需要特别注意的是。如果切片a中的元素是一个指针类型或包含指针字段的结构体类型（需要被垃圾回收），
				上面剪切和删除的示例代码会存在一个潜在的内存泄漏问题：一些具有值的元素仍被切片a引用，因此无法被垃圾回收机制回收掉。
				下面的代码可以解决这个问题。
			*/
			copy(target[i:], target[i+1:])
			target[len(target)-1] = 0 // 或类型T的零值
			target = target[:len(target)-1]
			break
		}
	}
	return target
}

//
//  CutElements
//  @Description: 按照给定的下标区间剪切数组(下标元素左包含右不包含)
//  @param from
//  @param to
//  @param target
//  @return []int
//
func CutElements(from, to int, target []int) []int {
	//一般方法，但如果内部元素是指针或者含有指针的结构体时，剪切过的老数组如果再调用被删除的数据会发生内存泄漏
	//	target = append(target[:from], target[to:]...)

	//放内存泄漏，剪切掉的元素补nil或者本元素0值
	copy(target[from:], target[to:])
	for k, n := len(target)-to+from, len(target); k < n; k++ {
		target[k] = 0 // nil或类型T的零值
	}
	target = target[:len(target)-to+from]
	return target
}
