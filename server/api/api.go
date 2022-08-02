package api

import (
	v1 "server/api/v1"

	"github.com/gin-gonic/gin"
)

func RegistryAll(g *gin.RouterGroup) {
	v1.Registry(g.Group("/v1"))
}
