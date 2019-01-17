// Package chazuke provides some utilities for JSON.
// It works without declaration of the struct.
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

type errContainer struct {
	err error
}

func (container *jsonContainer) Get(key string) Container {
	values, ok := container.values.(map[string]interface{})
	if !ok || values[key] == nil {
		return &errContainer{err: fmt.Errorf("json doesn't have key = %v", key)}
	}

	return &jsonContainer{values: values[key]}
}

func (container *jsonContainer) Value() (string, error) {
	return container.values.(string), nil
}

func (container *jsonContainer) Array() ([]Container, error) {
	values, ok := container.values.([]interface{})
	if !ok {
		return nil, fmt.Errorf("value is not array")
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
		// untested
		return "", err
	}

	return string(b), nil
}

func (container *errContainer) Get(key string) Container {
	return container
}

func (container *errContainer) Value() (string, error) {
	return "", container.err
}

func (container *errContainer) Array() ([]Container, error) {
	return nil, container.err
}

func (container *errContainer) JSON() (string, error) {
	return "", container.err
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
