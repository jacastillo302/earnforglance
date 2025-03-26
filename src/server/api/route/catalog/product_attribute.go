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

func ProductAttributeRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewProductAttributeRepository(db, domain.CollectionProductAttribute)
	lc := &controller.ProductAttributeController{
		ProductAttributeUsecase: usecase.NewProductAttributeUsecase(ur, timeout),
		Env:                     env,
	}

	itemGroup := group.Group("/api/v1/catalog")
	itemGroup.GET("/product_attributes", lc.Fetch)
	itemGroup.GET("/product_attribute", lc.FetchByID)
	itemGroup.POST("/product_attribute", lc.Create)
	itemGroup.POST("/product_attributes", lc.CreateMany)
	itemGroup.PUT("/product_attribute", lc.Update)
	itemGroup.DELETE("/product_attribute", lc.Delete)
}
