package main

import (
	"fmt"
	"time"
)

type Message struct {
	OrderId string
	Title   string
	Price   int
}

func buyTicket(channel chan<- Message, orders []Message) {
	for _, order := range orders {
		time.Sleep(time.Second * 1)
		fmt.Printf("Send Buy:::%s\n", order.OrderId)
		channel <- order // send order message
	}

	close(channel)
}

func cancelTicket(channel chan<- string, cancelOrder []string) {
	for _, orderId := range cancelOrder {
		time.Sleep(time.Second * 10) //delay
		fmt.Printf("Send cancel ticket:::%s\n", orderId)
		channel <- orderId // send order message
	}

	close(channel)
}

func handlerOrder(orderChannel <-chan Message, cancelChannel <-chan string) {
	for {
		order, orderOK := <-orderChannel // read goroutine order message
		if orderOK {
			fmt.Printf("Handler buy ticket::%s - %s - %d\n", order.OrderId, order.Title, order.Price)
		} else {
			fmt.Println("Order channel is closed.")
			break
		}

		cancel, cancelOK := <-cancelChannel // read goroutine cancel message
		if cancelOK {
			fmt.Printf("Handler cancel ticket::%s\n", cancel)
		} else {
			fmt.Println("Cancel channel is closed.")
			break
		}
	}
}

func handlerOrderWithSelect(orderChannel <-chan Message, cancelChannel <-chan string) {
	for {
		select {
		case order, orderOK := <-orderChannel: // read goroutine order message
			if orderOK {
				fmt.Printf("handler buy ticket::%s - %s - %d\n", order.OrderId, order.Title, order.Price)
			} else {
				fmt.Println("Order channel is closed.")
				orderChannel = nil
			}

		case cancel, cancelOK := <-cancelChannel: // read goroutine cancel message
			if cancelOK {
				fmt.Printf("handler cancel ticket::%s\n", cancel)
			} else {
				fmt.Println("Cancel channel is closed.")
				cancelChannel = nil
			}
		}

		// Exit when all channels are closed
		if orderChannel == nil && cancelChannel == nil {
			break
		}
	}
}

func main() {
	buyChannel := make(chan Message)
	cancelChannel := make(chan string)

	// 2 - Simulate orders
	buyOrders := []Message{
		{OrderId: "Ordeer-01", Title: "Tips Go", Price: 30},
		{OrderId: "Ordeer-02", Title: "Tips Nodejs", Price: 30},
		{OrderId: "Ordeer-03", Title: "Tips Java", Price: 50},
	}

	cancelOrders := []string{"Ordeer-01", "Ordeer-03"}

	go buyTicket(buyChannel, buyOrders)
	go cancelTicket(cancelChannel, cancelOrders)

	// handler order
	// go handlerOrder(buyChannel, cancelChannel)
	go handlerOrderWithSelect(buyChannel, cancelChannel)

	time.Sleep(time.Second * 25)
	fmt.Println("End buying and canceling..")
}
