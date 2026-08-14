package utils

import "github.com/pocketbase/pocketbase/tools/security"

// IDRandomAlphabet defines the alphabet used for generating random IDs, conforming to PB default validation.
const IDRandomAlphabet = "abcdefghijklmnopqrstuvwxyz0123456789"

func CreatePocketBaseIDString() string {
	return security.PseudorandomStringWithAlphabet(15, IDRandomAlphabet)
}
