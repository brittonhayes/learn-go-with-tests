package geometry

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArea(t *testing.T) {
	areaTests := []struct {
		name     string
		shape    Shape
		expected float64
	}{
		{"rectangle", Rectangle{12.0, 6.0}, 72.0},
		{"circle", Circle{10, 12}, 314.1592653589793},
		{"triangle", Triangle{10.0, 20.0, 10.0}, 200},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.shape.Area()
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestPerimeter(t *testing.T) {
	perimeterTests := []struct {
		name     string
		shape    Shape
		expected float64
	}{
		{"rectangle", Rectangle{12.0, 6.0}, 36},
		{"circle", Circle{10, 12}, 37.69911184307752},
		{"triangle", Triangle{10.0, 20.0, 10.0}, 40},
	}

	for _, tt := range perimeterTests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.shape.Perimeter()
			assert.Equal(t, tt.expected, actual)
		})
	}

}


func ExampleCircle_Area() {
	circle := Circle{10,12}
	area := circle.Area()
	fmt.Println(area)
	// Output: 314.1592653589793
}