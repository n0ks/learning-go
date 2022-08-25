package pointers

import (
	"errors"
	"fmt"
)

type ShitCoin int

type Wallet struct {
	balance ShitCoin
}

type Stringer interface {
	String() string
}

// Without the '*' `w` is just a copy of what is being passed in to the method
// "*" points to the original value of Wallet
func (w *Wallet) Deposit(value ShitCoin) {
	w.balance += value

}

func (w *Wallet) GetBalance() ShitCoin {
	return w.balance
}

func (s ShitCoin) String() string {
	return fmt.Sprintf("%d STC", s)

}
var ErrNoFunds = "cannot withdraw, insufficient funds"

func (w *Wallet) Withdraw(value ShitCoin) error {
	if value > w.balance {
    e:= errors.New(ErrNoFunds)
    return e
	}
	w.balance -= value
	return nil
}
