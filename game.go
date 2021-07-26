package main

import (
	"fmt"
	"os"
	"regexp"
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
	//game.snakes = append(game.snakes, NewSnake(13))
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

func (game *Game) Option() (string, string) {
	root := NewNode(nil, nil, "root", "", "")
	gamemode_node := NewNode(root, nil, "game mode", "", "")
	setting_node := NewNode(root, nil, "setting", "", "")
	history_note := NewNode(root, nil, "history", "2", "")
	solo_node := NewNode(gamemode_node, nil, "solo", "0_0", "")
	double_node := NewNode(gamemode_node, nil, "double", "0_1", "")
	keymapping_node := NewNode(setting_node, nil, "keymapping", "", "")
	body_char_node := NewNode(setting_node, nil, "body_char", "1_1", "")
	fps_node := NewNode(setting_node, nil, "fps", "1_2", "")
	snake_speed_node := NewNode(setting_node, nil, "snake_speed", "1_3", "")
	snake1_node := NewNode(keymapping_node, nil, "snake1", "1_0_0", "haha")
	snake2_node := NewNode(keymapping_node, nil, "snake2", "1_0_1", "")

	root.next = append(root.next, gamemode_node)
	root.next = append(root.next, setting_node)
	root.next = append(root.next, history_note)
	gamemode_node.next = append(gamemode_node.next, solo_node)
	gamemode_node.next = append(gamemode_node.next, double_node)
	setting_node.next = append(setting_node.next, keymapping_node)
	setting_node.next = append(setting_node.next, body_char_node)
	setting_node.next = append(setting_node.next, fps_node)
	setting_node.next = append(setting_node.next, snake_speed_node)
	keymapping_node.next = append(keymapping_node.next, snake1_node)
	keymapping_node.next = append(keymapping_node.next, snake2_node)

	index := 0
	old_index := 0
	cur_node := root //*/
	option_id := ""
	value := ""
	for i := 0; i < len(cur_node.next); i++ {
		fmt.Printf("  %s\t%s\n", cur_node.next[i].name, cur_node.next[i].value)
	}
	fmt.Printf("\033[%d;%dH", index, 0)
	fmt.Print(">")
	for {
		game.UpdateCurKey()
		if game.key_change_flag == 1 {
			game.key_change_flag = 0

			if game.cur_key == 'w' {
				index--
				if index < 0 {
					index = 0
				}
			}

			if game.cur_key == 's' {
				index++
				if index >= len(cur_node.next) {
					index = len(cur_node.next) - 1
				}
			}
			fmt.Printf("\033[%d;%dH", old_index+1, 0)
			fmt.Print(" ")
			fmt.Printf("\033[%d;%dH", index+1, 0)
			fmt.Print(">")

			if game.cur_key == 10 || game.cur_key == 127 {
				if game.cur_key == 10 {
					cur_node = cur_node.next[index]
					index = 0
				}

				if game.cur_key == 127 && cur_node.prev != nil {
					cur_node = cur_node.prev //only if cur_node.prev != nil
					index = 0
				}
				fmt.Printf("\033[%d;%dH", 1, 1)
				for i := 0; i < 10; i++ {
					fmt.Println("                                     ")
				}
				fmt.Printf("\033[%d;%dH", 1, 1)
				for i := 0; i < len(cur_node.next); i++ {
					fmt.Printf("  %s\t%s\n", cur_node.next[i].name, cur_node.next[i].value)
				}
				fmt.Printf("\033[%d;%dH", old_index+1, 0)
				fmt.Print(" ")
				fmt.Printf("\033[%d;%dH", index+1, 0)
				fmt.Print(">")
			}

			option_id = cur_node.id
			value = cur_node.value

			if cur_node.next == nil {
				break
			}
			old_index = index
		}
		time.Sleep(time.Duration(20) * time.Millisecond)
	}

	fmt.Printf("\033[%d;%dH", 0, 0)
	for i := 0; i < 10; i++ {
		fmt.Println("                                     ")
	}
	fmt.Printf("\033[%d;%dH", 0, 0)
	return option_id, value
}

func (game *Game) Run() {
	option_id, value := game.Option()

	is_match, _ := regexp.MatchString("1_0_", option_id)
	if is_match {
		idx := int(option_id[len(option_id)-1]) - 48
		fmt.Println(idx)
		saver := &Saver{}
		keymap := saver.Load()
		keymap[idx] = value
		saver.Save(keymap)
	}

	is_match, _ = regexp.MatchString("0_", option_id)
	if is_match {
		cnt := int(option_id[len(option_id)-1]) - 48 + 1
		for i := 0; i < cnt; i++ {
			game.snakes = append(game.snakes, NewSnake(13+i))
		}
	}

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
			}
		}
		game.showGameStatus(game.snakes)

		time.Sleep(time.Duration(game.interval) * time.Millisecond)
		game.t++
	}
	game.ED()
}

/**
* function takes Snake type as a parameter
* This way, we have access to all of the snake properties
* we can show any other thing related to the snake.E
 */
func (game *Game) showGameStatus(players []*Snake) {
	/*
	   ╔════════════════╗
	   ║  GAME STATUS   ║
	   ╠════════════════╣
	   ║ ** Player 1 ** ║
	   ║ Score: 3       ║
	   ║────────────────║
	   ║ ** Player 2 ** ║
	   ║ Score: 3       ║
	   ╚════════════════╝
	*/

	// players := [2]int{3, 2}
	// header
	screenRow := 8
	fmt.Printf("\033[%d;%dH", screenRow, 95)
	fmt.Printf("╔══════════════════╗")
	screenRow++

	fmt.Printf("\033[%d;%dH", screenRow, 95)
	fmt.Printf("║   GAME STATUS    ║")
	screenRow++

	fmt.Printf("\033[%d;%dH", screenRow, 95)
	fmt.Printf("╠══════════════════╣")
	screenRow++

	// players
	n := len(players)
	for i := 0; i < n; i++ {

		// player header
		fmt.Printf("\033[%d;%dH", screenRow, 95)
		fmt.Printf("║")

		fmt.Printf("\033[%d;%dH", screenRow, 97)
		fmt.Printf("** Player %d **", i)

		fmt.Printf("\033[%d;%dH", screenRow, 114)
		fmt.Printf("║")
		screenRow++

		// player score
		fmt.Printf("\033[%d;%dH", screenRow, 95)
		fmt.Printf("║")

		fmt.Printf("\033[%d;%dH", screenRow, 97)
		fmt.Printf("Score: %d", players[i].body.count)

		fmt.Printf("\033[%d;%dH", screenRow, 114)
		fmt.Printf("║")
		screenRow++

		//seperator
		if n > 1 && i != n-1 {
			fmt.Printf("\033[%d;%dH", screenRow, 95)
			fmt.Printf("║──────────────────║")
			screenRow++
		}
	}

	fmt.Printf("\033[%d;%dH", screenRow, 95)
	fmt.Printf("╚══════════════════╝")
	screenRow++
}
