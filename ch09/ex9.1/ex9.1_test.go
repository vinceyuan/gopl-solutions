package main

import (
	"fmt"
	"testing"
)

func TestBank(t *testing.T) {
	done := make(chan struct{})

	// Alice
	go func() {
		Deposit(200)
		if ok := Withdraw(20); ok {
			fmt.Println("Withdraw 20")
		} else {
			fmt.Println("Can't withdraw 20")
		}
		if ok := Withdraw(400); ok {
			fmt.Println("Withdraw 400")
		} else {
			fmt.Println("Can't withdraw 400")
		}
		fmt.Println("=", Balance())
		done <- struct{}{}
	}()

	// Bob
	go func() {
		Deposit(100)
		if ok := Withdraw(20); ok {
			fmt.Println("Withdraw 20")
		} else {
			fmt.Println("Can't withdraw 20")
		}
		if ok := Withdraw(400); ok {
			fmt.Println("Withdraw 400")
		} else {
			fmt.Println("Can't withdraw 400")
		}
		done <- struct{}{}
	}()

	// Wait for both transactions.
	<-done
	<-done

	if got, want := Balance(), 260; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}
}
