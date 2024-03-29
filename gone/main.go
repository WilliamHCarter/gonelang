// gone/main.go
package main

import (
	"log"
	"os"

	"github.com/williamhcarter/gonelang/gone/cmd"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: gone <command>")
	}

	command := os.Args[1]

	switch command {
	case "compile":
		cmd.Compile()
	case "generate":
		cmd.Generate()

	default:
		log.Fatalf("Unknown command: %s", command)
	}
}
