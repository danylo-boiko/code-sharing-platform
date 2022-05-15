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

	/*api := router.Group("/api")
	{
		supportedLanguages := api.Group("/supported_languages")
		{
			supportedLanguages.GET("/")
		}

		codeSnippets := api.Group("/code_snippets")
		{
			codeSnippets.GET("/")
			codeSnippets.GET("/:id")
			codeSnippets.POST("/")
			codeSnippets.PUT("/:id")
			codeSnippets.DELETE("/:id")
		}
	}*/

	return router
}
