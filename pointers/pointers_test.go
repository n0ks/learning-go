package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {

		wallet := Wallet{}
		wallet.Deposit(ShitCoin(10))

		want := ShitCoin(10)
		assert(t, wallet, want)

	})

	t.Run("Withdraw", func(t *testing.T) {

		wallet := Wallet{}
		wallet.Deposit(ShitCoin(10))

		err := wallet.Withdraw(ShitCoin(5))

		want := ShitCoin(5)
		assertNoError(t, err)
		assert(t, wallet, want)

	})

	t.Run("Withdraw with insufficient funds", func(t *testing.T) {

		wallet := Wallet{ShitCoin(20)}

		err := wallet.Withdraw(ShitCoin(99))

		assertError(t, err, ErrNoFunds)
		assert(t, wallet, ShitCoin(20))

	})
}

func assert(t testing.TB, wallet Wallet, want ShitCoin) {
	t.Helper()
	got := wallet.GetBalance()

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func assertError(t testing.TB, got error, want string) {
	t.Helper()

	if got == nil {
		t.Fatal("didn't get and error but wanted one")

	}

	if got.Error() != want {
		t.Errorf("got %q, want %q", got, want)
	}

}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("go and error but didn't want one")
	}

}
