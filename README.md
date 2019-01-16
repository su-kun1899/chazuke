# chazuke

[![CircleCI](https://circleci.com/gh/su-kun1899/chazuke.svg?style=svg)](https://circleci.com/gh/su-kun1899/chazuke)

chazuke is a tiny Go library that provides some utilities for JSON.

## Installation

```sh
$ go get github.com/su-kun1899/chazuke
```

## Usage

### Getting Value

```go
jsonVal := `
{
    "team": "FC Barcelona",
    "manager": {
        "name": "Ernest Valverde",
        "birthDay": "1964-02-09"
    }
}
`

container, _ := chazuke.New(jsonVal)

team, _ := container.Get("team").Value()
fmt.Println(team) // Should be "FC Barcelona"

birthDay, _ := container.Get("manager").Get("birthDay").Value()
fmt.Println(birthDay) // Should be "1964-02-09"
```

### With Array Value

```go
jsonVal := `
{
	"players":[
		{"name":"Messi", "position":"Forward"},
		{"name":"Coutinho", "position":"Midfielder"},
		{"name":"Pique", "position":"Defender"}
	]
}
`

container, _ := chazuke.New(jsonVal)

players, _ := container.Get("players").Array()
for _, p := range players {
	name, _ := p.Get("name").Value()
	fmt.Println(name) // Should be each player's name
}
```

## License

MIT

## Author

Koji Sudo (a.k.a su-kun1899)
