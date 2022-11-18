package app

import (
	"time"

	"go.temporal.io/sdk/workflow"
)

func GreetingWorkflow(ctx workflow.Context, name string) (string, error) {
	// create a new workflow options
	options := workflow.ActivityOptions{
		// set timeout to 5 second
		StartToCloseTimeout: time.Second * 5,
	}

	// add workflow option to the context
	ctx = workflow.WithActivityOptions(ctx, options)

	// prepare result variable
	var result string

	// execute activity with context and wait until the Future has error/result
	// ExecuteActivity can take multiple inputs, but it is recommended to only use one input
	err := workflow.ExecuteActivity(ctx, ComposeGreeting, name).Get(ctx, &result)

	// return result or error
	return result, err
}
