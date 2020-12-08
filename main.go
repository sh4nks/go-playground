package main

import (
	"html/template"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"github.com/unrolled/render"
)

func main() {
	rendr := render.New(render.Options{
		RenderPartialsWithoutPrefix: true,
		IsDevelopment:               true,
		Directory:                   "templates",
		Layout:                      "layout",
		Extensions:                  []string{".html"},
		Funcs: []template.FuncMap{
			template.FuncMap{"myFunc": func() string {
					return "My function"
				},
			},
		},
	})

	router := chi.NewRouter()
	router.Use(
		middleware.RequestID,
		middleware.Logger,
		middleware.RedirectSlashes,
		middleware.Recoverer,
	)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		data := map[string]string{"Content": "Test Content"}
		htmlOpts := render.HTMLOptions{
			Funcs: template.FuncMap{"requestFunc": func() string {
					return "My request function"
				},
			},
		}
		rendr.HTML(w, 200, "main", data, htmlOpts)
	})

	http.ListenAndServe(":8000", router)
}
