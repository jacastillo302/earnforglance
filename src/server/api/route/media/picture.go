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

func PictureRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewPictureRepository(db, domain.CollectionPicture)
	lc := &controller.PictureController{
		PictureUsecase: usecase.NewPictureUsecase(ur, timeout),
		Env:            env,
	}
	itemGroup := group.Group("/api/v1/media")
	itemGroup.GET("/pictures", lc.Fetch)
	itemGroup.GET("/picture", lc.FetchByID)
	itemGroup.POST("/picture", lc.Create)
	itemGroup.POST("/pictures", lc.CreateMany)
	itemGroup.PUT("/picture", lc.Update)
	itemGroup.DELETE("/picture", lc.Delete)
}
