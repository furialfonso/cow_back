package jobs

import (
	"context"
	"fmt"
	"time"

	"cow_back/pkg/config"
	"cow_back/pkg/services/user"

	"github.com/go-co-op/gocron"
)

type IJob interface {
	UserCache()
}

type job struct {
	userService user.IUserService
}

func InitLoadCacheHandler(userService user.IUserService) error {
	ctx := context.Background()
	return userService.UserLoad(ctx)
}

func NewJobHandler(userService user.IUserService) IJob {
	return &job{
		userService: userService,
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
		_, err := j.Cron(cronExpression).Do(jb.userService.UserLoad, ctx)
		if err != nil {
			fmt.Println("Error creating job", err)
		}
		j.StartAsync()
	}
}
