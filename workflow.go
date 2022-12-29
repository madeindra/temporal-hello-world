package app

import (
	"time"

	"go.temporal.io/sdk/workflow"
)

type RevertUsername struct {
}

func GreetingWorkflowFuture(ctx workflow.Context, name string) (string, error) {
	// create a new workflow options
	options := workflow.ActivityOptions{
		// set timeout to 5 second
		StartToCloseTimeout: time.Second * 5,
	}

	// add workflow option to the context
	ctx = workflow.WithActivityOptions(ctx, options)

	// add child ctx
	// childCtx, cancelHandler := workflow.WithCancel(ctx)

	signalChan := workflow.GetSignalChannel(ctx, "your-signal-name")

	// declare selector
	selector := workflow.NewSelector(ctx)
	selector.AddReceive(signalChan, func(c workflow.ReceiveChannel, more bool) {
		var revertUname RevertUsername

		c.Receive(ctx, &revertUname)
		// TODO: revert work
		workflow.GetLogger(ctx).Info("Username reverted")
	})

	// prepare result variable
	var result string

	selector.AddFuture(workflow.NewTimer(ctx, 10*time.Second), func(f workflow.Future) {
		ao := workflow.ActivityOptions{
			StartToCloseTimeout: 60 * time.Second,
		}

		ctx = workflow.WithActivityOptions(ctx, ao)

		err := workflow.ExecuteActivity(ctx, ComposeGreeting, name).Get(ctx, &result)
		if err != nil {
			return
		}

		// cancelHandler()
		workflow.GetLogger(ctx).Info("Username changed")
	})

	// This wait for receive or future to complete
	selector.Select(ctx)

	workflow.GetLogger(ctx).Info("Workflow completed")

	// return result or error
	return result, nil
}

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
