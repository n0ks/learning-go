package hello

import "testing"

func TestHello(t *testing.T) {

	t.Run("say hi to people", func(t *testing.T) {
		got := Hi("John", "")
		want := "Hi, John"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hi, fuck you' when empty string is provided", func(t *testing.T) {
		got := Hi("", "")
		want := "Hi, fuck you"
		assertCorrectMessage(t, got, want)
	})

	t.Run("talk in sPaNiSh", func(t *testing.T) {

		got := Hi("Cucaracha", "Spanish")
		want := "Hola, Cucaracha"

		assertCorrectMessage(t, got, want)
	})

}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("msg is %q, expected was %q", got, want)

	}

}
