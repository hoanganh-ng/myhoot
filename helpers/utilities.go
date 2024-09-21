package helpers

import (
	"math/rand"
	"time"
)

var RandSource = rand.NewSource(time.Now().Unix())
