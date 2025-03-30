package route

import (
	"time"

	controller "earnforglance/server/api/controller/customers"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/customers"

	repository "earnforglance/server/repository/customers"
	"earnforglance/server/service/data/mongo"
	usecase "earnforglance/server/usecase/customers"

	"github.com/gin-gonic/gin"
)

func ExternalAuthenticationRecordRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewExternalAuthenticationRecordRepository(db, domain.CollectionExternalAuthenticationRecord)
	lc := &controller.ExternalAuthenticationRecordController{
		ExternalAuthenticationRecordUsecase: usecase.NewExternalAuthenticationRecordUsecase(ur, timeout),
		Env:                                 env,
	}

	Group := group.Group("/api/v1/customers")

	group.GET("/external_authentication_records", lc.Fetch)
	group.GET("/external_authentication_record", lc.FetchByID)
	group.POST("/external_authentication_record", lc.Create)
	Group.POST("/external_authentication_records", lc.CreateMany)
	group.PUT("/external_authentication_record", lc.Update)
	group.DELETE("external_authentication_record", lc.Delete)
}
