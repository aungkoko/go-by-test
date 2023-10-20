package wallet

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	assertBalance := func(t testing.TB, w Wallet, want Bitcoin) {
		t.Helper()
		got := w.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	assertError := func(t testing.TB, got error, want string) {
		t.Helper()
		if got == nil {
			t.Fatal("didn't get an error but wanted one")
		}

		if got.Error() != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		want := Bitcoin(10)

		assertBalance(t, wallet, want)
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{Bitcoin(20)}

		err := wallet.Withdraw(Bitcoin(10))

		want := Bitcoin(10)
		assertNoError(t, err)

		assertBalance(t, wallet, want)
	})

	t.Run("withdrawinsufficientamount", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		fmt.Println(err)

		assertError(t, err, ErrInsufficientFunds.Error())
		assertBalance(t, wallet, startingBalance)
	})
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}
