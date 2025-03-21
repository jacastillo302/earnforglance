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

func PermissionRecordCustomerRoleMappingRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewPermissionRecordCustomerRoleMappingRepository(db, domain.CollectionPermissionRecordCustomerRoleMapping)
	lc := &controller.PermissionRecordCustomerRoleMappingController{
		PermissionRecordCustomerRoleMappingUsecase: usecase.NewPermissionRecordCustomerRoleMappingUsecase(ur, timeout),
		Env: env,
	}

	group.GET("/permission_record_customer_role_mappings", lc.Fetch)
	group.GET("/permission_record_customer_role_mapping", lc.FetchByID)
	group.POST("/permission_record_customer_role_mapping", lc.Create)
	group.PUT("/permission_record_customer_role_mapping", lc.Update)
	group.DELETE("/permission_record_customer_role_mapping", lc.Delete)
}
