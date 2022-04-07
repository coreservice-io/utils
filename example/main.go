package main

import (
	"log"

	"github.com/coreservice-io/UUtils/bytes_util"
	"github.com/coreservice-io/UUtils/hash_util"
	"github.com/coreservice-io/UUtils/rand_util"
	"github.com/coreservice-io/UUtils/time_util"
)

func main() {
	//time_util
	log.Println(time_util.GetUTCDate())
	log.Println(time_util.GetUTCDateTime())

	//bytes_util
	log.Println(bytes_util.Format(12312341344))
	bytesnum, err := bytes_util.Parse("11.47TB") //case insensitive
	if err != nil {
		log.Println(err.Error())
	} else {
		log.Println(bytesnum)
	}

	//rand_util
	log.Println("randstr:" + rand_util.GenRandStr(80))

	//hash_util
	log.Println(hash_util.MD5Hash([]byte("1234214")))
	log.Println(hash_util.MD5Hash([]byte("1234214")))
	log.Println(hash_util.MD5Hash_StringArray([]string{"123", "1234"}))
	log.Println(hash_util.MD5Hash_StringArray([]string{}))

	log.Println(hash_util.CRC32Hash([]byte("1234214")))
	log.Println(hash_util.CRC32Hash([]byte("1234214")))
	log.Println(hash_util.CRC32Hash([]byte("ffffffffffffffffffffffffffffffffffffffff")))
	log.Println(hash_util.CRC32Hash([]byte("ffffffffffffffffffffffffffffffffffffffff")))

}
