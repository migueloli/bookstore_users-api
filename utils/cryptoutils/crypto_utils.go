package cryptoutils

import (
	"crypto/md5"
	"encoding/hex"
)

// GetMd5 is a function to cryptograph a string.
func GetMd5(input string) string {
	hash := md5.New()
	hash.Reset()
	hash.Write([]byte(input))
	return hex.EncodeToString(hash.Sum(nil))
}
