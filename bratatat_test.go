package bratatat_test

import (
	"fmt"
	"testing"

	"github.com/nil2nekoni/bratatat"
)

func TestDo(t *testing.T) {
	n := 0
	bratatat.Do(func() {
		n++
	})

	expect := 3
	actual := n
	if actual != expect {
		t.Errorf("bratatat.Do should does something 3 times")
	}
}

func ExampleDo() {
	bratatat.Do(func() {
		fmt.Println("BRATATAT!")
	})
	// Output:
	// BRATATAT!
	// BRATATAT!
	// BRATATAT!
}
