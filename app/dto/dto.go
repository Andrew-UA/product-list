package dto

import "encoding/json"

type DTO interface {
}

type Nullable[T any] struct {
	value *T
	isSet bool
}

func NewNullable[T any](value T) Nullable[T] {
	return Nullable[T]{
		value: &value,
		isSet: true,
	}
}

func (n *Nullable[T]) Value() (*T, bool) {
	if n.isSet {
		return n.value, true
	}

	return nil, false
}

func (n *Nullable[T]) SetValue(value T) {
	n.value = &value
	n.isSet = true
}

func (n *Nullable[T]) SetNull() {
	n.isSet = true
	n.value = nil
}

func (n *Nullable[T]) Unset() {
	n.value = nil
	n.isSet = false
}

func (n *Nullable[T]) MarshalJSON() ([]byte, error) {
	if !n.isSet {
		return []byte("null"), nil
	}
	return json.Marshal(n.value)
}

func (n *Nullable[T]) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		n.value = nil
		n.isSet = true

		return nil
	}

	var v T
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	n.value = &v
	n.isSet = true

	return nil
}
