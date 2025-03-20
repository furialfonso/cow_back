package jobs

import (
	"context"
	"fmt"
	"time"

	"shared-wallet-service/infrastructure/config"
	"shared-wallet-service/usecases/user"

	"github.com/go-co-op/gocron"
)

type IJob interface {
	UserCache()
}

type job struct {
	userUseCase user.IUserUseCase
}

func InitLoadCacheHandler(userUseCase user.IUserUseCase) error {
	ctx := context.Background()
	return userUseCase.UserLoad(ctx)
}

func NewJobHandler(userUseCase user.IUserUseCase) IJob {
	return &job{
		userUseCase: userUseCase,
	}
}

func (jb *job) UserCache() {
	nameJob := "users"
	active := config.Get().UBool(fmt.Sprintf("job.%s.active", nameJob), false)
	cronExpression := config.Get().UString(fmt.Sprintf("job.%s.expression", nameJob))
	ctx := context.Background()
	if active {
		fmt.Println("user cache job running")
		j := gocron.NewScheduler(time.UTC)
		_, err := j.Cron(cronExpression).Do(jb.userUseCase.UserLoad, ctx)
		if err != nil {
			fmt.Println("Error creating job", err)
		}
		j.StartAsync()
	}
}
