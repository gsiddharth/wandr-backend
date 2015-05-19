package utils

import (
	"crypto/md5"
	"io"
	"math/rand"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var seed int64 = 0

func MD5Hash(str string) string {
	p_md5 := md5.New()
	io.WriteString(p_md5, str)
	return string(p_md5.Sum(nil))
}

func RandomString(n int) string {

	b := make([]rune, n)
	seed = (seed + 1) % 1000000

	rand.Seed(int64(time.Now().Nanosecond()) * seed)

	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}
