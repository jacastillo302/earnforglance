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

func SpecificationAttributeGroupRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewSpecificationAttributeGroupRepository(db, domain.CollectionSpecificationAttributeGroup)
	lc := &controller.SpecificationAttributeGroupController{
		SpecificationAttributeGroupUsecase: usecase.NewSpecificationAttributeGroupUsecase(ur, timeout),
		Env:                                env,
	}

	group.GET("/specification_attribute_groups", lc.Fetch)
	group.GET("/specification_attribute_group", lc.FetchByID)
	group.POST("/specification_attribute_group", lc.Create)
	group.POST("/specification_attribute_groups", lc.CreateMany)
	group.PUT("/specification_attribute_group", lc.Update)
	group.DELETE("/specification_attribute_group", lc.Delete)
}
