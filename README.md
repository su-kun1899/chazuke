# chazuke

[![CircleCI](https://circleci.com/gh/su-kun1899/chazuke.svg?style=svg)](https://circleci.com/gh/su-kun1899/chazuke)
[![codecov](https://codecov.io/gh/su-kun1899/chazuke/branch/master/graph/badge.svg)](https://codecov.io/gh/su-kun1899/chazuke)
[![GoDoc](https://godoc.org/github.com/su-kun1899/chazuke?status.svg)](http://godoc.org/github.com/su-kun1899/chazuke)
[![Go Report Card](https://goreportcard.com/badge/github.com/su-kun1899/chazuke)](https://goreportcard.com/report/github.com/su-kun1899/chazuke)

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

container, _ := chazuke.FromJSON(jsonVal)

team, _ := container.Get("team").Value()
fmt.Println(team) // Should be "FC Barcelona"

birthDay, _ := container.Get("manager").Get("birthDay").Value()
fmt.Println(birthDay) // Should be "1964-02-09"
```

### Array Value

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

container, _ := chazuke.FromJSON(jsonVal)

players, _ := container.Get("players").Array()
for _, p := range players {
	name, _ := p.Get("name").Value()
	fmt.Println(name) // Should be each player's name
}
```

### Partial JSON

```go
jsonVal := `
{
    "captain": {
        "name":"Messi", 
        "position":"Forward"
    }
}
`

container, _ := chazuke.FromJSON(jsonVal)

j, _ := container.Get("captain").JSON()
fmt.Println(j) // Should be {"name":"Messi","position":"Forward"}
```

## License

MIT

## Author

Koji Sudo (a.k.a su-kun1899)
