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

func ProductAttributeMappingRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewProductAttributeMappingRepository(db, domain.CollectionProductAttributeMapping)
	lc := &controller.ProductAttributeMappingController{
		ProductAttributeMappingUsecase: usecase.NewProductAttributeMappingUsecase(ur, timeout),
		Env:                            env,
	}

	itemGroup := group.Group("/api/v1/catalog")
	itemGroup.GET("/product_attribute_mappings", lc.Fetch)
	itemGroup.GET("/product_attribute_mapping", lc.FetchByID)
	itemGroup.POST("/product_attribute_mapping", lc.Create)
	itemGroup.POST("/product_attribute_mappings", lc.CreateMany)
	itemGroup.PUT("/product_attribute_mapping", lc.Update)
	itemGroup.DELETE("/product_attribute_mapping", lc.Delete)
}
