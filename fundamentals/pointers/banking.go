package banking

import (
	"errors"
	"fmt"
)

type Real int

type Wallet struct {
	balance Real
}

type Stringer interface {
	String() string
}

var ErrInsuficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Deposit(amount Real) {
	//fmt.Printf("address of balance in Deposit is %p \n", &w.balance)
	w.balance += amount
}
func (w *Wallet) Withdraw(amount Real) error {
	//fmt.Printf("address of balance in Deposit is %p \n", &w.balance)
	if amount > w.balance {
		return ErrInsuficientFunds
	}
	w.balance -= amount
	return nil
}
func (w *Wallet) Balance() Real {
	return w.balance
}

func (r Real) String() string {
	return fmt.Sprintf("R$%d", r)
}
