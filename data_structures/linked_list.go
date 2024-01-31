package datastructures

import "fmt"

type Node struct {
    Data string
    Next *Node
}

type LinkedList struct {
    Head *Node
}

func (list *LinkedList) Append(data string) {
    newNode := &Node{Data: data, Next: nil}

    if list.Head == nil {
        list.Head = newNode
        return
    }

    lastNode := list.Head
    for lastNode.Next != nil {
        lastNode = lastNode.Next
    }

    lastNode.Next = newNode
}

func (list *LinkedList) Display() {
    current := list.Head
    for current != nil {
        fmt.Printf("%s -> ", current.Data)
        current = current.Next
    }
    fmt.Println("nil")
}