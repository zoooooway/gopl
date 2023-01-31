package count

import (
	"fmt"
	"testing"
)

func TestCount(t *testing.T) {
	wc, ic := countWordsAndImages("https://vertx.io/")
	fmt.Printf("words: %d\n", wc)
	fmt.Printf("images: %d\n", ic)
}
