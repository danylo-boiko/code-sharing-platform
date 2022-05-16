package handlers

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateCodeSnippet(c *gin.Context) {
	OkResponse(c, "ok", nil)
}
