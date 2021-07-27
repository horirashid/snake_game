package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)
// Players struct used to hold the array of players
type Players struct {
	Players []Player `json:"players"`
}

// Player struct used to read the attributes of the player json file
type Player struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	HighestScore int    `json:"highest_score"`
}

func (p *Players) readJson() error {
	jsonFile, err := os.Open("players.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	// read Json file as byte array
	byteValue, _ := ioutil.ReadAll(jsonFile)

	err = json.Unmarshal(byteValue, &p)

	return err
}

func (p *Players) getHighestScore() ([]int, error) {

	p.readJson()
	highestScore := []int {0,0}
	for i := 0; i < len(p.Players); i++ {
		//fmt.Println("Player Id: " + users.Players[i].Id)
		//fmt.Println("Player Name: " + users.Players[i].Name)
		//fmt.Println("Player HighestScore: " + strconv.Itoa(users.Players[i].HighestScore))
		highestScore[i] = p.Players[i].HighestScore
	}

	return  highestScore, nil
}

func (p *Players) setHighestScore(playerId int, newScore int) error {
	p.readJson()
	p.Players[playerId].HighestScore = newScore

	// Convert golang object back to byte
	byteValue, err := json.MarshalIndent(p, "", "\t")
	if err != nil {
		return err
	}
	// Write back to file
	err = ioutil.WriteFile("players.json", byteValue, 0644)
	return err
}
func main() {
	var p Players
	err := p.setHighestScore(0, 6)
	if err != nil {
		fmt.Println(nil)
	}
	x, err := p.getHighestScore()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(x)
	}
}