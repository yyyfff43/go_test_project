/*
* @File : str_to_pinyin
* @Describe :字符串转拼音测试，使用github.com/mozillazg/go-pinyin包，范例：https://codingdict.com/os/software/44669
* @Author: Jerry Yang
* @Date : 2022/2/13 23:22
* @Software: GoLand
 */

package pinyin

import (
	"fmt"
	"github.com/mozillazg/go-pinyin"
	"reflect"
)

func ChStr2Pinyin() {
	hans := "1解放思想and实事求是do"
	a := pinyin.NewArgs()
	// 默认输出 [[zhong] [guo] [ren]]
	fmt.Println(pinyin.Pinyin(hans, a))

	// 包含声调 [[zhōng] [guó] [rén]]
	a.Style = pinyin.Tone
	fmt.Println(pinyin.Pinyin(hans, a))

	// 声调用数字表示 [[zho1ng] [guo2] [re2n]]
	a.Style = pinyin.Tone2
	fmt.Println(pinyin.Pinyin(hans, a))

	// 开启多音字模式 [[zhong zhong] [guo] [ren]]
	a = pinyin.NewArgs()
	a.Heteronym = true
	fmt.Println(pinyin.Pinyin(hans, a))
	// [[zho1ng zho4ng] [guo2] [re2n]]
	a.Style = pinyin.Tone2
	fmt.Println(pinyin.Pinyin(hans, a))

}

//
//  converterToFirstSpell
//  @Description: 将输入的带多音字的中文词语转化成拼音首字母字符串，有多少个多音字输出多少个组合，
//  例：输入 雪中悍刀行 输出 xzhdh,xzhdx
//  @param chines
//  @return string
//
func ConverterToFirstSpell(chines string) string {

	// 开启多音字模式，例：中国人 转 [[zhong zhong] [guo] [ren]]
	//py := pinyin.NewArgs()
	//py.Heteronym = true
	//py.Separator = ""
	//py.Style = pinyin.FirstLetter
	//pySlice := pinyin.Pinyin(chines, py)
	//cpPySlice := CartesianProductSlice(pySlice)
	//sort.Strings(cpPySlice)
	//cpPyDuplicate := SortDuplicate(cpPySlice)
	//newValue := cpPyDuplicate.([]interface{})
	//cpFinal := newValue.([]string)
	//return strings.Join(cpFinal, ",")
	return ""
}

func getFirstWordPinyin(str string) []string {

	return nil
}

func getPinyin(words string) []string {
	if words != "" {

	}
	return nil
}

//
//  CartesianProductSlice
//  @Description: 将传入的二维数组做笛卡尔积运算，返回一个组合的二维数组
//  @param allStr
//  @return res
//
func CartesianProductSlice(allStr [][]string) (res []string) {
	var baseSlices [][]string
	for i := 0; i < len(allStr); i++ {
		baseSlices = append(baseSlices, allStr[i])
	}
	res = baseSlices[0]

	for _, v := range baseSlices[1:] {
		res = makeData(res, v)
	}

	return
}

//
//  MakeData
//  @Description: 把新设定的切片连接到老切片上
//  @param base
//  @param makeData
//  @return str
//
func makeData(base []string, makeData []string) (str []string) {
	for _, v := range base {
		for _, makeDataValue := range makeData {
			str = append(str, v+makeDataValue)
		}
	}
	return
}

//
//  Duplicate
//  @Description: 将带有重复元素的字符串数组去重返回
//  @param strSlice
//  @return []string
//
//func Duplicate(strSlice []string) []string {
//	sort.Strings(strSlice)
//	SortDuplicate(strSlice)
//	fmt.Println(strSlice)
//	return strSlice
//}

//
//  SortDuplicate
//  @Description: 将sort排序后的数组元素去重
//  @param a
//  @return ret
//
func SortDuplicate(a interface{}) (ret []interface{}) {
	va := reflect.ValueOf(a)
	for i := 0; i < va.Len(); i++ {
		if i > 0 && reflect.DeepEqual(va.Index(i-1).Interface(), va.Index(i).Interface()) {
			continue
		}
		ret = append(ret, va.Index(i).Interface())
	}
	return ret
}
