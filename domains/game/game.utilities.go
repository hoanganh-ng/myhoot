package game

import (
	"math/rand"

	"github.com/hoanganh-ng/myhoot/helpers"
)

const pwdCharactors string = "1234567890"

func generatePassword(length int) string {
	pwd := ""
	for i := 0; i < length; i++ {
		charactorIndex := rand.New(helpers.RandSource).Intn(len(pwdCharactors))
		pwd += string(pwdCharactors[charactorIndex])
	}
	return pwd
}
