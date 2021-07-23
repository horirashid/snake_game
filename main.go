package main

import (
	"fmt"
	"os"
	"time"
)

var width int = 80
var height int = 30

func ClearScreen() {
	for i := 0; i < width+2; i++ {
		fmt.Printf("-")
	}
	fmt.Println()
	for i := 0; i < height+2; i++ {
		fmt.Print("|")
		for j := 0; j < width; j++ {
			fmt.Print(" ")
		}
		fmt.Println("|")
	}
	for i := 0; i < width+2; i++ {
		fmt.Printf("-")
	}
}

func MoveHead(last_x int, last_y int, cur_x int, cur_y int) {
	/*fmt.Printf("\033[%d;%dH", last_y, last_x)
	fmt.Printf("%c", ' ')

	fmt.Printf("\033[%d;%dH", cur_y, cur_x)
	fmt.Printf("%c", '*')*/

}

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

func main() {
	input := New(os.Stdin)

	ClearScreen()

	snake := NewSnake()

	for {

		//asdy -> asdy
		var last byte
		b, found := input.Inkey()
		if found {
			//log.Printf("key found: '%c' value=%d", b, b)
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

		snake.Move()
		if snake.isHitWall() {
			fmt.Printf("\033[%d;%dH", height+5, 0)
			fmt.Println("Hit Wall!")
			return
		}

		time.Sleep(100 * time.Millisecond)
	}
}
