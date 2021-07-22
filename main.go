package main

import (
	"fmt"
	"os"
	"time"
)

var width int = 40
var height int = 20

func ClearScreen() {
	fmt.Println()
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
	fmt.Printf("\033[%d;%dH", last_y, last_x)
	fmt.Printf("%c", ' ')

	fmt.Printf("\033[%d;%dH", cur_y, cur_x)
	fmt.Printf("%c", '*')

}

func isHitWall(cur_x int, cur_y int) bool {
	if cur_x <= 1 {
		return true
	}
	if cur_x >= width+2 {
		return true
	}
	if cur_y <= 2 {
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

	var last_x int
	var last_y int
	var cur_x int
	var cur_y int

	cur_x = 2
	cur_y = 3
	last_x = cur_x
	last_y = cur_y

	fmt.Printf("\033[%d;%dH", cur_y, cur_x)
	fmt.Printf("%c", '*')

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
						cur_y--
					} else if last == 's' {
						cur_y++
					} else if last == 'a' {
						cur_x--
					} else if last == 'd' {
						cur_x++
					}
					break
				}
			}
		}

		time.Sleep(100 * time.Millisecond)

		MoveHead(last_x, last_y, cur_x, cur_y)
		if isHitWall(cur_x, cur_y) {
			fmt.Printf("\033[%d;%dH", height+6, 0)
			fmt.Println("Hit Wall!")
			return
		}

		last_x = cur_x
		last_y = cur_y
	}
}
