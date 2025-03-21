package route

import (
	"time"

	controller "earnforglance/server/api/controller/orders"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/orders"

	"earnforglance/server/mongo"
	repository "earnforglance/server/repository/orders"
	usecase "earnforglance/server/usecase/orders"

	"github.com/gin-gonic/gin"
)

func BestSellersReportLineRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewBestSellersReportLineRepository(db, domain.CollectionBestSellersReportLine)
	lc := &controller.BestSellersReportLineController{
		BestSellersReportLineUsecase: usecase.NewBestSellersReportLineUsecase(ur, timeout),
		Env:                          env,
	}

	group.GET("/best_sellers_report_lines", lc.Fetch)
	group.GET("/best_sellers_report_line", lc.FetchByID)
	group.POST("/best_sellers_report_line", lc.Create)
	group.PUT("/best_sellers_report_line", lc.Update)
	group.DELETE("/best_sellers_report_line", lc.Delete)
}
