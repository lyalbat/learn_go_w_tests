package banking

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)
		//fmt.Printf("address of balance in test is %p \n", &wallet.balance)
		want := Real(10)
		assertCorrectBalance(t, wallet, want)
	})
	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Real(20)}
		err := wallet.Withdraw(10)
		want := Real(10)
		assertNoError(t, err)
		assertCorrectBalance(t, wallet, want)
	})
	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Real(20)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(100)
		assertErr(t, err, ErrInsuficientFunds)
		assertCorrectBalance(t, wallet, startingBalance)
	})
}

func assertErr(t testing.TB, got, want error) {
	t.Helper()
	if got == nil {
		//if we didn't pass the fatal, it would return the error and carry on
		//and it would try to get Error method from nil, which would case PANIC
		t.Fatal("wanted an error but didn't get one")
	}
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}
func assertCorrectBalance(t testing.TB, wallet Wallet, want Real) {
	t.Helper()
	got := wallet.Balance()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
