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

	group.GET("/related_products", lc.Fetch)
	group.GET("/related_product", lc.FetchByID)
	group.POST("/related_product", lc.Create)
	group.PUT("/related_product", lc.Update)
	group.DELETE("/related_product", lc.Delete)
}
