package routes

import (
	"github.com/prime-solution-dev/prime-erp-core/internal/utils"

	"github.com/gin-gonic/gin"
	priceService "github.com/prime-solution-dev/prime-erp-core/Internal/services/price-service"
)

func RegisterRoutes(ctx *gin.Engine) {

	//price
	price := ctx.Group("/price")

	price.POST("/get-config-next-process", func(c *gin.Context) {
		utils.ProcessRequest(c, priceService.GetPriceList)
	})

	//quotation

}
