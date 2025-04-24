package config

import (
	"fmt"
)

// ShellConfiguration representa a configuração de um shell para execução de comandos
type ShellConfiguration struct {
	command    string
	args       string
	workingDir string
}

// NewShellConfiguration cria uma nova instância de ShellConfiguration
func NewShellConfiguration(shell, shellArgs, workingDir string) (ShellConfiguration, error) {
	if shell == "" {
		return ShellConfiguration{}, NewErrInvalidShell("shell command cannot be empty")
	}
	if workingDir == "" {
		return ShellConfiguration{}, NewErrInvalidShell("working directory cannot be empty")
	}

	return ShellConfiguration{
		command:    shell,
		args:       shellArgs,
		workingDir: workingDir,
	}, nil
}

// Shell retorna o comando do shell
func (s ShellConfiguration) Shell() string {
	return s.command
}

// ShellArgs retorna os argumentos do shell
func (s ShellConfiguration) ShellArgs() string {
	return s.args
}

// WorkingDir retorna o diretório de trabalho
func (s ShellConfiguration) WorkingDir() string {
	return s.workingDir
}

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
