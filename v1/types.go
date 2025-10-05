package config

import "encoding/json"

type JSON[T any] struct {
	Value T
}

func (v *JSON[T]) Decode(value string) error {
	return json.Unmarshal([]byte(value), &v.Value)
}
