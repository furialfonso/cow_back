package dependencies

import (
	"docker-go-project/api/handlers"
	"docker-go-project/api/server"
	"docker-go-project/pkg/platform/database"
	"docker-go-project/pkg/repository"
	"docker-go-project/pkg/services"

	"go.uber.org/dig"
)

type Dependencies struct {
}

func BuildDependencies() *dig.Container {
	Container := dig.New()
	_ = Container.Provide(server.New)
	_ = Container.Provide(server.NewRouter)
	_ = Container.Provide(func() database.IDataBase {
		return database.NewDataBase("mysql")
	})
	_ = Container.Provide(repository.NewRepository)
	_ = Container.Provide(handlers.NewHandlerPing)
	_ = Container.Provide(handlers.NewGroupHandler)
	_ = Container.Provide(services.NewGroupService)

	return Container
}
