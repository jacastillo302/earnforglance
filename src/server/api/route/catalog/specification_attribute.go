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

func SpecificationAttributeRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewSpecificationAttributeRepository(db, domain.CollectionSpecificationAttribute)
	lc := &controller.SpecificationAttributeController{
		SpecificationAttributeUsecase: usecase.NewSpecificationAttributeUsecase(ur, timeout),
		Env:                           env,
	}

	group.GET("/specification_attributes", lc.Fetch)
	group.GET("/specification_attribute", lc.FetchByID)
	group.POST("/specification_attribute", lc.Create)
	group.POST("/specification_attributes", lc.CreateMany)
	group.PUT("/specification_attribute", lc.Update)
	group.DELETE("/specification_attribute", lc.Delete)
}
