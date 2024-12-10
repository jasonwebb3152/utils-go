package crypto

import (
	"crypto/md5"
	"encoding/hex"
)

func GetMd5(input string) string {
	/** Get Md5 Hash for a string */
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(input))
	return hex.EncodeToString(hash.Sum(nil))
}
