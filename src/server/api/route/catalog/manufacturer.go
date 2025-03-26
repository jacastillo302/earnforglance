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

	itemGroup := group.Group("/api/v1/catalog")
	itemGroup.GET("/manufacturers", lc.Fetch)
	itemGroup.GET("/manufacturer", lc.FetchByID)
	itemGroup.POST("/manufacturer", lc.Create)
	itemGroup.POST("/manufacturers", lc.CreateMany)
	itemGroup.PUT("/manufacturer", lc.Update)
	itemGroup.DELETE("/manufacturer", lc.Delete)
}
