package Hash

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
)

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func GetSha1Hash(text string) string {
	hasher := sha1.New()
	hasher.Write([]byte(text))
	encrypted := sha1.Sum(nil)
	return fmt.Sprintf("%x", encrypted)
}