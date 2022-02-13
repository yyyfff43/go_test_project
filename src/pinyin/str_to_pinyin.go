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
)

func ChStr2Pinyin(){
	hans := "中国人"
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
//	a = NewArgs()
//	a.Heteronym = true
//	fmt.Println(pinyin.Pinyin(hans, a))
//	// [[zho1ng zho4ng] [guo2] [re2n]]
//	a.Style = pinyin.Tone2
//	fmt.Println(pinyin.Pinyin(hans, a))
}


