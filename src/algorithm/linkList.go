package algorithm

import (
	"fmt"
	"github.com/isdamir/gotype"
)

//链表定义
type LNode struct {
	Data interface{}
	Next *LNode
}

func Reverse(node *LNode){
     if node == nil || node.Next == nil {
		 return
	 }
	 var pre *LNode //定义前驱结点
     var cur *LNode //定义当前节点
     next := node.Next
     for next != nil{
     	cur = next.Next
     	next.Next = pre
     	pre = next
     	next =pre
	 }
	 node.Next = pre
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
