package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
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
	return

	rand.Seed(time.Now().Unix())
	input := New(os.Stdin)

	ditu := NewMap(width, height)
	ditu.Show()

	snake := NewSnake(13)
	GenerateFood()

	for {
		var last byte
		b, found := input.Inkey()
		if found {
			last = b
			for {
				b, found := input.Inkey()
				if found {
					last = b
				} else {
					//log.Printf("key found: '%c' value=%d", last, last)
					if last == 'w' {
						snake.ChangeDirection('u')
						//cur_y--
					} else if last == 's' {
						snake.ChangeDirection('d')
						//cur_y++
					} else if last == 'a' {
						snake.ChangeDirection('l')
						//cur_x--
					} else if last == 'd' {
						snake.ChangeDirection('r')
						//cur_x++
					} else if last == 'q' {
						fmt.Printf("\033[%d;%dH", height+6, 0)
						fmt.Println("Quit")
						return
					}
					break
				}
			}
		}

		snake.DirectionFilter()
		if snake.Eat(food) {
			for {
				GenerateFood()
				flag := 0
				for _, j := range snake.body.pos[:len(snake.body.pos)-1] {
					if food.x == j.x && food.y == j.y {
						flag = 1
						break
					}
				}
				if flag == 0 {
					break
				}
			}
			score++
		}
		if snake.isEatSelf() {
			fmt.Printf("\033[%d;%dH", height+5, 0)
			fmt.Println("Eat Self!")
			return
		}

		snake.Move()

		if snake.isHitWall(ditu) {
			fmt.Printf("\033[%d;%dH", height+5, 0)
			fmt.Println("Hit Wall!")
			return
		}

		UpdateScore()

		time.Sleep(100 * time.Millisecond)
	}

}
