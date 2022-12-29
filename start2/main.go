package main

import (
	"context"
	"fmt"
	"log"

	"go.temporal.io/sdk/client"
)

type RevertUsername struct {
}

func main() {
	// create client
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()

	// create workflow options
	// options := client.StartWorkflowOptions{
	// 	ID:        "username-123",
	// 	TaskQueue: app.GreetingTaskQueue,
	// }

	// start workflow
	err = c.SignalWorkflow(context.Background(), "username-123", "", "your-signal-name", RevertUsername{})
	if err != nil {
		fmt.Printf("Error", err)
	}
	// get results
	// var greeting string
	// err = wf.Get(context.Background(), &greeting)
	// if err != nil {
	// 	log.Fatalln("unable to get Workflow result", err)
	// }

	// printResult(greeting, wf.GetID(), wf.GetRunID())
}

func printResult(greeting string, workflowID string, runID string) {
	fmt.Printf("\n WorkflowID: %s RunID: %s \n", workflowID, runID)
	fmt.Printf("\n %s \n\n", greeting)
}
