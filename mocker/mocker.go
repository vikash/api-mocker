package mocker

import (
	"fmt"
	"github.com/vikash/gofr/pkg/gofr"
	"github.com/vikash/gofr/pkg/gofr/logging"
	"io/ioutil"
	"strings"
)

type Mocker struct{
	entities  map[string]Structure
	server *gofr.App
	logger logging.Logger
}

func NewFromFolder(c *gofr.Context, dirname string) *Mocker{
	m := &Mocker{
		entities: make(map[string]Structure),
		server: gofr.New(),
		logger: c.Logger,
	}

	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		c.Error(err)
	}

	for _, file := range files {
		if !file.IsDir() {
			fullPath := fmt.Sprintf("%s/%s", dirname, file.Name())
			c.Debug("File Found: ", fullPath)
			entityName := strings.TrimSuffix(file.Name(), ".json")
			bytes, err := ioutil.ReadFile(fullPath)

			if err!=nil {
				c.Errorf("Error reading file %s. Error: %s", entityName, err)
			}

			m.entities[entityName] = jsonToStructure(bytes)

		}
	}

	return m
}

func (m *Mocker) Serve() {
	m.addRoutes()
	m.server.Run()
}

func (m *Mocker) addRoutes() {
	for entity, structure := range m.entities {
		// We need a local copy for handler to function properly - common map range bug alert!
		ls := structure
		e := entity

		m.logger.Debug("Adding route for", entity )
		m.server.GET("/"+e, func(c *gofr.Context) (interface{}, error) {
			count := 10
			response := make([]interface{},count)
			for i := 0; i < count; i ++ {
				response[i] = objectForStructure(ls)
			}
			return map[string]interface{}{
				e: response,
				"meta": map[string]int{
					"limit": 10,
					"offset": 0,
				},
			}, nil
		})
	}
}

