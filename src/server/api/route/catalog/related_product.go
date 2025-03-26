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

func RelatedProductRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewRelatedProductRepository(db, domain.CollectionRelatedProduct)
	lc := &controller.RelatedProductController{
		RelatedProductUsecase: usecase.NewRelatedProductUsecase(ur, timeout),
		Env:                   env,
	}

	itemGroup := group.Group("/api/v1/catalog")
	itemGroup.GET("/related_products", lc.Fetch)
	itemGroup.GET("/related_product", lc.FetchByID)
	itemGroup.POST("/related_product", lc.Create)
	itemGroup.POST("/related_products", lc.CreateMany)
	itemGroup.PUT("/related_product", lc.Update)
	itemGroup.DELETE("/related_product", lc.Delete)
}
