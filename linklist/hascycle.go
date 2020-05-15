package main

import (
	"linklist/baselist"
)

func hasCycle(head *baselist.MyLinkedList) bool {

	return true

}

func main() {

	head := baselist.Constructor()

	head.AddAtHead(1)
	head.AddAtHead(2)
	head.AddAtHead(3)

	baselist.PrintMyLinkList(head)

}
