package codelocationmatch

import (
	"fmt"
	"regexp"
	"testing"
)

func TestMatch01(t *testing.T) {
	s1 := `test.go#L12`
	s2 := `test.go#L12`
	ret := Match(s1, s2, true, nil)
	fmt.Println(ret)
}

func TestMatch02(t *testing.T) {
	s1 := `test.go:12`
	s2 := `/tmp/testeval-xxx/test.go#L12`
	ret := Match(s1, s2, true, regexp.MustCompile(`^.+?/testeval-[^/]+/(.+)$`))
	fmt.Println(ret)
}

func TestMatch03(t *testing.T) {
	s1 := `test.go#L12-20`
	s2 := `/tmp/testeval-xxx/test.go#L15`
	ret := Match(s1, s2, true, regexp.MustCompile(`^.+?/testeval-[^/]+/(.+)$`))
	fmt.Println(ret)
}
