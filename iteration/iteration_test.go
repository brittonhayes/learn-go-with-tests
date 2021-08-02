package iteration

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func TestRepeat(t *testing.T) {
	t.Run("Repeat 5 times", func(t *testing.T) {
		count := 5
		repeated := Repeat("a", count)
		expected := "aaaaa"

		assert.Equal(t, expected, repeated)
	})
}

func ExampleRepeat() {
	repeatCount := 5
	result := Repeat("A", repeatCount)
	fmt.Println(result)
	// Output: AAAAA
}