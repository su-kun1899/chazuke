package chazuke

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// TODO interfaceにする
type JsonContainer struct {
	values interface{}
}

func (container *JsonContainer) Get(key string) *JsonContainer {
	values, ok := container.values.(map[string]interface{})
	if !ok {
		// TODO errを管理する
		panic("mapじゃない")
	}

	return &JsonContainer{values: values[key]}
}

func (container *JsonContainer) Value() (string, error) {
	s, ok := container.values.(string)
	if !ok {
		return "", fmt.Errorf("container has illegal value = %v", container.values)
	}

	return s, nil
}

func (container *JsonContainer) Array() ([]*JsonContainer, error) {
	values, ok := container.values.([]interface{})
	if !ok {
		// TODO errを管理する
		panic("Arrayじゃない")
	}

	containers := make([]*JsonContainer, len(values))
	for i, v := range values {
		containers[i] = &JsonContainer{values: v}
	}

	return containers, nil
}

func (container *JsonContainer) JSON() (string, error) {
	b, err := json.Marshal(container.values)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func New(jsonVal string) (*JsonContainer, error) {
	var buf bytes.Buffer
	buf.Write([]byte(jsonVal))
	dec := json.NewDecoder(&buf)

	var values interface{}
	if err := dec.Decode(&values); err != nil {
		return nil, err
	}

	return &JsonContainer{values: values}, nil
}
