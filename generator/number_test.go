package generator

import (
	"testing"
)

func TestNumberRangeWithIntegers(t *testing.T) {
	min, max := 100, 1000

	// Run the same thing 1000 times to reduce probability of error.
	for i := 0 ; i <=1000 ; i ++ {
		generator := NewNumber(map[string]interface{}{
			"min": min,
			"max": max,
		})
		n := generator.Generate().(int)

		if n < min || n > max {
			t.Errorf("expected between %d and %d. got %d", min, max, n)
		}
	}
}