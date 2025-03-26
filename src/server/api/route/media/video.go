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

func VideoRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewVideoRepository(db, domain.CollectionVideo)
	lc := &controller.VideoController{
		VideoUsecase: usecase.NewVideoUsecase(ur, timeout),
		Env:          env,
	}
	itemGroup := group.Group("/api/v1/media")
	itemGroup.GET("/videos", lc.Fetch)
	itemGroup.GET("/video", lc.FetchByID)
	itemGroup.POST("/video", lc.Create)
	itemGroup.POST("/videos", lc.CreateMany)
	itemGroup.PUT("/video", lc.Update)
	itemGroup.DELETE("/video", lc.Delete)
}
