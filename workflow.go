package main

import (
	"context"
	"fmt"
	"time"

	"go.temporal.io/sdk/workflow"
)

func Workflow(ctx workflow.Context) error {
	options := workflow.ActivityOptions{
		TaskQueue: "activity",
		// StartToCloseTimeout: time.Second * 10,
		ScheduleToCloseTimeout: time.Second * 10,
		// ScheduleToStartTimeout: time.Second * 10,
		HeartbeatTimeout: time.Second * 30,
	}

	ctx = workflow.WithActivityOptions(ctx, options)

	err := workflow.ExecuteActivity(ctx, Activity).Get(ctx, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println("done")

	return nil
}

func Activity(ctx context.Context) error {
	fmt.Println("Activity")
	return nil
}
