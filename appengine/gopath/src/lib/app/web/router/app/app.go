package app

import (
	"lib/app/web/controllers/app"
	"lib/app/web/router"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Setup 開演
func Setup() {
	time.Local = time.FixedZone("Asia/Tokyo", 9*60*60)

	r := mux.NewRouter()

	// r.NotFoundHandler = NotFoundHandler

	router.InPublic(r, func(r *router.Router) {
		r.Get("/", app.Root)
		r.Post("/auto", app.Auto)
		r.Post("/manual", app.Manual)
	})

	handler := router.WithCSRF(router.WithContext(r))

	http.Handle("/", handler)
}
