package route

import (
	"time"

	controller "earnforglance/server/api/controller/logging"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/logging"

	"earnforglance/server/mongo"
	repository "earnforglance/server/repository/logging"
	usecase "earnforglance/server/usecase/logging"

	"github.com/gin-gonic/gin"
)

func LogRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewLogRepository(db, domain.CollectionLog)
	lc := &controller.LogController{
		LogUsecase: usecase.NewLogUsecase(ur, timeout),
		Env:        env,
	}

	group.GET("/logs", lc.Fetch)
	group.GET("/log", lc.FetchByID)
	group.POST("/log", lc.Create)
	group.PUT("/log", lc.Update)
	group.DELETE("/log", lc.Delete)
}
