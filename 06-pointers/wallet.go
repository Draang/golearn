package pointers

import (
	"errors"
	"fmt"
)

type Bitcoin int
type Stringer interface {
	String() string
}
type Wallet struct {
	balance Bitcoin
}
var ErrInsufficentFunds=errors.New("no puedes retirar mas de tu balance")
// El Asterisco representa que estas mandando la direccion de memoria un puntero de esa manera se cambia el mismo objeto
func (w *Wallet) Deposit(amount Bitcoin) {
	//w.balance+=amount <== metodo no efectivo por que la w es una copia de la wallet no modificamos la misma direccion de memoria
	/*
		(w Wallet)
		address of balance in Deposit is 0xc00010c1c8
		address of balance in test is 0xc00010c1c0
		fmt.Printf("address of balance in Deposit is %p \n", &w.balance)
	*/
	w.balance += amount
}
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficentFunds
	}
	w.balance -= amount
	return nil
}
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
