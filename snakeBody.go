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
	tail  int
	head  int
}

func (q *Queue) Push(p *Point) (err error) {
	if q.isEmpty() {
		q.resetQueue()
		q.tail++
	}
	q.head++
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
	p = q.pos[q.tail]
	q.tail++
	q.count--
	return p, nil
}

func (q *Queue) List() (err error) {
	if q.isEmpty() {
		q.resetQueue()
		fmt.Println("queue empty")
		err = errors.New("queue empty")
		return
	}
	for i := q.tail; i <= q.head; i++ {
		fmt.Printf("queue[%d]=%d\n", i, q.pos[i])
	}
	fmt.Println("end of queue")
	return
}

func (q *Queue) resetQueue() {
	q.head = -1
	q.tail = -1
	q.count = 0
	q.pos = q.pos[:0]
}
func (q *Queue) isEmpty() bool {
	return q.count == 0
}

func main() {
	p1 := &Point{5,5}
	p2 := &Point{4,5}
	p3 := &Point{3,5}
	p4 := &Point{2,5}

	q := &Queue{
		pos:   []*Point{p1,p2},
		head:  1,
		count: 2,
		tail:  0,
	}
	q.List()
	q.Push(p3)
	p,_ := q.Pop()
	fmt.Printf("the popped value is: x=%d, y=%d\n",p.x, p.y)
	q.List()
	q.Push(p4)
	q.Pop()
	q.List()
	q.resetQueue()
	q.List()
}
