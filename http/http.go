package http

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"html/template"
	"net/http"
)

var (
	martini_m *martini.ClassicMartini
)

func StartMartini() {
	// martini
	martini_m = martini.Classic()
	martini_m.Use(render.Renderer(render.Options{
		Funcs: []template.FuncMap{{
			"nl2br":      nl2br,
			"htmlquote":  htmlQuote,
			"str2html":   str2html,
			"dateformat": dateFormat,
		}},
	}))

	// cluster manager
	ClusterRouter()

	http.ListenAndServe(":3000", martini_m)
}
