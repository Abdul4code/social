package main

import (
	"fmt"

	"internal/env"
)

func main() {
	config := Config{
		Address: env.GetEnv("ADDRESS", ":8080"),
	}

	app := &Application{
		Config: config,
	}

	fmt.Printf("Starting server on %s\n", app.Config.Address)

	if err := app.Start(); err != nil {
		panic(err)
	}
}
