链表的分类：

* 单链表
* 双向链表
* 循环链表

# 1 带头结点的单链表

~~~go
type ListNode struct {
	Val  int
	Next *ListNode
}

func (linkedList *ListNode) createLinkedList(nums []int) {
	if len(nums) == 0 {
		return
	}

	lastNode := linkedList
	for _, ele := range nums {
		newNode := linkedList.appendNode(ele, lastNode)
		lastNode = newNode
	}
}

// 在 node 节点的后面新增结点
func (linkedList *ListNode) appendNode(value int, node *ListNode) *ListNode {
	newNode := &ListNode{
		Val:  value,
		Next: nil,
	}
	// 无需判断！是带有哨兵结点的
	// if node.Next != nil {
	// 	newNode.Next = node.Next
	// }
	newNode.Next = node.Next
	node.Next = newNode
	return newNode
}

func (linkedList *ListNode) printAll() {
	nextNode := linkedList.Next
	for nextNode != nil {
		fmt.Printf("node:%d\n", nextNode.Val)
		nextNode = nextNode.Next
	}
}
~~~

一个带有“哨兵”结点的单链表的 merge 操作：

~~~go
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// 带有 Head 结点
	if l1.Next == nil {
		return l2
	}
	if l2.Next == nil {
		return l1
	}

	l1 = l1.Next
	l2 = l2.Next

	mergeList := &ListNode{}
	flagPtr := mergeList
	for l1 != nil && l2 != nil {
		if l1.Val >= l2.Val {
			flagPtr = flagPtr.appendNode(l2.Val, flagPtr)
			l2 = l2.Next
		} else {
			flagPtr = flagPtr.appendNode(l1.Val, flagPtr)
			l1 = l1.Next
		}
	}

	for l1 != nil {
		flagPtr = flagPtr.appendNode(l1.Val, flagPtr)
		l1 = l1.Next
	}

	for l2 != nil {
		flagPtr = flagPtr.appendNode(l2.Val, flagPtr)
		l2 = l2.Next
	}

	return mergeList
}

func testLinkedList() {
	// 哨兵节点的指针，也可称之为 head
	l1 := &ListNode{}
	l1.createLinkedList([]int{1, 2, 3, 4})
	l1.printAll()

	l2 := &ListNode{}
	l2.createLinkedList([]int{0, 0, 0, 5})
	l2.printAll()

	merge := mergeTwoLists(l1, l2)
	merge.printAll()
}
~~~

