// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 261.
//!+

// Package bank provides a concurrency-safe bank with one account.
package bank

type withdrawal struct {
	amount  int
	success chan bool
}

var deposits = make(chan int)          // send amount to deposit
var balances = make(chan int)          // receive balance
var withdrawls = make(chan withdrawal) // send withdrawl amount

//Deposit deposits amount to account
func Deposit(amount int) { deposits <- amount }

//Balance returns current balance
func Balance() int { return <-balances }

//Withdraw attempts to withdraw amount from account
func Withdraw(amount int) bool {
	ch := make(chan bool)
	withdrawls <- withdrawal{amount, ch}
	return <-ch
}

//Monitor
func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case w := <-withdrawls:
			if w.amount <= balance {
				w.success <- true
			} else {
				w.success <- false
			}
		case balances <- balance:
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}

//!-
