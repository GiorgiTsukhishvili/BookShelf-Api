package middlewares

import (
	"github.com/gin-gonic/gin"
)

func LangCheck(ctx *gin.Context) {
	langHeader := ctx.GetHeader("Accept-Language")

	if langHeader == "" {
		langHeader = "en"
	}

	ctx.Set("lang", langHeader)

	ctx.Next()
}
