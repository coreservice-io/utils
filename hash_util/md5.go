package hash_util

import (
	"crypto/md5"
	"encoding/hex"
)

var md5hash = md5.New()

func MD5Hash(input []byte) string {
	md5hash.Reset()
	md5hash.Write(input)
	return hex.EncodeToString(md5hash.Sum(nil))
}

//combine all string elements and hash
func MD5Hash_StringArray(input []string) string {
	var merge string
	for i := 0; i < len(input); i++ {
		merge = merge + input[i]
	}
	return MD5Hash([]byte(merge))
}
