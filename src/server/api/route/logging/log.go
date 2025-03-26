package route

import (
	"time"

	controller "earnforglance/server/api/controller/logging"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/logging"

	repository "earnforglance/server/repository/logging"
	"earnforglance/server/service/data/mongo"
	usecase "earnforglance/server/usecase/logging"

	"github.com/gin-gonic/gin"
)

func LogRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewLogRepository(db, domain.CollectionLog)
	lc := &controller.LogController{
		LogUsecase: usecase.NewLogUsecase(ur, timeout),
		Env:        env,
	}
	itemGroup := group.Group("/api/v1/logging")
	itemGroup.GET("/logs", lc.Fetch)
	itemGroup.GET("/log", lc.FetchByID)
	itemGroup.POST("/log", lc.Create)
	itemGroup.POST("/logs", lc.CreateMany)
	itemGroup.PUT("/log", lc.Update)
	itemGroup.DELETE("/log", lc.Delete)
}
