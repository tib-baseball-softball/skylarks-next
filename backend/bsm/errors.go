package bsm

import "fmt"

type URLAllowlistError struct {
	url     string
	message string
}

func (e *URLAllowlistError) Error() string {
	return fmt.Sprintf("%s - %s", e.url, e.message)
}
