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

func ProductCategoryRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewProductCategoryRepository(db, domain.CollectionProductCategory)
	lc := &controller.ProductCategoryController{
		ProductCategoryUsecase: usecase.NewProductCategoryUsecase(ur, timeout),
		Env:                    env,
	}

	itemGroup := group.Group("/api/v1/catalog")
	itemGroup.GET("/product_categories", lc.Fetch)
	itemGroup.GET("/product_category", lc.FetchByID)
	itemGroup.POST("/product_category", lc.Create)
	itemGroup.POST("/product_categories", lc.CreateMany)
	itemGroup.PUT("/product_category", lc.Update)
	itemGroup.DELETE("/product_category", lc.Delete)
}
