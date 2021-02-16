package di

import (
	"go.uber.org/dig"
	"rest/db"
	"rest/repository"
)

var Container *dig.Container

func BuildDIContainer() {
	Container = dig.New()
	_ = Container.Provide(db.NewConfig)
	_ = Container.Provide(db.ConnectDatabase)
	_ = Container.Provide(repository.NewAdvertRepository)
}
