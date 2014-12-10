package controllers

import (
	// "bytes"
	"encoding/json"
	"errors"
	"github.com/chengz0/xengine/global"
	"github.com/deepglint/glog"
	"github.com/deepglint/go-dockerclient"
	"github.com/go-martini/martini"
	"net/http"
)

func GetImages(params martini.Params, resp http.ResponseWriter) {
	client, err := CheckParams(params)
	if err != nil {
		glog.Errorln(err.Error())
		resp.WriteHeader(400)
		return
	}
	images, err := client.ListImages(false)
	if err != nil {
		glog.Errorf("Can not list images: %s", err.Error())
		resp.WriteHeader(500)
		return
	}
	body, _ := json.Marshal(images)
	resp.Write(body)
	// resp.WriteHeader(200)
}

func DelImageById(params martini.Params, resp http.ResponseWriter) {
	client, err := CheckParams(params)
	if err != nil {
		glog.Errorln(err.Error())
		resp.WriteHeader(400)
		return
	}
	if id, ok := params["id"]; ok {
		err := client.RemoveImage(id)
		if err != nil {
			glog.Errorf("Image not found: %s", err)
			resp.WriteHeader(404)
			return
		}
		resp.WriteHeader(204)
		return
	}
	resp.WriteHeader(400)
}

func CheckParams(params martini.Params) (*docker.Client, error) {
	if hostip, ipok := params["hostip"]; ipok {
		if hostinfo, cok := global.HostsInfo[hostip]; cok {
			client := hostinfo.Client
			return client, nil
		}
		errclient := errors.New("Error client")
		glog.Errorln(errclient.Error())
		return nil, errclient
	}
	errhost := errors.New("Error hostip param!")
	glog.Errorln(errhost.Error())
	return nil, errhost
}

func CheckParamsWithId(params martini.Params) (*docker.Client, error) {
	hostip, ipok := params["hostip"]
	_, idok := params["id"]
	if ipok && idok {
		if hostinfo, cok := global.HostsInfo[hostip]; cok {
			client := hostinfo.Client
			return client, nil
		}
		errclient := errors.New("Error client")
		glog.Errorln(errclient.Error())
		return nil, errclient
	}
	errhost := errors.New("Error hostip & id params!")
	glog.Errorln(errhost.Error())
	return nil, errhost
}
