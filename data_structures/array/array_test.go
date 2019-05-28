package main

import (
	"reflect"
	"testing"
)

func createArray() (*Array, error) {
	actual := &Array{
		data:   make([]string, 3, 3),
		length: 0,
	}

	cases := []struct {
		index int
		value string
	}{
		{
			index: 0,
			value: "foo",
		},
		{
			index: 1,
			value: "bar",
		},
		{
			index: 2,
			value: "foobar",
		},
	}

	for _, c := range cases {
		if err := actual.insert(c.index, c.value); err != nil {
			return nil, err
		}
	}

	return actual, nil
}
func TestInsert(t *testing.T) {
	actual, err := createArray()
	if err != nil {
		t.Errorf("%v", err)
	}

	expected := &Array{
		data:   []string{"foo", "bar", "foobar"},
		length: 3,
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual:%v expected:%v\n", actual, expected)
	}
}

func TestDelete(t *testing.T) {
	actual, err := createArray()
	if err != nil {
		t.Errorf("%v", err)
	}

	s, err := actual.delete(1)
	if s == "" && err != nil {
		t.Errorf("%v", err)
	}

	expected := &Array{
		data:   []string{"foo", "foobar", ""},
		length: 2,
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual:%v expected:%v\n", actual, expected)
	}
}

func TestGet(t *testing.T) {
	actual, err := createArray()
	if err != nil {
		t.Errorf("%v", err)
	}

	s, err := actual.get(1)

	if s != "bar" {
		t.Errorf("%v is not equal to bar", s)
	}
}
