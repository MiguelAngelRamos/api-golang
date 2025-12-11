package notification

import "errors"

var (
	ErrEmptyDestination = errors.New("el destino no puede estar vacío")
	ErrEmptyMessage     = errors.New("el mensaje no puede estar vacío")
)

// var ErrEmptyDestination = errors.New("el destino no puede estar vacío")
// var ErrEmptyMessage     = errors.New("el mensaje no puede estar vacío")
