package cmd

import (
	"log"
	"os"

	"github.com/farazff/IoT-parking/cmd/serve"
)

// Execute - Entrypoint for cli
func Execute() {
	if err := serve.Register().Execute(); err != nil {
		log.Printf("failed to execute root command: %s\n", err.Error())
		os.Exit(1)
	}
}
