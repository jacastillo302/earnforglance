package route

import (
	"time"

	controller "earnforglance/server/api/controller/directory"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/directory"

	"earnforglance/server/mongo"
	repository "earnforglance/server/repository/directory"
	usecase "earnforglance/server/usecase/directory"

	"github.com/gin-gonic/gin"
)

func StateProvinceRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewStateProvinceRepository(db, domain.CollectionStateProvince)
	lc := &controller.StateProvinceController{
		StateProvinceUsecase: usecase.NewStateProvinceUsecase(ur, timeout),
		Env:                  env,
	}

	itemGroup := group.Group("/api/v1/directory")
	itemGroup.GET("/state_provinces", lc.Fetch)
	itemGroup.GET("/state_province", lc.FetchByID)
	itemGroup.POST("/state_province", lc.Create)
	itemGroup.POST("/state_provinces", lc.CreateMany)
	itemGroup.PUT("/state_province", lc.Update)
	itemGroup.DELETE("/state_province", lc.Delete)
}
