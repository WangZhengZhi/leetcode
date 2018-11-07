package main

import (
	"fmt"
)

type Element interface {
}
type LinkList struct {
	Data Element   //数据域
	Nest *LinkList //指针域
}

func (head *LinkList) GetLength() int {
	point := head
	var length int
	for point != nil {
		length++
		point = point.Nest

	}
	return length

} //获取链表的长度

func (head *LinkList) Traverse() {
	point := head
	for point != nil {
		fmt.Println(point.Data) //输出数据
		point = point.Nest

	}
} //遍历链表
func (head *LinkList) Delete(index int) Element {
	if index < 0 || index > head.GetLength() {
		fmt.Println("index不合法")
		return false

	} else {
		point := head
		for i := 0; i < index; i++ {
			point = point.Nest //移动位置

		}
		point.Nest = point.Nest.Nest
		data := point.Nest.Data
		return data
	}

} //删除指定位置的节点
func (head *LinkList) Insert(index int, data Element) {
	if index < 0 || index > head.GetLength() {
		fmt.Println("数据不合法")
	} else {
		point := head
		for i := 0; i < index; i++ {
			point = point.Nest //移动位置
		}
		var NewLinkList LinkList //新的节点
		NewLinkList.Data = data
		NewLinkList.Nest = point.Nest
		point.Nest = &NewLinkList
	}

} //在指定index位置插入数据
func (head *LinkList) Search(data Element) {
	point := head
	index := 0
	for point!= nil {
		if point.Data == data {
			fmt.Println(data, "show in at index: ", index)
			break
		} else {
			index++
			point = point.Nest
			if index > head.GetLength() {
				fmt.Println("no this data in the LinkList!!!")
				break
			}
			continue
		}
	}
} //搜寻data数据
func (head *LinkList) GetData(index int) Element {
	point := head
	if index < 0 || index > head.GetLength() {
		return false

	} else {

		for i := 0; i < index; i++ {
			point = point.Nest //移动位置
		}
		return point.Data
	}

} //拿取LinkList中index位置的元素
func (head *LinkList) Add(data Element) {
	point := head
	for point.Nest != nil {
		point = point.Nest //移动位置

	}
	var NewLinklist LinkList
	point.Nest = &NewLinklist //newlinklist接在point.nest之后
	NewLinklist.Data = data

} //给linklist增加数据

func main() {
	var head LinkList = LinkList{Data: 0, Nest: nil}
	var array []Element
	for i := 0; i < 9; i++ {
		array = append(array, Element(i+1))
		head.Add(array[i]) //增加数据

	}
	head.Traverse() //遍历linklist
	head.Search(9)  //查询数据
	fmt.Println("length:", head.GetLength())
	fmt.Println("data:", head.GetData(9))
	/*head.Delete(2) //删除索引为2 的数据
	head.Traverse()//遍历数据*/

}
