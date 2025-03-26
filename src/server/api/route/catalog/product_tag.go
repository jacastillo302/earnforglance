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

func ProductTagRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewProductTagRepository(db, domain.CollectionProductTag)
	lc := &controller.ProductTagController{
		ProductTagUsecase: usecase.NewProductTagUsecase(ur, timeout),
		Env:               env,
	}

	itemGroup := group.Group("/api/v1/catalog")
	itemGroup.GET("/product_tags", lc.Fetch)
	itemGroup.GET("/product_tag", lc.FetchByID)
	itemGroup.POST("/product_tag", lc.Create)
	itemGroup.POST("/product_tags", lc.CreateMany)
	itemGroup.PUT("/product_tag", lc.Update)
	itemGroup.DELETE("/product_tag", lc.Delete)
}
