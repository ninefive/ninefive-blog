package lib

import (
	"crypto/md5"
	"encoding/hex"
)

func Pwshash(str string) string {
	return Strtomd5(str)
}

func Strtomd5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	rs := hex.EncodeToString(h.Sum(nil))
	return rs
}
