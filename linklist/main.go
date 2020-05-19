package main

import (
	"fmt"
	"linklist/baselist"
)

func main() {

	//[3,2,0,-4]

	root := baselist.Constructor()
	root.AddAtHead(3)
	root.AddAtTail(2)
	root.AddAtTail(0)
	root.AddAtTail(-4)

	root.Next.Next.Next = root.Next

	//baselist.PrintMyLinkList(root)
	//fmt.Println(baselist.HasCycle(root))
	//fmt.Println(baselist.DetectCycle(root))

	delRoot := baselist.Constructor()
	delRoot.AddAtHead(1)
	delRoot.AddAtTail(2)
	delRoot.AddAtTail(3)
	delRoot.AddAtTail(4)
	delRoot.AddAtTail(5)

	//baselist.PrintMyLinkList(delRoot)

	delRoot = baselist.RemoveNthFromEnd(delRoot, 1)

	//baselist.PrintMyLinkList(delRoot)

	ReverseRoot := baselist.Constructor()
	ReverseRoot.AddAtHead(1)
	ReverseRoot.AddAtTail(2)
	ReverseRoot.AddAtTail(3)
	ReverseRoot.AddAtTail(4)
	ReverseRoot.AddAtTail(5)

	baselist.PrintMyLinkList(ReverseRoot)
	ReverseRoot = baselist.ReverseLinkList(ReverseRoot)
	if !baselist.HasCycle(ReverseRoot) {
		baselist.PrintMyLinkList(ReverseRoot)
	} else {
		fmt.Println("Has cycle")
	}

	ReverseRoot = baselist.Constructor()
	ReverseRoot.AddAtHead(1)
	baselist.PrintMyLinkList(ReverseRoot)

	ReverseRoot = baselist.ReverseLinkList(ReverseRoot)
	baselist.PrintMyLinkList(ReverseRoot)

	RemeElementsRoot := baselist.Constructor()
	RemeElementsRoot.AddAtHead(1)
	fmt.Println(baselist.RemoveElements(RemeElementsRoot, 1))

	oddEvenRoot := baselist.Constructor()
	oddEvenRoot.AddAtHead(1)
	oddEvenRoot.AddAtTail(2)
	oddEvenRoot.AddAtTail(3)
	oddEvenRoot.AddAtTail(4)
	oddEvenRoot.AddAtTail(5)
	oddEvenRoot.AddAtTail(6)
	oddEvenRoot.AddAtTail(7)
	oddEvenRoot.AddAtTail(8)
	baselist.PrintMyLinkList(oddEvenRoot)
	baselist.OddEvenList(oddEvenRoot)
	baselist.PrintMyLinkList(oddEvenRoot)
}
