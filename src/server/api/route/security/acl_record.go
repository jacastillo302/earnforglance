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

func AclRecordRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewAclRecordRepository(db, domain.CollectionAclRecord)
	lc := &controller.AclRecordController{
		AclRecordUsecase: usecase.NewAclRecordUsecase(ur, timeout),
		Env:              env,
	}
	itemGroup := group.Group("/api/v1/security")
	itemGroup.GET("/acl_records", lc.Fetch)
	itemGroup.GET("/acl_record", lc.FetchByID)
	itemGroup.POST("/acl_record", lc.Create)
	itemGroup.POST("/acl_records", lc.CreateMany)
	itemGroup.PUT("/acl_record", lc.Update)
	itemGroup.DELETE("/acl_record", lc.Delete)
}
