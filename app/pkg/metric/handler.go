package metric

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"os"
)

const (
	URL = "/api/heartbeat"
)

type Handler struct {
}

func (h *Handler) Register(router *httprouter.Router) {
	router.HandlerFunc(http.MethodGet, URL, h.Heartbeat)
}

// Heartbeat
// @Summary Heartbeat metrics
// @Tags metrics
// @Success 204
// @Failure 404
// @Router /api/heartbeat [get]
func (h *Handler) Heartbeat(w http.ResponseWriter, r *http.Request) {
	user := os.Getenv("USER")
	fmt.Fprintf(w, "Hello, %s!", user)

	//w.WriteHeader(http.StatusNoContent)
}
