package crypto

import "testing"

func TestMD5(t *testing.T) {
	if MD5("biezhi") != "b3b71cd2fbee70ae501d024fe12a8fba" {
		t.Error("MD5 错误")
	}
}
