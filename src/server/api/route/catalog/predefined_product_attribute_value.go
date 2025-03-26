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

func PredefinedProductAttributeValueRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewPredefinedProductAttributeValueRepository(db, domain.CollectionPredefinedProductAttributeValue)
	lc := &controller.PredefinedProductAttributeValueController{
		PredefinedProductAttributeValueUsecase: usecase.NewPredefinedProductAttributeValueUsecase(ur, timeout),
		Env:                                    env,
	}

	itemGroup := group.Group("/api/v1/catalog")
	itemGroup.GET("/predefined_product_attribute_values", lc.Fetch)
	itemGroup.GET("/predefined_product_attribute_value", lc.FetchByID)
	itemGroup.POST("/predefined_product_attribute_value", lc.Create)
	itemGroup.POST("/predefined_product_attribute_values", lc.CreateMany)
	itemGroup.PUT("/predefined_product_attribute_value", lc.Update)
	itemGroup.DELETE("/predefined_product_attribute_value", lc.Delete)
}
