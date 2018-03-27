package regex

import "testing"

func TestIsDomain(t *testing.T) {
	if !IsDomain("x.xx") || IsDomain("xx") || IsDomain(".xx") {
		t.Error("域名正则错误")
	}
}

func TestIsIPv4(t *testing.T) {
	if !IsIPv4("1.1.1.1") || IsIPv4("11") || IsIPv4(".1") {
		t.Error("IP正则错误")
	}
}
