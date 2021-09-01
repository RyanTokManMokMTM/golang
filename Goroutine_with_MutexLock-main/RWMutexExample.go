package main

import (
	"log"
	"sync"
	"time"
)

type Bank struct {
	Balance int
	mutx    sync.RWMutex //multiple mutex 1 write
}

func (bank *Bank) Deposit(amount int) {
	//wait here
	bank.mutx.Lock() //write lock
	time.Sleep(time.Second) //wait 1s
	bank.Balance += amount
	bank.mutx.Unlock() //write unlock
}

func (bank *Bank) AccountBalance() (accountBalance int) {
	bank.mutx.RLock() //read lock ->can't wirte
	time.Sleep(time.Second) //wait 1s
	accountBalance = bank.Balance
	bank.mutx.RUnlock() //read unlock
	return
}

func main() {
	myBank := new(Bank)
	waitGp := new(sync.WaitGroup)
	time := 5
	waitGp.Add(time)

	//wait
	for i := 0; i < time; i++ {
		go func() {
			myBank.Deposit(500)
			log.Printf("Write:My Bank Deposit amount %v", 1000)
			waitGp.Done()
		}()
	}

	//read
	waitGp.Add(time)
	for i := 0; i < time; i++ {
		go func() {
			log.Printf("Read:My Bank balance %v", myBank.AccountBalance())
			waitGp.Done()
		}()
	}

	waitGp.Wait()
}
