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

func SpecificationAttributeOptionRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewSpecificationAttributeOptionRepository(db, domain.CollectionSpecificationAttributeOption)
	lc := &controller.SpecificationAttributeOptionController{
		SpecificationAttributeOptionUsecase: usecase.NewSpecificationAttributeOptionUsecase(ur, timeout),
		Env:                                 env,
	}

	group.GET("/specification_attribute_options", lc.Fetch)
	group.GET("/specification_attribute_option", lc.FetchByID)
	group.POST("/specification_attribute_option", lc.Create)
	group.PUT("/specification_attribute_option", lc.Update)
	group.DELETE("/specification_attribute_option", lc.Delete)
}
