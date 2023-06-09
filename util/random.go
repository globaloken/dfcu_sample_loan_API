package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomAmount() int64 {
	rand.Seed(time.Now().UnixNano())

	amount := rand.Intn(4951) + 50 // generates random int between 50 and 5000 (inclusive)
	amount *= 1000
	return int64(amount)
}

// RandomInt generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomOwner generates random owner name
func RandomOwner() string {
	return RandomString(6)
}

// RandomMoney generates random amount of money
func RandomMoney() int64 {
	return RandomInt(0, 50000)
}

// Random currency generates a random currency
func RandomCurrency() string {
	currencies := []string{USD, UGX, EUR}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}

// Random email generates a random email
func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(6))
}
