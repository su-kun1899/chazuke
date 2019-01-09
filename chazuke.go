package chazuke

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type Container struct {
	values interface{}
}

func (container *Container) Get(key string) *Container {
	values, ok := container.values.(map[string]interface{})
	if !ok {
		// TODO errを管理する
		panic("mapじゃない")
	}

	return &Container{values: values[key]}
}

func (container *Container) Value() (string, error) {
	s, ok := container.values.(string)
	if !ok {
		return "", fmt.Errorf("Container has illegal value = %v.\nUse chazuke.New()", container.values)
	}

	return s, nil
}

func (container *Container) Array() ([]*Container, error) {
	values, ok := container.values.([]interface{})
	if !ok {
		// TODO errを管理する
		panic("Arrayじゃない")
	}

	containers := make([]*Container, len(values))
	for i, v := range values {
		containers[i] = &Container{values: v}
	}

	return containers, nil
}

func (container *Container) JSON() (string, error) {
	b, err := json.Marshal(container.values)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func New(jsonVal string) (*Container, error) {
	var buf bytes.Buffer
	buf.Write([]byte(jsonVal))
	dec := json.NewDecoder(&buf)

	var values interface{}
	if err := dec.Decode(&values); err != nil {
		return nil, err
	}

	return &Container{values: values}, nil
}
