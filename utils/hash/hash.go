package hash

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
)

// GetMD5Hash ... Get MD5 result from text
func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

// GetSha1Hash ... Get SHA1 result from text
func GetSha1Hash(text string) string {
	hasher := sha1.New()
	hasher.Write([]byte(text))
	encrypted := hasher.Sum(nil)
	return fmt.Sprintf("%x", encrypted)
}
