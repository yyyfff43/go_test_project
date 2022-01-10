package basic

import "testing"

//
//  TestRangeEachUtf8Ch
//  @Description: 测试utf8编码的中文字符串遍历
//  @param t 测试结构体
//
func TestRangeEachUtf8Ch(t *testing.T) {
	var (
		str string = "高度重视和善于总结历史经验、以史为鉴，是我们党的鲜明特点和历史经验。每到重要历史关头，我们党都会总结党的历史，从中吸取历史智慧，掌握历史主动。1945年4月党的六届七中全会通过了《关于若干历史问题的决议》，1981年党的十一届六中全会通过了《关于建国以来党的若干历史问题的决议》，都对一些重大事件和重要人物作出了评价，分清了是非，统一了全党思想，对推动党和人民事业发展产生了重大影响。"
	)

	RangeEachUtf8Ch(str)

}
