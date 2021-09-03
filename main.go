package main

import (
	"fmt"
)

const (
	N                    = 5
	NONE_NODE_FLAG_INDEX = -1
)

type StaticLinkedList struct {
	list [N]*Node

	first int // 数据链表的第一个结点 index 值
	last  int // 数据链表的最后一个结点 index 值
}

type Node struct {
	ele  rune // ele 默认值是 0
	next int  // 只要 next 值为 LAST_NODE_FLAG_INDEX，就表示是最后一个结点
}

func (list *StaticLinkedList) printAll() {
	// *Node 类型
	for index, node := range list.list {
		if node != nil {
			fmt.Printf("index:%d, ele:%c, next:%d\n", index, node.ele, node.next)
		}
	}
}

// printEle 从 list.first 开始的所有元素值
func (list *StaticLinkedList) printEle() {
	node := list.list[list.first]
	for {
		if node == nil {
			break
		}
		fmt.Printf("ele:%c, next:%d\n", node.ele, node.next)
		node = list.list[node.next]
	}
}

func (list *StaticLinkedList) create() {
	if N <= 1 {
		panic("静态链表元素个数 N 至少包括 2 个结点！")
	}
	node := &Node{
		next: 1, // 首个可用结点的 index 值
	}
	list.list[0] = node
	list.last = NONE_NODE_FLAG_INDEX
	// FIXME 待修订！
	list.first = 1
}

// addLast 从备用链表中取出一个可用 index，并创建 Node 放置于 index 处
func (list *StaticLinkedList) addLast(ele rune) bool {
	availableNode := list.list[0]
	noUsedIndex := availableNode.next
	if noUsedIndex == NONE_NODE_FLAG_INDEX {
		fmt.Println("已满！")
		return false
	}

	fmt.Printf("ele:%c, noUsedIndex:%d\n", ele, noUsedIndex)

	newNode := &Node{
		ele:  ele,
		next: NONE_NODE_FLAG_INDEX,
	}

	if list.last != NONE_NODE_FLAG_INDEX {
		lastNode := list.list[list.last]
		if lastNode == nil {
			panic("last node is nil!")
		}
		lastNode.next = noUsedIndex
	}

	list.list[noUsedIndex] = newNode
	//FIXME 找不到下一个 nextNoUsed
	list.list[0].next = nextNoUsed
	return true
}

func main() {
	list := &StaticLinkedList{}
	list.create()

	list.addLast('A')
	list.printAll()

	list.addLast('B')
	list.printAll()

	fmt.Println()
	list.delete('A')
	list.printAll()
}

func (list *StaticLinkedList) insert() {

}

func (list *StaticLinkedList) delete(ele rune) {
	node := list.list[list.first]
	index := list.first
	for {
		if node.ele == ele {
			// delete the special Node: node
			node.ele = 0
			if list.first == index {
				list.first = node.next
			}
			node.next = NONE_NODE_FLAG_INDEX

			list.list[list.last].next = index
			break
		}

		if node.next == NONE_NODE_FLAG_INDEX {
			break
		}
		index = node.next
		node = list.list[node.next]
	}
}
