package main

type WithdrawData struct {
	amount   int
	resultCh chan bool
}

var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balance
var withdrawCh = make(chan WithdrawData)

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }
func Withdraw(amount int) bool {
	ch := make(chan bool)
	wd := WithdrawData{amount, ch}
	withdrawCh <- wd
	return <-ch
}

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case withdrawData := <-withdrawCh:
			if balance >= withdrawData.amount {
				balance -= withdrawData.amount
				withdrawData.resultCh <- true
			} else {
				withdrawData.resultCh <- false
			}
		case balances <- balance:
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}
