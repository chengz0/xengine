package main

import (
	"flag"
	"github.com/chengz0/xengine/controllers"
	"github.com/chengz0/xengine/http"
	"github.com/deepglint/glog"
	"gopkg.in/mgo.v2"
	"regexp"
)

type Config struct {
	DBServer   string
	Database   string
	Collection string
	User       string
	Passwd     string
}

func main() {
	var config Config
	flag.StringVar(&config.DBServer, "server", "127.0.0.1:27017", "db server")
	flag.StringVar(&config.Database, "db", "bone", "database")
	flag.StringVar(&config.Collection, "c", "hosts", "collection")
	flag.StringVar(&config.User, "user", "bone", "username")
	flag.StringVar(&config.Passwd, "passwd", "deepdbdb", "password")
	flag.Parse()

	/*
		init host db
	*/
	glog.Info(config)
	session, err := mgo.Dial(config.DBServer)
	if err != nil {
		glog.Fatalf("Error dialing db: %s", err.Error())
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	db := session.DB(config.Database)
	err = db.Login(config.User, config.Passwd)
	if err != nil {
		glog.Fatalf("Error auth db: %s", err.Error())
	}
	controllers.HostsCollection = db.C(config.Collection)

	controllers.IPregex = regexp.MustCompile("^[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}$")

	// martini
	http.StartMartini()
}
