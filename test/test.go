package main

import (
	"fmt"
	"log"
	"net/http"

	// "strings"

	// "github.com/chengz0/xengine/models"
	// "gopkg.in/mgo.v2/bson"

	// "bytes"
	// "encoding/json"
	// "github.com/deepglint/go-dockerclient"
)

type HostIpId struct {
	Ip string
	Id string
}

func main() {
	// test
	client := &http.Client{}

	// host := new(models.HostModel)
	// host.Id = bson.NewObjectId()
	// host.HostIp = "192.168.2.112"
	// host.SensorId = "DG.BLADE.S12"
	// host.HostType = "sensor"
	// host.ParentId = "192.168.2.103"
	// body, _ := json.Marshal(host)
	// req, err := http.NewRequest("POST", "http://192.168.1.117:3000/host/add", bytes.NewBuffer(body))
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// req, err := http.NewRequest("GET", "http://192.168.1.117:3000/hosts", nil)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// ip := "192.168.2.109"
	// req, err := http.NewRequest("DELETE", "http://192.168.1.117:3000/host/del", strings.NewReader(ip))
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// // list images
	// req, err := http.NewRequest("GET", "http://192.168.1.117:3000/docker/images/192.168.2.106", nil)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// req, err := http.NewRequest("DELETE", "http://192.168.1.117:3000/docker/image/del/192.168.2.106/162fdef999f898c18e6bb432512a8cd6feb876dddb2a96ea1049c18d95f0943b", nil)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// // list containers
	// req, err := http.NewRequest("GET", "http://192.168.1.117:3000/docker/containers/192.168.2.106", nil)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// // stop container
	// req, err := http.NewRequest("GET", "http://192.168.1.117:3000/docker/container/stop/192.168.2.106/62aeef6218c5", nil)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// // kill container
	// req, err := http.NewRequest("GET", "http://192.168.1.117:3000/docker/container/kill/192.168.2.106/62aeef6218c5", nil)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// // pause container
	// req, err := http.NewRequest("GET", "http://192.168.1.117:3000/docker/container/pause/192.168.2.106/62aeef6218c5", nil)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// // unpause container
	// req, err := http.NewRequest("GET", "http://192.168.1.117:3000/docker/container/unpause/192.168.2.106/62aeef6218c5", nil)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// // remove container
	// req, err := http.NewRequest("DELETE", "http://192.168.1.117:3000/docker/container/remove/192.168.2.106/62aeef6218c5/true", nil)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// start container
	// opts := new(docker.HostConfig)
	// opts.Privileged = true
	// opts.RestartPolicy = docker.AlwaysRestart()
	// body, _ := json.Marshal(opts)
	// req, err := http.NewRequest("POST", "http://192.168.1.117:3000/docker/container/start/192.168.2.106/62aeef6218c5", bytes.NewReader(body))
	// if err != nil {
	// 	fmt.Println(err)
	// }

	response, err := client.Do(req)
	// buf := new(bytes.Buffer)
	// buf.ReadFrom(response.Body)
	// log.Println(response.Status, "======", buf.String())
	log.Println(response.StatusCode)
	defer response.Body.Close()

}
