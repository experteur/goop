package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/experteur/goop/internal/app"
)

func main() {
	projectsDir := getProjectsDir()
	application, err := app.New(projectsDir)
    if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
    }
	if err := application.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func getProjectsDir() string {
	cwd, err := os.Getwd()
	if err != nil {
		log.Panic(err)
	}
	projectsDir := filepath.Join(cwd, "example")

	return projectsDir
}
