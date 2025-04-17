package config

import "github.com/giovanoh/mcp-server-govbox/internal/domain/shellerrors"

// ShellConfiguration representa a configuração de um shell para execução de comandos
type ShellConfiguration struct {
	command    string
	args       string
	workingDir string
}

// NewShellConfiguration cria uma nova instância de ShellConfiguration
func NewShellConfiguration(shell, shellArgs, workingDir string) (ShellConfiguration, error) {
	if shell == "" {
		return ShellConfiguration{}, shellerrors.NewErrInvalidShell("shell command cannot be empty")
	}
	if workingDir == "" {
		return ShellConfiguration{}, shellerrors.NewErrInvalidShell("working directory cannot be empty")
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
