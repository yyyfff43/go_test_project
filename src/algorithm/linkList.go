package algorithm

import (
	"fmt"
)

//链表定义
type LNode struct {
	Data interface{}
	Next *LNode
}

//
//  Reverse
//  @Description: 节点反转(插入法，比掉头法效率高，不保存前驱结点地址，比递归法效率高，因为不需要递归调用消耗时间复杂度)
//  @param node 链表结点
//
func Reverse(node *LNode){
     if node == nil || node.Next == nil{
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

//创建链表
func CreateNode(node *LNode, max int) {
	cur := node
	for i := 1; i < max; i++ {
		cur.Next = &LNode{}
		cur.Next.Data = i
		cur = cur.Next
	}
}

//打印链表的方法
func PrintNode(info string, node *LNode) {
	fmt.Print(info)
	for cur := node.Next; cur != nil; cur = cur.Next {
		fmt.Print(cur.Data, " ")
	}
	fmt.Println()
}
