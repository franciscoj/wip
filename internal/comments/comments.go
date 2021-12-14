package comments

import (
	"context"
	"fmt"
	"sync"

	"github.com/franciscoj/wip/internal/client"
	"github.com/franciscoj/wip/internal/errors"
	"github.com/franciscoj/wip/internal/tasks"
)

const URL = "https://api.todoist.com/rest/v1/comments"

type Comment struct {
	ID      int    `json:"id"`
	TaskID  int    `json:"task_id"`
	Content string `json:"content"`
}

type Repo struct {
	comments []Comment
}

func (r Repo) GetByTaskID(taskID int) []Comment {
	var comments []Comment
	for _, comment := range r.comments {
		if comment.TaskID == taskID {
			comments = append(comments, comment)
		}
	}

	return comments
}

func LoadRepo(ctx context.Context, ts []tasks.Task) (Repo, error) {
	var wg sync.WaitGroup
	var mut sync.Mutex
	var comments []Comment
	var errs []error
	errCh := make(chan error, len(ts))

	for _, t := range ts {
		wg.Add(1)
		go func(task tasks.Task) {
			var c []Comment
			url := fmt.Sprintf("%s?task_id=%d", URL, task.ID)
			if err := client.Get(ctx, url, &c); err != nil {
				errCh <- fmt.Errorf("getting comments: %v, %w", task, err)
			}

			mut.Lock()
			comments = append(comments, c...)
			mut.Unlock()

			wg.Done()
		}(t)
	}

	wg.Wait()

	close(errCh)
	for err := range errCh {
		errs = append(errs, err)
	}
	if len(errs) != 0 {
		return Repo{}, errors.Multi{Errors: errs}
	}

	return Repo{comments}, nil
}
