package handler

import (
	"errors"
	"net/http"

	"github.com/ArtemST2006/Avito_internship/backend/internal/schemas"
	projerrors "github.com/ArtemST2006/Avito_internship/backend/pkg/errors"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetTeam(c *gin.Context) {
	teamName := c.Query("team_name")
	if teamName == "" {
		NewResponseError(c, http.StatusBadRequest, "BED_REQUEST", "have not team name")
		return
	}
	response, err := h.service.Team.GetTeam(teamName)
	if errors.Is(err, projerrors.ErrNotFound) {
		NewResponseError(c, http.StatusNotFound, "NOT_FOUND", err.Error())
		return
	} else if err != nil {
		NewResponseError(c, http.StatusInternalServerError, "SERVER_ERROR", err.Error())
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h *Handler) AddTeam(c *gin.Context) {
	request := new(schemas.CreateTeamRequest)
	if err := c.Bind(&request); err != nil {
		NewResponseError(c, http.StatusNotFound, "SERVER_ERROR", "have not team name")
		return
	}
	response, err := h.service.Team.AddTeam(*request)
	if errors.Is(err, projerrors.ErrAlreadyExist) {
		NewResponseError(c, http.StatusNotFound, "TEAM_EXISTS", err.Error())
		return
	} else if err != nil {
		NewResponseError(c, http.StatusInternalServerError, "SERVER_ERROR", err.Error())
		return
	}

	c.JSON(http.StatusOK, response)
}
