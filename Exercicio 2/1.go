package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
)

type Sincronizar struct {
	produto      int
	consumidores chan chan int
}

func (d *Sincronizar) Produtor(proChan chan int) {
	for {
		d.produto = <-proChan
		d.Cast()
	}
}

func (d *Sincronizar) Cast() {
	if len(d.consumidores) <= 1 {
		ch := <-d.consumidores
		ch <- d.produto
	} else {
		for len(d.consumidores) != 0 {
			ch := <-d.consumidores
			ch <- d.produto
		}
	}
}

func main() {

	wg := sync.WaitGroup{}

	sc := &Sincronizar{consumidores: make(chan chan int, 100)}

	seed := rand.NewSource(time.Now().UnixNano())
	gen := rand.New(seed)

	produtos := make(chan int)
	go sc.Produtor(produtos)

	reader := bufio.NewReader(os.Stdin)
INPUT:
	for {
		char, _, _ := reader.ReadRune()

		switch char {
		case 'c':
			go consumir(&wg, sc)
		case 'p':
			i := gen.Intn(100)
			fmt.Println("PRODUZINDO", i)
			produtos <- i
		case 'k':
			break INPUT
		}
	}
}

func consumir(wg *sync.WaitGroup, dt *Sincronizar) {
	wg.Add(1)
	fmt.Println("ADICIONANDO CONSUMIDOR")
	ch := make(chan int)
	dt.consumidores <- ch
	fmt.Println("RECEBEU", <-ch)
}
