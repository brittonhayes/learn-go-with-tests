package summing

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSum(t *testing.T) {
	type suite struct {
		name  string
		input []int
		expected  int
	}

	cases := []suite{
		{name: "sum slice of 5 integers", input: []int{1, 2, 3, 4, 5}, expected: 15},
		{name: "sum slice of 3 integers", input: []int{1, 2, 3}, expected: 6},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := Sum(c.input)
			assert.Equal(t, c.expected, actual)
		})
	}
}

func TestSumAll(t *testing.T) {
	type suite struct {
		name  string
		input [][]int
		expected  []int
	}

	cases := []suite{
		{name: "sum slice of slice of ints", input: [][]int{{1, 2, 3}, {1, 2}}, expected: []int{6, 3}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := SumAll(c.input...)
			assert.Equal(t, c.expected, actual)
		})
	}
}
