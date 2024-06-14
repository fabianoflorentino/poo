package main

import (
	"abstracao/usuario"
)

func main() {
	usuario_comum := usuario.RegistraUsuarioComum("Usuario Commum", "usuario_comum@email.com")
	usuario_administrador := usuario.RegistraUsuarioAdministrador("Usuario Administrador", "usuario_administrador@email.com")

	usuario.PrintUsuario(usuario_comum)
	usuario.PrintUsuario(usuario_administrador)
}
