package route

import (
	"time"

	"earnforglance/server/api/controller"
	"earnforglance/server/bootstrap"
	"earnforglance/server/domain"
	"earnforglance/server/mongo"
	"earnforglance/server/repository"
	"earnforglance/server/usecase"

	"github.com/gin-gonic/gin"
)

func NewTaskRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	tr := repository.NewTaskRepository(db, domain.CollectionTask)
	tc := &controller.TaskController{
		TaskUsecase: usecase.NewTaskUsecase(tr, timeout),
	}
	group.GET("/task", tc.Fetch)
	group.POST("/task", tc.Create)
}
