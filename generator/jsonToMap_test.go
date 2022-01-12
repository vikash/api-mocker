package generator

import "testing"

func TestJSONToObject(t *testing.T) {
	json := `{
	  "id": {
		"type": "number",
		"min": 100,
		"required": true
	  },
      "count": {
		"type": "number",
		"max": 20,
		"required": true
	  },
      "friends": {
		"type": "number"
      }
	}`

	s := jsonToStructure([]byte(json))
	o := ObjectForStructure(s)

	if _, ok := o["id"]; !ok {
		t.Error("Map was expected to have id but not found.")
	}
}
