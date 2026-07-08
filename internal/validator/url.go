package validator

import (
	"errors"
	"net/url"
)

// ValidateURL checks if provided URL can be parsed.
func ValidateURL(link string) error {
	url, err := url.ParseRequestURI(link)
	if err != nil {
		return errors.Join(ErrInvalidURL, err)
	}

	if url.Host == "" || url.Scheme == "" {
		return ErrInvalidURL
	}

	return nil
}
