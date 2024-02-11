package handlers

import (
	_ "code-sharing-platform/docs"
	"code-sharing-platform/pkg/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"os"
)

type Handler struct {
	services *services.Service
}

func NewHandler(services *services.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.Use(CORSMiddleware())

	swaggerUrl := ginSwagger.URL(fmt.Sprintf("http://%s:%s/swagger/doc.json", os.Getenv("APP_DOMAIN"), viper.GetString("app.port")))
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, swaggerUrl))

	auth := router.Group("/auth")
	{
		auth.POST("/sign-in", h.SignIn)
		auth.POST("/sign-up", h.SignUp)
	}

	api := router.Group("/api")
	{
		supportedLanguages := api.Group("/supported-languages", h.UserIdentity)
		{
			supportedLanguages.GET("/", h.GetSupportedLanguages)
		}

		codeSnippets := api.Group("/code-snippets")
		{
			codeSnippets.GET("/:id", h.AnonymousUserIdentity, h.GetCodeSnippetById)
			codeSnippets.POST("/", h.UserIdentity, h.CreateCodeSnippet)
			codeSnippets.PUT("/:id", h.UserIdentity, h.UpdateCodeSnippet)
			codeSnippets.DELETE("/:id", h.UserIdentity, h.DeleteCodeSnippet)
		}
	}

	return router
}
