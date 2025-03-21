package route

import (
	"time"

	controller "earnforglance/server/api/controller/common"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/common"

	"earnforglance/server/mongo"
	repository "earnforglance/server/repository/common"
	usecase "earnforglance/server/usecase/common"

	"github.com/gin-gonic/gin"
)

func PdfSettingsRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewPdfSettingsRepository(db, domain.CollectionPdfSettings)
	lc := &controller.PdfSettingsController{
		PdfSettingsUsecase: usecase.NewPdfSettingsUsecase(ur, timeout),
		Env:                env,
	}

	group.GET("/pdf_settingss", lc.Fetch)
	group.GET("/pdf_settings", lc.FetchByID)
	group.POST("/pdf_settings", lc.Create)
	group.PUT("/pdf_settings", lc.Update)
	group.DELETE("/pdf_settings", lc.Delete)
}
