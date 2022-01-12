package models

type Structure map[string]map[string]interface{}

func (s Structure) GetDefinitionForField(fieldName string) map[string]interface{} {
	return s[fieldName]

}

type ModelName string

type modelStore struct {
	store map[ModelName]Structure
}

type ModelStore interface {
	AddModel(model ModelName, str Structure)
	StructureForModel(model ModelName) Structure
	GetModels() map[ModelName]Structure
}

func NewModelStore() ModelStore {
	return &modelStore{store: make(map[ModelName]Structure)}
}

func (m modelStore) AddModel(model ModelName, str Structure) {
	m.store[model] = str
}

func (m modelStore) StructureForModel(model ModelName) Structure {
	return m.store[model]
}

func (m modelStore) GetModels() map[ModelName]Structure {
	return m.store
}
