package integers

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdder(t *testing.T) {
	actual := Add(2, 3)
	expected := 5
	assert.Equal(t, expected, actual)
}

func ExampleAdd() {
	actual := Add(2,5)
	fmt.Println(actual)
	// Output: 7
}