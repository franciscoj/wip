package app

import (
	"context"
	"fmt"

	"github.com/franciscoj/wip/internal/comments"
	"github.com/franciscoj/wip/internal/labels"
	"github.com/franciscoj/wip/internal/projects"
	"github.com/franciscoj/wip/internal/sections"
	"github.com/franciscoj/wip/internal/tasks"
)

type App struct {
	projects projects.Repo
	labels   labels.Repo
	sections sections.Repo
	tasks    tasks.Repo
	comments comments.Repo
}

func (a App) Print() {
	bySection := make(map[int][]tasks.Task)
	for _, t := range a.tasks.All() {
		bySection[t.SectionID] = append(bySection[t.SectionID], t)
	}

	for sectionID, sectionTasks := range bySection {
		section, _ := a.sections.Get(sectionID)
		fmt.Printf("### %s\n\n", section.Name)

		byLabel := make(map[int][]tasks.Task)
		for _, t := range sectionTasks {
			byLabel[t.GetLabelID()] = append(byLabel[t.GetLabelID()], t)
		}

		for labelID, tasks := range byLabel {
			label, _ := a.labels.Get(labelID)
			fmt.Printf("\n#### %s\n\n", label.Name)
			for _, t := range tasks {
				fmt.Printf("- %s\n", t.Content)
				for _, c := range a.comments.GetByTaskID(t.ID) {
					fmt.Printf("  - %s\n", c.Content)
				}
			}
		}
	}
}

func Load(ctx context.Context) (App, error) {
	pr, err := projects.LoadRepo(ctx)
	if err != nil {
		return App{}, err
	}
	frs := pr.GetWIP()
	if len(frs) != 1 {
		return App{}, fmt.Errorf("there can only be 1 WIP project")
	}
	p := frs[0]

	sr, err := sections.LoadRepo(ctx, p)
	if err != nil {
		return App{}, err
	}

	tr, err := tasks.LoadRepo(ctx, p)
	if err != nil {
		return App{}, err
	}

	cr, err := comments.LoadRepo(ctx, tr.All())
	if err != nil {
		return App{}, err
	}

	lr, err := labels.LoadRepo(ctx)
	if err != nil {
		return App{}, err
	}

	return App{pr, lr, sr, tr, cr}, err
}
