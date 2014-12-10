package http

import (
	"github.com/chengz0/xengine/controllers"
	"github.com/go-martini/martini"
)

func ContainerRouter() {
	martini_m.Group("/docker/container", func(router martini.Router) {
		//
		router.Get("/inspect/:hostip/:id", controllers.InspectContainerById)

		router.Post("/start/:hostip/:id", controllers.StartContainerById)

		router.Get("/restart/:hostip/:id", controllers.RestartContainerById)
		router.Get("/stop/:hostip/:id", controllers.StopContainerById)
		router.Get("/kill/:hostip/:id", controllers.KillContainerById)
		router.Delete("/remove/:hostip/:id/:force", controllers.DelContainerById)

		router.Get("/pause/:hostip/:id", controllers.PauseContainerById)
		router.Get("/unpause/:hostip/:id", controllers.UnpauseContainerById)
	})
	martini_m.Get("/docker/containers/:hostip", controllers.ListContainers)
}
