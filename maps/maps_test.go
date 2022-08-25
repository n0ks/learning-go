package maps

import "testing"

func TestSearch(t *testing.T) {

	dict := Dictionary{"test": "test"}

	t.Run("get word", func(t *testing.T) {

		got, _ := dict.Search("test")

		want := "test"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {

		_, got := dict.Search("unknown")

		want := NotFoundWordErr

		assertError(t, got, want)
	})

}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {

		dict := Dictionary{}
		word := "test"
		definition := "test"

		dict.Add(word, definition)

		assertDefinition(t, dict, word, definition)
	})

	t.Run("existing word", func(t *testing.T) {

		word := "test"
		definition := "test"

		dict := Dictionary{word: definition}
		err := dict.Add(word, "new test")
		assertError(t, err, WordExistsErr)
		assertDefinition(t, dict, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "test"

		dict := Dictionary{word: definition}
		newDef := "test new def"

		err := dict.Update(word, newDef)

		assertError(t, err, nil)

		assertDefinition(t, dict, word, newDef)
	})

	t.Run("new word", func(t *testing.T) {

		word := "test"
		definition := "test"

		dict := Dictionary{}

		err := dict.Update(word, definition)

		assertError(t, err, DoestNotExistErr)

	})

}

func TestDelete(t *testing.T) {
	word := "test"
	dict := Dictionary{word: "test definition"}

	dict.Delete(word)

	_, err := dict.Search(word)
	if err != NotFoundWordErr {
		t.Errorf("Expected %q to be deleted", word)
	}
}

func assertDefinition(t testing.TB, dict Dictionary, word, definition string) {
	t.Helper()

	got, err := dict.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if definition != got {
		t.Errorf("got %q want %q", got, definition)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}

}
