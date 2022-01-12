package generator

import (
	"encoding/json"
	"strings"
)

type Structure map[string]map[string]interface{}

func JSONToStructure(bytes []byte) Structure {
	structure := make(map[string]map[string]interface{})
	json.Unmarshal(bytes, &structure)
	return structure
}

func ObjectForStructure(structure Structure) map[string]interface{} {

	object := make(map[string]interface{})

	for name, def := range structure {
		objectType := def["type"].(string)
		if objectType == "" {
			objectType = "string"
		}

		object[name] = GetValue(objectType, def)
	}

	return object
}

func GetValue(objectType string, def map[string]interface{}) interface{} {
	var g Generator

	switch strings.ToLower(objectType) {
	case "number":
		g = NewNumber(def)
	case "string":
		g = String(def)
	case "image":
		g = NewImage(def)
	case "array":
		g = NewArray(def)
	default:
		g = String(def)
	}

	return g.Generate()
}
