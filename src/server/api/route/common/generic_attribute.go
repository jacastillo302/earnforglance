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

func GenericAttributeRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewGenericAttributeRepository(db, domain.CollectionGenericAttribute)
	lc := &controller.GenericAttributeController{
		GenericAttributeUsecase: usecase.NewGenericAttributeUsecase(ur, timeout),
		Env:                     env,
	}

	group.GET("/generic_attributes", lc.Fetch)
	group.GET("/generic_attribute", lc.FetchByID)
	group.POST("/generic_attribute", lc.Create)
	group.POST("/generic_attributes", lc.CreateMany)
	group.PUT("/generic_attribute", lc.Update)
	group.DELETE("/generic_attribute", lc.Delete)
}
