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

func PictureBinaryRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewPictureBinaryRepository(db, domain.CollectionPictureBinary)
	lc := &controller.PictureBinaryController{
		PictureBinaryUsecase: usecase.NewPictureBinaryUsecase(ur, timeout),
		Env:                  env,
	}
	itemGroup := group.Group("/api/v1/media")
	itemGroup.GET("/picture_binaries", lc.Fetch)
	itemGroup.GET("/picture_binary", lc.FetchByID)
	itemGroup.POST("/picture_binary", lc.Create)
	itemGroup.POST("/picture_binaries", lc.CreateMany)
	itemGroup.PUT("/picture_binary", lc.Update)
	itemGroup.DELETE("/picture_binary", lc.Delete)
}
