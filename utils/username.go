package utils

import (
	"math/rand"
	"strings"
	"time"
)

func RandomUsername() string {
	r := rand.New(rand.NewSource(time.Now().Unix()))

	in := "abcdefghijklmnopqrstuvwxyz0123456789"

	inRune := []rune(in)
	r.Shuffle(len(inRune), func(i, j int) {
		inRune[i], inRune[j] = inRune[j], inRune[i]
	})
	return strings.TrimSpace(string(inRune)[0:10])
}
