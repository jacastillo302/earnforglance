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
	itemGroup := group.Group("/api/v1/schedule_tasks")
	itemGroup.GET("/schedule_tasks", lc.Fetch)
	itemGroup.GET("/schedule_task", lc.FetchByID)
	itemGroup.POST("/schedule_task", lc.Create)
	itemGroup.POST("/schedule_tasks", lc.CreateMany)
	itemGroup.PUT("/schedule_task", lc.Update)
	itemGroup.DELETE("/schedule_task", lc.Delete)
}
