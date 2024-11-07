package ciclos

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	reap := Repeat("a", 6)
	exp := "aaaaaa"
	if reap != exp {
		t.Errorf("expected %q but got %q", exp, reap)
	}
}
func ExampleRepeat() {
	reap := Repeat("b", 1)
	fmt.Println(reap)
	//Output: b
}

// Test benchmark mide el run time del codigo
// para correr benchmark: go test -bench="."
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 6)
	}
}
