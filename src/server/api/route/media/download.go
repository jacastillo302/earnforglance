package route

import (
	"time"

	controller "earnforglance/server/api/controller/media"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/media"

	"earnforglance/server/mongo"
	repository "earnforglance/server/repository/media"
	usecase "earnforglance/server/usecase/media"

	"github.com/gin-gonic/gin"
)

func DownloadRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewDownloadRepository(db, domain.CollectionDownload)
	lc := &controller.DownloadController{
		DownloadUsecase: usecase.NewDownloadUsecase(ur, timeout),
		Env:             env,
	}

	group.GET("/downloads", lc.Fetch)
	group.GET("/download", lc.FetchByID)
	group.POST("/download", lc.Create)
	group.PUT("/download", lc.Update)
	group.DELETE("/download", lc.Delete)
}
