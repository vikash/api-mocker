package generator

import (
	"math/rand"

	"github.com/mitchellh/mapstructure"
)

type Number struct {
	Min      int
	Max      int
	Required bool
}

func NewNumber(def map[string]interface{}) Generator {
	num := Number{
		Min: 1,
		Max: 999999999,
	}
	mapstructure.Decode(def, &num)
	return &num
}

func (n *Number) Generate() interface{} {
	return randBetween(n.Min, n.Max)
}

func randBetween(min int, max int) int {
	return rand.Intn(max-min) + min
}
