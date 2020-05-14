package main

import (
	"fmt"
)

/*实现一种算法，删除单向链表中间的某个节点（除了第一个和最后一个节点，不一定是中间节点），假定你只能访问该节点。
示例：
输入：单向链表a->b->c->d->e->f中的节点c
[4,5,1,9]
结果：不返回任何数据，但该链表变为a->b->d->e->f
[4,1,9]


*/
type ListNode struct {
	Data int
	Next *ListNode
}

func InsertNode(node *ListNode, data int)  *ListNode {

	node.Data = data
	newNode := &ListNode{}
	node.Next = newNode
	return newNode

}

func DeleteNode(node *ListNode){
	*node = *node.Next
}

func PrintNode(node *ListNode){
	if node.Next == nil {
		fmt.Println("")
		return
	}

	fmt.Printf("%d->",node.Data)
	PrintNode(node.Next)
}

func SearchNode (node *ListNode,data int) *ListNode{

	if node.Next == nil {
		return  nil
	}


	if node.Data == data {
		return  node
	}

	return SearchNode(node.Next,data)

}

func SearchNextNode (node *ListNode,data int) *ListNode{

	if node.Next == nil {
		return  nil
	}


	if node.Data == data {
		return  node.Next
	}

	return SearchNextNode(node.Next,data)

}



func main()  {

	root := &ListNode{}

	strArr := [...]int{4,5,1,9}

	roof := root

	for _,v := range strArr{

		str := v
		roof = InsertNode(roof,str)

	}

	PrintNode(root)

	delNode := SearchNode(root,5)

	if delNode != nil {
		DeleteNode(delNode)
	}

	PrintNode(root)



}
