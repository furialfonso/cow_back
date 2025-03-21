package dependencies

import (
	"context"
	"log"
	"time"

	"shared-wallet-service/infrastructure/database/connections"
	iread "shared-wallet-service/infrastructure/database/interfaces/read"
	iwrite "shared-wallet-service/infrastructure/database/interfaces/write"
	repositoryBudget "shared-wallet-service/infrastructure/repositories/budget"
	repositoryUser "shared-wallet-service/infrastructure/repositories/user"

	repositoryTeam "shared-wallet-service/infrastructure/repositories/team"
	"shared-wallet-service/interfaces/handlers"
	handlerBudget "shared-wallet-service/interfaces/handlers/budget"
	handlerTeam "shared-wallet-service/interfaces/handlers/team"
	useCaseBudget "shared-wallet-service/usecases/budget"

	useCaseTeam "shared-wallet-service/usecases/team"

	"shared-wallet-service/api/server"
	"shared-wallet-service/infrastructure/cache"

	"shared-wallet-service/infrastructure/external/keycloak"
	"shared-wallet-service/infrastructure/external/restful"
	"shared-wallet-service/infrastructure/jobs"
	useCaseUser "shared-wallet-service/usecases/user"

	"go.uber.org/dig"
)

func BuildDependencies() *dig.Container {
	container := dig.New()
	registerProviders(container, getProviders())
	ctx := context.Background()
	initializeCacheHandler(ctx, container)
	return container
}

func getProviders() []any {
	return []any{
		server.New,
		server.NewRouter,
		func() iread.IReadDataBase {
			return connections.NewReadDataBase("mysql")
		},
		func() iwrite.IWriteDataBase {
			return connections.NewWriteDataBase("mysql")
		},
		handlers.NewHandlerPing,
		handlerBudget.NewBudgetHandler,
		useCaseBudget.NewBudgetUseCase,
		repositoryBudget.NewBudgetRepository,
		handlerTeam.NewTeamHandler,
		useCaseTeam.NewTeamUseCase,
		repositoryTeam.NewTeamRepository,
		jobs.NewJobHandler,
		useCaseUser.NewUserUseCase,
		repositoryUser.NewKeycloakRepository,
		keycloak.NewKeycloakClient,
		repositoryUser.NewCacheRepository,
		cache.NewCacheClient,
		restful.NewRestClient,
	}
}

func registerProviders(container *dig.Container, providers []any) {
	for _, provider := range providers {
		if err := container.Provide(provider); err != nil {
			log.Fatalf("Critical error providing dependency: %v", err)
		}
	}
}

func initializeCacheHandler(ctx context.Context, container *dig.Container) {
	for retries := 0; retries < 3; retries++ {
		if err := container.Invoke(func(userUseCase useCaseUser.IUserUseCase) error {
			return initializeUserCache(ctx, userUseCase)
		}); err != nil {
			log.Printf("Error invoking InitLoadCacheHandler (attempt %d): %v", retries+1, err)
			if retries == 2 {
				log.Fatalf("Critical error initializing cache handler after 3 attempts: %v", err)
			}
			time.Sleep(2 * time.Second)
			continue
		}
		break
	}
}

func initializeUserCache(ctx context.Context, userUseCase useCaseUser.IUserUseCase) error {
	return userUseCase.UserLoad(ctx)
}
