package generator

type CustomType struct {
	itemType string
}

func NewCustomType(itemType string) Generator {
	g := &CustomType{itemType: itemType}
	return g
}

func (t *CustomType) Generate() interface{} {
	return nil
	//objectType := t.ItemType
	//return GenerateValueForConfig(objectType, a.Item)
}
