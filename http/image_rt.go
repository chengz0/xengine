package http

import (
	"github.com/chengz0/xengine/controllers"
	"github.com/go-martini/martini"
)

func ImageRouter() {
	martini_m.Group("/docker/image", func(router martini.Router) {
		// router.Post("/push", controllers.PushImage)
		router.Delete("/del/:hostip/:id", controllers.DelImageById)
		// router.Post("/tag", controllers.TagImage)
	})
	martini_m.Get("/docker/images/:hostip", controllers.GetImages)
}
