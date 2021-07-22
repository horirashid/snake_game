package main

import (
	"container/list"
	"fmt"
)

type Point struct {
	x, y int
}

func main() {
	p1 := &Point{5,5}
	p2 := &Point{4,5}
	p3 := &Point{3,5}
	p4 := &Point{2,5}
	queue := list.New()
	queue.PushBack(p1) // push new Point object into the queue
	queue.PushBack(p2)
	queue.PushBack(p3)
	queue.PushBack(p4)

	// change value of one pointer
	tmp := queue.Front().Value.(*Point)
	tmp.x = tmp.x+1

	// iterate through the queue
	// print all values
	// remove all elements
	for queue.Len() > 0 {
		obj := queue.Front() // view the first object from the queue
		xy := obj.Value.(*Point) // convert it to the Point object
		fmt.Printf("x: %d, y: %d\n", xy.x, xy.y)
		queue.Remove(obj) // Dequeue
	}

}
