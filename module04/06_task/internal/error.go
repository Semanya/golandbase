package internal

import (
	"fmt"
)

type ErrParser struct {
	Message string
	Limit int
	LastString string
}

func (e *ErrParser) Error() string {
	limitMessage := fmt.Sprintf("%s, limit: %d, last string: %s", e.Message, e.Limit, e.LastString)
	return limitMessage
}