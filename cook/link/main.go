package link

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * 反转链表
 * @param head ListNode类 链表
 * @return ListNode类
 */
func reverseList(head *ListNode) *ListNode {
	// write code here
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	pre := head
	cur := head.Next
	head.Next = nil
	for true {
		tmp := cur.Next
		cur.Next = pre
		if tmp == nil {
			break
		}
		pre = cur
		cur = tmp
	}
	return cur
}
