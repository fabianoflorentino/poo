package main

import "fmt"

// Veiculo é uma interface que define um método dados que retorna uma string
type Veiculo interface {
	dados() string
}

// Carro é uma struct que representa um carro
type Carro struct {
	marca  string
	modelo string
}

// dados é um método que retorna uma string com os dados do carro
func (c Carro) dados() string {
	return fmt.Sprintf("Marca: %s, Modelo: %s", c.marca, c.modelo)
}

// Hatch é uma struct que representa um carro do tipo hatch
type Hatch struct {
	Carro
	portas int
}

// dados é um método que retorna uma string com os dados do carro hatch
func (h Hatch) dados() string {
	return fmt.Sprintf("Marca: %s, Modelo: %s, Portas: %d", h.marca, h.modelo, h.portas)
}

// Sedan é uma struct que representa um carro do tipo sedan
type Sedan struct {
	Carro
	portaMalas int
}

// dados é um método que retorna uma string com os dados do carro sedan
func (s Sedan) dados() string {
	return fmt.Sprintf("Marca: %s, Modelo: %s, Porta Malas: %d", s.marca, s.modelo, s.portaMalas)
}

type Conversivel struct {
	Carro
	capota bool
}

func (c Conversivel) dados() string {
	return fmt.Sprintf("%s, Capota: %t", c.Carro.dados(), c.capota)
}

// imprimirDados é uma função que recebe um veículo e imprime os dados do veículo
func imprimirDados(v Veiculo) {
	fmt.Println(v.dados())
}

func main() {
	// Acessando atributos
	hatch := Hatch{Carro{"Chevrolet", "Onix"}, 4}
	sedan := Sedan{Carro{"Honda", "Civic"}, 500}

	// Acessando métodos dados da struct Carro de forma explícita
	conversivel := Conversivel{Carro{"Fiat", "Spyder"}, true}

	imprimirDados(hatch)
	imprimirDados(sedan)
	imprimirDados(conversivel)
}
