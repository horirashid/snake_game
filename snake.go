package main

import (
	"fmt"
)

func NewSnake(ypos int) *Snake {
	s := &Snake{
		body:        &Queue{},
		dir:         'r',
		temp_dir:    'r',
		speed_scale: 8,
		keymap:      "wsad",
		body_char:   '@',
	}

	for i := 18; i <= 20; i++ {
		n := &Point{x: i, y: ypos}
		s.body.Push(n)
	}
	s.head = &Point{x: 20, y: ypos}

	return s
}

func NewSnakeByArray(body_pos []*Point, is_reverse int) *Snake {
	s := &Snake{
		body:        &Queue{},
		dir:         'r',
		temp_dir:    'r',
		speed_scale: 8,
		keymap:      "wsad",
		body_char:   '@',
	}

	for i := 0; i < len(body_pos); i++ {
		if is_reverse == 0 {
			s.body.Push(body_pos[i])
		} else {
			s.body.Push(body_pos[len(body_pos)-1-i])
		}
	}

	return s
}

type Snake struct {
	body        *Queue
	head        *Point
	dir         rune // u d l r
	temp_dir    rune
	speed_scale int
	keymap      string
	body_char   byte
}

func (snake *Snake) Show() {
	for _, j := range snake.body.pos[:len(snake.body.pos)] {
		fmt.Printf("\033[%d;%dH", j.y, j.x)
		fmt.Printf("%c", snake.body_char)
	}
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
	fmt.Printf("%c", snake.body_char)
	snake.body.Push(snake.head)
}

func (snake *Snake) DirectionFilter() {
	if snake.temp_dir == 'u' && snake.dir == 'd' {
		return
	} else if snake.temp_dir == 'd' && snake.dir == 'u' {
		return
	} else if snake.temp_dir == 'l' && snake.dir == 'r' {
		return
	} else if snake.temp_dir == 'r' && snake.dir == 'l' {
		return
	}
	snake.dir = snake.temp_dir
}

func (snake *Snake) ChangeDirectionByKey(key byte) {
	if key == snake.keymap[0] {
		snake.ChangeDirection('u')
	} else if key == snake.keymap[1] {
		snake.ChangeDirection('d')
	} else if key == snake.keymap[2] {
		snake.ChangeDirection('l')
	} else if key == snake.keymap[3] {
		snake.ChangeDirection('r')
	}
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
	snake.temp_dir = new_direction
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
	fmt.Printf("%c", snake.body_char)
}

func (snake *Snake) Eat(fpos Point) bool {
	flag := 0
	if snake.dir == 'r' {
		if snake.head.x+1 == fpos.x && snake.head.y == fpos.y {
			flag = 1
		}
	} else if snake.dir == 'l' {
		if snake.head.x-1 == fpos.x && snake.head.y == fpos.y {
			flag = 1
		}
	} else if snake.dir == 'u' {
		if snake.head.x == fpos.x && snake.head.y-1 == fpos.y {
			flag = 1
		}
	} else if snake.dir == 'd' {
		if snake.head.x == fpos.x && snake.head.y+1 == fpos.y {
			flag = 1
		}
	}
	if flag == 1 {
		np := &Point{x: fpos.x, y: fpos.y}
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
