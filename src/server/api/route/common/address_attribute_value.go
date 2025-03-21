package route

import (
	"time"

	controller "earnforglance/server/api/controller/common"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/common"

	"earnforglance/server/mongo"
	repository "earnforglance/server/repository/common"
	usecase "earnforglance/server/usecase/common"

	"github.com/gin-gonic/gin"
)

func AddressAttributeValueRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewAddressAttributeValueRepository(db, domain.CollectionAddressAttributeValue)
	lc := &controller.AddressAttributeValueController{
		AddressAttributeValueUsecase: usecase.NewAddressAttributeValueUsecase(ur, timeout),
		Env:                          env,
	}

	group.GET("/address_attribute_values", lc.Fetch)
	group.GET("/address_attribute_value", lc.FetchByID)
	group.POST("/address_attribute_value", lc.Create)
	group.PUT("/address_attribute_value", lc.Update)
	group.DELETE("address_attribute_value", lc.Delete)
}
