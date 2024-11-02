package game

import (
	"fmt"
	"strings"
	"testing"
)

func TestGenerateGamePassword(t *testing.T) {
	var needLength int = 5
	pwd := generatePassword(needLength)
	fmt.Println("generated password:", pwd)
	if len(pwd) != 5 {
		t.Fatalf("generatePassword(%v) = %v, want length = %v, got length = %v\n", needLength, pwd, needLength, len(pwd))
	}
	for i := 0; i < len(pwd); i++ {
		b := pwd[i]
		if strings.IndexByte(pwdCharactors, b) == -1 {
			t.Fatalf("generatePassword(%v) = %v, %v charactor is not existed in source", needLength, pwd, string(b))
		}
	}
}
