package domain

import (
	"net/url"
)

// ValidateURL checks if provided URL can be parsed.
func ValidateURL(link string) error {
	_, err := url.Parse(link)
	if err != nil {
		return ErrInvalidURL
	}

	return nil
}
