package internal

import (
	"fmt"
)

type ErrParser struct {
	message string
	limit int
	lastString string
}

func (e *ErrParser) Error() string {
	limitMessage := fmt.Sprintf("%s, limit: %d, last string: %s", e.message, e.limit, e.lastString)
	return limitMessage
}