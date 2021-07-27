package main

type Menu struct {
	root *Node
}

func NewMenu() *Menu {
	menu := &Menu{}

	menu.root = NewNode(nil, nil, "root", "", "")
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
}

func (menu *Menu) Select() string {

}
