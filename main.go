package main

import (
	"fmt"
	"math/rand"
)

var width int = 80
var height int = 30

var food Point
var score int = 0

func isHitWall(cur_x int, cur_y int) bool {
	if cur_x <= 1 {
		return true
	}
	if cur_x >= width+2 {
		return true
	}
	if cur_y <= 1 {
		return true
	}
	if cur_y >= height+5 {
		return true
	}
	return false
}

func GenerateFood() {
	food.x = rand.Intn(width-2) + 2
	food.y = rand.Intn(height-2) + 2
	fmt.Printf("\033[%d;%dH", food.y, food.x)
	fmt.Printf("%c", asd)
}

func UpdateScore() {
	fmt.Printf("\033[%d;%dH", 10, 100)
	fmt.Printf("   ")
	fmt.Printf("\033[%d;%dH", 10, 100)
	fmt.Printf("%d", score)
}

func main() {
	game := NewGame(80, 30, 100) //width, height, fps
	game.Run()
}
