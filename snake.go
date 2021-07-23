package main

import (
	"fmt"
)

const (
	asd = '@'
)

func NewSnake() *Snake {
	s := &Snake{body: &Queue{}, dir: 'r'}

	for i := 18; i <= 20; i++ {
		n := &Point{x: i, y: 13}
		fmt.Printf("\033[%d;%dH", n.y, n.x)
		fmt.Printf("%c", asd)
		s.body.Push(n)
	}
	s.head = &Point{x: 20, y: 13}

	return s
}

type Snake struct {
	body *Queue
	head *Point
	dir  rune // u d l r
}

func (snake *Snake) Move() {
	p, _ := snake.body.Pop()
	fmt.Printf("\033[%d;%dH", p.y, p.x)
	fmt.Printf("%c", ' ')
	if snake.dir == 'r' {
		snake.head = &Point{x: snake.head.x + 1, y: snake.head.y}
	} else if snake.dir == 'l' {
		snake.head = &Point{x: snake.head.x - 1, y: snake.head.y}
	} else if snake.dir == 'u' {
		snake.head = &Point{x: snake.head.x, y: snake.head.y - 1}
	} else if snake.dir == 'd' {
		snake.head = &Point{x: snake.head.x, y: snake.head.y + 1}
	}
	fmt.Printf("\033[%d;%dH", snake.head.y, snake.head.x)
	fmt.Printf("%c", asd)
	snake.body.Push(snake.head)
}

func (snake *Snake) ChangeDirection(new_direction rune) {
	if new_direction == 'u' && (snake.dir == 'd' || snake.dir == 'u') {
		return
	} else if new_direction == 'd' && (snake.dir == 'u' || snake.dir == 'd') {
		return
	} else if new_direction == 'l' && (snake.dir == 'r' || snake.dir == 'l') {
		return
	} else if new_direction == 'r' && (snake.dir == 'l' || snake.dir == 'r') {
		return
	}
	snake.dir = new_direction
	//fmt.Printf("%c\n", snake.dir)
}

func (snake *Snake) isHitWall(m *Map) bool {
	if snake.head.x <= 1 {
		return true
	}
	if snake.head.x >= m.wall_right {
		return true
	}
	if snake.head.y <= 1 {
		return true
	}
	if snake.head.y >= m.wall_down {
		return true
	}
	return false
}

func (snake *Snake) PrintPos(pos Point) {
	fmt.Printf("\033[%d;%dH", pos.y, pos.x)
	fmt.Printf("%c", asd)
}

func (snake *Snake) Eat() bool {
	flag := 0
	if snake.dir == 'r' {
		if snake.head.x+1 == food.x && snake.head.y == food.y {
			flag = 1
		}
	} else if snake.dir == 'l' {
		if snake.head.x-1 == food.x && snake.head.y == food.y {
			flag = 1
		}
	} else if snake.dir == 'u' {
		if snake.head.x == food.x && snake.head.y-1 == food.y {
			flag = 1
		}
	} else if snake.dir == 'd' {
		if snake.head.x == food.x && snake.head.y+1 == food.y {
			flag = 1
		}
	}
	if flag == 1 {
		np := &Point{x: food.x, y: food.y}
		snake.body.Push(np)
		snake.head = np
		return true
	}
	return false
}

func (snake *Snake) isEatSelf() bool {

	for _, j := range snake.body.pos[:len(snake.body.pos)-1] {
		if snake.dir == 'r' {
			if snake.head.x+1 == j.x && snake.head.y == j.y {
				return true
			}
		} else if snake.dir == 'l' {
			if snake.head.x-1 == j.x && snake.head.y == j.y {
				return true
			}
		} else if snake.dir == 'u' {
			if snake.head.x == j.x && snake.head.y-1 == j.y {
				return true
			}
		} else if snake.dir == 'd' {
			if snake.head.x == j.x && snake.head.y+1 == j.y {
				return true
			}
		}
	}
	return false
}
