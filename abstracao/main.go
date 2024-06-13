package main

import (
	"abstracao/usuario"
	"fmt"
)

func main() {
	usuario_comum := usuario.RegistraUsuarioComum("Usuario Commum", "usuario_comum@email.com")
	fmt.Println("Usuario Comum")
	fmt.Printf("Nome: %s, Email: %s\n\n", usuario_comum.Nome(), usuario_comum.Email())

	usuario_administrador := usuario.RegistraUsuarioAdministrador("Usuario Administrador", "usuario_administrador@email.com")

	fmt.Println("Usuario Administrador")
	fmt.Printf("Nome: %s, Email: %s\n", usuario_administrador.Nome(), usuario_administrador.Email())
}
