package main

import (
	"errors"
	"fmt"
)

type Point struct {
	x, y int
}
type Queue struct {
	pos   []*Point
	count int
	head  int
}

func (q *Queue) Push(p *Point) (err error) {
	if q.isEmpty() {
		q.resetQueue()
		q.tail++
	}
	q.count++
	q.pos = append(q.pos, p)
	return
}

func (q *Queue) Pop() (p *Point, err error) {
	if q.isEmpty() {
		q.resetQueue()
		fmt.Println("queue empty")
		err = errors.New("queue empty")
		return nil, err
	}
	p = q.pos[0]
	for i := 0; i < q.count-1; i++ {
		q.pos[i] = q.pos[i+1]
	}
	q.count--
	q.pos = append(q.pos[:q.count-1], q.pos[q.count:]...)
	return p, nil
}

func (q *Queue) List() (err error) {
	if q.isEmpty() {
		q.resetQueue()
		fmt.Println("queue empty")
		err = errors.New("queue empty")
		return
	}
	for i := 0; i < q.count; i++ {
		fmt.Printf("queue[%d]=%d\n", i, q.pos[i])
	}
	return
}

func (q *Queue) resetQueue() {
	q.head = -1
	q.count = 0
	q.pos = q.pos[:0]
}
func (q *Queue) isEmpty() bool {
	return q.count == 0
}

/*func main() {
	p1 := &Point{5, 5}
	p2 := &Point{4, 5}
	p3 := &Point{3, 5}
	p4 := &Point{2, 5}

	q := &Queue{
		pos:   []*Point{p1, p2},
		count: 2,
		head:  0,
	}

	q.List()
	fmt.Print("---push (3,5)\n")
	q.Push(p3)
	q.List()
	fmt.Print("---pop\n")
	q.Pop()
	q.List()
	fmt.Print("---push (2,5)\n")
	q.Push(p4)
	q.List()
	fmt.Print("---pop\n")
	q.Pop()
	q.List()
	fmt.Print("---reset\n")
	q.resetQueue()
	q.List()
}*/
