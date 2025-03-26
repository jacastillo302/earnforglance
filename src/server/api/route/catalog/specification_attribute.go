package route

import (
	"time"

	controller "earnforglance/server/api/controller/catalog"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/catalog"

	repository "earnforglance/server/repository/catalog"
	"earnforglance/server/service/data/mongo"
	usecase "earnforglance/server/usecase/catalog"

	"github.com/gin-gonic/gin"
)

func SpecificationAttributeRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewSpecificationAttributeRepository(db, domain.CollectionSpecificationAttribute)
	lc := &controller.SpecificationAttributeController{
		SpecificationAttributeUsecase: usecase.NewSpecificationAttributeUsecase(ur, timeout),
		Env:                           env,
	}

	itemGroup := group.Group("/api/v1/catalog")
	itemGroup.GET("/specification_attributes", lc.Fetch)
	itemGroup.GET("/specification_attribute", lc.FetchByID)
	itemGroup.POST("/specification_attribute", lc.Create)
	itemGroup.POST("/specification_attributes", lc.CreateMany)
	itemGroup.PUT("/specification_attribute", lc.Update)
	itemGroup.DELETE("/specification_attribute", lc.Delete)
}
