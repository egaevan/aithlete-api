package handler

import (
	"net/http"

	"github.com/aithlete/aithlete-api/infrastructure/logger"
	"github.com/aithlete/aithlete-api/interfaces/router"
)

var h http.Handler

func init() {
	log := logger.New()
	h = router.New(log)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	h.ServeHTTP(w, r)
}
