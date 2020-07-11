package main

import (
	"fmt"
)

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func main() {

	root := sliceToLinkedList([]int{1, 2, 3, 4, 5, 6})
	root.Next.Next.Child = sliceToLinkedList([]int{7, 8, 9, 10})
	root.Next.Next.Child.Next.Child = sliceToLinkedList([]int{11, 12})

	fmt.Println("\n\nNode: root\n")

	prettyPrint(root, "")

	fmt.Println("\n\nNode: flat\n")

	flat := flatten(root)
	prettyPrint(flat, "")

}

func sliceToLinkedList(s []int) *Node {
	root := &Node{Val: s[0]}
	current := root
	for i := range s {
		if i == 0 {
			continue
		}
		node := &Node{Val: s[i]}
		current.Next, node.Prev = node, current
		current = node
	}
	return root
}

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Prev *Node
 *     Next *Node
 *     Child *Node
 * }
 */

func flatten(root *Node) *Node {
	current := root
	for {
		if current == nil {
			break
		}

		if current.Child == nil {
			current = current.Next
			continue
		}

		current.Child.Prev = current
		childCurrent := current.Child
		for {
			if childCurrent.Next != nil {
				childCurrent = childCurrent.Next
				continue
			}
			childCurrent.Next = current.Next
			current.Next.Prev = childCurrent
			break
		}
		current.Next = current.Child
		current.Child = nil
		current = current.Next
	}
	return root
}

func prettyPrint(node *Node, prefix string) {
	if node == nil {
		return
	}
	current := node
	for {
		fmt.Printf("%s[Val: %d, ", prefix, current.Val)
		if current.Prev != nil {
			fmt.Printf("PrevVal: %d, ", current.Prev.Val)
		}

		if current.Next != nil {
			fmt.Printf("NextVal: %d, ", current.Next.Val)
		}

		if current.Child != nil {
			fmt.Print("Child: \n")
			prettyPrint(current.Child, "    "+prefix)
		}
		fmt.Print("]\n")

		if current.Next == nil {
			break
		}

		current = current.Next

	}

}
