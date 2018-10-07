package main

import (
	"fmt"
)

// ------------------

type Wallet struct {
	Cash int
}

func (w *Wallet) Pay(amount int) error {
	if w.Cash < amount {
		return fmt.Errorf("Not enoght money")
	}
	w.Cash -= amount
	return nil
}

// ------------------

type Card struct {
	Balance    int
	ValidUntil string
	Cardholder string
	CVV        string
	Number     string
}

func (c *Card) Pay(amount int) error {
	if c.Balance < amount {
		return fmt.Errorf("Not enoght money")
	}
	c.Balance -= amount
	return nil
}

// ------------------

type ApplePay struct {
	Money   int
	AppleID string
}

func (a *ApplePay) Pay(amount int) error {
	if a.Money < amount {
		return fmt.Errorf("Not enoght money")
	}
	a.Money -= amount
	return nil
}

// ------------------

type Payer interface {
	Pay(int) error
}

func Buy(p Payer) {
	err := p.Pay(10)
	if err != nil {
		fmt.Printf("Error with pay %T: %v\n\n", p, err)
		return
	}
	fmt.Printf("Thanks for payment", p)
}

// ------------------

func main() {

	myWallet := &Wallet{Cash: 100}
	Buy(myWallet)

	var myMoney Payer
	myMoney = &Card{Balance: 100, Cardholder: "vengerov"}
	Buy(myMoney)

	myMoney = &ApplePay{Money: 9}
	Buy(myMoney)
}
