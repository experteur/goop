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
    projectDir := filepath.Join(cwd, "example")
    howdyDir := filepath.Join(projectDir, "howdy")
    projectPath := filepath.Join(howdyDir, "project.md")

    project, err := markdown.LoadProject(projectPath)
    if err != nil {
        log.Panic(err)
    }
    fmt.Printf("%+v\n", project)
}
