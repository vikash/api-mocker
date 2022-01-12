package main

import (
	"github.com/vikash/api-mocker/mocker"
	"github.com/vikash/gofr/pkg/gofr"
)

func main() {
	app := gofr.NewCMD()

	app.SubCommand("", func(c *gofr.Context) (interface{}, error) {
		m := mocker.NewFromFolder(c, "./models")
		m.Serve()
		return nil, nil
	})

	app.Run()
}
