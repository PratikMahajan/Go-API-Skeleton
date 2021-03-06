package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Noah-Huppert/golog"


	"github.com/PratikMahajan/Go-Twitter-Downloader-Bot/config"
	"github.com/PratikMahajan/Go-Twitter-Downloader-Bot/metrics"
)

// BaseHandler provides helper methods and commonly used variables for API endpoints to base
// their http.Handlers off

type BaseHandler struct {
	// Ctx is the application context
	Ctx context.Context

	// Logger logs information
	Logger golog.Logger

	// Cfg is the application configuration
	Cfg *config.Config

	// Metrics holds Prometheus internal metrics recorders
	Metrics metrics.Metrics

}

// GetChild makes a child instance of the base handler with a prefix
func (h BaseHandler) GetChild(prefix string) BaseHandler {
	h.Logger.GetChild(prefix)

	return h
}

// RespondJSON sends an object as a JSON encoded response
func (h BaseHandler) RespondJSON(w http.ResponseWriter, status int, resp interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	encoder := json.NewEncoder(w)
	if err := encoder.Encode(resp); err != nil {
		panic(fmt.Errorf("failed to encode response as JSON: %s", err.Error()))
	}
}

// ParseJSON parses a request body as JSON
func (h BaseHandler) ParseJSON(r *http.Request, dest interface{}) {
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(dest); err != nil {
		panic(fmt.Errorf("failed to decode request body as JSON: %s", err.Error()))
	}
}

// RespondTEXT sends an object as a TEXT encoded response
func (h BaseHandler) RespondTEXT(w http.ResponseWriter, status int, resp string) {
	w.Header().Set("Content-Type", "plain/text")
	w.WriteHeader(status)

	if _, err := fmt.Fprintf(w, resp); err != nil {
		panic(fmt.Errorf("failed to write bash to body: %s", err.Error()))
	}

}