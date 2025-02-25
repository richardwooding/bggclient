package customerrors

import "fmt"

type RatingOutOfBoundsError struct {
	Rating int
}

func (e RatingOutOfBoundsError) Error() string {
	return fmt.Sprintf("rating out of bounds (1-10): %d", e.Rating)
}
