package route

import (
	"time"

	controller "earnforglance/server/api/controller/catalog"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/catalog"

	repository "earnforglance/server/repository/catalog"
	"earnforglance/server/service/data/mongo"
	usecase "earnforglance/server/usecase/catalog"

	"github.com/gin-gonic/gin"
)

func ProductVideoRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewProductVideoRepository(db, domain.CollectionProductVideo)
	lc := &controller.ProductVideoController{
		ProductVideoUsecase: usecase.NewProductVideoUsecase(ur, timeout),
		Env:                 env,
	}

	itemGroup := group.Group("/api/v1/catalog")
	itemGroup.GET("/product_videos", lc.Fetch)
	itemGroup.GET("/product_video", lc.FetchByID)
	itemGroup.POST("/product_video", lc.Create)
	itemGroup.POST("/product_videos", lc.CreateMany)
	itemGroup.PUT("/product_video", lc.Update)
	itemGroup.DELETE("/product_video", lc.Delete)
}
