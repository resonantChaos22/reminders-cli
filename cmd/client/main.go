package main

import (
	// "client"
	"flag"
	"fmt"
	"github.com/resonantChaos22/reminders-cli/client"
	"os"
)

var (
	backendURIFlag = flag.String("backend", "http://localhost:8080", "Backend API URL")
	helpFlag       = flag.Bool("help", false, "Display a helpful message")
)

func main() {
	flag.Parse()
	s := client.NewSwitch(*backendURIFlag)

	if *helpFlag || len(os.Args) == 1 {
		s.Help()
		return
	}

	err := s.Switch()
	if err != nil {
		fmt.Printf("cmd switch error: %s", err)
		os.Exit(2)
	}
}