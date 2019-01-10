package chazuke_test

import (
	"reflect"
	"testing"

	"github.com/su-kun1899/chazuke"
)

func TestContainer_Value(t *testing.T) {
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
				t.Errorf("Container.Value() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Container.Value() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainer_Value_WithoutNew(t *testing.T) {
	container := &chazuke.JsonContainer{}

	got, err := container.Value()
	if !(got == "" && err != nil) {
		t.Errorf("Container.Value() got = %v, error = %v", got, err)
		return
	}
}

func TestContainer_NestedValue(t *testing.T) {
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
		name      string
		parentKey string
		childKey  string
		want      string
		wantErr   bool
	}{
		{
			name:      "Get manager's name value",
			parentKey: "manager",
			childKey:  "name",
			want:      "Ernest Valverde",
			wantErr:   false,
		},

		{
			name:      "Get manager's birthDay value",
			parentKey: "manager",
			childKey:  "birthDay",
			want:      "1964-02-09",
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			container, err := chazuke.New(jsonVal)
			if err != nil {
				t.Fatal("unexpected error:", err)
			}

			got, err := container.Get(tt.parentKey).Get(tt.childKey).Value()
			if (err != nil) != tt.wantErr {
				t.Errorf("Container.Value() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Container.Value() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainer_Array(t *testing.T) {
	jsonVal := `
		{
			"cars":[
        		"FIAT 500",
        		"RENAULT KANGOO",
        		"MINI CROSSOVER"
		  	]
		}
	`

	tests := []struct {
		name     string
		arrayKey string
		want     []string
		wantErr  bool
	}{
		{
			name:     "Get cars",
			arrayKey: "cars",
			want:     []string{"FIAT 500", "RENAULT KANGOO", "MINI CROSSOVER"},
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			container, err := chazuke.New(jsonVal)
			if err != nil {
				t.Fatal("unexpected error:", err)
			}

			got, err := container.Get(tt.arrayKey).Array()
			if (err != nil) != tt.wantErr {
				t.Errorf("Container.Array() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if len(got) != len(tt.want) {
				t.Errorf("len(Container.Array()) = %v, want %v", len(got), len(tt.want))
			}

			for i, jc := range got {
				v, err := jc.Value()
				if err != nil {
					t.Fatal("unexpected error:", err)
				}

				if !reflect.DeepEqual(v, tt.want[i]) {
					t.Errorf("Container.Value() = %v, want %v", v, tt.want[i])
				}
			}
		})
	}
}

func TestContainer_Array_NestedValue(t *testing.T) {
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
		name     string
		arrayKey string
		key      string
		want     []string
		wantErr  bool
	}{
		{
			name:     "Get players' name",
			arrayKey: "players",
			key:      "name",
			want:     []string{"Messi", "Coutinho", "Pique"},
			wantErr:  false,
		},
		{
			name:     "Get players' position",
			arrayKey: "players",
			key:      "position",
			want:     []string{"Forward", "Midfielder", "Defender"},
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			container, err := chazuke.New(jsonVal)
			if err != nil {
				t.Fatal("unexpected error:", err)
			}

			got, err := container.Get(tt.arrayKey).Array()
			if (err != nil) != tt.wantErr {
				t.Errorf("Container.Array() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if len(got) != len(tt.want) {
				t.Errorf("len(Container.Array()) = %v, want %v", len(got), len(tt.want))
			}

			for i, jc := range got {
				v, err := jc.Get(tt.key).Value()
				if err != nil {
					t.Fatal("unexpected error:", err)
				}

				if !reflect.DeepEqual(v, tt.want[i]) {
					t.Errorf("Container.Value() = %v, want %v", v, tt.want[i])
				}
			}
		})
	}
}

func TestContainer_JSON(t *testing.T) {
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
		want    string
		wantErr bool
	}{
		{
			name:    "Get manager's JSON",
			key:     "manager",
			want:    `{"name": "Ernest Valverde","birthDay": "1964-02-09"}`,
			wantErr: false,
		},
		{
			name: "Get players' JSON",
			key:  "players",
			want: `[
        		{"name":"Messi", "position":"Forward"}, 
        		{"name":"Coutinho", "position":"Midfielder"},
        		{"name":"Pique", "position":"Defender"}
		  	]`,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			container, err := chazuke.New(jsonVal)
			if err != nil {
				t.Fatal("unexpected error:", err)
			}

			got, err := container.Get(tt.key).JSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("Container.Value() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			wantV, err := chazuke.New(tt.want)
			if err != nil {
				t.Fatal("unexpected error:", err)
			}
			gotV, err := chazuke.New(got)
			if err != nil {
				t.Fatal("unexpected error:", err)
			}

			if !reflect.DeepEqual(gotV, wantV) {
				t.Errorf("Container.JSON() = %v, want %v", gotV, wantV)
			}
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		jsonVal string
	}
	tests := []struct {
		name    string
		args    args
		want    *chazuke.JsonContainer
		wantErr bool
	}{
		{
			name:    "Illegal JSON",
			args:    args{jsonVal: "This is Illegal JSON."},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := chazuke.New(tt.args.jsonVal)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
