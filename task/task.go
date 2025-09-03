package task

import (
	"context"
	"io"
	"log"
	"os"
	"time"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"github.com/google/uuid"
	specs "github.com/opencontainers/image-spec/specs-go/v1"
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
	State State
	Image string
	Memory int
	Disk int
	ExposedPorts nat.PortSet
	PortBindings map[string]string
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

type Config struct {
	Name string
	AttachStdin bool
	AttachStdout bool
	AttachStderr bool
	Cmd []string
	Image string
	Memory int64
	Disk int64
	Env []string
	RestartPolicy string
}

type Docker struct {
	Client *client.Client
	Config Config
	ContainerId string
}

type DockerResult struct {
	Error error
	Action string
	ContainerId string
	Result string
}

func (d *Docker) Run() DockerResult {
	ctx := context.Background()
	reader, err := d.Client.ImagePull(
		ctx, d.Config.Image, image.PullOptions{})
	if err != nil {
		log.Printf("Error pulling image %s: %v\n", d.Config.Image, err)
		return DockerResult{Error: err}
	}
	io.Copy(os.Stdout, reader)
	return DockerResult{}
}

type Client struct {
    *client.Client
}

func (cli *Client) ContainerCreate(
ctx context.Context,
config *container.Config,
hostConfig *container.HostConfig,
networkingConfig *network.NetworkingConfig,
platform *specs.Platform,
containerName string) (container.CreateResponse, error)