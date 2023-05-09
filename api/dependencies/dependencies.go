package dependencies

import (
	"docker-go-project/api/handlers"
	"docker-go-project/api/server"
	"docker-go-project/pkg/services"

	"go.uber.org/dig"
)

type Dependencies struct {
}

func BuildDependencies() *dig.Container {
	Container := dig.New()
	_ = Container.Provide(server.New)
	_ = Container.Provide(server.NewRouter)
	_ = Container.Provide(handlers.NewHandlerPing)
	_ = Container.Provide(handlers.NewAirHandler)
	_ = Container.Provide(services.NewAirService)

	return Container
}
