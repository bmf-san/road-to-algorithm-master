package main

import "testing"

func TestGCD(t *testing.T) {
	cases := []struct {
		a        int
		b        int
		expected int
	}{
		{
			a:        1,
			b:        2,
			expected: 1,
		},
		{
			a:        5,
			b:        5,
			expected: 5,
		},
		{
			a:        24,
			b:        50,
			expected: 2,
		},
		{
			a:        630,
			b:        300,
			expected: 30,
		},
	}

	for _, c := range cases {
		actual := gcd(c.a, c.b)

		if actual != c.expected {
			t.Errorf("actual:%v expected:%v", actual, c.expected)
		}
	}
}
