package replace

import (
	"fmt"
	"testing"
)

func TestReplace(t *testing.T) {
	replaced := Replace("Bibi", "i", "o")
	expected := "Bobo"

	if replaced != expected {
		t.Errorf("expected %q but got %q", expected, replaced)
	}
}

func ExampleReplace() {
	replaced := Replace("Bibi", "i", "o")
	fmt.Println(replaced)
	// Output: Bobo
}

func BenchmarkReplace(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Replace("Bibi", "i", "o")
	}
}
