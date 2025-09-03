package utils

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetQueryInt(ctx *gin.Context, key string, defaultVal int) int {
	val, ok := ctx.GetQuery(key)
	if !ok {
		return defaultVal
	}
	i, err := strconv.Atoi(val)
	if err != nil {
		return defaultVal
	}
	return i
}