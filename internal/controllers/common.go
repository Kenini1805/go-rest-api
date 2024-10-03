package controllers

import (
	"net/http"

	"github.com/Kenini1805/go-rest-api/pkg/converter"
	httperrors "github.com/Kenini1805/go-rest-api/pkg/http_errors"
	"github.com/gin-gonic/gin"
)

func BindRequest(ctx *gin.Context, req interface{}) error {
	err := ctx.ShouldBind(req)
	if err != nil {
		response := converter.BuildErrorResponse(httperrors.ErrBadRequestMessage, err.Error(), converter.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return err
	}
	return nil
}
