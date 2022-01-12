package generator

import (
	"math/rand"

	"github.com/mitchellh/mapstructure"
)

type Number struct {
	Min int
	Max int
	Required bool
}

func NewNumber(def map[string]interface{}) Generator {
	num := Number{
		Min: 1,
		Max: 999999999,
	}
	mapstructure.Decode(def, &num )
	return &num
}

func (n *Number) Generate() interface{}{
	return rand.Intn(n.Max - n.Min) + n.Min
}
