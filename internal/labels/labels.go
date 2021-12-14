package labels

import (
	"context"
	"fmt"

	"github.com/franciscoj/wip/internal/client"
)

type Label struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Repo struct {
	labels []Label
}

func (r Repo) Get(id int) (Label, bool) {
	for _, label := range r.labels {
		if label.ID == id {
			return label, true
		}
	}

	return Label{}, false
}

const URL = "https://api.todoist.com/rest/v1/labels"

func LoadRepo(ctx context.Context) (Repo, error) {
	var labels []Label
	if err := client.Get(ctx, URL, &labels); err != nil {
		return Repo{}, fmt.Errorf("getting project: %w", err)
	}

	return Repo{labels}, nil
}
