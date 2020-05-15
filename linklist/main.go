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

	fmt.Println(baselist.DetectCycle(root))
	//fmt.Println(baselist.HasCycle(root))
	baselist.DetectEchoCycle(root)

}
