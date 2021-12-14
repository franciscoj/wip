package errors

import "fmt"

type Multi struct {
	Errors []error
}

func (e Multi) Error() string {
	return fmt.Sprintf("%d errors: %v", len(e.Errors), e.Errors)
}
