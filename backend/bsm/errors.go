package bsm

import "fmt"

type URLWhitelistError struct {
	url     string
	message string
}

func (e *URLWhitelistError) Error() string {
	return fmt.Sprintf("%s - %s", e.url, e.message)
}
