package utils

import (
	"math/rand"
	"strings"
	"time"
)

const aplhabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomNumber generates a random number between min and max
func RandomNumber(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(aplhabet)

	for i := 0; i < n; i++ {
		c := aplhabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomOwner generates a random owner name
func RandomOwner() string {
	return RandomString(8)
}

// RandomCurrency generates a random currency from the set available
func RandomCurrency() string {
	currencies := []string{"USD", "EUR", "CAD"}
	return currencies[rand.Intn(len(currencies))]
}
