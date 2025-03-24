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

	group.GET("/state_provinces", lc.Fetch)
	group.GET("/state_province", lc.FetchByID)
	group.POST("/state_province", lc.Create)
	group.POST("/state_provinces", lc.CreateMany)
	group.PUT("/state_province", lc.Update)
	group.DELETE("/state_province", lc.Delete)
}
