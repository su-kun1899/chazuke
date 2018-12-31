package chazuke_test

import (
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
