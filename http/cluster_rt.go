package http

import (
	"github.com/chengz0/xengine/controllers"
	"github.com/go-martini/martini"
)

func ClusterRouter() {
	martini_m.Group("/hosts", func(router martini.Router) {
		router.Post("/add", controllers.AddHost)
		router.Delete("/del", controllers.DelHost)
	})
	martini_m.Get("/hosts", controllers.GetHosts)
}
