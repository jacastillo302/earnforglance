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

func AddressRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewAddressRepository(db, domain.CollectionAddress)
	lc := &controller.AddressController{
		AddressUsecase: usecase.NewAddressUsecase(ur, timeout),
		Env:            env,
	}

	itemGroup := group.Group("/api/v1/common")
	itemGroup.GET("/addresses", lc.Fetch)
	itemGroup.GET("/address", lc.FetchByID)
	itemGroup.POST("/address", lc.Create)
	itemGroup.POST("/addresses", lc.CreateMany)
	itemGroup.PUT("/address", lc.Update)
	itemGroup.DELETE("/address", lc.Delete)
}
