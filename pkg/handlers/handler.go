package handlers

import (
	"code-sharing-platform/pkg/services"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *services.Service
}

func NewHandler(services *services.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-in", h.SignIn)
		auth.POST("/sign-up", h.SignUp)
	}

	api := router.Group("/api", h.UserIdentity)
	{
		supportedLanguages := api.Group("/supported-languages")
		{
			supportedLanguages.GET("/", h.GetSupportedLanguages)
		}

		codeSnippets := api.Group("/code-snippets")
		{
			codeSnippets.POST("/", h.CreateCodeSnippet)
		}
	}

	return router
}
