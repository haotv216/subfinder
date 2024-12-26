package runner

import (
	"encoding/json"
	"net/http"

	"github.com/projectdiscovery/gologger"
)

// APIServer defines the structure for the API server
type APIServer struct {
	runner *Runner
}

// NewAPIServer creates a new API server
func NewAPIServer(runner *Runner) *APIServer {
	return &APIServer{runner: runner}
}

// Start starts the API server
func (api *APIServer) Start(port string) {
	http.HandleFunc("/dns/api/enumerate", api.handleEnumerate)

	gologger.Info().Msgf("API server is running on port %s", port)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		gologger.Fatal().Msgf("Failed to start server: %s", err)
	}
}

// handleEnumerate handles the enumeration endpoint
func (api *APIServer) handleEnumerate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var request struct {
		Domains []string `json:"domains"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if len(request.Domains) == 0 {
		http.Error(w, "No domains provided", http.StatusBadRequest)
		return
	}

	// Process the domains
	output := make(map[string]interface{})
	for _, domain := range request.Domains {
		results := api.enumerateDomain(domain)
		output[domain] = results
	}

	response, err := json.Marshal(output)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

// enumerateDomain processes a single domain (stub example)
func (api *APIServer) enumerateDomain(domain string) interface{} {
	// Replace with actual enumeration logic
	return map[string]string{
		"domain": domain,
		"status": "success",
	}
}
