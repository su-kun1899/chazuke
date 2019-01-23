package main

import (
	"fmt"
	"github.com/su-kun1899/chazuke"
)

func main() {
	jsonVal := `
	{
		"team": "FC Barcelona",
		"manager": {
			"name": "Ernest Valverde",
  			"birthDay": "1964-02-09"
		},
		"players":[
			{"name":"Messi", "position":"Forward"},
			{"name":"Coutinho", "position":"Midfielder"},
			{"name":"Pique", "position":"Defender"}
		]
	}
	`

	container, _ := chazuke.FromJSON(jsonVal)

	team, _ := container.Get("team").Value()
	fmt.Println(team) // Should be "FC Barcelona"

	birthDay, _ := container.Get("manager").Get("birthDay").Value()
	fmt.Println(birthDay) // Should be "1964-02-09"

	players, _ := container.Get("players").Array()
	for _, p := range players {
		name, _ := p.Get("name").Value()
		fmt.Println(name) // Should be each player's name
	}

	json, _ := players[0].JSON()
	fmt.Println(json) // Should be {"name":"Messi","position":"Forward"}
}
