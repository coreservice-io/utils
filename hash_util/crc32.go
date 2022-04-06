package hash_util

import (
	"fmt"
	"hash/crc32"
)

var crc32q = crc32.New(crc32.MakeTable(0xD5828281))

func CRC32Hash(input []byte) string {
	crc32q.Reset()
	crc32q.Write(input)
	return fmt.Sprint(crc32q.Sum32())
}

//combine all string elements and hash
func CRC32Hash_StringArray(input []string) string {
	var merge string
	for i := 0; i < len(input); i++ {
		merge = merge + input[i]
	}
	return CRC32Hash([]byte(merge))
}
