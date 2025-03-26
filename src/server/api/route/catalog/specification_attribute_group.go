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

func SpecificationAttributeGroupRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewSpecificationAttributeGroupRepository(db, domain.CollectionSpecificationAttributeGroup)
	lc := &controller.SpecificationAttributeGroupController{
		SpecificationAttributeGroupUsecase: usecase.NewSpecificationAttributeGroupUsecase(ur, timeout),
		Env:                                env,
	}

	itemGroup := group.Group("/api/v1/catalog")
	itemGroup.GET("/specification_attribute_groups", lc.Fetch)
	itemGroup.GET("/specification_attribute_group", lc.FetchByID)
	itemGroup.POST("/specification_attribute_group", lc.Create)
	itemGroup.POST("/specification_attribute_groups", lc.CreateMany)
	itemGroup.PUT("/specification_attribute_group", lc.Update)
	itemGroup.DELETE("/specification_attribute_group", lc.Delete)
}
