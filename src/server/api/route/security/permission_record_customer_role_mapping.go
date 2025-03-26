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
	itemGroup := group.Group("/api/v1/security")
	itemGroup.GET("/permission_record_customer_role_mappings", lc.Fetch)
	itemGroup.GET("/permission_record_customer_role_mapping", lc.FetchByID)
	itemGroup.POST("/permission_record_customer_role_mapping", lc.Create)
	itemGroup.POST("/permission_record_customer_role_mappings", lc.CreateMany)
	itemGroup.PUT("/permission_record_customer_role_mapping", lc.Update)
	itemGroup.DELETE("/permission_record_customer_role_mapping", lc.Delete)
}
