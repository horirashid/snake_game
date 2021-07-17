package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func ClearScreen() {
	fmt.Println()
	for i := 0; i < 50; i++ {
		fmt.Printf("-")
	}
	for i := 0; i < 20; i++ {
		fmt.Println()
	}
	for i := 0; i < 50; i++ {
		fmt.Printf("-")
	}
}

var last_x int
var last_y int
var cur_x int
var cur_y int

func main() {
	input := New(os.Stdin)

	//ClearScreen()

	cur_x = 1
	cur_y = 3
	last_x = cur_x
	last_y = cur_y

	//fmt.Printf("\033[%d;%dH", cur_y, cur_x)
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
						fmt.Println("you pressed w")
					} else if last == 's' {
						fmt.Println("you pressed s")
					} else if last == 'a' {
						fmt.Println("you pressed a")
					} else if last == 'd' {
						fmt.Println("you pressed d")
					}
					break
				}
			}
		}
		fmt.Printf("fresh\n")
		time.Sleep(100 * time.Millisecond)

	}

	log.Printf("done")
}
