package sections

import (
	"context"
	"fmt"

	"github.com/franciscoj/wip/internal/client"
	"github.com/franciscoj/wip/internal/projects"
)

const URL = "https://api.todoist.com/rest/v1/sections"

type Section struct {
	ID   int    `json:"id"`
	Name string `json:"Name"`
}

type Repo struct {
	sections []Section
}

func (r Repo) Get(id int) (Section, bool) {
	for _, section := range r.sections {
		if section.ID == id {
			return section, true
		}
	}

	return Section{}, false
}

func LoadRepo(ctx context.Context, p projects.Project) (Repo, error) {
	var sections []Section

	url := fmt.Sprintf("%s?project_id=%d", URL, p.ID)
	if err := client.Get(ctx, url, &sections); err != nil {
		return Repo{}, fmt.Errorf("getting sections: %w", err)
	}

	return Repo{sections}, nil
}
