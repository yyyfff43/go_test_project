package algorithm

import (
	"fmt"
)

//链表结点定义
type LNode struct {
	Data interface{}
	Next *LNode
}

//
//  Reverse
//  @Description: 带头结点的节点反转(插入法，比掉头法效率高，不保存前驱结点地址，比递归法效率高，因为不需要递归调用消耗时间复杂度)
//  @param node 链表结点
//
func Reverse(node *LNode) {
	if node == nil || node.Next == nil {
		return
	}
	var cur *LNode
	var next *LNode
	cur = node.Next.Next
	//设置链表第一个结点为尾结点
	node.Next.Next = nil
	//把遍历到的结点插入到头结点的后面
	for cur != nil {
		next = cur.Next
		cur.Next = node.Next
		node.Next = cur
		cur = next
	}
}

//
//  CreateNode
//  @Description: 创建带头结点引导的链表，即head->1->2->3......
//  @param node 结点结构体
//  @param max 最大节点数
//
func CreateNode(node *LNode, max int) {
	cur := node //定义头结点（注意不是那个1的结点，而是设定的head引导结点，但head结点也包括在最终节点链表里）
	for i := 1; i < max; i++ {
		cur.Next = &LNode{}
		cur.Next.Data = i
		cur = cur.Next
	}
}

//
//  CreateNode
//  @Description: 打印链表
//  @param string 提示信息
//  @param node 链表结构体
//
func PrintNode(info string, node *LNode) {
	fmt.Print(info)
	//因为头结点head不输出，所以cur直接赋值为头结点的下一个结点，即值为1那个结点
	for cur := node.Next; cur != nil; cur = cur.Next {
		fmt.Print(cur.Data, " ")
	}
	fmt.Println()
}
