package route

import (
	"time"

	controller "earnforglance/server/api/controller/gdpr"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/gdpr"

	"earnforglance/server/mongo"
	repository "earnforglance/server/repository/gdpr"
	usecase "earnforglance/server/usecase/gdpr"

	"github.com/gin-gonic/gin"
)

func GdprLogRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewGdprLogRepository(db, domain.CollectionGdprLog)
	lc := &controller.GdprLogController{
		GdprLogUsecase: usecase.NewGdprLogUsecase(ur, timeout),
		Env:            env,
	}
	itemGroup := group.Group("/api/v1/gdpr")
	itemGroup.GET("/gdpr_logs", lc.Fetch)
	itemGroup.GET("/gdpr_log", lc.FetchByID)
	itemGroup.POST("/gdpr_log", lc.Create)
	itemGroup.POST("/gdpr_logs", lc.CreateMany)
	itemGroup.PUT("/gdpr_log", lc.Update)
	itemGroup.DELETE("/gdpr_log", lc.Delete)
}
