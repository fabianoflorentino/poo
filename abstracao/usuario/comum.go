package usuario

// Struct que representa um usuário comum
type UsuarioComum struct {
	nome  string
	email string
}

// Método que retorna o nome do usuário comum
func (uc UsuarioComum) Nome() string {
	return uc.nome
}

// Método que retorna o email do usuário comum
func (uc UsuarioComum) Email() string {
	return uc.email
}

// Função que registra um usuário comum
func RegistraUsuarioComum(nome, email string) Usuario {
	return UsuarioComum{nome, email}
}
