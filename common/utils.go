package comm

import (
	"math/rand"
	"time"
)

var R *rand.Rand

// 1:力量；2：敏捷
// 101：虚弱；102：易伤
const (
	Buff_Power    = 1
	Buff_Dex      = 2
	Debuff_Weak   = 101
	Debuff_Vulner = 102
)

func IsInSlice[T comparable](slice []T, item T) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func init() {
	R = rand.New(rand.NewSource(time.Now().UnixNano()))
}
