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

func ManufacturerRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewManufacturerRepository(db, domain.CollectionManufacturer)
	lc := &controller.ManufacturerController{
		ManufacturerUsecase: usecase.NewManufacturerUsecase(ur, timeout),
		Env:                 env,
	}

	group.GET("/manufacturers", lc.Fetch)
	group.GET("/manufacturer", lc.FetchByID)
	group.POST("/manufacturer", lc.Create)
	group.PUT("/manufacturer", lc.Update)
	group.DELETE("/manufacturer", lc.Delete)
}
