package manufacture

import (
	"github.com/google/uuid"
)

// NewUUID return a random uuid.
func NewUUID() (string, error) {
	u, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return u.String(), nil
}
