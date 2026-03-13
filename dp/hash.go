package dp

import (
	"crypto/md5"
	"crypto/sha3"
	"encoding/hex"
)

// GetMD5Hash returns the MD5 hash of the given string.
// MARK: do not use for anything security-related as this algorithm is considered insecure. Only used for request caching mechanism.
func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

// GetSHA3Hash returns the SHA3 hash of the given string.
func GetSHA3Hash(text string) string {
	hash := sha3.New224().Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}