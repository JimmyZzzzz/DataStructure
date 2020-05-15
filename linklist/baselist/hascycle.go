package baselist

func HasCycle(head *MyLinkedList) bool {
	if head == nil {
		return false
	}

	fast, slow := head, head

	for fast != nil && slow != nil && fast.Next != nil {

		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			return true
		}

	}

	return false

}

func DetectCycle(head *MyLinkedList) *MyLinkedList {
	if head == nil {
		return nil
	}

	//链表头a到环口b距离为x,快指针和慢指针相遇点c,b到c为y,c到b为z
	//快指针运行长度 x + y + z + y =  慢指针运行长度 x+y 的两倍
	//x+y+z+y = 2(x+y)  结论 x=z，运行第三个慢指针 相遇地为环入口

	slow, fast := head, head

	for slow != nil && fast != nil && fast.Next != nil {

		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {

			slow2 := head

			for slow != slow2 {
				slow = slow.Next
				slow2 = slow2.Next

			}

			return slow2

		}

	}
	return nil
}

func GetIntersectionNode(headA *MyLinkedList, headB *MyLinkedList) *MyLinkedList {

	if headA == nil || headB == nil {
		return nil
	}

	a, b := headA, headB

	for a != b {

		if a == nil {
			a = headB
		} else {
			a = a.Next
		}

		if b == nil {
			b = headA
		} else {
			b = b.Next
		}

	}

	return a

}

func removeNthFromEnd(head *MyLinkedList, n int) *MyLinkedList {

	if head == nil {
		return nil
	}

	i := 0
	lHead, dHead := head, head

	for lHead != nil {
		lHead = lHead.Next
		i++
	}

	index := n - i

	for dHead != nil {
		if index == 0 { //删头

		} else if index == i { //删尾

		} else {

		}
	}

}
