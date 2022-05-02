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
	bytes_num, err := bytes_util.Parse("11.47TB") //case insensitive
	if err != nil {
		log.Println(err.Error())
	} else {
		log.Println(bytes_num)
	}

	//rand_util
	log.Println("randstr:" + rand_util.GenRandStr(80))

	//hash_util
	log.Println(hash_util.MD5Hash([]byte("123123")))
	log.Println(hash_util.MD5HashString("123123"))

	log.Println(hash_util.SHA256([]byte("123123")))
	log.Println(hash_util.SHA256String("123123"))

	log.Println(hash_util.CRC32HashString("123123"))
	log.Println(hash_util.CRC32HashString("123123"))
	log.Println(hash_util.CRC32HashArrayString([]string{"1", "2", "3", "1", "2", "3"}))
	log.Println(hash_util.CRC32HashArrayString([]string{"1", "2", "3", "1", "2", "3"}))

}
