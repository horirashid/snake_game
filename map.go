package main

import (
	"fmt"
)

type Map struct {
	width      int
	height     int
	wall_up    int
	wall_down  int
	wall_left  int
	wall_right int
	foods      []*Point
}

func NewMap(w int, h int) *Map {
	m := &Map{
		width:      w,
		height:     h,
		wall_up:    1,
		wall_down:  height + 4,
		wall_left:  1,
		wall_right: width + 2,
		foods:      []*Point{},
	}
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
