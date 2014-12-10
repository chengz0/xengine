package global

import (
	"github.com/chengz0/xengine/models"
	"github.com/deepglint/glog"
	"github.com/deepglint/go-dockerclient"
	"gopkg.in/mgo.v2"
	"strings"
)

var (
	// cluster_db
	HostsCollection *mgo.Collection

	// host_clients
	// Clients map[string]*docker.Client

	// host_containers
	HostsInfo map[string]*models.HostInfoBinding

	// containerId_status
	ContainerStatus map[string]models.ContainerStatusBinding
)

func InitGlobal(db *mgo.Database, collection string) {

	// db collection
	HostsCollection = db.C(collection)

	// new client
	hosts, err := models.Hosts(HostsCollection)
	if err != nil {
		glog.Fatalf("Error getting hosts: %s", err)
	}

	// init host_containers
	HostsInfo = make(map[string]*models.HostInfoBinding)

	// init container_status
	ContainerStatus = make(map[string]models.ContainerStatusBinding)

	for _, host := range hosts {
		go InitClient(host.HostIp)
	}
}

func InitClient(hostip string) {
	glog.Infof("Parsing new host: %s", hostip)
	// new client
	hostinfo := new(models.HostInfoBinding)
	// hostinfo status
	hostinfo.Status = true

	// init host_client
	client, err := docker.NewClient("http://" + hostip + ":4243")
	if err != nil {
		glog.Errorf("Error on host: %s creating client: %s", hostip, err.Error())
		hostinfo.Status = false
		HostsInfo[hostip] = hostinfo
		return
	}
	hostinfo.Client = client

	apicontainers, _ := client.ListContainers(docker.ListContainersOptions{
		All: true,
	})
	containers := make([]*docker.Container, 0)
	for _, apicontainers := range apicontainers {
		container, _ := client.InspectContainer(apicontainers.ID)

		//containerstatus of host
		containerstatus := new(models.ContainerStatusBinding)
		containerstatus.HostIp = hostip
		containerstatus.ContainerName = strings.Replace(container.Name, "/", "", -1)
		if container.State.Running {
			containerstatus.Status = models.CONTAINER_START
		} else {
			containerstatus.Status = models.CONTAINER_STOP
			hostinfo.Status = false
		}
		ContainerStatus[container.ID] = *containerstatus

		// hostinfo containers
		containers = append(containers, container)
	}
	hostinfo.Containers = containers

	HostsInfo[hostip] = hostinfo
}
