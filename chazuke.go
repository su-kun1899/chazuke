package chazuke

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type Container interface {
	Get(key string) Container
	Value() (string, error)
	Array() ([]Container, error)
	JSON() (string, error)
}

type jsonContainer struct {
	values interface{}
}

func (container *jsonContainer) Get(key string) Container {
	values, ok := container.values.(map[string]interface{})
	if !ok {
		// TODO errの中身が分かるようにしたい
		return &jsonContainer{values: nil}
	}

	return &jsonContainer{values: values[key]}
}

func (container *jsonContainer) Value() (string, error) {
	s, ok := container.values.(string)
	if !ok {
		// TODO ここ通れるケースあるかな？
		return "", fmt.Errorf("container has illegal value = %v", container.values)
	}

	return s, nil
}

func (container *jsonContainer) Array() ([]Container, error) {
	values, ok := container.values.([]interface{})
	if !ok {
		// TODO errを管理する
		panic("Arrayじゃない")
	}

	containers := make([]Container, len(values))
	for i, v := range values {
		containers[i] = &jsonContainer{values: v}
	}

	return containers, nil
}

func (container *jsonContainer) JSON() (string, error) {
	b, err := json.Marshal(container.values)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func New(jsonVal string) (Container, error) {
	var buf bytes.Buffer
	buf.Write([]byte(jsonVal))
	dec := json.NewDecoder(&buf)

	var values interface{}
	if err := dec.Decode(&values); err != nil {
		return nil, err
	}

	return &jsonContainer{values: values}, nil
}
