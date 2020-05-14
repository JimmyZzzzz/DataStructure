package baselist

import (
	"fmt"
)

/**
在链表类中实现这些功能：
get(index)：获取链表中第 index 个节点的值。如果索引无效，则返回-1。
addAtHead(val)：在链表的第一个元素之前添加一个值为 val 的节点。插入后，新节点将成为链表的第一个节点。
addAtTail(val)：将值为 val 的节点追加到链表的最后一个元素。
addAtIndex(index,val)：在链表中的第 index 个节点之前添加值为 val  的节点。如果 index 等于链表的长度，则该节点将附加到链表的末尾。如果 index 大于链表长度，则不会插入节点。如果index小于0，则在头部插入节点。
deleteAtIndex(index)：如果索引 index 有效，则删除链表中的第 index 个节点。

MyLinkedList linkedList = new MyLinkedList();
linkedList.addAtHead(1);
linkedList.addAtTail(3);
linkedList.addAtIndex(1,2);   //链表变为1-> 2-> 3
linkedList.get(1);            //返回2
linkedList.deleteAtIndex(1);  //现在链表是1-> 3
linkedList.get(1);            //返回3

["addAtHead","addAtTail","addAtHead","addAtTail","addAtHead","addAtHead","get","addAtHead","get","get","addAtTail"]
[    [7],        [7],        [9],        [8],        [6],        [0],      [5],    [0],     [2],   [5],   [4] ]

*/

type MyLinkedList struct {
	Val  int
	Next *MyLinkedList
}

func Constructor() *MyLinkedList {
	return &MyLinkedList{Val: -2}
}

//func main() {
//
//	root := Constructor()
//	root.addAtHead(7)
//	root.addAtTail(7)
//	root.addAtHead(9)
//	root.addAtTail(8)
//	root.addAtHead(6)
//	root.addAtHead(0)
//	fmt.Println(root.get(5))
//	PrintMyLinkList(root)
//	root.addAtHead(0)
//	PrintMyLinkList(root)
//
//
//
//}

func PrintMyLinkList(node *MyLinkedList) {

	fmt.Printf("%d ", node.Val)

	if node.Next == nil {
		fmt.Println(" ")
		return
	}

	PrintMyLinkList(node.Next)
}

func (root *MyLinkedList) Get(searchInx int) int {
	i := 0
	n := root

	for {

		if i == searchInx {
			return n.Val
		}

		if n.Next == nil {
			break
		}

		n = n.Next

		i++

	}

	return -1

}

func (root *MyLinkedList) AddAtHead(val int) {

	if root.Val == -2 {
		root.Val = val
		return
	}

	newNode := &MyLinkedList{Val: val} //新节点
	tmp := *root                       //复制出HeaderNode数据
	newNode.Next = &tmp                //把复制HeaderNode数据地址复制给新节点next
	*root = *newNode                   //HeaderNode偏移

}

func (root *MyLinkedList) AddAtTail(val int) {

	n := root
	newNode := &MyLinkedList{Val: val}

	for {
		if n.Next != nil {
			n = n.Next
		} else {
			n.Next = newNode
			break
		}
	}

}

func (root *MyLinkedList) AddAtIndex(index int, val int) {

	n := root
	newNode := &MyLinkedList{Val: val}
	i := 0

	for {
		if i < index && n.Next == nil {
			//索引超过长度
			n.Next = newNode
			break
		} else if i == index { //当前索引位置

			*newNode = *n
			n.Next = newNode
			n.Val = val
			break

		}

		n = n.Next
		i++
	}

}

func (root *MyLinkedList) DeleteAtIndex(index int) {

	i := 0
	n := root
	prev := root

	for {

		if i == index && n.Next != nil {
			*n = *n.Next
			return
		} else if i == index && n.Next == nil {
			prev.Next = nil
			return
		} else if i != index && n.Next == nil {
			return
		}
		prev = n
		n = n.Next
		i++

	}

}
