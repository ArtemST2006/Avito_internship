package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Statistic(c *gin.Context) {
	response, err := h.service.Statistic.Statistic()
	if err != nil {
		NewResponseError(c, http.StatusInternalServerError, "SERVER_ERROR", err.Error())
	}

	c.JSON(http.StatusOK, response)
}
