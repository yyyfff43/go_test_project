package basic

import (
	"fmt"
	"reflect"
	"testing"
)

//
//  TestRangeEachUtf8Ch
//  @Description: 测试utf8编码的中文字符串遍历
//  @param t 测试结构体
//
func TestRangeEachUtf8Ch(t *testing.T) {
	var (
		str = "高度重视和善于总结历史经验、以史为鉴，是我们党的鲜明特点和历史经验。每到重要历史关头，我们党都会总结党的历史，从中吸取历史智慧，掌握历史主动。1945年4月党的六届七中全会通过了《关于若干历史问题的决议》，1981年党的十一届六中全会通过了《关于建国以来党的若干历史问题的决议》，都对一些重大事件和重要人物作出了评价，分清了是非，统一了全党思想，对推动党和人民事业发展产生了重大影响。"
	)

	RangeEachUtf8Ch(str)

}

//测试DealForumNum函数
func TestDealForumNum(t *testing.T) {
	res := DealForumNum(183386)
	fmt.Println(res)
}

//测试Split函数，因为slice不能比较直接，借助反射包中的方法比较
func TestSplit(t *testing.T) { // 测试函数名必须以Test开头，必须接收一个*testing.T类型参数
	got := Split("a:b:c", ":")         // 程序输出的结果
	want := []string{"a", "b", "c"}    // 期望的结果
	if !reflect.DeepEqual(want, got) { // 因为slice不能比较直接，借助反射包中的方法比较
		t.Errorf("expected:%v, got:%v", want, got) // 测试失败输出错误提示
	}
}
