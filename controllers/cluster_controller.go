package controllers

import (
	"bytes"
	"encoding/json"
	"github.com/chengz0/xengine/models"
	"github.com/deepglint/glog"
	// "github.com/go-martini/martini"
	"gopkg.in/mgo.v2"
	// "gopkg.in/mgo.v2/bson"
	"net/http"
	"regexp"
)

var (
	HostsCollection *mgo.Collection
	IPregex         *regexp.Regexp
)

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
	err = host.Save(HostsCollection)
	if err != nil {
		glog.Errorf("Error saving host: %s", err.Error())
		resp.WriteHeader(500)
		return
	}
	resp.WriteHeader(204)
}

func GetHosts(resp http.ResponseWriter) {
	hosts, err := models.Hosts(HostsCollection)
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
	if flag {
		host, err := models.HostByHostIp(HostsCollection, param)
		if err != nil {
			glog.Errorf("Error hostip: %s", err)
			resp.WriteHeader(500)
			return
		}
		host.Delete(HostsCollection)
		resp.WriteHeader(204)
	} else {
		host, err := models.HostBySensorId(HostsCollection, param)
		if err != nil {
			glog.Errorf("Error sensorid: %s", err)
			resp.WriteHeader(500)
			return
		}
		host.Delete(HostsCollection)
		resp.WriteHeader(204)
	}
}
