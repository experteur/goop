package markdown

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/experteur/goop/internal/domain"
)

func LoadProject(path string) (*domain.Project, error) {
    data, err := os.ReadFile(path)
    if err != nil {
        return nil, err
    }
    fmt.Printf("%+v\n", data)

    content := string(data)

    splitString := strings.Split(content, "\n\n")

    title := splitString[0]
    description := splitString[1]

    project := &domain.Project{
    	Name:        title,
    	Description: description,
    	Status:      "",
    	Owner:       "",
    	Path:        path,
    	LastUpdated: time.Time{},
    	Tags:        []string{},
    }

    return project, nil
}
