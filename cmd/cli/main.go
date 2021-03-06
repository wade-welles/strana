package main

import (
	"log"

	// Initialize processor constructors
	_ "github.com/blushft/strana/modules/init"
	_ "github.com/blushft/strana/processor/init"

	_ "github.com/joho/godotenv/autoload"

	"github.com/blushft/strana/cmd/cli/cli"
)

func main() {
	if err := cli.Execute(); err != nil {
		log.Fatal(err)
	}
}
