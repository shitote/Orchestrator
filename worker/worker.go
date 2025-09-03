package worker

import (
	"fmt"
	"orchestrator/task"

	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
)


type Worker struct {
	Name string
	Queue queue.Queue
	Db map[uuid.UUID]task.Task
	TaskCount int
}

func (w *Worker) CollectStats() {
	fmt.Println("I eill collect stats")
}

func (w *Worker) RunTask() {
	fmt.Println("I will start or stop a task")
}

func (w *Worker) StartTask() {
	fmt.Println("Start the task")
}

func (w *Worker) StopTask() {
	fmt.Println("I will stop a task")
}