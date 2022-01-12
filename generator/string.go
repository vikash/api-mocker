package generator

import (
	"math/rand"

	"github.com/mitchellh/mapstructure"
)

type StringGenerator struct {
	MaxLength int
	MinLength int
	Required bool
}

func String(def map[string]interface{}) Generator {
	sg := &StringGenerator{
		MaxLength: 20,
		MinLength: 4,
	}

	mapstructure.Decode(def, &sg )
	return sg
}

func (s *StringGenerator) Generate() interface{}{
	length := rand.Intn(s.MaxLength - s.MinLength) + s.MinLength
	return randSeq(length)
}

var letters = []rune("abcdefghijklmnopqrstuvwxyz ABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}