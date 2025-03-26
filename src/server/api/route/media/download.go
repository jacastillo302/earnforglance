package route

import (
	"time"

	controller "earnforglance/server/api/controller/media"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/media"

	repository "earnforglance/server/repository/media"
	"earnforglance/server/service/data/mongo"
	usecase "earnforglance/server/usecase/media"

	"github.com/gin-gonic/gin"
)

func DownloadRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewDownloadRepository(db, domain.CollectionDownload)
	lc := &controller.DownloadController{
		DownloadUsecase: usecase.NewDownloadUsecase(ur, timeout),
		Env:             env,
	}
	itemGroup := group.Group("/api/v1/media")
	itemGroup.GET("/downloads", lc.Fetch)
	itemGroup.GET("/download", lc.FetchByID)
	itemGroup.POST("/download", lc.Create)
	itemGroup.POST("/downloads", lc.CreateMany)
	itemGroup.PUT("/download", lc.Update)
	itemGroup.DELETE("/download", lc.Delete)
}
