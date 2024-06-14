package usuario

import "fmt"

// PrintUsuarioComum imprime os dados de um usuario comum
func PrintUsuario(u Usuario) {
	fmt.Printf("Nome: %s, Email: %s\n", u.Nome(), u.Email())
}
