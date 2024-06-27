package main

import "fmt"

// Conta é uma estrutura que representa uma conta bancária.
type Conta struct {
	numero int
	saldo  float64
}

// Depositar é um método que adiciona um valor ao saldo da conta.
func (c *Conta) Depositar(valor float64) {
	c.saldo += valor
}

// Sacar é um método que subtrai um valor do saldo da conta.
func (c *Conta) Sacar(valor float64) {
	c.saldo -= valor
}

// Saldo é um método que retorna o saldo da conta.
func (c *Conta) Saldo() float64 {
	return c.saldo
}

func main() {
	// Cria uma nova conta com saldo inicial de 100.0.
	conta := Conta{numero: 123, saldo: 100.0}

	// Tell, Don't Ask - Princípio de design de software.
	// Depositar e Sacar são métodos que informam a conta sobre as ações a serem realizadas.
	conta.Depositar(50.0)
	conta.Sacar(25.0)

	fmt.Println(conta.Saldo())
}
