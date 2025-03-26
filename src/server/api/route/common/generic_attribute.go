package route

import (
	"time"

	controller "earnforglance/server/api/controller/common"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/common"

	repository "earnforglance/server/repository/common"
	"earnforglance/server/service/data/mongo"
	usecase "earnforglance/server/usecase/common"

	"github.com/gin-gonic/gin"
)

func GenericAttributeRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewGenericAttributeRepository(db, domain.CollectionGenericAttribute)
	lc := &controller.GenericAttributeController{
		GenericAttributeUsecase: usecase.NewGenericAttributeUsecase(ur, timeout),
		Env:                     env,
	}

	itemGroup := group.Group("/api/v1/common")
	itemGroup.GET("/generic_attributes", lc.Fetch)
	itemGroup.GET("/generic_attribute", lc.FetchByID)
	itemGroup.POST("/generic_attribute", lc.Create)
	itemGroup.POST("/generic_attributes", lc.CreateMany)
	itemGroup.PUT("/generic_attribute", lc.Update)
	itemGroup.DELETE("/generic_attribute", lc.Delete)
}
