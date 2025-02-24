package customerrors

import "fmt"

type PriorityOutOfBoundsError struct {
	Priority int
}

func (e PriorityOutOfBoundsError) Error() string {
	return fmt.Sprintf("priority out of bounds (1-5): %d", e.Priority)
}
