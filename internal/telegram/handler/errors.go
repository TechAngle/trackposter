package handler

import "errors"

// Connector error.
var (
	ErrConnectorInternal  = errors.New("connector err")
	ErrRepositoryInternal = errors.New("repository err")
)

// Handler internal errors.
var (
	ErrEmptyQueue = errors.New("empty queue")
	ErrNilClient  = errors.New("nil client")
)
