package projects

import (
	"context"
	"fmt"
	"regexp"

	"github.com/franciscoj/wip/internal/client"
)

type Project struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (p *Project) isEmpty() bool {
	return p.ID == 0
}

func (p *Project) String() string {
	return fmt.Sprintf("%d: %s", p.ID, p.Name)
}

type Repo struct {
	projects []Project
}

var frRegex = regexp.MustCompile(`WIP - \d{4}-\d{2}-\d{2}`)

func (r *Repo) GetWIP() []Project {
	var fr []Project

	for _, p := range r.projects {
		if frRegex.MatchString(p.Name) {
			fr = append(fr, p)
		}
	}

	return fr
}

const URL = "https://api.todoist.com/rest/v1/projects"

func LoadRepo(ctx context.Context) (Repo, error) {
	var projects []Project
	if err := client.Get(ctx, URL, &projects); err != nil {
		return Repo{}, fmt.Errorf("getting project: %w", err)
	}

	return Repo{projects}, nil
}
