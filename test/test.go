package main

import (
	"fmt"
	"log"
	"net/http"

	// "bytes"

	"strings"

	// "encoding/json"
	// "github.com/chengz0/xengine/models"
	// "gopkg.in/mgo.v2/bson"
)

func main() {
	// test
	client := &http.Client{}

	// host := new(models.HostModel)
	// host.Id = bson.NewObjectId()
	// host.HostIp = "192.168.2.107"
	// host.SensorId = "DG.BLADE.S07"
	// host.HostType = "sensor"
	// host.ParentId = "192.168.2.103"
	// body, _ := json.Marshal(host)
	// req, err := http.NewRequest("POST", "http://192.168.1.117:3000/hosts/add", bytes.NewBuffer(body))
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// req, err := http.NewRequest("GET", "http://192.168.1.117:3000/hosts", nil)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	ip := "192.168.2.109"
	req, err := http.NewRequest("DELETE", "http://192.168.1.117:3000/hosts/del", strings.NewReader(ip))
	if err != nil {
		fmt.Println(err)
	}

	response, err := client.Do(req)
	// buf := new(bytes.Buffer)
	// buf.ReadFrom(response.Body)
	// log.Println(response.Status, "======", buf.String())
	log.Println(response.StatusCode)
	defer response.Body.Close()

}
