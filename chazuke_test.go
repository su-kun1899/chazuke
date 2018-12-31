package chazuke_test

import (
	"testing"

	"github.com/su-kun1899/chazuke"
)

func TestJSONContainer_Value(t *testing.T) {
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

	tests := []struct {
		name    string
		cz      *chazuke.JSONContainer
		want    string
		wantErr bool
	}{
		{
			name:    "Get title value",
			cz:      chazuke.New(json).Get("title"),
			want:    "example",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.cz.Value()
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
