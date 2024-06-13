package server

import (
	"github.com/azar-writes-code/traefik-temporal-poc/internal/handlers"
	"github.com/gin-gonic/gin"
	"go.temporal.io/sdk/client"
)

func SetupRouter(c client.Client) *gin.Engine {
    router := gin.Default()

    router.POST("/start-workflow", handlers.StartWorkflowHandler(c))

    return router
}
