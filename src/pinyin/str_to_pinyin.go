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
	"sort"

	"strings"
)

func ChStr2Pinyin(){
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
func ConverterToFirstSpell(chines string) []string {
	// 开启多音字模式，例：中国人 转 [[zhong zhong] [guo] [ren]]
	py := pinyin.NewArgs()
	py.Heteronym = true
	py.Separator = ""
	py.Style = pinyin.FirstLetter
	pySlice := pinyin.Pinyin(chines, py)
	if len(pySlice) > 0 {
		cpPySlice := CartesianProductSlice(pySlice)
		sort.Strings(cpPySlice)
		cpPyDuplicate := SortDuplicate(cpPySlice)
		return cpPyDuplicate
	}
	return nil
}

//
//  ConverterToSpell
//  @Description: 将输入的带多音字的中文词语转化成拼音字符串，有多少个多音字输出多少个组合
//  例：输入 雪中悍刀行 输出 xuezhonghandaohang,xuezhonghandaoheng,xuezhonghandaoxing
//  @param chines
//  @return string
//
func ConverterToSpell(chines string) []string {
	// 开启多音字模式，例：中国人 转 [[zhong zhong] [guo] [ren]]
	py := pinyin.NewArgs()
	py.Heteronym = true
	py.Separator = ""
	pySlice := pinyin.Pinyin(chines, py)
	if len(pySlice) > 0 {
		cpPySlice := CartesianProductSlice(pySlice)
		sort.Strings(cpPySlice)
		cpPyDuplicate := SortDuplicate(cpPySlice)
		return cpPyDuplicate
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
//  SortDuplicate
//  @Description: 将sort排序后的数组元素去重
//  @param a
//  @return ret
//
func SortDuplicate(s []string) (ret []string) {
	tmpM := make(map[string]string) // key的类型要和切片中的数据类型一致
	for _, v := range s {
		tmpM[v] = "1"
	}
	// 先清空s
	s = []string{}
	for i, _ := range tmpM {
		s = append(s, i)
	}
	return s
}

//
//  GetPinyin
//  @Description: 获取汉字词语的拼音字符串，返回过滤英文字符后的首字母，包含多音字的全词拼音，小写的关键字本身等
//  例：x麻辣的粉E 返回 [mldf maladifen maladefen x麻辣的粉E x麻辣的粉e]
//  @param words
//  @return []string
//
func GetPinyin(words string) []string {
	if words!="" {
		firstPinyin := ConverterToFirstSpell(words)
		fullPinyin := ConverterToSpell(words)
		wordsLowC := strings.ToLower(words)
		res := append(firstPinyin, fullPinyin...)
		res = append(res, words, wordsLowC)
		return res
	}
	return nil
}


