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
	s, ok := jc.values.(string)
	if !ok {
		// TODO errを管理する
		panic("stringじゃない")
	}

	return s, nil
}
func (jc *JSONContainer) Array() ([]*JSONContainer, error) {
	values, ok := jc.values.([]interface{})
	if !ok {
		// TODO errを管理する
		panic("Arrayじゃない")
	}

	containers := make([]*JSONContainer, len(values))
	for i, v := range values {
		containers[i] = &JSONContainer{values: v}
	}

	return containers, nil
}

func New(jsonVal string) (*JSONContainer, error) {
	var buf bytes.Buffer
	buf.Write([]byte(jsonVal))
	dec := json.NewDecoder(&buf)

	var values interface{}
	if err := dec.Decode(&values); err != nil {
		return nil, err
	}

	return &JSONContainer{values: values}, nil
}
