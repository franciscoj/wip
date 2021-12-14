package tasks

import (
	"context"
	"fmt"

	"github.com/franciscoj/wip/internal/client"
	"github.com/franciscoj/wip/internal/projects"
)

const URL = "https://api.todoist.com/rest/v1/tasks"

type Task struct {
	ID        int    `json:"id"`
	SectionID int    `json:"section_id"`
	LabelIDs  []int  `json:"label_ids"`
	Content   string `json:"content"`
}

func (t Task) GetLabelID() int {
	return t.LabelIDs[0]
}

type Repo struct {
	tasks []Task
}

func (r Repo) All() []Task {
	return r.tasks
}

func (r Repo) Get(id int) (Task, bool) {
	for _, task := range r.tasks {
		if task.ID == id {
			return task, true
		}
	}

	return Task{}, false
}

func LoadRepo(ctx context.Context, p projects.Project) (Repo, error) {
	var tasks []Task

	url := fmt.Sprintf("%s?project_id=%d", URL, p.ID)
	if err := client.Get(ctx, url, &tasks); err != nil {
		return Repo{}, fmt.Errorf("getting items: %w", err)
	}

	return Repo{tasks}, nil
}
