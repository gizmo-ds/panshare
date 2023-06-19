package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func ToMD5String(v []byte) string {
	h := md5.New()
	h.Write(v)
	return hex.EncodeToString(h.Sum(nil))
}
