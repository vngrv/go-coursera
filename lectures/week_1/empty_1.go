package main

import (
	"fmt"
	"strconv"
)

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

func (w *Wallet) String() string {
	return "Кошелек в котором " + strconv.Itoa(w.Cash) + " денег"
}

// -----

func main() {
	// myWallet := &Wallet{Cash: 100}
	// fmt.Printf("Raw payment: %#v\n", myWallet)
	// fmt.Printf("Cпособ оплаты: %s\n", myWallet)

	myVal, ok := emptyInterfaceVal.(int)

}
