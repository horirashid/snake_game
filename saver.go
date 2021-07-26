package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

type Saver struct {
	keys []string
}

func (saver *Saver) Save(keys []string) {
	keystr := ""
	/*for _, i := range keys {
			keystr = keystr + i + "\n"
	}*/
	// fmt.Println("len : ", len(keys))
	for i := 0; i < len(keys); i++ {
		keystr += keys[i] + "\n"
	}
	content := []byte(keystr)
	// fmt.Println(keystr)
	err := ioutil.WriteFile("keymap.txt", content, 0644)
	if err != nil {
		panic(err)
	}
}

func (saver *Saver) Load() []string {
	saver.useNewReader("keymap.txt")
	return saver.keys
}

func (saver *Saver) useNewReader(filename string) {
	var count int = 0

	fin, error := os.OpenFile(filename, os.O_RDONLY, 0)
	if error != nil {
		fmt.Println("keymap.txt not found !!!")
		panic(error)
	}
	defer fin.Close()

	/*create a Reader*/
	rd := bufio.NewReader(fin)

	/*read the file and stop when meet err or EOF*/
	for {
		line, err := rd.ReadString('\n')
		if err != nil || err == io.EOF {
			break
		}
		count++
		line = strings.Replace(line, "\f", "", -1)
		line = strings.Replace(line, "\n", "", -1)
		saver.keys = append(saver.keys, line)
	}
}

//初始化
// sa := &Saver{}

//读取
// k := sa.Load()

// 存储
// sa.Save([]string{"qwer", "tyui"})
/*func main() {
	sa := &Saver{}
	k := sa.Load()
	k[0] = "fuck"
	sa.Save(k)
}*/
