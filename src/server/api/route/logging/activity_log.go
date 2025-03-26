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

func ActivityLogRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewActivityLogRepository(db, domain.CollectionActivityLog)
	lc := &controller.ActivityLogController{
		ActivityLogUsecase: usecase.NewActivityLogUsecase(ur, timeout),
		Env:                env,
	}
	itemGroup := group.Group("/api/v1/logging")
	itemGroup.GET("/activity_logs", lc.Fetch)
	itemGroup.GET("/activity_log", lc.FetchByID)
	itemGroup.POST("/activity_log", lc.Create)
	itemGroup.POST("/activity_logs", lc.CreateMany)
	itemGroup.PUT("/activity_log", lc.Update)
	itemGroup.DELETE("/activity_log", lc.Delete)
}
