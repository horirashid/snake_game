package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
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
	body_char       byte
	speed           int
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
	return game
}

func (game *Game) Waitkey() {
	for {
		_, found := game.input.Inkey()
		if found {
			break
		}
	}
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
	snake_g := NewSnakeByArray([]*Point{
		&Point{20, 10},
		&Point{19, 9},
		&Point{18, 9},
		&Point{17, 9},
		&Point{16, 9},
		&Point{15, 9},
		&Point{14, 9},
		&Point{13, 9},
		&Point{12, 9},
		&Point{11, 10},
		&Point{10, 11},
		&Point{10, 12},
		&Point{10, 13},
		&Point{10, 14},
		&Point{10, 15},
		&Point{10, 16},
		&Point{10, 17},
		&Point{10, 18},
		&Point{10, 19},
		&Point{10, 20},
		&Point{11, 21},
		&Point{12, 22},
		&Point{13, 22},
		&Point{14, 22},
		&Point{15, 22},
		&Point{16, 22},
		&Point{17, 22},
		&Point{18, 22},
		&Point{19, 22},
		&Point{20, 21},
		&Point{20, 20},
		&Point{20, 19},
		&Point{20, 18},
		&Point{20, 17},
		&Point{20, 16},
		&Point{19, 16},
		&Point{18, 16},
		&Point{17, 16},
		&Point{16, 16},
	}, 1)
	snake_g.head = &Point{20, 10}
	snake_g.temp_dir = 'd'
	snake_g.dir = 'd'

	snake_r := NewSnakeByArray([]*Point{
		&Point{38, 22},
		&Point{37, 21},
		&Point{36, 20},
		&Point{35, 19},
		&Point{34, 18},
		&Point{33, 17},
		&Point{32, 16},
		&Point{32, 15},
		&Point{33, 15},
		&Point{34, 15},
		&Point{35, 15},
		&Point{36, 15},
		&Point{37, 14},
		&Point{38, 13},
		&Point{38, 12},
		&Point{38, 11},
		&Point{37, 10},
		&Point{36, 9},
		&Point{35, 9},
		&Point{34, 9},
		&Point{33, 9},
		&Point{32, 9},
		&Point{31, 9},
		&Point{30, 9},
		&Point{29, 9},
		&Point{28, 9},
		&Point{28, 10},
		&Point{28, 11},
		&Point{28, 12},
		&Point{28, 13},
		&Point{28, 14},
		&Point{28, 15},
		&Point{28, 16},
		&Point{28, 17},
		&Point{28, 18},
		&Point{28, 19},
		&Point{28, 20},
		&Point{28, 21},
		&Point{28, 22},
	}, 1)
	snake_r.head = &Point{38, 22}

	snake_o := NewSnakeByArray([]*Point{
		&Point{52, 9},
		&Point{51, 9},
		&Point{50, 9},
		&Point{48, 10},
		&Point{47, 11},
		&Point{46, 12},
		&Point{45, 13},
		&Point{45, 14},
		&Point{45, 15},
		&Point{45, 16},
		&Point{45, 17},
		&Point{45, 18},
		&Point{46, 19},
		&Point{47, 20},
		&Point{48, 21},
		&Point{50, 22},
		&Point{51, 22},
		&Point{52, 22},
		&Point{54, 21},
		&Point{55, 20},
		&Point{56, 19},
		&Point{57, 18},
		&Point{57, 17},
		&Point{57, 16},
		&Point{57, 15},
		&Point{57, 14},
		&Point{57, 13},
		&Point{56, 12},
		&Point{55, 11},
		&Point{54, 10},
	}, 1)
	snake_o.head = &Point{52, 9}

	snake_u := NewSnakeByArray([]*Point{
		&Point{61, 9},
		&Point{61, 10},
		&Point{61, 11},
		&Point{61, 12},
		&Point{61, 13},
		&Point{61, 14},
		&Point{61, 15},
		&Point{61, 16},
		&Point{61, 17},
		&Point{61, 18},
		&Point{61, 19},
		&Point{61, 20},
		&Point{61, 21},
		&Point{61, 22},
		&Point{62, 22},
		&Point{63, 22},
		&Point{64, 22},
		&Point{65, 22},
		&Point{66, 22},
		&Point{67, 22},
		&Point{68, 22},
		&Point{69, 22},
		&Point{70, 22},
		&Point{71, 22},
		&Point{71, 21},
		&Point{71, 20},
		&Point{71, 19},
		&Point{71, 18},
		&Point{71, 17},
		&Point{71, 16},
		&Point{71, 15},
		&Point{71, 14},
		&Point{71, 13},
		&Point{71, 12},
		&Point{71, 11},
		&Point{71, 10},
		&Point{71, 9},
	}, 1)
	snake_u.head = &Point{61, 9}

	snake_p := NewSnakeByArray([]*Point{
		&Point{76, 15},
		&Point{77, 15},
		&Point{78, 15},
		&Point{79, 15},
		&Point{80, 15},
		&Point{81, 15},
		&Point{82, 15},
		&Point{83, 15},
		&Point{84, 14},
		&Point{85, 13},
		&Point{86, 12},
		&Point{85, 11},
		&Point{84, 10},
		&Point{83, 9},
		&Point{82, 9},
		&Point{81, 9},
		&Point{80, 9},
		&Point{79, 9},
		&Point{78, 9},
		&Point{77, 9},
		&Point{76, 9},
		&Point{75, 9},
		&Point{75, 10},
		&Point{75, 11},
		&Point{75, 12},
		&Point{75, 13},
		&Point{75, 14},
		&Point{75, 15},
		&Point{75, 16},
		&Point{75, 17},
		&Point{75, 18},
		&Point{75, 19},
		&Point{75, 20},
		&Point{75, 21},
		&Point{75, 22},
	}, 1)
	snake_p.head = &Point{76, 15}

	snake_5 := NewSnakeByArray([]*Point{
		&Point{101, 9},
		&Point{100, 9},
		&Point{99, 9},
		&Point{98, 9},
		&Point{97, 9},
		&Point{96, 9},
		&Point{95, 9},
		&Point{94, 9},
		&Point{93, 9},
		&Point{93, 10},
		&Point{93, 11},
		&Point{93, 12},
		&Point{93, 13},
		&Point{96, 14},
		&Point{98, 15},
		&Point{100, 16},
		&Point{102, 17},
		&Point{102, 18},
		&Point{102, 19},
		&Point{100, 20},
		&Point{98, 21},
		&Point{96, 22},
		&Point{95, 22},
	}, 1)
	snake_5.head = &Point{101, 9}

	snake_g.Show()
	snake_r.Show()
	snake_o.Show()
	snake_u.Show()
	snake_p.Show()
	snake_5.Show()

	game.Waitkey()

	for i := 0; i < 250; i++ {
		snake_g.Move()
		snake_r.Move()
		snake_o.Move()
		snake_u.Move()
		snake_p.Move()
		snake_5.Move()
		time.Sleep(time.Duration(10) * time.Millisecond)
	}

	/*fmt.Printf("\033[%d;%dH", 1, 1)
	for i := 0; i < 120+2; i++ {
		fmt.Printf("-")
	}
	fmt.Println()
	for i := 0; i < 30+2; i++ {
		fmt.Print("|")
		for j := 0; j < 120; j++ {
			fmt.Print(" ")
		}
		fmt.Println("|")
	}
	for i := 0; i < 120+2; i++ {
		fmt.Printf("-")
	}*/
	fmt.Printf("\033[%d;%dH", 1, 1)
}

