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

func ProductAttributeValueRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewProductAttributeValueRepository(db, domain.CollectionProductAttributeValue)
	lc := &controller.ProductAttributeValueController{
		ProductAttributeValueUsecase: usecase.NewProductAttributeValueUsecase(ur, timeout),
		Env:                          env,
	}

	group.GET("/product_attribute_values", lc.Fetch)
	group.GET("/product_attribute_value", lc.FetchByID)
	group.POST("/product_attribute_value", lc.Create)
	group.POST("/product_attribute_values", lc.CreateMany)
	group.PUT("/product_attribute_value", lc.Update)
	group.DELETE("/product_attribute_value", lc.Delete)
}
