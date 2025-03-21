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

	group.GET("/addresses", lc.Fetch)
	group.GET("/address", lc.FetchByID)
	group.POST("/address", lc.Create)
	group.PUT("/address", lc.Update)
	group.DELETE("/address", lc.Delete)
}
