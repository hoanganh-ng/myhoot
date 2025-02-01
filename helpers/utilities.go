package helpers

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

var RandSource = rand.NewSource(time.Now().Unix())

func PrettyPrint(o ...interface{}) {
	if len(o) == 1 {
		b, _ := json.MarshalIndent(o[0], "", "  ")
		fmt.Println(string(b))
		return
	}
	b, _ := json.MarshalIndent(o, "", "  ")
	fmt.Println(string(b))
}
