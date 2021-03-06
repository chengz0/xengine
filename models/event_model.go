package models

import (
	"github.com/deepglint/go-dockerclient"
	"sync"
)

type HostInfoBinding struct {
	Client     *docker.Client
	Containers []*docker.Container
	Listener   chan *docker.APIEvents
	SyncMutex  *sync.Mutex
	Status     bool
}

type ContainerStatusBinding struct {
	HostIp        string
	ContainerName string
	Status        string
}

//
const (
	CONTAINER_MISS    = 0
	CONTAINER_START   = 1
	CONTAINER_STOP    = 2
	CONTAINER_PAUSE   = 3
	CONTAINER_UNPAUSE = 4
	CONTAINER_DIE     = 5
	CONTAINER_RESTART = 6
)
