package main

import (
	"log"

	"github.com/ultimathul3/sea-battle-server/internal/app"
	"github.com/ultimathul3/sea-battle-server/internal/config"
)

func main() {
	cfg, err := config.ReadEnvFile()
	if err != nil {
		log.Fatal(err)
	}

	if err = app.Run(cfg); err != nil {
		log.Fatal(err)
	}
}
