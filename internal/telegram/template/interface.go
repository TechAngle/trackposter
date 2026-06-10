// Package template provides basic types used in templates.
package template

// MessageTemplate represents type for Telegram messages.
type MessageTemplate string

// String returns default representation for type.
func (t MessageTemplate) String() string {
	return string(t)
}
