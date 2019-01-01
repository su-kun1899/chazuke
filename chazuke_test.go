package chazuke_test

import (
	"reflect"
	"testing"

	"github.com/su-kun1899/chazuke"
)

func TestJSONContainer_Value(t *testing.T) {
	jsonVal := `
		{
		  "title": "example",
		  "description": "this is example.",
		  "friends":[
		    {"firstName":"Taro", "lastName":"Yamada"}, 
		    {"firstName":"Jiro", "lastName":"Sato"},
		    {"firstName":"Hanako", "lastName":"Tanaka"}
		  ]
		}
	`

	tests := []struct {
		name    string
		key     string
		want    string
		wantErr bool
	}{
		{
			name:    "Get title value",
			key:     "title",
			want:    "example",
			wantErr: false,
		},
		{
			name:    "Get description value",
			key:     "description",
			want:    "this is example.",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			container, err := chazuke.New(jsonVal)
			if err != nil {
				t.Fatal("unexpected error:", err)
			}

			got, err := container.Get(tt.key).Value()
			if (err != nil) != tt.wantErr {
				t.Errorf("JSONContainer.Value() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("JSONContainer.Value() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJSONContainer_Values(t *testing.T) {
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

	tests := []struct {
		name    string
		key     string
		want    map[string]interface{}
		wantErr bool
	}{
		{
			name: "Get manager values",
			key:  "manager",
			want: map[string]interface{}{
				"name":     "Ernest Valverde",
				"birthDay": "1964-02-09",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			container, err := chazuke.New(jsonVal)
			if err != nil {
				t.Fatal("unexpected error:", err)
			}

			got, err := container.Get(tt.key).Values()
			if (err != nil) != tt.wantErr {
				t.Errorf("JSONContainer.Values() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JSONContainer.Values() = %v, want %v", got, tt.want)
			}
		})
	}
}
