package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// SnakePlayers struct used to hold the array of players
type SnakePlayers struct {
	Players []Player `json:"players"`
}

// Player struct used to read the attributes of the player json file
type Player struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	HighestScore int    `json:"highest_score"`
}

func (p *SnakePlayers) readJson() error {
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

func (p *SnakePlayers) getHighestScore() ([]int, error) {

	err := p.readJson()
	if err != nil {
		return nil, err
	}
	highestScore := make([]int, len(p.Players))
	for i := 0; i < len(p.Players); i++ {
		//fmt.Println("Player Id: " + users.SnakePlayers[i].Id)
		//fmt.Println("Player Name: " + users.SnakePlayers[i].Name)
		//fmt.Println("Player HighestScore: " + strconv.Itoa(users.SnakePlayers[i].HighestScore))
		highestScore[i] = p.Players[i].HighestScore
	}

	return highestScore, nil
}

func (p *SnakePlayers) setHighestScore(playerId int, newScore int) error {
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

/*func main() {
	var p SnakePlayers
	err := p.setHighestScore(0, 6)
	if err != nil {
		fmt.Println(err)
	}
	x, err := p.getHighestScore()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(x)
	}
}*/
