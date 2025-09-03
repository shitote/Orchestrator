package task

import (
	"time"

	"github.com/docker/go-connections/nat"
	"github.com/google/uuid"
)

type State int

const (
	Peding State = iota
	Scheduled
	Running
	Completed
	Failed
)

type Task struct {
	ID uuid.UUID
	Name string
	state State
	Image string
	Memoru int
	Disk int
	ExposedPorts nat.PortSet
	portBindings map[string]string
	RestartPolicy string
	StartTime time.Time
	FInishTIme time.Time
}

type TaskEvent struct {
	ID uuid.UUID
	State State
	TimeStamp time.Time
	Task Task
}