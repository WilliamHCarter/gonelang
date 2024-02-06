// gone/main.go
package main

import (
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: gone <command>")
	}

	command := os.Args[1]

	switch command {
	case "compile":
		gone.Compile()
	// Add cases for other commands if needed.
	default:
		log.Fatalf("Unknown command: %s", command)
	}
}
