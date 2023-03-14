package main

import (
	"context"
	"log"

	"go.temporal.io/sdk/client"
)

func main() {
	c, err := client.Dial(client.Options{
		HostPort:  "localhost:7234",
		Namespace: "testing",
	})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	workflowOptions := client.StartWorkflowOptions{
		TaskQueue: "workflow",
		ID:        "something",
	}
	we, err := c.ExecuteWorkflow(context.Background(), workflowOptions, Workflow)
	if err != nil {
		log.Fatalln("Unable to execute workflow", err)
	}
	log.Println("Started Workflow", we.GetID(), "RunID", we.GetRunID())
}
