package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No command provided")
		os.Exit(2)
	}

	cmd := os.Args[1]
	switch cmd {
	case "greet":
		msg := "REMINDERS CLI - CLI BASICS"
		if len(os.Args) == 3 {
			f := strings.Split(os.Args[2], "=")
			if len(f) == 2 && f[0] == "--msg" {
				msg = f[1]
			}
		}
		fmt.Printf("Hello and Welcome: %s\n", msg)
	case "help":
		fmt.Println("Some help message")
	default:
		fmt.Printf("Unknown Command: %s\n", cmd)
	}
}
