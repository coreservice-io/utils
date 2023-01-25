package main

import (
	"fmt"
	"log"
	"time"

	"github.com/coreservice-io/utils/bytes_util"
	"github.com/coreservice-io/utils/hash_util"
	"github.com/coreservice-io/utils/rand_util"
	"github.com/coreservice-io/utils/time_util"
	"github.com/coreservice-io/utils/token_util"
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

	//token util
	time_now := time.Now().Unix()
	tutil := token_util.NewTokenUtil("test salt")
	fmt.Println("new token util pass secs:", time.Now().Unix()-time_now)
	//
	tutil.AppendWhiteListToken(map[string]interface{}{"token_w": struct{}{}})
	//

	log.Println("token_util instance:", tutil)
	token1 := tutil.GenToken()
	token2 := tutil.GenToken()
	token3 := tutil.GenSuperToken()

	token1_, token1_s := tutil.CheckToken(token1)
	log.Println("check token1 is token:", token1_, " is super token:", token1_s)

	token2_, token2_s := tutil.CheckToken(token2)
	log.Println("check token2 is token:", token2_, " is super token:", token2_s)

	token3_, token3_s := tutil.CheckToken(token3)
	log.Println("check token3 is token:", token3_, " is super token:", token3_s)

	tokenw_, tokenw_s := tutil.CheckToken("token_w")
	log.Println("check tokenw is token:", tokenw_, " is super token:", tokenw_s)

	tokene_, tokene_s := tutil.CheckToken("token_err")
	log.Println("check token_err is token:", tokene_, " is super token:", tokene_s)

	log.Println("get white list token:", tutil.GetWhiteListToken("token_w"))
	log.Println("get white list token:", tutil.GetWhiteListToken("token_err"))

}
