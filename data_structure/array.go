package data_structure

import (
	"fmt"
)

type User struct {
	Id int
	name string
	age int
}

func SetUser() []User  {
	userlist := []User{{1,"admin",10},{2,"小红",11}}
	return userlist
}
func (user User) ReadUser () int  {
	return user.Id
}


func SetArray() []int {
	arr := []int{2,3,6,5,7}
	return arr
}
func ReadArray(arr []int)  {
	for k, v := range arr  {
		fmt.Println(k,"=>",v)
	}
}
func ReadArrayByKey(arr []int,key int) (int,bool)  {
	num := len(arr)
	if key>=num {
		return 0,false
	}
	if key<0 {
		return 0,false
	}
	return arr[key],true
}
func UpdateArray(arr []int,key int,val int)([]int,bool)  {
	num := len(arr)
	if key>=num {
		return arr,false
	}
	if key<0 {
		return arr,false
	}
	arr[key] = val
	return arr,true
}
func DelArray(arr []int,key int)([]int,bool)  {
	num := len(arr)
	if key>=num {
		return arr,false
	}
	if key<0 {
		return arr,false
	}
	return append(arr[:key],arr[key+1:]...) ,true
}
