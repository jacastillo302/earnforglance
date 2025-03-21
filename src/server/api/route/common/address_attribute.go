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

	group.GET("/address_attributes", lc.Fetch)
	group.GET("/address_attribute", lc.FetchByID)
	group.POST("/address_attribute", lc.Create)
	group.PUT("/address_attribute", lc.Update)
	group.DELETE("/address_attribute", lc.Delete)
}
