package bank

import (
	"errors"
	"strconv"
)

var errNotEnoughMoney = errors.New("not enough money")

type Account struct {
	owner   string
	balance int
}

func NewAccount(owner string) *Account {
	return &Account{owner, 0}
}

func (a *Account) Deposit(amount int) {
	a.balance += amount
}

func (a *Account) Withdraw(amount int) error {
	if temp := a.balance; temp-amount < 0 {
		return errNotEnoughMoney
	}
	a.balance -= amount
	return nil
}

func (a *Account) ChangeOwner(owner string) {
	a.owner = owner
}

func (a *Account) Owner() string {
	return a.owner
}

func (a *Account) Balance() int {
	return a.balance
}

func (a *Account) String() (summary string) { // Java의 toString(), override한 것임
	summary = "owner: " + a.owner + "balance: " + strconv.Itoa(a.balance)
	return
}
