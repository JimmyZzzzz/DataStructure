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

func RemoveNthFromEnd(head *MyLinkedList, n int) *MyLinkedList {

	if head == nil {
		return nil
	}

	i := 0
	lHead, dHead := head, head

	for lHead != nil {
		lHead = lHead.Next
		i++
	}

	index := i - n
	prev := dHead

	i-- //变成索引最长刻度

	//删头
	if index == 0 && dHead.Next != nil { //头下一节不为空
		*dHead = *dHead.Next
		return dHead
	} else if index == 0 && dHead.Next == nil { //头下一节点为空
		dHead = nil
		return dHead
	}

	didx := 0

	for dHead != nil {

		if didx == index {

			if index == i { //删尾
				prev.Next = nil
				break
			} else {
				*dHead = *dHead.Next
				break
			}

		}

		prev = dHead
		dHead = dHead.Next
		didx++

	}

	return head

}

func ReverseLinkList(head *MyLinkedList) *MyLinkedList {

	if head == nil {
		return nil
	}

	curr, next := head, head

	var prev *MyLinkedList = nil

	for curr != nil {

		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next

	}

	return prev

}
