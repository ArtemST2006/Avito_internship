package handler

import (
	"errors"
	"net/http"

	"github.com/ArtemST2006/Avito_internship/backend/internal/schemas"
	projerrors "github.com/ArtemST2006/Avito_internship/backend/pkg/errors"
	"github.com/gin-gonic/gin"
)

func (h *Handler) SetIsActive(c *gin.Context) {
	request := new(schemas.ActieveUserRequest)
	if err := c.Bind(&request); err != nil {
		NewResponseError(c, http.StatusBadRequest, "BAD_REQUEST", err.Error())
		return
	}

	response, err := h.service.User.SetIsActive(*request)

	if err != nil {
		if errors.Is(err, projerrors.ErrNotFound) {
			NewResponseError(c, http.StatusConflict, "NOT_FOUND", err.Error())
			return
		}
		NewResponseError(c, http.StatusInternalServerError, "SERVER_ERROR", err.Error())
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h *Handler) GetUserReview(c *gin.Context) {
	userId := c.Query("user_id")
	if userId == "" {
		NewResponseError(c, http.StatusBadRequest, "BAD_REQUEST", "no params")
		return
	}

	response, err := h.service.User.GetUserReview(userId)

	if err != nil {
		if errors.Is(err, projerrors.ErrNotFound) {
			NewResponseError(c, http.StatusConflict, "NOT_FOUND", err.Error())
			return
		}
		NewResponseError(c, http.StatusInternalServerError, "SERVER_ERROR", err.Error())
		return
	}

	c.JSON(http.StatusOK, response)
}
