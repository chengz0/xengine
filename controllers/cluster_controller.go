package controllers

import (
	"bytes"
	"encoding/json"
	"github.com/chengz0/xengine/global"
	"github.com/chengz0/xengine/models"
	"github.com/deepglint/glog"
	"github.com/go-martini/martini"
	"net/http"
	"regexp"
	// "sync"
	// "strings"
)

var (
	IPregex *regexp.Regexp
)

type hoststatus struct {
	ContainerId string
	Name        string
	Status      string
}

type clusterstatus struct {
	HostIp   string
	SensorId string
	Status   bool
}

func AddHost(resp http.ResponseWriter, req *http.Request) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(req.Body)
	host := new(models.HostModel)
	err := json.Unmarshal(buf.Bytes(), host)
	if err != nil {
		glog.Errorf("Error parsing host: %s", err.Error())
		resp.WriteHeader(400)
		return
	}
	err = host.Save(global.HostsCollection)
	if err != nil {
		glog.Errorf("Error saving host: %s", err.Error())
		resp.WriteHeader(500)
		return
	}
	// add new hostinfo
	global.ClusterGroup.Add(1)
	global.InitClient(host.HostIp)
	global.ClusterGroup.Wait()
	// add new hostmodel
	global.ClusterHosts = append(global.ClusterHosts, *host)
	// add new listener
	go InitEventListener(global.HostsInfo[host.HostIp])

	resp.WriteHeader(204)
}

func GetHosts(resp http.ResponseWriter) {
	hosts, err := models.Hosts(global.HostsCollection)
	if err != nil {
		glog.Errorf("Error getting host by id: %s", err)
		resp.WriteHeader(500)
		return
	}
	resp.WriteHeader(200)
	body, _ := json.Marshal(hosts)
	resp.Write(body)
}

func DelHost(req *http.Request, resp http.ResponseWriter) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(req.Body)
	param := buf.String()
	flag := IPregex.MatchString(param)
	var host models.HostModel
	var err error
	if flag {
		host, err = models.HostByHostIp(global.HostsCollection, param)
		if err != nil {
			glog.Errorf("Error hostip: %s", err)
			resp.WriteHeader(500)
			return
		}
	} else {
		host, err = models.HostBySensorId(global.HostsCollection, param)
		if err != nil {
			glog.Errorf("Error sensorid: %s", err)
			resp.WriteHeader(500)
			return
		}
	}

	host.Delete(global.HostsCollection)
	resp.WriteHeader(204)
}

func GetHostStatus(params martini.Params, resp http.ResponseWriter) {
	hostip := params["hostip"]
	host := make([]hoststatus, 0)
	for _, container := range global.HostsInfo[hostip].Containers {
		hs := new(hoststatus)
		hs.ContainerId = container.ID
		hs.Name = global.ContainerStatus[container.ID].ContainerName
		hs.Status = global.ContainerStatus[container.ID].Status
		host = append(host, *hs)
	}
	body, _ := json.Marshal(host)
	resp.Write(body)
}

func GetClusterStatus(resp http.ResponseWriter) {
	cluster := make([]clusterstatus, 0)
	for _, host := range global.ClusterHosts {
		cs := new(clusterstatus)
		cs.HostIp = host.HostIp
		cs.SensorId = host.SensorId
		cs.Status = global.HostsInfo[host.HostIp].Status
		cluster = append(cluster, *cs)
	}
	body, _ := json.Marshal(cluster)
	resp.Write(body)
}
