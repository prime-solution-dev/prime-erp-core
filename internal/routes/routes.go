package routes

import (
	"prime-erp-core/internal/utils"

	approvalService "prime-erp-core/internal/services/approval-service"
	creditService "prime-erp-core/internal/services/credit-service"
	priceService "prime-erp-core/internal/services/price-service"
	quotationService "prime-erp-core/internal/services/quotation-service"
	saleService "prime-erp-core/internal/services/sale-service"

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

	quotation.POST("/GetQuotation", func(c *gin.Context) {
		utils.ProcessRequest(c, quotationService.GetQuotation)
	})
	quotation.POST("/CreateQuotation", func(c *gin.Context) {
		utils.ProcessRequest(c, quotationService.CreateQuotation)
	})

	//sale
	sale := ctx.Group("/sale")
	sale.POST("/GetSale", func(c *gin.Context) {
		utils.ProcessRequest(c, saleService.GetSale)
	})

	//approval
	approval := ctx.Group("/approval")
	approval.POST("/VerifyApprove", func(c *gin.Context) {
		utils.ProcessRequest(c, approvalService.VerifyApprove)
	})
	approval.POST("/GetApproval", func(c *gin.Context) {
		utils.ProcessRequest(c, approvalService.GetApproval)
	})
	approval.POST("/CreateApproval", func(c *gin.Context) {
		utils.ProcessRequest(c, approvalService.CreateApproval)
	})
	approval.POST("/UpdateApproval", func(c *gin.Context) {
		utils.ProcessRequest(c, approvalService.UpdateApproval)
	})
	//credit
	credit := ctx.Group("/credit")
	credit.POST("/GetCreditCurrent", func(c *gin.Context) {
		utils.ProcessRequest(c, creditService.GetCreditCurrentAPI)
	})

}
