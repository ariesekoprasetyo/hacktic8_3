package main

import (
	"fmt"
	"time"
)

func customer(id int, barberCh, customerCh chan int) {
	select {
	case barberCh <- id:
		fmt.Printf("Customer %d: entering the barber shop\n", id)
	default:
		fmt.Printf("Customer %d: barber shop is full. I'll come later\n", id)
		return
	}
	<-customerCh
	fmt.Printf("Customer %d: getting a haircut\n", id)
	time.Sleep(2 * time.Second)
	fmt.Printf("Customer %d: leaving the barber shop\n", id)
}

func barber(barberCh, customerCh chan int) {
	for {
		select {
		case id := <-barberCh:
			fmt.Printf("Barber: starting to cut hair for customer %d\n", id)
			time.Sleep(3 * time.Second)
			fmt.Printf("Barber: finished cutting hair for customer %d\n", id)
			customerCh <- 1
		default:
			fmt.Printf("Barber: no customers. I'm taking a nap.\n")
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	barberCh := make(chan int, 2)
	customerCh := make(chan int, 2)
	go barber(barberCh, customerCh)

	for i := 1; i <= 6; i++ {
		time.Sleep(500 * time.Millisecond)
		go customer(i, barberCh, customerCh)
	}

	time.Sleep(10 * time.Second)
	fmt.Println("Done")
}
