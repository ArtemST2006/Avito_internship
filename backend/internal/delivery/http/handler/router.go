package handler

import (
	"time"

	"github.com/ArtemST2006/Avito_internship/backend/internal/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost", "http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// team := router.Group("/team")
	{
		// team.POST("/add", h.AddTeam)
		// team.GET("/get", h.GetTeam)
	}
	// users := router.Group("/users")
	{
		// users.POST(atePullRequest)
		// pullRequest.POST("/merge", h.MergePullRequest)
		// pullRequest.POST("/reassign", h.ReassignPullRequest)"/setIsActive", h.SetUserActive)
		// users.GET("/getReview", h.GetUserReview)
	}
	// pullRequest := router.Group("/pullRequest")
	{
		// pullRequest.POST("/create", h.CreatePullRequest)
		// pullRequest.POST("/merge", h.MergePullRequest)
		// pullRequest.POST("/reassign", h.ReassignPullRequest)
	}

	return router
}
