package main

import (
	"github.com/projectdiscovery/gologger"
	"github.com/projectdiscovery/subfinder/v2/pkg/runner"
)

func main() {
	// Parse the command line flags and read config files
	options := runner.ParseOptions()

	// Tạo một instance của Runner
	newRunner, err := runner.NewRunner(options)
	if err != nil {
		gologger.Fatal().Msgf("Could not create runner: %s\n", err)
	}

	// Tạo một API server sử dụng Runner
	apiServer := runner.NewAPIServer(newRunner)

	// Khởi động API server trên cổng 8080
	gologger.Info().Msg("Starting API server on port 8080")
	apiServer.Start("8080")
}
