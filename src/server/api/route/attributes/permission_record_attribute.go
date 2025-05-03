package route

import (
	"time"

	controller "earnforglance/server/api/controller/attributes"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/attributes"

	repository "earnforglance/server/repository/attributes"
	"earnforglance/server/service/data/mongo"
	usecase "earnforglance/server/usecase/attributes"

	"github.com/gin-gonic/gin"
)

func PermisionRecordAttributeRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewPermisionRecordAttributeRepository(db, domain.CollectionPermisionRecordAttribute)
	lc := &controller.PermisionRecordAttributeController{
		PermisionRecordAttributeUsecase: usecase.NewPermisionRecordAttributeUsecase(ur, timeout),
		Env:                             env,
	}

	itemGroup := group.Group("/api/v1/attributes")
	itemGroup.GET("/permission_record_attributes", lc.Fetch)
	itemGroup.GET("/permission_record_attribute", lc.FetchByID)
	itemGroup.POST("/permission_record_attribute", lc.Create)
	itemGroup.POST("/permission_record_attributes", lc.CreateMany)
	itemGroup.PUT("/permission_record_attribute", lc.Update)
	itemGroup.DELETE("/permission_record_attribute", lc.Delete)
}
