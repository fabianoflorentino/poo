package usuario

// Struct que representa um usuário administrador
type UsuarioAdministrador struct {
	nome  string
	email string
}

// Método que retorna o nome do usuário administrador
func (ua UsuarioAdministrador) Nome() string {
	return ua.nome
}

// Método que retorna o email do usuário administrador
func (ua UsuarioAdministrador) Email() string {
	return ua.email
}

// Função que registra um usuário administrador
func RegistraUsuarioAdministrador(nome, email string) Usuario {
	return UsuarioAdministrador{nome, email}
}
