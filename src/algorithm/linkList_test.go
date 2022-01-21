package algorithm

import (
	"testing"
)


//测试链表定义
func TestCreateNode(t *testing.T) {
	node := LNode{}
	CreateNode(&node,10)
	PrintNode("遍历链表：",&node)
}

//测试打印链表
func TestPrintNode(t *testing.T) {
	node := LNode{}
	CreateNode(&node,20)
	PrintNode("遍历链表：",&node)
}
