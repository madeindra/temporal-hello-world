package main

import (
	"log"

	app "github.com/madeindra/temporal-hello-world"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	// create client
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()

	// create new worker for task queue
	w := worker.New(c, app.GreetingTaskQueue, worker.Options{})

	// register workflow and activity to worker
	w.RegisterWorkflow(app.GreetingWorkflowFuture)
	w.RegisterWorkflow(app.GreetingWorkflow)
	w.RegisterActivity(app.ComposeGreeting)

	// run worker & listen to task queue
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("unable to start Worker", err)
	}
}
