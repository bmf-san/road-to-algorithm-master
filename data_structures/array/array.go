package main

import (
	"errors"
	"fmt"
)

// A Array is array implemented by slice.
type Array struct {
	data   []string
	length int // Keep a array memory size
}

// insert insert a data to array.
func (a *Array) insert(index int, value string) error {
	if a.length == int(cap(a.data)) {
		return errors.New("a array is full")
	}

	if index != a.length && index >= a.length {
		return errors.New("out of index range")
	}

	// shift data
	for i := a.length; i > index; i-- {
		a.data[i] = a.data[i-1]
	}

	// insert a value to target index
	a.data[index] = value

	// update the length
	a.length++

	return nil
}

// delete delete a target data by index.
func (a *Array) delete(index int) (string, error) {
	if index >= a.length {
		return "", errors.New("out of index range")
	}

	// target value for deleting
	v := a.data[index]

	for i := index; i < a.length-1; i++ {
		a.data[i] = a.data[i+1]
	}

	// unset
	a.data[a.length-1] = ""

	// update the length
	a.length--

	return v, nil
}

// get get a target data by index.
func (a *Array) get(index int) (string, error) {
	if index >= a.length {
		return "", errors.New("out of index range")
	}

	// random access
	return a.data[index], nil
}

func main() {
	a := &Array{
		data:   make([]string, 10, 10),
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
		if err := a.insert(c.index, c.value); err != nil {
			fmt.Printf("index: %v value: %v is error. %v\n", c.index, c.value, err)
		}
	}

	if s, err := a.delete(2); err != nil {
		fmt.Printf("index: 0 is error. %v\n", err)
	} else {
		fmt.Printf("%v is deleted.", s)
	}

	if r, err := a.get(0); err != nil {
		fmt.Printf("index: 0 is error. %v", err)
	} else {
		fmt.Printf("%v", r)
	}
}
