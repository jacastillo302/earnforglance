package route

import (
	"time"

	controller "earnforglance/server/api/controller/customers"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/customers"

	"earnforglance/server/mongo"
	repository "earnforglance/server/repository/customers"
	usecase "earnforglance/server/usecase/customers"

	"github.com/gin-gonic/gin"
)

func CustomerRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewCustomerRepository(db, domain.CollectionCustomer)
	lc := &controller.CustomerController{
		CustomerUsecase: usecase.NewCustomerUsecase(ur, timeout),
		Env:             env,
	}

	group.GET("/customers", lc.Fetch)
	group.GET("/customer", lc.FetchByID)
	group.POST("/customer", lc.Create)
	group.POST("/customers", lc.CreateMany)
	group.PUT("/customer", lc.Update)
	group.DELETE("/customer", lc.Delete)
}
