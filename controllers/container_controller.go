package controllers

import (
	"encoding/json"
	"github.com/deepglint/glog"
	"github.com/deepglint/go-dockerclient"
	"github.com/go-martini/martini"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

/*
	list containers of host
*/
func ListContainers(params martini.Params, resp http.ResponseWriter) {
	client, err := CheckParams(params)
	if err != nil {
		glog.Errorln(err.Error())
		resp.WriteHeader(400)
		return
	}
	opts := new(docker.ListContainersOptions)
	opts.All = true
	for key, value := range params {
		switch strings.ToLower(key) {
		case "all":
			all, err := strconv.ParseBool(value)
			if err != nil {
				resp.WriteHeader(400)
				return
			}
			opts.All = all
		case "size":
			size, err := strconv.ParseBool(value)
			if err != nil {
				resp.WriteHeader(400)
				return
			}
			opts.Size = size
		case "limit":
			limit, err := strconv.Atoi(value)
			if err != nil {
				resp.WriteHeader(400)
				return
			}
			opts.Limit = limit
		case "since":
			opts.Since = value
		case "before":
			opts.Before = value
		default:
			continue
		}
	}
	containers, err := client.ListContainers(*opts)
	if err != nil {
		resp.WriteHeader(500)
		return
	}
	body, _ := json.Marshal(containers)
	resp.Write(body)
}

/*
	inspect container of host by id
*/
func InspectContainerById(params martini.Params, resp http.ResponseWriter) {
	client, err := CheckParamsWithId(params)
	if err != nil {
		glog.Errorln(err.Error())
		resp.WriteHeader(400)
		return
	}
	id := params["id"]
	container, err := client.InspectContainer(id)
	if err != nil {
		glog.Errorf("Error inspecting container: %s", err)
		resp.WriteHeader(500)
		return
	}
	body, _ := json.Marshal(container)
	resp.Write(body)
}

/*
	start container of host by id
*/
func StartContainerById(params martini.Params, req *http.Request, resp http.ResponseWriter) {
	client, err := CheckParamsWithId(params)
	if err != nil {
		glog.Errorln(err.Error())
		resp.WriteHeader(400)
		return
	}
	id := params["id"]
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		glog.Errorf("Error req: %s", err.Error())
		resp.WriteHeader(400)
		return
	}

	if len(body) == 0 {
		err = client.StartContainer(id, nil)
		if err != nil {
			glog.Errorf("Error hostconfig: %s", err.Error())
			resp.WriteHeader(500)
			return
		}
		resp.WriteHeader(204)
		return
	}

	opts := new(docker.HostConfig)
	err = json.Unmarshal(body, &opts)
	if err != nil {
		glog.Errorf("Error starting opts: %s", err)
		resp.WriteHeader(400)
		return
	}
	err = client.StartContainer(id, opts)
	if err != nil {
		glog.Errorf("Error starting container: %s", err.Error())
		resp.WriteHeader(500)
		return
	}
	resp.WriteHeader(204)
}

/*
	stop container of host by id
*/
func StopContainerById(params martini.Params, resp http.ResponseWriter) {
	client, err := CheckParamsWithId(params)
	if err != nil {
		glog.Errorln(err.Error())
		resp.WriteHeader(400)
		return
	}
	id := params["id"]
	var timeout uint
	if timeoutstr, ok := params["timeout"]; ok {
		timeoutint, err := strconv.Atoi(timeoutstr)
		if err != nil {
			glog.Errorf("Error timeout: %s", err)
			resp.WriteHeader(400)
			return
		}
		timeout = uint(timeoutint)
	} else {
		timeout = 0
	}

	err = client.StopContainer(id, timeout)
	if err != nil {
		glog.Errorf("Error stopping container: %s", err)
		resp.WriteHeader(500)
		return
	}
	resp.WriteHeader(204)
}

/*
	kill container of host by id
*/
func KillContainerById(params martini.Params, resp http.ResponseWriter) {
	client, err := CheckParamsWithId(params)
	if err != nil {
		glog.Errorln(err.Error())
		resp.WriteHeader(400)
		return
	}
	id := params["id"]
	opts := new(docker.KillContainerOptions)
	opts.ID = id
	opts.Signal = docker.SIGKILL
	err = client.KillContainer(*opts)
	if err != nil {
		glog.Errorf("Error killing container: %s", err)
		resp.WriteHeader(500)
		return
	}
	resp.WriteHeader(204)
}

/*
	remove container of host by id
*/
func DelContainerById(params martini.Params, resp http.ResponseWriter) {
	client, err := CheckParamsWithId(params)
	if err != nil {
		glog.Errorln(err.Error())
		resp.WriteHeader(400)
		return
	}

	opts := new(docker.RemoveContainerOptions)
	opts.RemoveVolumes = false
	opts.Force = true
	for key, value := range params {
		switch strings.ToLower(key) {
		case "id":
			opts.ID = value
		case "removevolumes":
			removevolumes, err := strconv.ParseBool(value)
			if err != nil {
				glog.Errorf("Error removevolumes: %s", err)
				resp.WriteHeader(400)
				return
			}
			opts.RemoveVolumes = removevolumes
		case "force":
			force, err := strconv.ParseBool(value)
			if err != nil {
				glog.Errorf("Error force: %s", err)
				resp.WriteHeader(400)
				return
			}
			opts.Force = force
		default:
			continue
		}
	}
	err = client.RemoveContainer(*opts)
	if err != nil {
		glog.Errorf("Error removing running container: %s", err.Error())
		resp.WriteHeader(500)
		return
	}
	resp.WriteHeader(204)
}

/*
	restart container of host by id
*/
func RestartContainerById(params martini.Params, resp http.ResponseWriter) {
	client, err := CheckParamsWithId(params)
	if err != nil {
		glog.Errorln(err.Error())
		resp.WriteHeader(400)
		return
	}
	id := params["id"]
	var timeout uint
	if timeoutstr, ok := params["timeout"]; ok {
		timeoutint, err := strconv.Atoi(timeoutstr)
		if err != nil {
			glog.Errorf("Error timeout: %s", err)
			resp.WriteHeader(400)
			return
		}
		timeout = uint(timeoutint)
	} else {
		timeout = 0
	}
	err = client.RestartContainer(id, timeout)
	if err != nil {
		glog.Errorf("Error restarting container: %s", err.Error())
		resp.WriteHeader(500)
		return
	}
	resp.WriteHeader(204)
}

/*
	pause container of host by id
*/
func PauseContainerById(params martini.Params, resp http.ResponseWriter) {
	client, err := CheckParamsWithId(params)
	if err != nil {
		glog.Errorln(err.Error())
		resp.WriteHeader(400)
		return
	}
	id := params["id"]
	err = client.PauseContainer(id)
	if err != nil {
		glog.Errorf("Error pausing container: %s", err.Error())
		resp.WriteHeader(500)
		return
	}
	resp.WriteHeader(204)
}

/*
	unpause container of host by id
*/
func UnpauseContainerById(params martini.Params, resp http.ResponseWriter) {
	client, err := CheckParamsWithId(params)
	if err != nil {
		glog.Errorln(err.Error())
		resp.WriteHeader(400)
		return
	}
	id := params["id"]
	err = client.UnpauseContainer(id)
	if err != nil {
		glog.Errorf("Error unpausing container: %s", err.Error())
		resp.WriteHeader(500)
		return
	}
	resp.WriteHeader(204)
}
