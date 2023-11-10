package requestDto

type GetVoucherReq struct {
	VoucherCode string `json:"voucher_code" form:"voucher_code"  validate:"required" `
}
