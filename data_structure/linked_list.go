package data_structure

import "fmt"

type Node struct {
	Val int
	Next *Node
}

func InsertLinkedListNode( head *Node, node *Node)  {
	temp := head
	for  {
		if temp.Next == nil {
			break
		}
		temp = temp.Next
	}
	temp.Next = node
}
func ShowLinkedListNode(node *Node)  {
	list := node
	if list.Next == nil {
		fmt.Println("链表为空")
	}
	for {
		if list.Next == nil {
			fmt.Printf("val is %x \n",list.Val)
			fmt.Println("end")
			break
		}
		fmt.Printf("val is %x \n",list.Val)
		list = list.Next
	}

}