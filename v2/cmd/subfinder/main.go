package main

import (
	"github.com/projectdiscovery/gologger"
	"github.com/projectdiscovery/subfinder/v2/pkg/runner"
)

func main() {
	// Parse the command line flags and read config files
	options := runner.ParseOptions()

	// Create an instance of Runner
	newRunner, err := runner.NewRunner(options)
	if err != nil {
		gologger.Fatal().Msgf("Could not create runner: %s\n", err)
	}

	// Create an API server using Runner
	apiServer := runner.NewAPIServer(newRunner)

	// Start the API server on port 8080
	gologger.Info().Msg("Starting API server on port 8080")
	apiServer.Start("8080")
}
