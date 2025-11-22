package handler

import "github.com/gin-gonic/gin"

// type errorResponse struct {
// 	Error string `json:"error"`
// }

func NewResponseError(c *gin.Context, status int, code, message string) {
	c.AbortWithStatusJSON(status, gin.H{
		"error": gin.H{
			"code":    code,
			"message": message,
		},
	})
}
