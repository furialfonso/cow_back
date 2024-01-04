package dependencies

import (
	"cow_back/api/handlers"
	groupHandler "cow_back/api/handlers/group"
	teamHandler "cow_back/api/handlers/team"
	"cow_back/api/jobs"
	"cow_back/api/server"
	"cow_back/pkg/platform/cache"
	"cow_back/pkg/platform/database"
	"cow_back/pkg/platform/restful"
	groupRepository "cow_back/pkg/repository/group"
	teamRepository "cow_back/pkg/repository/team"
	groupService "cow_back/pkg/services/group"
	teamService "cow_back/pkg/services/team"
	"cow_back/pkg/services/user"

	"go.uber.org/dig"
)

var countErrors int = 0

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
