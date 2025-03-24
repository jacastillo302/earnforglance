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

	group.GET("/product_specification_attributes", lc.Fetch)
	group.GET("/product_specification_attribute", lc.FetchByID)
	group.POST("/product_specification_attribute", lc.Create)
	group.POST("/product_specification_attributes", lc.CreateMany)
	group.PUT("/product_specification_attribute", lc.Update)
	group.DELETE("/product_specification_attribute", lc.Delete)
}
