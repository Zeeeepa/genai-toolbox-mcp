package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

// webRouter creates a router that represents the routes under /ui
func webRouter() (http.Handler, error) {
	r := chi.NewRouter()

	r.Use(middleware.AllowContentType("application/json"))
	r.Use(middleware.StripSlashes)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Get("/homepage", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("UI Homepage"))
	})

	r.Get("/sources", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("Sources View"))
	})

	r.Get("/tools", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("Tools View"))
	})

	return r, nil
}
