package route

import (
	"time"

	controller "earnforglance/server/api/controller/security"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/security"

	repository "earnforglance/server/repository/security"
	"earnforglance/server/service/data/mongo"
	usecase "earnforglance/server/usecase/security"

	"github.com/gin-gonic/gin"
)

func PermissionRecordRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewPermissionRecordRepository(db, domain.CollectionPermissionRecord)
	lc := &controller.PermissionRecordController{
		PermissionRecordUsecase: usecase.NewPermissionRecordUsecase(ur, timeout),
		Env:                     env,
	}
	itemGroup := group.Group("/api/v1/security")
	itemGroup.GET("/permission_records", lc.Fetch)
	itemGroup.GET("/permission_record", lc.FetchByID)
	itemGroup.POST("/permission_record", lc.Create)
	itemGroup.POST("/permission_records", lc.CreateMany)
	itemGroup.PUT("/permission_record", lc.Update)
	itemGroup.DELETE("/permission_record", lc.Delete)
}
