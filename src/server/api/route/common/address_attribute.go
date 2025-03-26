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

func AddressAttributeRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewAddressAttributeRepository(db, domain.CollectionAddressAttribute)
	lc := &controller.AddressAttributeController{
		AddressAttributeUsecase: usecase.NewAddressAttributeUsecase(ur, timeout),
		Env:                     env,
	}

	itemGroup := group.Group("/api/v1/common")
	itemGroup.GET("/address_attributes", lc.Fetch)
	itemGroup.GET("/address_attribute", lc.FetchByID)
	itemGroup.POST("/address_attribute", lc.Create)
	itemGroup.POST("/address_attributes", lc.CreateMany)
	itemGroup.PUT("/address_attribute", lc.Update)
	itemGroup.DELETE("/address_attribute", lc.Delete)
}
