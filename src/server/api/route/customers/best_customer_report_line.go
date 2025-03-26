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

func BestCustomerReportLineRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewBestCustomerReportLineRepository(db, domain.CollectionBestCustomerReportLine)
	lc := &controller.BestCustomerReportLineController{
		BestCustomerReportLineUsecase: usecase.NewBestCustomerReportLineUsecase(ur, timeout),
		Env:                           env,
	}

	itemGroup := group.Group("/api/v1/customers")
	itemGroup.GET("/best_customer_report_lines", lc.Fetch)
	itemGroup.GET("/best_customer_report_line", lc.FetchByID)
	itemGroup.POST("/best_customer_report_line", lc.Create)
	itemGroup.POST("/best_customer_report_lines", lc.CreateMany)
	itemGroup.PUT("/best_customer_report_line", lc.Update)
	itemGroup.DELETE("/best_customer_report_line", lc.Delete)
}
