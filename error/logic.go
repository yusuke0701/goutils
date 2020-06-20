package error

import "strings"

type ErrConflict struct {
	IDs []string
	Msg string
}

func (e *ErrConflict) Error() string {
	if len(e.IDs) == 0 {
		return e.Msg
	}
	return strings.Join(e.IDs, ",")
}

type ErrNotFound struct {
	IDs []string
	Msg string
}

func (e *ErrNotFound) Error() string {
	if len(e.IDs) == 0 {
		return e.Msg
	}
	return strings.Join(e.IDs, ",")
}
