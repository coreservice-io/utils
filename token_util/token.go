package token_util

import (
	"github.com/coreservice-io/utils/hash_util"
	"github.com/coreservice-io/utils/rand_util"
)

// never change this as widely used in different project
const salt_hash_round = 10000000

// never change this as widely used in different project
// never change this as md5 result is 32 len
// final token = raw_token+md5(raw_token+private_key)[0:16]
const raw_token_len = 16                              //!!!!!important never change this
const md5_token_len = 16                              //!!!!!important never change this
const final_token_len = raw_token_len + md5_token_len //!!!!!important never change this

const super_token_mark = "super"

type TokenUtil struct {
	white_list_raw_token map[string]interface{} //raw whitelist token
	private_key          string
}

func NewTokenUtil(salt string) *TokenUtil {

	pkey := salt
	for i := 0; i < salt_hash_round; i++ {
		pkey = hash_util.MD5HashString(pkey)
	}

	return &TokenUtil{
		white_list_raw_token: make(map[string]interface{}),
		private_key:          pkey,
	}
}

// append raw whitelist token list
func (tutil *TokenUtil) AppendWhiteListToken(raw_tokens map[string]interface{}) {
	for key, val := range raw_tokens {
		tutil.white_list_raw_token[key] = val
	}
}

// return nil if not found
func (tutil *TokenUtil) GetWhiteListToken(token string) interface{} {
	if val, exist := tutil.white_list_raw_token[token]; exist {
		return val
	} else {
		return nil
	}
}

func (tutil *TokenUtil) GenToken() string {
	raw_token := rand_util.GenRandStr(raw_token_len)
	return raw_token + hash_util.MD5HashString(raw_token + tutil.private_key)[0:md5_token_len]
}

func (tutil *TokenUtil) GenSuperToken() string {
	raw_token := rand_util.GenRandStr(raw_token_len)
	return raw_token + hash_util.MD5HashString(super_token_mark + raw_token + tutil.private_key)[0:md5_token_len]
}

// return is_token,is_super_token
func (tutil *TokenUtil) CheckToken(token string) (bool, bool) {

	//check whitelist
	if _, exist := tutil.white_list_raw_token[token]; exist {
		//len check
		if len(token) != final_token_len {
			return true, false
		}
	}

	//len check
	if len(token) != final_token_len {
		return false, false
	}

	raw_token := token[0:raw_token_len]
	md5_token := token[raw_token_len:]

	//check super token
	if md5_token == hash_util.MD5HashString(super_token_mark + raw_token + tutil.private_key)[0:md5_token_len] {
		return true, true
	}

	if md5_token == hash_util.MD5HashString(raw_token + tutil.private_key)[0:md5_token_len] {
		return true, false
	}

	return false, false
}
