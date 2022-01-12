package models

import (
	"encoding/json"
)

func JSONToStructure(bytes []byte) Structure {
	structure := make(map[string]map[string]interface{})
	json.Unmarshal(bytes, &structure)
	return structure
}
