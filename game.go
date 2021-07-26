package main

import (
	"fmt"
	"os"
	"time"
)

type Game struct {
	snakes          []*Snake
	ditu            *Map
	cur_key         byte
	key_change_flag int
	input           *Inkey
	interval        int
	t               int
}

func NewGame(w int, h int, fps int) *Game {
	game := &Game{
		ditu:            NewMap(w, h),
		cur_key:         ' ',
		key_change_flag: 0,
		input:           New(os.Stdin),
		interval:        1000 / fps,
		t:               0,
	}
	game.snakes = append(game.snakes, NewSnake(13))
	game.snakes = append(game.snakes, NewSnake(16))
	//game.snakes = append(game.snakes, NewSnake(19))
	/*game.snakes = append(game.snakes, NewSnake(18))
	game.snakes = append(game.snakes, NewSnake(20))
	game.snakes = append(game.snakes, NewSnake(22))
	game.snakes = append(game.snakes, NewSnake(24))*/
	return game
}

func (game *Game) UpdateCurKey() {
	var last byte
	b, found := game.input.Inkey()
	if found {
		last = b
		for {
			b, found := game.input.Inkey()
			if found {
				last = b
			} else {
				//log.Printf("key found: '%c' value=%d", last, last)
				/*if last == 'w' {
					game.cur_key = 'w'
					game.key_change_flag = 1
				} else if last == 's' {
					game.cur_key = 's'
					game.key_change_flag = 1
				} else if last == 'a' {
					game.cur_key = 'a'
					game.key_change_flag = 1
				} else if last == 'd' {
					game.cur_key = 'd'
					game.key_change_flag = 1
				} else if last == 'q' {
					game.cur_key = 'q'
					game.key_change_flag = 1
				}*/
				game.cur_key = last
				game.key_change_flag = 1
				break
			}
		}
	}
}

func (game *Game) OP() {
	fmt.Println("snake!  <press any key to start>")
	for {
		_, found := game.input.Inkey()
		if found {
			break
		}
	}
	fmt.Printf("\033[%d;%dH", 0, 0)
}

func (game *Game) ED() {
	fmt.Printf("\033[%d;%dH", game.ditu.height+5, 0)
	fmt.Println("Quit")
}

func (game *Game) Prepare() {
	game.ditu.Show()
	for i := 0; i < len(game.snakes); i++ {
		game.snakes[i].Show()
	}
	game.ditu.GenerateFood()
}

func (game *Game) Run() {
	game.OP()
	game.Prepare()
	for {
		game.UpdateCurKey()

		//if key is pressed
		if game.key_change_flag == 1 {
			game.key_change_flag = 0

			//quit
			if game.cur_key == 'q' {
				break
			}

			//change directions of snakes
			for i := 0; i < len(game.snakes); i++ {
				if i == 0 {
					if game.cur_key == 'w' {
						game.snakes[i].ChangeDirection('u')
					} else if game.cur_key == 's' {
						game.snakes[i].ChangeDirection('d')
					} else if game.cur_key == 'a' {
						game.snakes[i].ChangeDirection('l')
					} else if game.cur_key == 'd' {
						game.snakes[i].ChangeDirection('r')
					}
				} else if i == 2 {
					if game.cur_key == 't' {
						game.snakes[i].ChangeDirection('u')
					} else if game.cur_key == 'g' {
						game.snakes[i].ChangeDirection('d')
					} else if game.cur_key == 'f' {
						game.snakes[i].ChangeDirection('l')
					} else if game.cur_key == 'h' {
						game.snakes[i].ChangeDirection('r')
					}
				} else if i == 1 {
					if game.cur_key == 'i' {
						game.snakes[i].ChangeDirection('u')
					} else if game.cur_key == 'k' {
						game.snakes[i].ChangeDirection('d')
					} else if game.cur_key == 'j' {
						game.snakes[i].ChangeDirection('l')
					} else if game.cur_key == 'l' {
						game.snakes[i].ChangeDirection('r')
					}
				}
			}
		}

		for i := 0; i < len(game.snakes); i++ {
			if game.t%game.snakes[i].speed_scale == 0 {

				game.snakes[i].DirectionFilter()

				if game.snakes[i].Eat(game.ditu.food) {
					for {
						game.ditu.GenerateFood()
						flag := 0
						for _, j := range game.snakes[i].body.pos[:len(game.snakes[i].body.pos)-1] {
							if game.ditu.food.x == j.x && game.ditu.food.y == j.y {
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

				if game.snakes[i].isEatSelf() {
					fmt.Printf("\033[%d;%dH", game.ditu.height+5, 0)
					fmt.Println("Eat Self!")
					return
				}

				game.snakes[i].Move()

				if game.snakes[i].isHitWall(game.ditu) {
					fmt.Printf("\033[%d;%dH", height+5, 0)
					fmt.Println("Hit Wall!")
					return
				}

				fmt.Printf("\033[%d;%dH", 10+i, 100)
				fmt.Printf("len%d :     ", i+1)
				fmt.Printf("\033[%d;%dH", 10+i, 106)
				fmt.Printf("%d", game.snakes[i].body.count)
			}
		}

		time.Sleep(time.Duration(game.interval) * time.Millisecond)
		game.t++
	}
	game.ED()
}
