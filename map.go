package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Map struct {
	width      int
	height     int
	wall_up    int
	wall_down  int
	wall_left  int
	wall_right int
	food       Point
}

func NewMap(w int, h int) *Map {
	m := &Map{
		width:      w,
		height:     h,
		wall_up:    1,
		wall_down:  h + 4,
		wall_left:  1,
		wall_right: w + 2,
	}
	rand.Seed(time.Now().Unix())
	return m
}

func (m *Map) Show() {
	for i := 0; i < m.width+2; i++ {
		fmt.Printf("-")
	}
	fmt.Println()
	for i := 0; i < m.height+2; i++ {
		fmt.Print("|")
		for j := 0; j < m.width; j++ {
			fmt.Print(" ")
		}
		fmt.Println("|")
	}
	for i := 0; i < m.width+2; i++ {
		fmt.Printf("-")
	}
}

func (m *Map) GenerateFood() {
	m.food.x = rand.Intn(m.width-2) + 2
	m.food.y = rand.Intn(m.height-2) + 2
	fmt.Printf("\033[%d;%dH", m.food.y, m.food.x)
	fmt.Printf("%c", '*')
}
