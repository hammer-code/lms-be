package hash

import (
	"crypto/sha1"
	"encoding/hex"
)

func GenerateHash(material string) string {
	hash := sha1.New()
	hash.Write([]byte(material))
	return hex.EncodeToString(hash.Sum(nil))
}
