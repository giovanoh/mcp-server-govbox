package shellerrors

import "fmt"

// ErrInvalidShell representa um erro de configuração de shell inválida.
type ErrInvalidShell struct {
	Message string
}

// NewErrInvalidShell cria um novo erro de configuração de shell inválida.
func NewErrInvalidShell(message string) *ErrInvalidShell {
	return &ErrInvalidShell{
		Message: message,
	}
}

// Error retorna a mensagem de erro.
func (e *ErrInvalidShell) Error() string {
	return fmt.Sprintf("invalid shell: %s", e.Message)
}
