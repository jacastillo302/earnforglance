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

func ActivityLogTypeRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewActivityLogTypeRepository(db, domain.CollectionActivityLogType)
	lc := &controller.ActivityLogTypeController{
		ActivityLogTypeUsecase: usecase.NewActivityLogTypeUsecase(ur, timeout),
		Env:                    env,
	}

	group.GET("/activity_log_types", lc.Fetch)
	group.GET("/activity_log_type", lc.FetchByID)
	group.POST("/activity_log_type", lc.Create)
	group.PUT("/activity_log_type", lc.Update)
	group.DELETE("/activity_log_type", lc.Delete)
}
