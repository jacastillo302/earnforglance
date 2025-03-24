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

func PictureRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewPictureRepository(db, domain.CollectionPicture)
	lc := &controller.PictureController{
		PictureUsecase: usecase.NewPictureUsecase(ur, timeout),
		Env:            env,
	}

	group.GET("/pictures", lc.Fetch)
	group.GET("/picture", lc.FetchByID)
	group.POST("/picture", lc.Create)
	group.POST("/pictures", lc.CreateMany)
	group.PUT("/picture", lc.Update)
	group.DELETE("/picture", lc.Delete)
}
