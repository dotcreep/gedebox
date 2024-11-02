package handler

import "fmt"

type HandlerOperation struct {
	message string
}

func (e *HandlerOperation) Error() string {
	return fmt.Sprintf("Error: Operation '%s' is not supported", e.message)
}

func OpError(message string) error {
	return &HandlerOperation{message: message}
}

type HandlerDistro struct {
	message string
}

func (e *HandlerDistro) Error() string {
	return fmt.Sprintf("Error: Distro '%s' is not supported", e.message)
}

func DistError(message string) error {
	return &HandlerDistro{message: message}
}
