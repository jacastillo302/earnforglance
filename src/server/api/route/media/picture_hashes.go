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

	group.GET("/picture_hashes", lc.Fetch)
	group.GET("/picture_hash", lc.FetchByID)
	group.POST("/picture_hash", lc.Create)
	group.POST("/picture_hashes", lc.CreateMany)
	group.PUT("/picture_hash", lc.Update)
	group.DELETE("/picture_hash", lc.Delete)
}
