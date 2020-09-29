package main

import (
	"github.com/mpludowski/bWljaGFsLXBsdWRvd3NraQo/app"
	"github.com/mpludowski/bWljaGFsLXBsdWRvd3NraQo/model"
	"github.com/mpludowski/bWljaGFsLXBsdWRvd3NraQo/worker"
)

func main() {
	model.Db = app.ConnectDb()
	defer app.CloseDb(model.Db)

	go worker.Start()

	a := app.NewApp("bWljaGFsLXBsdWRvd3NraQo=")
	a.Run(":8080")
}
