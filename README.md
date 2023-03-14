# Testing `ScheduleToClose` Time Out
1. Start temporal server
```bash
# cli-0
temporalite start --ephemeral --namespace testing --port 7234
```

1. Register worker that registers the workflow
```bash
# cli-1
go run worker_workflow.go workflow.go
```

1. Register worker that registers the activity
```bash
# cli-2
go run worker_activity.go workflow.go
```

1. Once the activity has been registered in the temporal server, just close the worker so it doesn't pick the tasks from the `activity` queue
```bash
# cli-2
# kill process
```

1. Run the starter
```bash
# cli-3
go run starter.go workflow.go
```

## Expected
It should either fail with `ScheduleToClose` Timeout, or not fail.

## Actual
Getting `ScheduleToStart` Timeout

```
2023/03/14 11:27:50 DEBUG ExecuteActivity Namespace testing TaskQueue workflow WorkerID 32992@DD-Miquel-Puig-Mena.local@ WorkflowType Workflow WorkflowID something RunID 32685a11-843c-4203-bab5-708f0982e970 Attempt 1 ActivityID 5 ActivityType Activity
2023/03/14 11:28:00 ERROR Workflow panic Namespace testing TaskQueue workflow WorkerID 32992@DD-Miquel-Puig-Mena.local@ WorkflowType Workflow WorkflowID something RunID 32685a11-843c-4203-bab5-708f0982e970 Attempt 1 Error activity error (type: Activity, scheduledEventID: 5, startedEventID: 0, identity: ): activity ScheduleToStart timeout (type: ScheduleToStart) StackTrace coroutine root [panic]:
main.Workflow({0x104c6f648?, 0x140002d49c0?})
        /Users/miquel/Documents/repos/poc-scheduleToClose-timeout/workflow.go:24 +0x13c
reflect.Value.call({0x104b31e00?, 0x104c62348?, 0x12c9d5228?}, {0x10498f89f, 0x4}, {0x1400031a3a8, 0x1, 0x0?})
        /usr/local/go/src/reflect/value.go:586 +0x838
reflect.Value.Call({0x104b31e00?, 0x104c62348?, 0x140004a0cd8?}, {0x1400031a3a8?, 0x140004a0cf8?, 0x104961730?})
        /usr/local/go/src/reflect/value.go:370 +0x90
go.temporal.io/sdk/internal.executeFunction({0x104b31e00, 0x104c62348}, {0x140004a0de8, 0x1, 0x18?})
        /Users/miquel/go/pkg/mod/go.temporal.io/sdk@v1.21.1/internal/internal_worker.go:1729 +0x2d0
go.temporal.io/sdk/internal.(*workflowEnvironmentInterceptor).ExecuteWorkflow(0x140000ac960, {0x104c6f4c0?, 0x14000094840}, 0x1400031a390)
        /Users/miquel/go/pkg/mod/go.temporal.io/sdk@v1.21.1/internal/workflow.go:506 +0x13c
go.temporal.io/sdk/internal.(*workflowExecutor).Execute(0x140000b65c0, {0x104c6f4c0, 0x14000094840}, 0x25?)
        /Users/miquel/go/pkg/mod/go.temporal.io/sdk@v1.21.1/internal/internal_worker.go:780 +0x230
go.temporal.io/sdk/internal.(*syncWorkflowDefinition).Execute.func1({0x104c6f648, 0x140000a5620})
        /Users/miquel/go/pkg/mod/go.temporal.io/sdk@v1.21.1/internal/internal_workflow.go:507 +0xd0
2023/03/14 11:28:00 WARN  Failed to process workflow task. Namespace testing TaskQueue workflow WorkerID 32992@DD-Miquel-Puig-Mena.local@ WorkflowType Workflow WorkflowID something RunID 32685a11-843c-4203-bab5-708f0982e970 Attempt 1 Error activity error (type: Activity, scheduledEventID: 5, startedEventID: 0, identity: ): activity ScheduleToStart timeout (type: ScheduleToStart)
2023/03/14 11:28:00 ERROR Workflow panic Namespace testing TaskQueue workflow WorkerID 32992@DD-Miquel-Puig-Mena.local@ WorkflowType Workflow WorkflowID something RunID 32685a11-843c-4203-bab5-708f0982e970 Attempt 1 Error activity error (type: Activity, scheduledEventID: 5, startedEventID: 0, identity: ): activity ScheduleToStart timeout (type: ScheduleToStart) StackTrace coroutine root [panic]:
main.Workflow({0x104c6f648?, 0x1400032f2c0?})
        /Users/miquel/Documents/repos/poc-scheduleToClose-timeout/workflow.go:24 +0x13c
reflect.Value.call({0x104b31e00?, 0x104c62348?, 0x12c9d5228?}, {0x10498f89f, 0x4}, {0x1400031aff0, 0x1, 0x0?})
        /usr/local/go/src/reflect/value.go:586 +0x838
reflect.Value.Call({0x104b31e00?, 0x104c62348?, 0x140004a1cd8?}, {0x1400031aff0?, 0x140004a1cf8?, 0x104961730?})
        /usr/local/go/src/reflect/value.go:370 +0x90
go.temporal.io/sdk/internal.executeFunction({0x104b31e00, 0x104c62348}, {0x140004a1de8, 0x1, 0x18?})
        /Users/miquel/go/pkg/mod/go.temporal.io/sdk@v1.21.1/internal/internal_worker.go:1729 +0x2d0
go.temporal.io/sdk/internal.(*workflowEnvironmentInterceptor).ExecuteWorkflow(0x14000337680, {0x104c6f4c0?, 0x14000317380}, 0x1400031afd8)
        /Users/miquel/go/pkg/mod/go.temporal.io/sdk@v1.21.1/internal/workflow.go:506 +0x13c
go.temporal.io/sdk/internal.(*workflowExecutor).Execute(0x14000332fc0, {0x104c6f4c0, 0x14000317380}, 0x25?)
        /Users/miquel/go/pkg/mod/go.temporal.io/sdk@v1.21.1/internal/internal_worker.go:780 +0x230
go.temporal.io/sdk/internal.(*syncWorkflowDefinition).Execute.func1({0x104c6f648, 0x1400032f1a0})
        /Users/miquel/go/pkg/mod/go.temporal.io/sdk@v1.21.1/internal/internal_workflow.go:507 +0xd0
2023/03/14 11:28:00 WARN  Failed to process workflow task. Namespace testing TaskQueue workflow WorkerID 32992@DD-Miquel-Puig-Mena.local@ WorkflowType Workflow WorkflowID something RunID 32685a11-843c-4203-bab5-708f0982e970 Attempt 1 Error activity error (type: Activity, scheduledEventID: 5, startedEventID: 0, identity: ): activity ScheduleToStart timeout (type: ScheduleToStart)
2023/03/14 11:28:00 ERROR Workflow panic Namespace testing TaskQueue workflow WorkerID 32992@DD-Miquel-Puig-Mena.local@ WorkflowType Workflow WorkflowID something RunID 32685a11-843c-4203-bab5-708f0982e970 Attempt 2 Error activity error (type: Activity, scheduledEventID: 5, startedEventID: 0, identity: ): activity ScheduleToStart timeout (type: ScheduleToStart) StackTrace coroutine root [panic]:
main.Workflow({0x104c6f648?, 0x1400036e7b0?})
        /Users/miquel/Documents/repos/poc-scheduleToClose-timeout/workflow.go:24 +0x13c
reflect.Value.call({0x104b31e00?, 0x104c62348?, 0x10552bfc8?}, {0x10498f89f, 0x4}, {0x1400000c3a8, 0x1, 0x0?})
        /usr/local/go/src/reflect/value.go:586 +0x838
reflect.Value.Call({0x104b31e00?, 0x104c62348?, 0x1400049dcd8?}, {0x1400000c3a8?, 0x1400049dcf8?, 0x104961730?})
        /usr/local/go/src/reflect/value.go:370 +0x90
go.temporal.io/sdk/internal.executeFunction({0x104b31e00, 0x104c62348}, {0x1400049dde8, 0x1, 0x18?})
        /Users/miquel/go/pkg/mod/go.temporal.io/sdk@v1.21.1/internal/internal_worker.go:1729 +0x2d0
go.temporal.io/sdk/internal.(*workflowEnvironmentInterceptor).ExecuteWorkflow(0x140002fa0f0, {0x104c6f4c0?, 0x140002d65a0}, 0x1400000c390)
        /Users/miquel/go/pkg/mod/go.temporal.io/sdk@v1.21.1/internal/workflow.go:506 +0x13c
go.temporal.io/sdk/internal.(*workflowExecutor).Execute(0x1400004a200, {0x104c6f4c0, 0x140002d65a0}, 0x25?)
        /Users/miquel/go/pkg/mod/go.temporal.io/sdk@v1.21.1/internal/internal_worker.go:780 +0x230
go.temporal.io/sdk/internal.(*syncWorkflowDefinition).Execute.func1({0x104c6f648, 0x1400036e690})
        /Users/miquel/go/pkg/mod/go.temporal.io/sdk@v1.21.1/internal/internal_workflow.go:507 +0xd0
2023/03/14 11:28:00 WARN  Failed to process workflow task. Namespace testing TaskQueue workflow WorkerID 32992@DD-Miquel-Puig-Mena.local@ WorkflowType Workflow WorkflowID something RunID 32685a11-843c-4203-bab5-708f0982e970 Attempt 2 Error activity error (type: Activity, scheduledEventID: 5, startedEventID: 0, identity: ): activity ScheduleToStart timeout (type: ScheduleToStart)
2023/03/14 11:28:10 ERROR Workflow panic Namespace testing TaskQueue workflow WorkerID 32992@DD-Miquel-Puig-Mena.local@ WorkflowType Workflow WorkflowID something RunID 32685a11-843c-4203-bab5-708f0982e970 Attempt 3 Error activity error (type: Activity, scheduledEventID: 5, startedEventID: 0, identity: ): activity ScheduleToStart timeout (type: ScheduleToStart) StackTrace coroutine root [panic]:
main.Workflow({0x104c6f648?, 0x1400036f620?})
        /Users/miquel/Documents/repos/poc-scheduleToClose-timeout/workflow.go:24 +0x13c
reflect.Value.call({0x104b31e00?, 0x104c62348?, 0x10552bfc8?}, {0x10498f89f, 0x4}, {0x1400000c930, 0x1, 0x0?})
        /usr/local/go/src/reflect/value.go:586 +0x838
reflect.Value.Call({0x104b31e00?, 0x104c62348?, 0x1400049ccd8?}, {0x1400000c930?, 0x1400049ccf8?, 0x104961730?})
        /usr/local/go/src/reflect/value.go:370 +0x90
go.temporal.io/sdk/internal.executeFunction({0x104b31e00, 0x104c62348}, {0x1400049cde8, 0x1, 0x18?})
        /Users/miquel/go/pkg/mod/go.temporal.io/sdk@v1.21.1/internal/internal_worker.go:1729 +0x2d0
go.temporal.io/sdk/internal.(*workflowEnvironmentInterceptor).ExecuteWorkflow(0x140000adae0, {0x104c6f4c0?, 0x140000953e0}, 0x1400000c918)
        /Users/miquel/go/pkg/mod/go.temporal.io/sdk@v1.21.1/internal/workflow.go:506 +0x13c
go.temporal.io/sdk/internal.(*workflowExecutor).Execute(0x140000b7800, {0x104c6f4c0, 0x140000953e0}, 0x25?)
        /Users/miquel/go/pkg/mod/go.temporal.io/sdk@v1.21.1/internal/internal_worker.go:780 +0x230
go.temporal.io/sdk/internal.(*syncWorkflowDefinition).Execute.func1({0x104c6f648, 0x140001eb7a0})
        /Users/miquel/go/pkg/mod/go.temporal.io/sdk@v1.21.1/internal/internal_workflow.go:507 +0xd0
2023/03/14 11:28:10 WARN  Failed to process workflow task. Namespace testing TaskQueue workflow WorkerID 32992@DD-Miquel-Puig-Mena.local@ WorkflowType Workflow WorkflowID something RunID 32685a11-843c-4203-bab5-708f0982e970 Attempt 3 Error activity error (type: Activity, scheduledEventID: 5, startedEventID: 0, identity: ): activity ScheduleToStart timeout (type: ScheduleToStart)
2023/03/14 11:28:20 ERROR Workflow panic Namespace testing TaskQueue workflow WorkerID 32992@DD-Miquel-Puig-Mena.local@ WorkflowType Workflow WorkflowID something RunID 32685a11-843c-4203-bab5-708f0982e970 Attempt 4 Error activity error (type: Activity, scheduledEventID: 5, startedEventID: 0, identity: ): activity ScheduleToStart timeout (type: ScheduleToStart) StackTrace coroutine root [panic]:
main.Workflow({0x104c6f648?, 0x1400018f5f0?})
        /Users/miquel/Documents/repos/poc-scheduleToClose-timeout/workflow.go:24 +0x13c
reflect.Value.call({0x104b31e00?, 0x104c62348?, 0x12c901368?}, {0x10498f89f, 0x4}, {0x140000b08e8, 0x1, 0x0?})
        /usr/local/go/src/reflect/value.go:586 +0x838
reflect.Value.Call({0x104b31e00?, 0x104c62348?, 0x14000111cd8?}, {0x140000b08e8?, 0x14000111cf8?, 0x104961730?})
        /usr/local/go/src/reflect/value.go:370 +0x90
go.temporal.io/sdk/internal.executeFunction({0x104b31e00, 0x104c62348}, {0x14000111de8, 0x1, 0x18?})
        /Users/miquel/go/pkg/mod/go.temporal.io/sdk@v1.21.1/internal/internal_worker.go:1729 +0x2d0
go.temporal.io/sdk/internal.(*workflowEnvironmentInterceptor).ExecuteWorkflow(0x14000504870, {0x104c6f4c0?, 0x1400018c840}, 0x140000b08d0)
        /Users/miquel/go/pkg/mod/go.temporal.io/sdk@v1.21.1/internal/workflow.go:506 +0x13c
go.temporal.io/sdk/internal.(*workflowExecutor).Execute(0x14000482700, {0x104c6f4c0, 0x1400018c840}, 0x25?)
        /Users/miquel/go/pkg/mod/go.temporal.io/sdk@v1.21.1/internal/internal_worker.go:780 +0x230
go.temporal.io/sdk/internal.(*syncWorkflowDefinition).Execute.func1({0x104c6f648, 0x1400018f4d0})
        /Users/miquel/go/pkg/mod/go.temporal.io/sdk@v1.21.1/internal/internal_workflow.go:507 +0xd0
2023/03/14 11:28:20 WARN  Failed to process workflow task. Namespace testing TaskQueue workflow WorkerID 32992@DD-Miquel-Puig-Mena.local@ WorkflowType Workflow WorkflowID something RunID 32685a11-843c-4203-bab5-708f0982e970 Attempt 4 Error activity error (type: Activity, scheduledEventID: 5, startedEventID: 0, identity: ): activity ScheduleToStart timeout (type: ScheduleToStart)
2023/03/14 11:28:35 ERROR Workflow panic Namespace testing TaskQueue workflow WorkerID 32992@DD-Miquel-Puig-Mena.local@ WorkflowType Workflow WorkflowID something RunID 32685a11-843c-4203-bab5-708f0982e970 Attempt 5 Error activity error (type: Activity, scheduledEventID: 5, startedEventID: 0, identity: ): activity ScheduleToStart timeout (type: ScheduleToStart) StackTrace coroutine root [panic]:
main.Workflow({0x104c6f648?, 0x140005d0ae0?})
        /Users/miquel/Documents/repos/poc-scheduleToClose-timeout/workflow.go:24 +0x13c
reflect.Value.call({0x104b31e00?, 0x104c62348?, 0x12c9d7848?}, {0x10498f89f, 0x4}, {0x140002ffc38, 0x1, 0x0?})
        /usr/local/go/src/reflect/value.go:586 +0x838
reflect.Value.Call({0x104b31e00?, 0x104c62348?, 0x1400010ecd8?}, {0x140002ffc38?, 0x1400010ecf8?, 0x104961730?})
        /usr/local/go/src/reflect/value.go:370 +0x90
go.temporal.io/sdk/internal.executeFunction({0x104b31e00, 0x104c62348}, {0x1400010ede8, 0x1, 0x18?})
        /Users/miquel/go/pkg/mod/go.temporal.io/sdk@v1.21.1/internal/internal_worker.go:1729 +0x2d0
go.temporal.io/sdk/internal.(*workflowEnvironmentInterceptor).ExecuteWorkflow(0x140000adf40, {0x104c6f4c0?, 0x140000956e0}, 0x140002ffc20)
        /Users/miquel/go/pkg/mod/go.temporal.io/sdk@v1.21.1/internal/workflow.go:506 +0x13c
go.temporal.io/sdk/internal.(*workflowExecutor).Execute(0x140005d2000, {0x104c6f4c0, 0x140000956e0}, 0x25?)
        /Users/miquel/go/pkg/mod/go.temporal.io/sdk@v1.21.1/internal/internal_worker.go:780 +0x230
go.temporal.io/sdk/internal.(*syncWorkflowDefinition).Execute.func1({0x104c6f648, 0x140005d09c0})
        /Users/miquel/go/pkg/mod/go.temporal.io/sdk@v1.21.1/internal/internal_workflow.go:507 +0xd0
2023/03/14 11:28:35 WARN  Failed to process workflow task. Namespace testing TaskQueue workflow WorkerID 32992@DD-Miquel-Puig-Mena.local@ WorkflowType Workflow WorkflowID something RunID 32685a11-843c-4203-bab5-708f0982e970 Attempt 5 Error activity error (type: Activity, scheduledEventID: 5, startedEventID: 0, identity: ): activity ScheduleToStart timeout (type: ScheduleToStart)
2023/03/14 11:28:53 ERROR Workflow panic Namespace testing TaskQueue workflow WorkerID 32992@DD-Miquel-Puig-Mena.local@ WorkflowType Workflow WorkflowID something RunID 32685a11-843c-4203-bab5-708f0982e970 Attempt 6 Error activity error (type: Activity, scheduledEventID: 5, startedEventID: 0, identity: ): activity ScheduleToStart timeout (type: ScheduleToStart) StackTrace coroutine root [panic]:
main.Workflow({0x104c6f648?, 0x1400036ec30?})
        /Users/miquel/Documents/repos/poc-scheduleToClose-timeout/workflow.go:24 +0x13c
reflect.Value.call({0x104b31e00?, 0x104c62348?, 0x12c9d7848?}, {0x10498f89f, 0x4}, {0x140002fe210, 0x1, 0x0?})
        /usr/local/go/src/reflect/value.go:586 +0x838
reflect.Value.Call({0x104b31e00?, 0x104c62348?, 0x1400010ccd8?}, {0x140002fe210?, 0x1400010ccf8?, 0x104961730?})
        /usr/local/go/src/reflect/value.go:370 +0x90
go.temporal.io/sdk/internal.executeFunction({0x104b31e00, 0x104c62348}, {0x1400010cde8, 0x1, 0x18?})
        /Users/miquel/go/pkg/mod/go.temporal.io/sdk@v1.21.1/internal/internal_worker.go:1729 +0x2d0
go.temporal.io/sdk/internal.(*workflowEnvironmentInterceptor).ExecuteWorkflow(0x1400007e5a0, {0x104c6f4c0?, 0x14000094600}, 0x140002fe1f8)
        /Users/miquel/go/pkg/mod/go.temporal.io/sdk@v1.21.1/internal/workflow.go:506 +0x13c
go.temporal.io/sdk/internal.(*workflowExecutor).Execute(0x140005d2140, {0x104c6f4c0, 0x14000094600}, 0x25?)
        /Users/miquel/go/pkg/mod/go.temporal.io/sdk@v1.21.1/internal/internal_worker.go:780 +0x230
go.temporal.io/sdk/internal.(*syncWorkflowDefinition).Execute.func1({0x104c6f648, 0x1400036eae0})
        /Users/miquel/go/pkg/mod/go.temporal.io/sdk@v1.21.1/internal/internal_workflow.go:507 +0xd0
2023/03/14 11:28:53 WARN  Failed to process workflow task. Namespace testing TaskQueue workflow WorkerID 32992@DD-Miquel-Puig-Mena.local@ WorkflowType Workflow WorkflowID something RunID 32685a11-843c-4203-bab5-708f0982e970 Attempt 6 Error activity error (type: Activity, scheduledEventID: 5, startedEventID: 0, identity: ): activity ScheduleToStart timeout (type: ScheduleToStart)
2023/03/14 11:29:20 ERROR Workflow panic Namespace testing TaskQueue workflow WorkerID 32992@DD-Miquel-Puig-Mena.local@ WorkflowType Workflow WorkflowID something RunID 32685a11-843c-4203-bab5-708f0982e970 Attempt 7 Error activity error (type: Activity, scheduledEventID: 5, startedEventID: 0, identity: ): activity ScheduleToStart timeout (type: ScheduleToStart) StackTrace coroutine root [panic]:
main.Workflow({0x104c6f648?, 0x1400036fe60?})
        /Users/miquel/Documents/repos/poc-scheduleToClose-timeout/workflow.go:24 +0x13c
reflect.Value.call({0x104b31e00?, 0x104c62348?, 0x12c9d7848?}, {0x10498f89f, 0x4}, {0x140002fef00, 0x1, 0x0?})
        /usr/local/go/src/reflect/value.go:586 +0x838
reflect.Value.Call({0x104b31e00?, 0x104c62348?, 0x1400010ccd8?}, {0x140002fef00?, 0x1400010ccf8?, 0x104961730?})
        /usr/local/go/src/reflect/value.go:370 +0x90
go.temporal.io/sdk/internal.executeFunction({0x104b31e00, 0x104c62348}, {0x1400010cde8, 0x1, 0x18?})
        /Users/miquel/go/pkg/mod/go.temporal.io/sdk@v1.21.1/internal/internal_worker.go:1729 +0x2d0
go.temporal.io/sdk/internal.(*workflowEnvironmentInterceptor).ExecuteWorkflow(0x1400007f3b0, {0x104c6f4c0?, 0x14000094de0}, 0x140002feee8)
        /Users/miquel/go/pkg/mod/go.temporal.io/sdk@v1.21.1/internal/workflow.go:506 +0x13c
go.temporal.io/sdk/internal.(*workflowExecutor).Execute(0x140005d2740, {0x104c6f4c0, 0x14000094de0}, 0x25?)
        /Users/miquel/go/pkg/mod/go.temporal.io/sdk@v1.21.1/internal/internal_worker.go:780 +0x230
go.temporal.io/sdk/internal.(*syncWorkflowDefinition).Execute.func1({0x104c6f648, 0x1400036fd40})
        /Users/miquel/go/pkg/mod/go.temporal.io/sdk@v1.21.1/internal/internal_workflow.go:507 +0xd0
2023/03/14 11:29:20 WARN  Failed to process workflow task. Namespace testing TaskQueue workflow WorkerID 32992@DD-Miquel-Puig-Mena.local@ WorkflowType Workflow WorkflowID something RunID 32685a11-843c-4203-bab5-708f0982e970 Attempt 7 Error activity error (type: Activity, scheduledEventID: 5, startedEventID: 0, identity: ): activity ScheduleToStart timeout (type: ScheduleToStart)
2023/03/14 11:30:03 ERROR Workflow panic Namespace testing TaskQueue workflow WorkerID 32992@DD-Miquel-Puig-Mena.local@ WorkflowType Workflow WorkflowID something RunID 32685a11-843c-4203-bab5-708f0982e970 Attempt 8 Error activity error (type: Activity, scheduledEventID: 5, startedEventID: 0, identity: ): activity ScheduleToStart timeout (type: ScheduleToStart) StackTrace coroutine root [panic]:
main.Workflow({0x104c6f648?, 0x14000181320?})
        /Users/miquel/Documents/repos/poc-scheduleToClose-timeout/workflow.go:24 +0x13c
reflect.Value.call({0x104b31e00?, 0x104c62348?, 0x12c9d5228?}, {0x10498f89f, 0x4}, {0x1400031a618, 0x1, 0x0?})
        /usr/local/go/src/reflect/value.go:586 +0x838
reflect.Value.Call({0x104b31e00?, 0x104c62348?, 0x14000565cd8?}, {0x1400031a618?, 0x14000565cf8?, 0x104961730?})
        /usr/local/go/src/reflect/value.go:370 +0x90
go.temporal.io/sdk/internal.executeFunction({0x104b31e00, 0x104c62348}, {0x14000565de8, 0x1, 0x18?})
        /Users/miquel/go/pkg/mod/go.temporal.io/sdk@v1.21.1/internal/internal_worker.go:1729 +0x2d0
go.temporal.io/sdk/internal.(*workflowEnvironmentInterceptor).ExecuteWorkflow(0x140002faf50, {0x104c6f4c0?, 0x140002d6d20}, 0x1400031a4f8)
        /Users/miquel/go/pkg/mod/go.temporal.io/sdk@v1.21.1/internal/workflow.go:506 +0x13c
go.temporal.io/sdk/internal.(*workflowExecutor).Execute(0x140003326c0, {0x104c6f4c0, 0x140002d6d20}, 0x25?)
        /Users/miquel/go/pkg/mod/go.temporal.io/sdk@v1.21.1/internal/internal_worker.go:780 +0x230
go.temporal.io/sdk/internal.(*syncWorkflowDefinition).Execute.func1({0x104c6f648, 0x14000181170})
        /Users/miquel/go/pkg/mod/go.temporal.io/sdk@v1.21.1/internal/internal_workflow.go:507 +0xd0
2023/03/14 11:30:03 WARN  Failed to process workflow task. Namespace testing TaskQueue workflow WorkerID 32992@DD-Miquel-Puig-Mena.local@ WorkflowType Workflow WorkflowID something RunID 32685a11-843c-4203-bab5-708f0982e970 Attempt 8 Error activity error (type: Activity, scheduledEventID: 5, startedEventID: 0, identity: ): activity ScheduleToStart timeout (type: ScheduleToStart)
```