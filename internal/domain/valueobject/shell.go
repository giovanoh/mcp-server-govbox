package valueobject

import "github.com/giovanoh/mcp-server-govbox/internal/domain/shellerrors"

// Shell representa a configuração de um shell para execução de comandos
type Shell struct {
	command    string
	args       string
	workingDir string
}

// NewShell cria uma nova instância de Shell
func NewShell(shell, shellArgs, workingDir string) (Shell, error) {
	if shell == "" {
		return Shell{}, shellerrors.NewErrInvalidShell("shell command cannot be empty")
	}
	if workingDir == "" {
		return Shell{}, shellerrors.NewErrInvalidShell("working directory cannot be empty")
	}

	return Shell{
		command:    shell,
		args:       shellArgs,
		workingDir: workingDir,
	}, nil
}

// Shell retorna o comando do shell
func (s Shell) Shell() string {
	return s.command
}

// ShellArgs retorna os argumentos do shell
func (s Shell) ShellArgs() string {
	return s.args
}

// WorkingDir retorna o diretório de trabalho
func (s Shell) WorkingDir() string {
	return s.workingDir
}
