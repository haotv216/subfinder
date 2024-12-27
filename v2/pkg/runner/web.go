package runner

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
	router := gin.Default()

	// Define routes
	router.POST("/dns/api/enumerate", api.handleEnumerate)

	gologger.Info().Msgf("API server is running on port %s", port)

	if err := router.Run(":" + port); err != nil {
		gologger.Fatal().Msgf("Failed to start server: %s", err)
	}
}

// handleEnumerate handles the enumeration endpoint
func (api *APIServer) handleEnumerate(c *gin.Context) {
	var request struct {
		Domains []string `json:"domains"`
	}

	// Bind JSON body to the struct
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if len(request.Domains) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No domains provided"})
		return
	}

	// Process the domains
	output := make(map[string]interface{})
	for _, domain := range request.Domains {
		results := api.enumerateDomain(domain)
		output[domain] = results
	}

	c.JSON(http.StatusOK, output)
}

// enumerateDomain processes a single domain (stub example)
func (api *APIServer) enumerateDomain(domain string) interface{} {
	// Replace with actual enumeration logic
	return map[string]string{
		"domain": domain,
		"status": "success",
	}
}
