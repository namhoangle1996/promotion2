package httpErrors

import "github.com/gin-gonic/gin"

func ErrorCtxResponse(ctx *gin.Context, err error) {
	httpCode, code, message := parseErrors(err)
	ctx.JSON(httpCode, gin.H{"code": code, "message": message})
}

func parseErrors(err error) (httpCode, code int, message string) {
	switch true {

	}

	return 1, 1, ""
}
