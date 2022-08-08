package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	b := NewBarbearia(3)
	wg := sync.WaitGroup{}

	go b.Barbeiro(&wg)
	go b.AbrirBarbearia(&wg)

	for i := 0; i < 10; i++ {
		c := Cliente{id: i}
		wg.Add(1)
		go c.Entrar(b.entrada)
	}
	wg.Wait()
}

type Barbearia struct {
	cadeiras         int
	ocupadas         int
	barbeiroAcordado bool
	entrada          chan *Cliente
	cortando         chan *Cliente
}

func NewBarbearia(capacidade int) *Barbearia {
	b := Barbearia{
		cadeiras:         capacidade,
		ocupadas:         0,
		barbeiroAcordado: false,
		entrada:          make(chan *Cliente),
		cortando:         make(chan *Cliente),
	}

	return &b
}

func (b *Barbearia) Barbeiro(wg *sync.WaitGroup) {
	fmt.Println("Barbeiro chegou para trabalhar")
	for {
		if b.ocupadas == 0 {
			fmt.Println("Barbeiro foi dormir")
			b.barbeiroAcordado = false
		}
		Cliente := <- b.cortando
		if !b.barbeiroAcordado {
			fmt.Println("Cliente", Cliente.id, "acordou o barbeiro")
		}
		b.barbeiroAcordado = true
		b.CortarCabelo(*Cliente)
		wg.Done()
	}
}

func (b *Barbearia) CortarCabelo(cl Cliente) {
	b.ocupadas--
	fmt.Println("Cliente", cl.id, "esta cortando o cabelo")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("Cliente", cl.id, "acabou de cortar o cabelo")
}

func (b *Barbearia) AbrirBarbearia(wg *sync.WaitGroup) {
	for Cliente := range b.entrada {
		if b.ocupadas < b.cadeiras {
			b.ocupadas++
			go Cliente.Sentar(b.cortando)
		} else {
			fmt.Println("Cliente", Cliente.id, "tentou entrar mas estava lotado")
			wg.Done()
		}
	}

}

type Cliente struct {
	id int
}

func (c *Cliente) Entrar(entrada chan *Cliente) {
	fmt.Println("Cliente", c.id, "tentou entrar")
	entrada <- c
}

func (c *Cliente) Sentar(cortar chan *Cliente) {
	fmt.Println("Cliente", c.id, "sentou na espera")
	c.DesejoCortarCabelo(cortar)
}

func (c *Cliente) DesejoCortarCabelo(cortar chan *Cliente) {
	cortar <- c
}
