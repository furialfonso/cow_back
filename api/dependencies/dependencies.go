package dependencies

import (
	"docker-go-project/api/handlers"
	groupHandler "docker-go-project/api/handlers/group"
	userHandler "docker-go-project/api/handlers/user"
	"docker-go-project/api/server"
	"docker-go-project/pkg/platform/database"
	groupRpository "docker-go-project/pkg/repository/group"
	groupService "docker-go-project/pkg/services/group"

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
	_ = Container.Provide(groupRpository.NewGroupRepository)
	_ = Container.Provide(handlers.NewHandlerPing)
	_ = Container.Provide(groupHandler.NewGroupHandler)
	_ = Container.Provide(groupService.NewGroupService)
	_ = Container.Provide(userHandler.NewUserHandler)

	return Container
}
