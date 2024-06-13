package main

import (
	"log"

	"github.com/azar-writes-code/traefik-temporal-poc/internal/server"
	"go.temporal.io/sdk/client"
)

func main() {
    c, err := client.Dial(client.Options{})
    if err != nil {
        log.Fatalln("Unable to create Temporal client", err)
    }
    defer c.Close()

    r := server.SetupRouter(c)
    if err := r.Run(":7200"); err != nil {
        log.Fatalln("Unable to start Gin server", err)
    }
}
