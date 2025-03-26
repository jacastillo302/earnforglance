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

func ProductSpecificationAttributeRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewProductSpecificationAttributeRepository(db, domain.CollectionProductSpecificationAttribute)
	lc := &controller.ProductSpecificationAttributeController{
		ProductSpecificationAttributeUsecase: usecase.NewProductSpecificationAttributeUsecase(ur, timeout),
		Env:                                  env,
	}

	itemGroup := group.Group("/api/v1/catalog")
	itemGroup.GET("/product_specification_attributes", lc.Fetch)
	itemGroup.GET("/product_specification_attribute", lc.FetchByID)
	itemGroup.POST("/product_specification_attribute", lc.Create)
	itemGroup.POST("/product_specification_attributes", lc.CreateMany)
	itemGroup.PUT("/product_specification_attribute", lc.Update)
	itemGroup.DELETE("/product_specification_attribute", lc.Delete)
}
