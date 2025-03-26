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

func PictureHashesRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewPictureHashesRepository(db, domain.CollectionPictureHashes)
	lc := &controller.PictureHashesController{
		PictureHashesUsecase: usecase.NewPictureHashesUsecase(ur, timeout),
		Env:                  env,
	}
	itemGroup := group.Group("/api/v1/media")
	itemGroup.GET("/picture_hashes", lc.Fetch)
	itemGroup.GET("/picture_hash", lc.FetchByID)
	itemGroup.POST("/picture_hash", lc.Create)
	itemGroup.POST("/picture_hashes", lc.CreateMany)
	itemGroup.PUT("/picture_hash", lc.Update)
	itemGroup.DELETE("/picture_hash", lc.Delete)
}
