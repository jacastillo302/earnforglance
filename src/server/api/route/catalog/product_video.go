package route

import (
	"time"

	controller "earnforglance/server/api/controller/catalog"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/catalog"

	"earnforglance/server/mongo"
	repository "earnforglance/server/repository/catalog"
	usecase "earnforglance/server/usecase/catalog"

	"github.com/gin-gonic/gin"
)

func ProductVideoRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewProductVideoRepository(db, domain.CollectionProductVideo)
	lc := &controller.ProductVideoController{
		ProductVideoUsecase: usecase.NewProductVideoUsecase(ur, timeout),
		Env:                 env,
	}

	group.GET("/product_videos", lc.Fetch)
	group.GET("/product_video", lc.FetchByID)
	group.POST("/product_video", lc.Create)
	group.POST("/product_videos", lc.CreateMany)
	group.PUT("/product_video", lc.Update)
	group.DELETE("/product_video", lc.Delete)
}
