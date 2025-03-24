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

	group.GET("/gdpr_logs", lc.Fetch)
	group.GET("/gdpr_log", lc.FetchByID)
	group.POST("/gdpr_log", lc.Create)
	group.POST("/gdpr_logs", lc.CreateMany)
	group.PUT("/gdpr_log", lc.Update)
	group.DELETE("/gdpr_log", lc.Delete)
}
