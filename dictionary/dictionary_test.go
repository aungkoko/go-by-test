package dictionary

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "This is a test"}

	t.Run("known", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "This is a test"

		assertString(t, got, want)
	})

	t.Run("unknown", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		want := ErrNotFound
		if err == nil {
			t.Fatal("expected to get an error")
		}

		assertError(t, err, want)
	})
}

func assertString(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("we expected %s we got %s", want, got)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want error %q", got, want)
	}
}

func TestAdd(t *testing.T) {
	d := Dictionary{}
	word := "test"
	definition := "this is a test"
	d.Add(word, definition)
	assertDefinition(t, d, word, definition)
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()
	want := "this is a test"
	got, err := dictionary.Search("test")

	if err != nil {
		t.Error("should fine added word:", err)
	}

	assertString(t, got, want)
}
