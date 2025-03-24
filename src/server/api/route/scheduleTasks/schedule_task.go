package route

import (
	"time"

	controller "earnforglance/server/api/controller/scheduleTasks"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/scheduleTasks"

	"earnforglance/server/mongo"
	repository "earnforglance/server/repository/scheduleTasks"
	usecase "earnforglance/server/usecase/scheduleTasks"

	"github.com/gin-gonic/gin"
)

func ScheduleTaskRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewScheduleTaskRepository(db, domain.CollectionScheduleTask)
	lc := &controller.ScheduleTaskController{
		ScheduleTaskUsecase: usecase.NewScheduleTaskUsecase(ur, timeout),
		Env:                 env,
	}

	group.GET("/schedule_tasks", lc.Fetch)
	group.GET("/schedule_task", lc.FetchByID)
	group.POST("/schedule_task", lc.Create)
	group.POST("/schedule_tasks", lc.CreateMany)
	group.PUT("/schedule_task", lc.Update)
	group.DELETE("/schedule_task", lc.Delete)
}
