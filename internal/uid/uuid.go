// Package uid defines logic for generating UUID.
package uid

import (
	"errors"

	"github.com/google/uuid"
	"trackposter/internal/domain"
)

// New generates new UUID.
// If something went wrong - returns empty string and error.
func New() (string, error) {
	uid, err := uuid.NewRandom()
	if err != nil {
		return "", errors.Join(domain.ErrUUID, err)
	}

	return uid.String(), nil
}
