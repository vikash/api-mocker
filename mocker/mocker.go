package mocker

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/vikash/api-mocker/generator"

	"github.com/vikash/api-mocker/models"

	"github.com/vikash/gofr/pkg/gofr"
	"github.com/vikash/gofr/pkg/gofr/logging"
)

type Mocker struct {
	store  models.ModelStore
	server *gofr.App
	logger logging.Logger
}

func NewFromFolder(c *gofr.Context, dirname string) *Mocker {
	m := &Mocker{
		store:  models.NewModelStore(),
		server: gofr.New(),
		logger: c.Logger,
	}

	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		c.Error(err)
	}

	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".json") {
			fullPath := fmt.Sprintf("%s/%s", dirname, file.Name())
			c.Debug("File Found: ", fullPath)
			entityName := strings.TrimSuffix(file.Name(), ".json")
			bytes, err := ioutil.ReadFile(fullPath)

			if err != nil {
				c.Errorf("Error reading file %s. Error: %s", entityName, err)
			}

			m.store.AddModel(models.ModelName(entityName), models.JSONToStructure(bytes))
		}
	}

	return m
}

func (m *Mocker) Serve() {
	m.addRoutes()
	m.server.Run()
}

func (m *Mocker) addRoutes() {
	for entity, _ := range m.store.GetModels() {
		// We need a local copy for handler to function properly - common map range bug alert!
		e := entity

		m.logger.Debug("Adding route for", entity)
		m.server.GET("/"+string(e), func(c *gofr.Context) (interface{}, error) {
			count := 10
			response := make([]interface{}, count)
			for i := 0; i < count; i++ {
				response[i] = generator.GenerateObject(e, m.store)
			}
			return map[models.ModelName]interface{}{
				e: response,
				"meta": map[string]int{
					"limit":  10,
					"offset": 0,
				},
			}, nil
		})
	}
}
