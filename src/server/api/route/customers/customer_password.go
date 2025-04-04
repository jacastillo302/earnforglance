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

func CustomerPasswordRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewCustomerPasswordRepository(db, domain.CollectionCustomerPassword)
	lc := &controller.CustomerPasswordController{
		CustomerPasswordUsecase: usecase.NewCustomerPasswordUsecase(ur, timeout),
		Env:                     env,
	}

	Group := group.Group("/api/v1/customers")

	Group.GET("/customer_passwords", lc.Fetch)
	Group.GET("/customer_password", lc.FetchByID)
	Group.POST("/customer_password", lc.Create)
	Group.POST("/customer_passwords", lc.CreateMany)
	Group.PUT("/customer_password", lc.Update)
	Group.DELETE("customer_password", lc.Delete)
}
