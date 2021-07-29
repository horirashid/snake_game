package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Saver struct {
	keys  []string
	body  rune
	speed int
}

func (saver *Saver) allLoad() {
	saver.GetKeyMap()
	saver.GetBody()
	saver.GetSpeed()
}

func (saver *Saver) allSave() {
	keystr := ""
	for i := 0; i < len(saver.keys); i++ {
		keystr += saver.keys[i] + "\n"
	}
	keystr += string(saver.body) + "\n"
	keystr += strconv.Itoa(saver.speed) + "\n"

	content := []byte(keystr)
	// fmt.Println(keystr)
	err := ioutil.WriteFile("keymap.txt", content, 0644)
	if err != nil {
		panic(err)
	}
}
func (saver *Saver) SaveKeyMap(keys []string) {
	saver.allLoad()
	saver.keys = keys
	saver.allSave()
}

func (saver *Saver) SaveBody(body rune) {
	saver.allLoad()
	saver.body = body
	saver.allSave()
}

func (saver *Saver) SaveSpeed(speed int) {
	saver.allLoad()
	saver.speed = speed
	saver.allSave()
}

func (saver *Saver) GetKeyMap() []string {
	all := saver.useNewReader("keymap.txt")
	saver.keys = all[:len(all)-2]
	return saver.keys
}

func (saver *Saver) GetBody() rune {
	all := saver.useNewReader("keymap.txt")
	k := []rune(all[len(all)-2])
	saver.body = k[0]
	return saver.body
}

func (saver *Saver) GetSpeed() int {
	all := saver.useNewReader("keymap.txt")
	k := all[len(all)-1]
	i, _ := strconv.Atoi(k)
	saver.speed = i
	return saver.speed
}

func (saver *Saver) useNewReader(filename string) []string {
	var count int = 0

	fin, error := os.OpenFile(filename, os.O_RDONLY, 0)
	if error != nil {
		fmt.Println("keymap.txt not found !!!")
		panic(error)
	}
	defer fin.Close()

	/*create a Reader*/
	rd := bufio.NewReader(fin)
	strs := []string{}
	/*read the file and stop when meet err or EOF*/
	for {
		line, err := rd.ReadString('\n')
		if err != nil || err == io.EOF {
			break
		}
		count++
		line = strings.Replace(line, "\f", "", -1)
		line = strings.Replace(line, "\n", "", -1)
		strs = append(strs, line)
	}
	return strs
}

//初始化
// sa := &Saver{}

//读取
// k := sa.Load()

// 存储
// sa.Save([]string{"qwer", "tyui"})
// func main() {
// 	sa := &Saver{}
// 	sa.SaveBody('@')

// 	// fmt.Println(k)

// }
