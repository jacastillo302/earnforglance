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

	group.GET("/picture_binaries", lc.Fetch)
	group.GET("/picture_binary", lc.FetchByID)
	group.POST("/picture_binary", lc.Create)
	group.POST("/picture_binaries", lc.CreateMany)
	group.PUT("/picture_binary", lc.Update)
	group.DELETE("/picture_binary", lc.Delete)
}
