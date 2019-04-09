package main

import (
	"fmt"
	"sort"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := &ListNode{
		Val: 1,
	}

	var l2 ListNode

	fmt.Println(toArray(mergeTwoLists(l1, &l2), []int{}))
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	l3 := &ListNode{}

	arr := []int{}
	arr = toArray(l1, arr)
	arr = toArray(l2, arr)
	sort.Ints(arr)

	l3 = toNode(arr, l3)

	return l3
}

func toArray(node *ListNode, arr []int) []int {
	if node == nil {
		return nil
	}

	arr = append(arr, node.Val)
	if node.Next != nil {
		return toArray(node.Next, arr)
	}

	return arr
}

func toNode(arr []int, node *ListNode) *ListNode {
	if len(arr) < 1 || (node == &ListNode{}) {
		return nil
	}

	node.Val = arr[0]
	if len(arr) > 1 {
		node.Next = toNode(arr[1:len(arr)], &ListNode{})
	}

	return node
}
