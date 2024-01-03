package dependencies

import (
	"docker-go-project/api/handlers"
	groupHandler "docker-go-project/api/handlers/group"
	teamHandler "docker-go-project/api/handlers/team"
	"docker-go-project/api/jobs"
	"docker-go-project/api/server"
	"docker-go-project/pkg/platform/cache"
	"docker-go-project/pkg/platform/database"
	"docker-go-project/pkg/platform/restful"
	groupRepository "docker-go-project/pkg/repository/group"
	teamRepository "docker-go-project/pkg/repository/team"
	groupService "docker-go-project/pkg/services/group"
	teamService "docker-go-project/pkg/services/team"
	"docker-go-project/pkg/services/user"

	"go.uber.org/dig"
)

var (
	countErrors int = 0
)

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
	_ = Container.Provide(teamHandler.NewTeamHandler)
	_ = Container.Provide(teamService.NewTeamService)
	_ = Container.Provide(teamRepository.NewTeamRepository)
	_ = Container.Provide(jobs.NewJobHandler)
	_ = Container.Provide(user.NewUserService)
	_ = Container.Provide(cache.NewCache)
	_ = Container.Provide(restful.NewRestfulService)
	if err := Container.Invoke(jobs.InitLoadCacheHandler); err != nil {
		countErrors++
		if countErrors == 3 {
			panic(err)
		}
		BuildDependencies()
	}

	return Container
}
