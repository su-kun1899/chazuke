package chazuke_test

import (
	"github.com/su-kun1899/chazuke"
	"testing"
)

func TestValue(t *testing.T) {
	json := `
		{
		  "title": "example",
		  "friends":[
		    {"firstName":"Taro", "lastName":"Yamada"}, 
		    {"firstName":"Jiro", "lastName":"Sato"},
		    {"firstName":"Hanako", "lastName":"Tanaka"}
		  ]
		}
	`

	got, err := chazuke.New(json).Get("title").Value()
	if err != nil {
		t.Fatal("unexpected error:", err)
		return
	}

	want := "example"
	if got != want {
		t.Errorf("chazuke.Value() = %v, want %v", got, want)
		return
	}
}
