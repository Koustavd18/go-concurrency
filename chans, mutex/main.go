package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/fatih/color"
)

const NumberOfPizza = 10

var pizzaMade, pizzaFailed, total int

type Producer struct {
	data chan PizzaOrder
	quit chan chan error
}

type PizzaOrder struct {
	pizzaNumber int
	message     string
	success     bool
}

func (p *Producer) Close() error {
	ch := make(chan error)

	p.quit <- ch

	return <-ch
}

func makePizza(pizzaNum int) *PizzaOrder {
	pizzaNum++

	if pizzaNum <= NumberOfPizza {
		delay := rand.Intn(5) + 1
		fmt.Printf("Recieved Order #%d\n", pizzaNum)

		rnd := rand.Intn(12) + 1
		msg := ""
		success := false

		if rnd < 5 {
			pizzaFailed++
		} else {
			pizzaMade++
		}
		total++

		fmt.Printf("Making pizza num #%d will take %d secs...\n", pizzaNum, delay)

		time.Sleep(time.Duration(delay) * time.Second)

		if rnd <= 2 {
			msg = fmt.Sprintf(" we ran out of stuffs for pizza #%d", pizzaNum)
		} else if rnd <= 4 {
			msg = fmt.Sprintf("cook quitted while making pizza #%d", pizzaNum)
		} else {
			success = true
			msg = fmt.Sprintf("pizza #%d is ready", pizzaNum)
		}

		p := PizzaOrder{
			pizzaNum,
			msg,
			success,
		}

		return &p
	}

	return &PizzaOrder{
		pizzaNumber: pizzaNum,
	}
}

func pizzeria(pizzaMaker *Producer) {

	i := 0

	for {
		currentPizza := makePizza(i)
		if currentPizza != nil {
			i = currentPizza.pizzaNumber
			select {
			case pizzaMaker.data <- *currentPizza:

			case quitCh := <-pizzaMaker.quit:
				close(pizzaMaker.data)
				close(quitCh)
				return
			}
		}
	}

}

func main() {

	rand.Seed(time.Now().UnixNano())

	color.Cyan("The Pizzeria is open for business")
	color.Cyan("---------------------------------")

	pizzaJob := &Producer{
		make(chan PizzaOrder),
		make(chan chan error),
	}

	go pizzeria(pizzaJob)

	for i := range pizzaJob.data {
		if i.pizzaNumber <= NumberOfPizza {

			if i.success {
				color.Green(i.message)
				color.Green("OFD #%d\n", i.pizzaNumber)
			} else {
				color.Red(i.message)
				color.Red("customer is mad")
			}
		} else {
			color.Cyan("Done making round breads")
			err := pizzaJob.Close()
			if err != nil {
				color.Red("error closing ch")
			}
		}
	}
	color.Cyan("----------------")
	color.Cyan("Done for the day")

	color.Cyan("we made %d , failed %d, total %d", pizzaMade, pizzaFailed, total)

}

// -------------------------------------

type Income struct {
	source string
	amount int
}

var wg sync.WaitGroup

func balaceBank() {

	bankBalance := 0
	var balance sync.Mutex

	incomes := []Income{
		{"job", 500},
		{"freelance", 50},
		{"rent", 100},
		{"gift", 10},
	}

	fmt.Printf("Initial Bank balance %d \n", bankBalance)

	wg.Add(len(incomes))

	for _, income := range incomes {

		go func(income Income) {
			defer wg.Done()

			for i := 1; i < 53; i++ {
				balance.Lock()
				temp := bankBalance
				temp += income.amount
				bankBalance = temp
				balance.Unlock()
			}

		}(income)
	}
	wg.Wait()

	fmt.Printf("Final account balance %d\n", bankBalance)

}
