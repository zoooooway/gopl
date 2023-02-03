package expand

import (
	"fmt"
	"strings"
	"testing"
)

func TestExpand(t *testing.T) {
	s := "jjjjfoojjjjfoofjfooo"
	fmt.Println(expand(s, func(s string) string { return strings.ToUpper(s) }))
}
