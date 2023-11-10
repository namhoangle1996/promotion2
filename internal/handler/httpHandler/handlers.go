package httpHandler

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"kk/internal/dto/requestDto"
	"kk/internal/usecase/usecaseInterface"
	"kk/utils/httpErrors"
	"net/http"
)

type promotionHandler struct {
	group    *gin.RouterGroup
	promoUc  usecaseInterface.PromotionServiceInterface
	validate *validator.Validate
}

func newPromotionHandler(group *gin.RouterGroup, promoUc usecaseInterface.PromotionServiceInterface, validate *validator.Validate) *promotionHandler {
	return &promotionHandler{group: group, promoUc: promoUc, validate: validate}
}

func (h *promotionHandler) Create(c *gin.Context) {
	err := h.promoUc.Create(c)
	if err != nil {
		httpErrors.ErrorCtxResponse(c, err)
		return
	}
}

func (h *promotionHandler) GetVoucherByCode(c *gin.Context) {
	var req requestDto.GetVoucherReq
	if err := c.BindQuery(&req); err != nil {
		httpErrors.ErrorCtxResponse(c, err)
		return
	}

	promo, err := h.promoUc.GetVoucherByCode(c, req.VoucherCode)
	if err != nil {
		httpErrors.ErrorCtxResponse(c, err)
		return
	}

	c.JSON(http.StatusOK, promo)
}
