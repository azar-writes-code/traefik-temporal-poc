package handlers

import (
	"context"
	"net/http"

	"github.com/azar-writes-code/traefik-temporal-poc/internal/workflows"
	"github.com/gin-gonic/gin"
	"go.temporal.io/sdk/client"
)

type StartWorkflowRequest struct {
    Name string `json:"name" binding:"required"`
}

func StartWorkflowHandler(c client.Client) gin.HandlerFunc {
    return func(ctx *gin.Context) {
        var req StartWorkflowRequest

        if err := ctx.ShouldBindJSON(&req); err != nil {
            ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        workflowOptions := client.StartWorkflowOptions{
            ID:        "sample-workflow",
            TaskQueue: "sample-task-queue",
        }

        we, err := c.ExecuteWorkflow(context.Background(), workflowOptions, workflows.SampleWorkflow, req.Name)
        if err != nil {
            ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        ctx.JSON(http.StatusOK, gin.H{"workflowID": we.GetID(), "runID": we.GetRunID()})
    }
}
