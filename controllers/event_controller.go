package controllers

import (
	// "encoding/json"
	"github.com/chengz0/xengine/global"
	"github.com/chengz0/xengine/models"
	"github.com/deepglint/glog"
	// "sync"
)

func InitAllListener() {
	glog.Infoln("Initing docker listener...")
	// var hostgroup sync.WaitGroup
	// hostgroup.Add(len(global.HostsInfo))
	for _, hostinfo := range global.HostsInfo {
		// go InitEventListener(hostinfo, hostgroup)
		go InitEventListener(hostinfo)
	}
	// hostgroup.Done()
}

func InitEventListener(hostinfo *models.HostInfoBinding) {
	// wg.Done()

	// add listener
	err := hostinfo.Client.AddEventListener(hostinfo.Listener)
	if err != nil {
		glog.Errorf("Error adding docker listener: %s", err.Error())
		return
	}
	defer func() {
		err = hostinfo.Client.RemoveEventListener(hostinfo.Listener)
		if err != nil {
			glog.Errorf("Error removing docker listener: %s", err.Error())
		}
	}()
	glog.Infoln(hostinfo.Status)
	// listening events
	for {
		select {
		case msg := <-hostinfo.Listener:
			glog.Infoln(msg)
			// body, err := json.Marshal(msg)
			// if err != nil {
			// 	glog.Errorf("Can not parse apievent: %s", err.Error())
			// 	break
			// }
			hostinfo.SyncMutex.Lock()
			global.ContainerStatus[msg.ID].Status = msg.Status
			hostinfo.SyncMutex.Unlock()
		case <-global.Timeout:
			break
		}
	}
}
