package models

import (
	"github.com/deepglint/go-dockerclient"
)

type HostInfoBinding struct {
	Client     *docker.Client
	Containers []*docker.Container
	Status     bool
}

type ContainerStatusBinding struct {
	HostIp        string
	ContainerName string
	Status        int
}

const (
	CONTAINER_MISS    = 0
	CONTAINER_START   = 1
	CONTAINER_STOP    = 2
	CONTAINER_PAUSE   = 3
	CONTAINER_UNPAUSE = 4
	CONTAINER_DIE     = 5
	CONTAINER_RESTART = 6
)
