package api

import (
	"net/http"

	"github.com/mundipagg/boleto-api/config"
	"github.com/mundipagg/boleto-api/log"

	"github.com/mundipagg/boleto-api/models"
	gin "gopkg.in/gin-gonic/gin.v1"
)

//InstallRestAPI "instala" e sobe o servico de rest
func InstallRestAPI(router *gin.Engine) {
	router.Use(gin.Recovery())
	router.Use(executionController())
	if config.Get().DevMode && !config.Get().MockMode {
		router.Use(gin.Logger())
	}
	InstallV1(router)
	router.GET("/boleto", getBoleto)
}

func checkError(c *gin.Context, err error, l *log.Log) bool {

	if err != nil {
		errResp := models.BoletoResponse{
			Errors: models.NewErrors(),
		}
		switch v := err.(type) {
		case models.IErrorResponse:
			errResp.Errors.Append(v.ErrorCode(), v.Error())
			c.JSON(http.StatusBadRequest, errResp)
		case models.IHttpNotFound:
			errResp.Errors.Append("not_found", v.Error())
			c.JSON(http.StatusNotFound, errResp)
		case models.IFormatError:
			errResp.Errors.Append("bad_request", v.Error())
			c.JSON(http.StatusBadRequest, errResp)
		case models.IServerError:
			errResp.Errors.Append("api_error", "internal error")
			l.Fatal(v.Error(), v.Message())
			c.JSON(http.StatusInternalServerError, errResp)
		default:
			l.Fatal(err.Error(), "")
			errResp.Errors.Append("api_error", "internal error")
			c.JSON(http.StatusInternalServerError, errResp)
		}
		return true
	}
	return false
}
