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

func PredefinedProductAttributeValueRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewPredefinedProductAttributeValueRepository(db, domain.CollectionPredefinedProductAttributeValue)
	lc := &controller.PredefinedProductAttributeValueController{
		PredefinedProductAttributeValueUsecase: usecase.NewPredefinedProductAttributeValueUsecase(ur, timeout),
		Env:                                    env,
	}

	group.GET("/predefined_product_attribute_values", lc.Fetch)
	group.GET("/predefined_product_attribute_value", lc.FetchByID)
	group.POST("/predefined_product_attribute_value", lc.Create)
	group.PUT("/predefined_product_attribute_value", lc.Update)
	group.DELETE("/predefined_product_attribute_value", lc.Delete)
}
