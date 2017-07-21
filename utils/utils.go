package utils

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetParamInt8(ctx *gin.Context, param string, defalt int8) int8 {
	val := ctx.Request.FormValue(param)
	rvl, err := strconv.Atoi(val)
	if err != nil {
		return defalt
	}
	return int8(rvl)
}

func GetParamString(ctx *gin.Context, param string, defaultValue string) string {
	val := ctx.Request.FormValue(param)
	if val == "" {
		return defaultValue
	}
	return val
}
