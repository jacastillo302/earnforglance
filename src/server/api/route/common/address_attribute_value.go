package route

import (
	"time"

	controller "earnforglance/server/api/controller/common"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/common"

	repository "earnforglance/server/repository/common"
	"earnforglance/server/service/data/mongo"
	usecase "earnforglance/server/usecase/common"

	"github.com/gin-gonic/gin"
)

func AddressAttributeValueRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewAddressAttributeValueRepository(db, domain.CollectionAddressAttributeValue)
	lc := &controller.AddressAttributeValueController{
		AddressAttributeValueUsecase: usecase.NewAddressAttributeValueUsecase(ur, timeout),
		Env:                          env,
	}

	itemGroup := group.Group("/api/v1/common")
	itemGroup.GET("/address_attribute_values", lc.Fetch)
	itemGroup.GET("/address_attribute_value", lc.FetchByID)
	itemGroup.POST("/address_attribute_value", lc.Create)
	itemGroup.POST("/address_attribute_values", lc.CreateMany)
	itemGroup.PUT("/address_attribute_value", lc.Update)
	itemGroup.DELETE("/address_attribute_value", lc.Delete)
}
