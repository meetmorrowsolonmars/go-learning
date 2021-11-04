package main

import "fmt"

func main() {
	fmt.Println("Linked list exercise")

	var list *Node

	PushBack(&list, 1)
	PushBack(&list, 2)
	PushBack(&list, 3)
	PushBack(&list, 4)
	PushBack(&list, 5)

	// Prints [1, 2, 3, 4, 5]
	fmt.Println("\nPush back")
	for iterator := list; iterator != nil; iterator = iterator.Next {
		fmt.Printf("%+v\n", iterator)
	}

	PushFront(&list, 0)
	PushFront(&list, -1)
	PushFront(&list, -2)

	// Prints [-2, -1, 0, 1, 2, 3, 4, 5]
	fmt.Println("\nPush front")
	for iterator := list; iterator != nil; iterator = iterator.Next {
		fmt.Printf("%+v\n", iterator)
	}

	Pop(&list, 0)
	Pop(&list, 1)
	Pop(&list, 2)

	// Prints [-2, -1, 3, 4, 5]
	fmt.Println("\nPop")
	for iterator := list; iterator != nil; iterator = iterator.Next {
		fmt.Printf("%+v\n", iterator)
	}

	DeleteList(&list)

	// The loop is not executed even once.
	// Prints "nil"
	fmt.Println("\nDelete list")
	for iterator := list; iterator != nil; iterator = iterator.Next {
		fmt.Printf("%+v\n", iterator)
	}
	fmt.Printf("%+v\n", list)
}

type Node struct {
	Value int
	Next  *Node
}

func PushBack(head **Node, value int) {
	if head == nil {
		panic("the pointer where the list will be created or where the list " +
			"already created cannot be empty")
	}

	if *head == nil {
		*head = &Node{Value: value, Next: nil}
		return
	}

	iterator := *head
	for iterator.Next != nil {
		iterator = iterator.Next
	}

	iterator.Next = &Node{Value: value, Next: nil}
}

func PushFront(head **Node, value int) {
	// your code
}

func Pop(head **Node, value int) *Node {
	// your code
	return nil
}

func DeleteList(head **Node) {
	// your code
}
