package dependencies

import (
	"docker-go-project/api/handlers"
	groupHandler "docker-go-project/api/handlers/group"
	teamHandler "docker-go-project/api/handlers/team"
	userHandler "docker-go-project/api/handlers/user"
	"docker-go-project/api/server"
	"docker-go-project/pkg/platform/database"
	groupRepository "docker-go-project/pkg/repository/group"
	teamRepository "docker-go-project/pkg/repository/team"
	userRepository "docker-go-project/pkg/repository/user"
	groupService "docker-go-project/pkg/services/group"
	teamService "docker-go-project/pkg/services/team"
	userService "docker-go-project/pkg/services/user"

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
	_ = Container.Provide(handlers.NewHandlerPing)
	_ = Container.Provide(groupHandler.NewGroupHandler)
	_ = Container.Provide(groupService.NewGroupService)
	_ = Container.Provide(groupRepository.NewGroupRepository)
	_ = Container.Provide(userHandler.NewUserHandler)
	_ = Container.Provide(userService.NewUserService)
	_ = Container.Provide(userRepository.NewUserRepository)
	_ = Container.Provide(teamHandler.NewTeamHandler)
	_ = Container.Provide(teamService.NewTeamService)
	_ = Container.Provide(teamRepository.NewTeamRepository)

	return Container
}
