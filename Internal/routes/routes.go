package routes

import (
	"github.com/prime-solution-dev/prime-erp-core/internal/utils"

	"github.com/gin-gonic/gin"
	priceService "github.com/prime-solution-dev/prime-erp-core/Internal/services/price-service"
	quotationService "github.com/prime-solution-dev/prime-erp-core/Internal/services/quotation-service"
)

func RegisterRoutes(ctx *gin.Engine) {

	//price
	price := ctx.Group("/price")

	price.POST("/GetPriceList", func(c *gin.Context) {
		utils.ProcessRequest(c, priceService.GetPriceList)
	})

	//quotation
	quotation := ctx.Group("/quotation")

	quotation.POST("/GetQoutation", func(c *gin.Context) {
		utils.ProcessRequest(c, quotationService.GetQuotation)
	})
}
