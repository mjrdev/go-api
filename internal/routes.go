package router

import (
	"github.com/go-chi/chi/v5"
	apiController "github.com/mjrdev/api-go/internal/controller"
)

func Listen(r *chi.Mux) {
	r.Get("/", apiController.Hello)
}
