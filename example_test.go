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

	container, err := chazuke.New(jsonVal)
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
