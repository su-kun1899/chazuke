package chazuke

import (
	"bytes"
	"encoding/json"
)

type JSONContainer struct {
	values interface{}
}

func (jc *JSONContainer) Get(key string) *JSONContainer {
	values, ok := jc.values.(map[string]interface{})
	if !ok {
		// TODO errを管理する
		panic("mapじゃない")
	}

	return &JSONContainer{values: values[key]}
}

func (jc *JSONContainer) Value() (string, error) {
	return jc.values.(string), nil
}

func New(jsonVal string) *JSONContainer {
	var buf bytes.Buffer
	buf.Write([]byte(jsonVal))
	dec := json.NewDecoder(&buf)

	var values interface{}
	if err := dec.Decode(&values); err != nil {
		// TODO errを戻り値にする
		panic(err)
	}

	return &JSONContainer{values: values}
}
