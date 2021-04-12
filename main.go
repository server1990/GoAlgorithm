package main

import (
	"fmt"
	"goalgorithm/data_structure"
)

func main()  {
	head := &data_structure.Node{}
	stud1 := &data_structure.Node{
		Val: 12,
	}
	data_structure.InsertLinkedListNode(head,stud1)
	stud2 := &data_structure.Node{
		Val: 13,
	}
	data_structure.InsertLinkedListNode(head,stud2)
	stud3 := &data_structure.Node{
		Val: 14,
	}
	data_structure.InsertLinkedListNode(head,stud3)
	fmt.Println(head.Next.Val)

	data_structure.ShowLinkedListNode(head)
	//arr := data_structure.SetArray()
	////fmt.Println(data_structure.ReadArrayByKey(arr,-1))
	////fmt.Println(data_structure.UpdateArray(arr,-1,7))
	//arra,isBoll := data_structure.DelArray(arr,2)
	//if isBoll == false {
	//	fmt.Println("数据错误")
	//}
	//data_structure.ReadArray(arra)
	//var user []data_structure.User
	//user = data_structure.SetUser()
	//for key,val := range user  {
	//	fmt.Println(key,val.ReadUser())
	//}
	//fmt.Println(user)
}