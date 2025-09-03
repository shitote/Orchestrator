package manager

import (
	"fmt"
	"orchestrator/task"

	"github.com/docker/distribution/uuid"
	"github.com/golang-collections/collections/queue"
)

type Manager struct {
	Pending queue.Queue
	TaskDb map[string][]task.Task
	EventDb map[string][]task.TaskEvent
	Workers []string
	WorkerTaskMap map[string][]uuid.UUID
	TaskWorkerMap map[uuid.UUID]string
}

func (m *Manager) SelectWorker() {
	fmt.Println("I will select an approproate workde")
}

func (m *Manager) UpdateTasks() {
	fmt.Println("I will update the tasls")
}

func (m *Manager) SendWork() {
	fmt.Println("I will send work to workers")
}