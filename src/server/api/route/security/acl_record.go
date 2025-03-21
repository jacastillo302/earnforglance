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

func AclRecordRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewAclRecordRepository(db, domain.CollectionAclRecord)
	lc := &controller.AclRecordController{
		AclRecordUsecase: usecase.NewAclRecordUsecase(ur, timeout),
		Env:              env,
	}

	group.GET("/acl_records", lc.Fetch)
	group.GET("/acl_record", lc.FetchByID)
	group.POST("/acl_record", lc.Create)
	group.PUT("/acl_record", lc.Update)
	group.DELETE("/acl_record", lc.Delete)
}
