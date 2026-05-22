package handler

import (
	"net/http"

	"github.com/aithlete/aithlete-api/infrastructure/logger"
	"github.com/aithlete/aithlete-api/interfaces/router"
	"github.com/aithlete/aithlete-api/pkg/app"
)

var h http.Handler

func init() {
	log := logger.New()
	deps := app.Bootstrap(log)
	h = router.New(log, deps.Handlers)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	h.ServeHTTP(w, r)
}
