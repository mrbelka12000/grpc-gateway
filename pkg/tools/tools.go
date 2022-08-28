package tools

import (
	"fmt"
	"math/rand"
	"time"
)

func GetRandomString() string {
	rand.Seed(time.Now().Unix())
	return fmt.Sprintf("%016d", rand.Int63n(1e16))
}
