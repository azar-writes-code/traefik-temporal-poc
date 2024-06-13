package workflows

import (
	"time"

	"go.temporal.io/sdk/workflow"
)

// SampleWorkflow is a sample Temporal workflow definition.
func SampleWorkflow(ctx workflow.Context, name string) error {
    logger := workflow.GetLogger(ctx)
    logger.Info("Workflow started", "name", name)

    // Simulate some work with a sleep
    err := workflow.Sleep(ctx, 10*time.Second)
    if err != nil {
        return err
    }

    logger.Info("Workflow completed", "name", name)
    return nil
}
