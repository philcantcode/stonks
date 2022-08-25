package utils

import (
	"crypto/sha256"
	"fmt"
)

func HashStruct(o interface{}) string {
	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("%v", o)))

	return fmt.Sprintf("%x", h.Sum(nil))
}
