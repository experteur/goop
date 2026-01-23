package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/experteur/goop/internal/markdown"
)
func main() {
	cwd, err := os.Getwd()
    if err != nil {
        log.Panic(err)
    }
    projectsDir := filepath.Join(cwd, "example")

    projects, err := markdown.LoadProjects(projectsDir)
    if err != nil {
        log.Panic(err)
    }
    for _, project := range projects{
        fmt.Printf("%+v\n", project)
    }
}
