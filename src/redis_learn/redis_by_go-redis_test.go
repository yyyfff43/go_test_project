/*
* @File : Jerry Yang
* @Describe :
* @Author: 248037973@qq.com
* @Date : 2022/3/20 8:45
* @Software: GoLand
 */
package redis_learn

import "testing"

//测试初始化范例
func TestV8Example(t *testing.T) {
	V8Example()
}

//测试字符串操作
func TestStringOperation(t *testing.T) {
	StringOperation()
}

//测试列表
func TestListOperation(t *testing.T) {
	ListOperation()
}

//测试去重集合
func TestSetOperation(t *testing.T) {
	SetOperation()
}

//测试hash
func TestHashOperation(t *testing.T) {
	HashOperation()
}
