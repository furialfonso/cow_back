package dependencies

import (
	repositoryBudget "shared-wallet-service/infrastructure/repositories/budget"
	repositoryCache "shared-wallet-service/infrastructure/repositories/cache"
	repositoryKeycloak "shared-wallet-service/infrastructure/repositories/keycloak"
	repositoryTeam "shared-wallet-service/infrastructure/repositories/team"
	"shared-wallet-service/interfaces/handlers"
	handlerBudget "shared-wallet-service/interfaces/handlers/budget"
	handlerTeam "shared-wallet-service/interfaces/handlers/team"
	useCaseBudget "shared-wallet-service/usecases/budget"

	useCaseTeam "shared-wallet-service/usecases/team"

	"shared-wallet-service/api/server"
	"shared-wallet-service/infrastructure/cache"
	"shared-wallet-service/infrastructure/database"
	"shared-wallet-service/infrastructure/external/keycloak"
	"shared-wallet-service/infrastructure/external/restful"
	"shared-wallet-service/infrastructure/jobs"
	useCaseUser "shared-wallet-service/usecases/user"

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
	_ = Container.Provide(handlerBudget.NewBudgetHandler)
	_ = Container.Provide(useCaseBudget.NewBudgetUseCase)
	_ = Container.Provide(repositoryBudget.NewBudgetRepository)

	_ = Container.Provide(handlerTeam.NewTeamHandler)
	_ = Container.Provide(useCaseTeam.NewTeamUseCase)
	_ = Container.Provide(repositoryTeam.NewTeamRepository)

	_ = Container.Provide(jobs.NewJobHandler)
	_ = Container.Provide(useCaseUser.NewUserUseCase)

	_ = Container.Provide(repositoryKeycloak.NewKeycloakRepository)
	_ = Container.Provide(keycloak.NewKeycloakClient)

	_ = Container.Provide(repositoryCache.NewCacheRepository)
	_ = Container.Provide(cache.NewCacheClient)

	_ = Container.Provide(restful.NewRestClient)

	if err := Container.Invoke(jobs.InitLoadCacheHandler); err != nil {
		countErrors++
		if countErrors == 3 {
			panic(err)
		}
		BuildDependencies()
	}

	return Container
}
