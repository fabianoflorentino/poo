package main

import "fmt"

type Conta struct {
	numero int
	saldo  float64
}

func (c *Conta) Depositar(valor float64) {
	c.saldo += valor
}

func (c *Conta) Sacar(valor float64) {
	c.saldo -= valor
}

func (c *Conta) Saldo() float64 {
	return c.saldo
}

func main() {
	conta := Conta{numero: 123, saldo: 100.0}
	conta.Depositar(50.0)
	conta.Sacar(25.0)
	fmt.Println(conta.Saldo())
}
