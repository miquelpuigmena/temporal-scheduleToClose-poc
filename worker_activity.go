package main

import (
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
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

	actWorker := worker.New(c, "activity", worker.Options{})
	actWorker.RegisterActivity(Activity)

	err = actWorker.Run(worker.InterruptCh())
	if err != nil {
		panic(err)
	}
}
