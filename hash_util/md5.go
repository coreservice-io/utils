package hash_util

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5Hash(input []byte) []byte {
	var md5hash = md5.New()
	md5hash.Write(input)
	return md5hash.Sum(nil)
}

func MD5HashString(input string) string {
	return hex.EncodeToString(MD5Hash([]byte(input)))
}

//combine all string elements and hash
// func MD5HashArray(input [][]byte) []byte {
// 	var merge []byte
// 	for i := 0; i < len(input); i++ {
// 		merge = append(merge, input[i]...)
// 	}
// 	return MD5Hash([]byte(merge))
// }

// func MD5HashArrayString(input []string) string {
// 	var merge string
// 	for i := 0; i < len(input); i++ {
// 		merge = merge + "_" + input[i]
// 	}
// 	return MD5HashString(merge)
// }
