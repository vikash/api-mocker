package generator

import (
	"github.com/mitchellh/mapstructure"
)

type Array struct {
	Item      map[string]interface{}
	MaxLength int
	MinLength int
}

func NewArray(def map[string]interface{}) Generator {
	g := &Array{
		Item:      nil,
		MaxLength: 5,
		MinLength: 0,
	}

	mapstructure.Decode(def, &g)
	return g
}

func (a *Array) Generate() interface{} {
	size := randBetween(a.MinLength, a.MaxLength)
	resp := make([]interface{}, size)

	for i := 0; i < size; i++ {
		v, _ := GenerateValueForConfig(a.Item)
		resp[i] = v
	}
	return resp

}
