package global

import (
	"github.com/chengz0/xengine/models"
	"github.com/deepglint/glog"
	"github.com/deepglint/go-dockerclient"
	"gopkg.in/mgo.v2"
	"strings"
	"sync"
	"time"
)

var (
	// global timeout
	Timeout = time.After(2 * time.Second)

	// hosts in cluster
	ClusterHosts []models.HostModel

	// cluster_db
	HostsCollection *mgo.Collection

	// host_clients
	// Clients map[string]*docker.Client

	// host_containers
	HostsInfo map[string]*models.HostInfoBinding

	// containerId_status
	ContainerStatus map[string]*models.ContainerStatusBinding

	ClusterGroup sync.WaitGroup
)

func InitGlobal(db *mgo.Database, collection string) {

	// db collection
	HostsCollection = db.C(collection)

	// init hosts
	ClusterHosts, _ = models.Hosts(HostsCollection)
	// glog.Infoln(ClusterHosts)

	// init host_containers
	HostsInfo = make(map[string]*models.HostInfoBinding)

	// init container_status
	ContainerStatus = make(map[string]*models.ContainerStatusBinding)

	ClusterGroup.Add(len(ClusterHosts))
	for _, host := range ClusterHosts {
		go InitClient(host.HostIp)
	}
	ClusterGroup.Wait()
	glog.Infoln("Done with client.")
}

func InitClient(hostip string) {
	glog.Infof("Parsing new host: %s", hostip)
	defer ClusterGroup.Done()

	// new client
	hostinfo := new(models.HostInfoBinding)
	// hostinfo status
	hostinfo.Status = true

	// init host_client for hostinfo
	client, err := docker.NewClient("http://" + hostip + ":4243")
	if err != nil {
		glog.Errorf("Error on host: %s creating client: %s", hostip, err.Error())
		hostinfo.Status = false
		HostsInfo[hostip] = hostinfo
		return
	}
	hostinfo.Client = client

	// hostinfo containers
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
			containerstatus.Status = "start"
		} else {
			containerstatus.Status = "stop"
			hostinfo.Status = false
		}
		ContainerStatus[container.ID] = containerstatus

		containers = append(containers, container)
	}
	hostinfo.Containers = containers

	// hostinfo listener
	listener := make(chan *docker.APIEvents)
	hostinfo.Listener = listener

	// hostinfo sync
	syncmutex := new(sync.Mutex)
	hostinfo.SyncMutex = syncmutex

	HostsInfo[hostip] = hostinfo
	glog.Infoln(ClusterGroup)
}
