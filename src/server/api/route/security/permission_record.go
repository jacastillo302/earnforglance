package route

import (
	"time"

	controller "earnforglance/server/api/controller/security"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/security"

	"earnforglance/server/mongo"
	repository "earnforglance/server/repository/security"
	usecase "earnforglance/server/usecase/security"

	"github.com/gin-gonic/gin"
)

func PermissionRecordRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewPermissionRecordRepository(db, domain.CollectionPermissionRecord)
	lc := &controller.PermissionRecordController{
		PermissionRecordUsecase: usecase.NewPermissionRecordUsecase(ur, timeout),
		Env:                     env,
	}

	group.GET("/permission_records", lc.Fetch)
	group.GET("/permission_record", lc.FetchByID)
	group.POST("/permission_record", lc.Create)
	group.PUT("/permission_record", lc.Update)
	group.DELETE("/permission_record", lc.Delete)
}
