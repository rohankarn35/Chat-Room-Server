package authentication

import (
	"fmt"
	"math/rand"
)

func GenerateUniqueID(newrng *rand.Rand) string {
	return fmt.Sprintf("%08d", newrng.Intn(100000000))
}
