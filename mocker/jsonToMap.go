package mocker

import (
	"encoding/json"
	"github.com/vikash/api-mocker/generator"
	"strings"
)

type Structure map[string]map[string]interface{}

func jsonToStructure(bytes []byte) Structure{
	structure := make(map[string]map[string]interface{})
	json.Unmarshal(bytes, &structure)
	return structure
}

func objectForStructure(structure Structure) map[string]interface{} {

	object := make(map[string]interface{})

	for name, def := range structure {
		objectType := def["type"].(string)
		if objectType == "" {
			objectType = "string"
		}

		object[name] = getValue(objectType, def)
	}

	return object
}


func getValue(objectType string, def map[string]interface{}) interface{} {
	var g generator.Generator

	switch strings.ToLower(objectType) {
	case "number":
		g = generator.NewNumber(def)
	case "string":
		g = generator.String(def)
	case "image":
		g = generator.NewImage(def)
	default:
		g = generator.String(def)
	}

	return g.Generate()
}


