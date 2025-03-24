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

func VideoRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewVideoRepository(db, domain.CollectionVideo)
	lc := &controller.VideoController{
		VideoUsecase: usecase.NewVideoUsecase(ur, timeout),
		Env:          env,
	}

	group.GET("/videos", lc.Fetch)
	group.GET("/video", lc.FetchByID)
	group.POST("/video", lc.Create)
	group.POST("/videos", lc.CreateMany)
	group.PUT("/video", lc.Update)
	group.DELETE("/video", lc.Delete)
}
