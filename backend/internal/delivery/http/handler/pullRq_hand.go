package handler

import (
	"errors"
	"net/http"

	"github.com/ArtemST2006/Avito_internship/backend/internal/schemas"
	projerrors "github.com/ArtemST2006/Avito_internship/backend/pkg/errors"
	"github.com/gin-gonic/gin"
)

func (h *Handler) MergePullRequest(c *gin.Context) {
	request := new(schemas.PullRqMergeRequest)
	if err := c.Bind(&request); err != nil {
		NewResponseError(c, http.StatusBadRequest, "BAD_REQUEST", err.Error())
		return
	}

	response, err := h.service.PullRequest.MergePR(*request)
	if errors.Is(err, projerrors.ErrNotFound) {
		NewResponseError(c, http.StatusNotFound, "NOT_FOUND", err.Error())
		return
	} else if err != nil {
		NewResponseError(c, http.StatusInternalServerError, "SERVER_ERROR", err.Error())
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h *Handler) CreatePullRequest(c *gin.Context) {
	request := new(schemas.CreatePullRequestRequest)
	if err := c.Bind(&request); err != nil {
		NewResponseError(c, http.StatusBadRequest, "BAD_REQUEST", err.Error())
		return
	}

	response, err := h.service.PullRequest.CreatePR(*request)
	if err != nil {
		if errors.Is(err, projerrors.ErrNotFound) {
			NewResponseError(c, http.StatusNotFound, "NOT_FOUND", err.Error())
			return
		} else if errors.Is(err, projerrors.ErrAlreadyExist) {
			NewResponseError(c, http.StatusConflict, "PR_EXISTS", err.Error())
			return
		} else {
			NewResponseError(c, http.StatusInternalServerError, "SERVER_ERROR", err.Error())
			return
		}
	}

	c.JSON(http.StatusOK, response)
}

func (h *Handler) ReassignPullRequest(c *gin.Context) {
	request := new(schemas.PRChangeAuthorRequest)
	if err := c.Bind(&request); err != nil {
		NewResponseError(c, http.StatusBadRequest, "BAD_REQUEST", err.Error())
		return
	}

	response, err := h.service.PullRequest.ChangeAuthorPR(*request)
	if err != nil {
		if errors.Is(err, projerrors.ErrNotFound) {
			NewResponseError(c, http.StatusNotFound, "NOT_FOUND", err.Error())
			return
		} else if errors.Is(err, projerrors.ErrNoCandidate) {
			NewResponseError(c, http.StatusConflict, "NO_CANDIDATE", err.Error())
			return
		} else if errors.Is(err, projerrors.ErrNoAssign) {
			NewResponseError(c, http.StatusConflict, "NOT_ASSIGNED", err.Error())
			return
		} else if errors.Is(err, projerrors.ErrAlreadyMerged) {
			NewResponseError(c, http.StatusConflict, "PR_MERGED", err.Error())
			return
		} else {
			NewResponseError(c, http.StatusInternalServerError, "SERVER_ERROR", err.Error())
			return
		}
	}

	c.JSON(http.StatusOK, response)
}
