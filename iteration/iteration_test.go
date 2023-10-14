package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	got := Repeat("a", 5)
	expected := "aaaaa"

	assertCorrectMessage(t, got, expected)

	t.Run("test diff char and time", func(t *testing.T) {
		got := Repeat("ab", 3)
		expected := "ababab"

		assertCorrectMessage(t, got, expected)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func ExampleRepeat() {
	got := Repeat("b", 6)
	fmt.Println(got)
	// Outputtt: bbbbbb
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
