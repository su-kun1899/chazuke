package chazuke_test

import (
	"fmt"
	"github.com/su-kun1899/chazuke"
)

func ExampleContainer_Value() {
	jsonVal := `
    {
        "team": "FC Barcelona",
        "manager": {
            "name": "Ernest Valverde",
            "birthDay": "1964-02-09"
        }
    }
    `

	container, err := chazuke.FromJSON(jsonVal)
	if err != nil {
		panic(err.Error())
	}

	team, err := container.Get("team").Value()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(team)

	birthDay, err := container.Get("manager").Get("birthDay").Value()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(birthDay)

	// Output:
	// FC Barcelona
	// 1964-02-09
}

func ExampleContainer_Array() {
	jsonVal := `
    {
        "players":[
            {"name":"Messi", "position":"Forward"},
            {"name":"Coutinho", "position":"Midfielder"},
            {"name":"Pique", "position":"Defender"}
        ]
    }
    `

	container, err := chazuke.FromJSON(jsonVal)
	if err != nil {
		panic(err.Error())
	}

	players, err := container.Get("players").Array()
	if err != nil {
		panic(err.Error())
	}

	for _, p := range players {
		name, err := p.Get("name").Value()
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(name)
	}

	// Output:
	// Messi
	// Coutinho
	// Pique
}

func ExampleContainer_JSON() {
	jsonVal := `
    {
        "captain": {
            "name":"Messi", 
            "position":"Forward"
        }
    }
    `

	container, err := chazuke.FromJSON(jsonVal)
	if err != nil {
		panic(err.Error())
	}

	j, err := container.Get("captain").JSON()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(j)
	// Output:
	// {"name":"Messi","position":"Forward"}
}
