package routes

import (
	"prime-erp-core/internal/utils"

	approvalService "prime-erp-core/Internal/services/approval-service"
	priceService "prime-erp-core/Internal/services/price-service"
	quotationService "prime-erp-core/Internal/services/quotation-service"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(ctx *gin.Engine) {

	//price
	price := ctx.Group("/price")

	price.POST("/GetPriceListGroup", func(c *gin.Context) {
		utils.ProcessRequest(c, priceService.GetPriceListGroup)
	})
	price.POST("/GetPaymentTerm", func(c *gin.Context) {
		utils.ProcessRequest(c, priceService.GetPaymentTerm)
	})
	price.POST("/GetComparePrice", func(c *gin.Context) {
		utils.ProcessRequest(c, priceService.GetComparePrice)
	})

	//quotation
	quotation := ctx.Group("/quotation")

	quotation.POST("/GetQoutation", func(c *gin.Context) {
		utils.ProcessRequest(c, quotationService.GetQuotation)
	})
	quotation.POST("/CreateQuotation", func(c *gin.Context) {
		utils.ProcessRequest(c, quotationService.CreateQuotation)
	})

	//approval
	approval := ctx.Group("/approval")
	approval.POST("/VerifyApprove", func(c *gin.Context) {
		utils.ProcessRequest(c, approvalService.VerifyApprove)
	})
}