func (game *Game) ED() {
	fmt.Printf("\033[%d;%dH", game.ditu.height+5, 1)
	fmt.Println("Quit")
}

func (game *Game) Prepare() {
	game.ditu.Show()
	for i := 0; i < len(game.snakes); i++ {
		game.snakes[i].Show()
	}
	game.ditu.GenerateFood()
}

func (game *Game) Select() (string, string) {
	var option_id string
	var value string
	for {
		option_id, value = game.Option()

		//start game
		if option_id[0] == '0' {
			cnt := int(option_id[len(option_id)-1]) - 48 + 1
			for i := 0; i < cnt; i++ {
				snake := NewSnake(13 + i)
				if i == 0 {
					snake.keymap = "wsad"
				} else if i == 1 {
					snake.keymap = "ikjl"
				}
				game.snakes = append(game.snakes, snake)
			}
			break
		}
	}
	return option_id, value
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
	snake_speed_node := NewNode(setting_node, nil, "snake_speed", "1_3", "")
	saver := &Saver{}
	keymap := saver.GetKeyMap()

	var snake_nodes []*Node
	for i := 0; i < len(keymap); i++ {
		snake_node := NewNode(keymapping_node, nil, "snake"+strconv.Itoa(i), "1_0_"+strconv.Itoa(i), keymap[i])
		snake_nodes = append(snake_nodes, snake_node)
	}

	root.next = append(root.next, gamemode_node)
	root.next = append(root.next, setting_node)
	root.next = append(root.next, history_note)
	gamemode_node.next = append(gamemode_node.next, solo_node)
	gamemode_node.next = append(gamemode_node.next, double_node)
	setting_node.next = append(setting_node.next, keymapping_node)
	setting_node.next = append(setting_node.next, body_char_node)
	setting_node.next = append(setting_node.next, snake_speed_node)
	for i := 0; i < len(snake_nodes); i++ {
		keymapping_node.next = append(keymapping_node.next, snake_nodes[i])
	}

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

				if game.cur_key == 10 || game.cur_key == 'q' {
					if game.cur_key == 10 {
						cur_node = cur_node.next[index]
						index = 0
					}

					if game.cur_key == 'q' && cur_node.prev != nil {
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
		//start game
		if option_id[0] == '0' {
			cnt := int(option_id[len(option_id)-1]) - 48 + 1
			for i := 0; i < cnt; i++ {
				snake := NewSnake(13 + i)
				snake.speed_scale = game.speed
				snake.body_char = game.body_char
				snake.keymap = snake_nodes[i].value
				game.snakes = append(game.snakes, snake)
			}
			break
		} else {
			//0.change keymapping
			is_match, _ := regexp.MatchString("1_0_*", option_id)
			if is_match {
				idx := option_id[4] - '0'
				fmt.Printf("\033[%d;%dH", 1, 1)
				for i := 0; i < 10; i++ {
					fmt.Println("                                     ")
				}
				fmt.Printf("\033[%d;%dH", 1, 1)
				x_idx := 1
				key_mapping := ""
				for {
					game.UpdateCurKey()
					if game.key_change_flag == 1 {
						game.key_change_flag = 0
						if game.cur_key == 10 {
							break
						}
						if x_idx <= 4 {
							fmt.Printf("\033[%d;%dH", 1, x_idx)
							fmt.Printf("%c", game.cur_key)
							key_mapping += string(game.cur_key)
							x_idx++
						}
					}
				}
				saver := &Saver{}
				keymap := saver.GetKeyMap()
				keymap[idx] = key_mapping
				saver.SaveKeyMap(keymap)
				snake_nodes[idx].value = key_mapping
			}

			//1.change body_char
			is_match, _ = regexp.MatchString("1_1", option_id)
			if is_match {
				fmt.Printf("\033[%d;%dH", 1, 1)
				for i := 0; i < 10; i++ {
					fmt.Println("                                     ")
				}
				fmt.Printf("\033[%d;%dH", 1, 1)
				saver := &Saver{}
				body_char := saver.GetBody()
				fmt.Printf("\033[%d;%dH", 1, 1)
				fmt.Printf("%c", body_char)
				fmt.Printf("\033[%d;%dH", 1, 1)
				for {
					game.UpdateCurKey()
					if game.key_change_flag == 1 {
						game.key_change_flag = 0
						if game.cur_key == 10 {
							break
						}
						fmt.Printf("\033[%d;%dH", 1, 1)
						fmt.Printf("%c", game.cur_key)
						body_char = rune(game.cur_key)
					}
				}
				game.body_char = byte(body_char)
				saver.SaveBody(body_char)
			}

			//2.change fps
			is_match, _ = regexp.MatchString("1_2", option_id)
			if is_match {
				fmt.Printf("\033[%d;%dH", 1, 1)
				for i := 0; i < 10; i++ {
					fmt.Println("                                     ")
				}
				fmt.Printf("\033[%d;%dH", 1, 1)
				fmt.Println("change fps")
				game.Waitkey()
			}

			//3.change snake_speed
			is_match, _ = regexp.MatchString("1_3", option_id)
			if is_match {
				fmt.Printf("\033[%d;%dH", 1, 1)
				for i := 0; i < 10; i++ {
					fmt.Println("                                     ")
				}
				fmt.Printf("\033[%d;%dH", 1, 1)

				saver := &Saver{}
				speed := saver.GetSpeed()
				fmt.Print(speed)
				fmt.Printf("\033[%d;%dH", 1, 1)

				x_idx := 1
				num := ""
				for {
					game.UpdateCurKey()
					if game.key_change_flag == 1 {
						game.key_change_flag = 0
						if game.cur_key == 10 {
							break
						}
						if x_idx == 1 {
							fmt.Println("                                     ")
							fmt.Printf("\033[%d;%dH", 1, 1)
						}
						if x_idx < 4 {
							if game.cur_key >= '0' && game.cur_key <= '9' {
								fmt.Printf("\033[%d;%dH", 1, x_idx)
								fmt.Printf("%c", game.cur_key)
								x_idx++
								num += string(game.cur_key)
							}
						}
					}
				}
				game.speed, _ = strconv.Atoi(num)
				saver.SaveSpeed(game.speed)
			}

			//4.see history
			if option_id == "2" {
				fmt.Printf("\033[%d;%dH", 1, 1)
				for i := 0; i < 10; i++ {
					fmt.Println("                                     ")
				}
				fmt.Printf("\033[%d;%dH", 1, 1)
				fmt.Println("see history")
				game.Waitkey()
			}

			//return to prev
			cur_node = cur_node.prev //only if cur_node.prev != nil
			index = 0
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
			old_index = index
		}
	}
	fmt.Printf("\033[%d;%dH", 0, 0)
	for i := 0; i < 10; i++ {
		fmt.Println("                                     ")
	}
	fmt.Printf("\033[%d;%dH", 0, 0)
	return option_id, value
}

func (game *Game) Init() {
	saver := &Saver{}
	game.speed = saver.GetSpeed()
	game.body_char = byte(saver.GetBody())
}

func (game *Game) Run() {
	game.Init()
	game.Option()
	game.OP()
	game.Prepare()
	game.Waitkey()

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
				game.snakes[i].ChangeDirectionByKey(game.cur_key)
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
					//score++
				}

				if game.snakes[i].isEatSelf() {
					fmt.Printf("\033[%d;%dH", game.ditu.height+5, 0)
					fmt.Println("Eat Self!")
					return
				}

				game.snakes[i].Move()

				if game.snakes[i].isHitWall(game.ditu) {
					fmt.Printf("\033[%d;%dH", game.ditu.height+5, 0)
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

	var p SnakePlayers
	highest, err := p.getHighestScore()
	if err != nil {
		fmt.Printf("Error getting the highest score for players. Error output: %s", err)
	}
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

		// player highest score
		fmt.Printf("\033[%d;%dH", screenRow, 95)
		fmt.Printf("║")

		fmt.Printf("\033[%d;%dH", screenRow, 97)
		fmt.Printf("Highest Score: %d", highest[i])

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
