package main

import (
	"container/list"
	"fmt"
)

func Push(elem interface{}, queue *list.List) {
	// ...
	queue.PushBack(elem)
}

func Pop(queue *list.List) interface{} {
	// ...
	return queue.Remove(queue.Front())
}

func printQueue(queue *list.List) {
	// ...
	for temp := queue.Front(); temp != nil; temp = temp.Next() {
		fmt.Printf("%v", temp.Value)
	}
}

func main() {
	queue := list.New()

	Push(1, queue)

	Push(2, queue)

	Push(3, queue)

	printQueue(queue) // 123

	Pop(queue)

	printQueue(queue) // в ту же строку: 23

}
