package models

import (
	"github.com/deepglint/glog"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"strings"
)

type HostModel struct {
	Id       bson.ObjectId "_id"
	HostIp   string
	SensorId string
	HostType string
	ParentId string
	// Status   bool
}

/*
	controller provider for controller
*/
func Hosts(db *mgo.Collection) ([]HostModel, error) {
	var hosts []HostModel
	err := db.Find(bson.M{}).All(&hosts)
	if err != nil {
		glog.Errorf("Error listing hosts: %s", err.Error())
		return hosts, err
	}
	return hosts, nil
}

func HostByHostIp(db *mgo.Collection, hostip string) (HostModel, error) {
	var host HostModel
	err := db.Find(bson.M{"hostip": hostip}).One(&host)
	if err != nil {
		glog.Errorf("Error finding hostip: %s", err.Error())
		return host, nil
	}
	return host, nil
}

func HostBySensorId(db *mgo.Collection, sensorid string) (HostModel, error) {
	var host HostModel
	err := db.Find(bson.M{"sensorid": sensorid}).One(&host)
	if err != nil {
		glog.Errorf("Error finding hostip: %s", err.Error())
		return host, nil
	}
	return host, nil
}

func HostsByConditions(db *mgo.Collection, conditions map[string]interface{}) ([]HostModel, error) {
	var selector []bson.M
	for key, value := range conditions {
		if value == nil {
			continue
		}
		switch strings.ToLower(key) {
		case "hostip":
			selector = append(selector, bson.M{"hostip": value.(string)})
		case "sensorid":
			selector = append(selector, bson.M{"sensorid": value.(string)})
		case "hosttype":
			selector = append(selector, bson.M{"hosttype": value.(string)})
		case "parentid":
			selector = append(selector, bson.M{"parentid": value.(string)})
		// case "status":
		// 	selector = append(selector, bson.M{"status": value.(bool)})
		default:
			continue
		}
	}
	query := make(map[string]interface{})
	if len(selector) > 0 {
		query["$and"] = selector
	}
	glog.Infof("Query map: %#v", query)
	var hosts []HostModel
	err := db.Find(query).All(&hosts)
	if err != nil {
		glog.Errorf("Error querying hosts: %s", err.Error())
		return nil, err
	}
	return hosts, nil
}

/*
	db basic CURD
*/
func (this *HostModel) Save(db *mgo.Collection) error {
	if this.Id == "" {
		this.Id = bson.NewObjectId()
	}
	err := db.Insert(&this)
	if err != nil {
		glog.Errorf("Error inserting host: %s", err.Error())
		return err
	}
	return nil
}

func (this *HostModel) Update(db *mgo.Collection) error {
	selector := bson.M{"_id": this.Id}
	update := bson.M{"$set": bson.M{"HostIp": this.HostIp, "SensorId": this.SensorId, "HostType": this.HostType, "ParentId": this.ParentId}}
	err := db.Update(selector, update)
	if err != nil {
		glog.Errorf("Error updating host: %s", err.Error())
		return err
	}
	return nil
}

func (this *HostModel) Delete(db *mgo.Collection) error {
	err := db.RemoveId(this.Id)
	if err != nil {
		glog.Errorf("Error deleting host: %s", err.Error())
		return err
	}
	this.Id = ""
	return nil
}
