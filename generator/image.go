package generator

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"math/rand"
	"strconv"
	"strings"
)

type Image struct {
	MinWidth int
	MaxWidth int
	AspectRatio string
	Size string
}

func NewImage(def map[string]interface{}) Generator {
	i := Image{
		MinWidth: 160,
		MaxWidth: 1000,
		AspectRatio: "1:1",
	}
	mapstructure.Decode(def, &i )
	return &i
}

func (i *Image) getAspectNumbers() (int, int) {

	aspects := strings.Split(i.AspectRatio, ":")
	if len(aspects)!=2 {
		return 1,1
	}

	var a, b int
	var err error

	a, err = strconv.Atoi(aspects[0])
	if err!=nil || a == 0 {
		a = 1
	}

	b, err = strconv.Atoi(aspects[1])
	if err!=nil || b == 0 {
		b = 1
	}

	return a,b
}

func (i *Image)Generate() interface{} {
	width := rand.Intn(i.MaxWidth - i.MinWidth) + i.MinWidth

	a,b := i.getAspectNumbers()
	height := (b * width) / a

	return fmt.Sprintf("https://via.placeholder.com/%dX%d", width, height)
}