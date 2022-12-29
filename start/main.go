package main

import (
	"context"
	"fmt"
	"log"

	app "github.com/madeindra/temporal-hello-world"
	"go.temporal.io/sdk/client"
)

func main() {
	// create client
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()

	// create workflow options
	options := client.StartWorkflowOptions{
		ID:        "username-123",
		TaskQueue: app.GreetingTaskQueue,
	}

	// start workflow
	name := "World"
	wf, err := c.ExecuteWorkflow(context.Background(), options, app.GreetingWorkflowFuture, name)
	if err != nil {
		log.Fatalln("unable to complete Workflow", err)
	}

	// get results
	var greeting string
	err = wf.Get(context.Background(), &greeting)
	if err != nil {
		log.Fatalln("unable to get Workflow result", err)
	}

	printResult(greeting, wf.GetID(), wf.GetRunID())
}

func printResult(greeting string, workflowID string, runID string) {
	fmt.Printf("\n WorkflowID: %s RunID: %s \n", workflowID, runID)
	fmt.Printf("\n %s \n\n", greeting)
}
