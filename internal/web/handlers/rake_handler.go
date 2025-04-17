package handlers

import (
	"context"
	"fmt"
	"io"
	"os/exec"
	"strings"
	"sync"

	"github.com/giovanoh/mcp-server-govbox/internal/config"
	"github.com/mark3labs/mcp-go/mcp"
)

// RakeHandler é um manipulador para executar comandos Rake.
type RakeHandler struct {
	shell  config.ShellConfiguration
	action string
}

// NewRakeHandler cria um novo manipulador para executar comandos Rake.
func NewRakeHandler(shell config.ShellConfiguration, action string) *RakeHandler {
	return &RakeHandler{shell: shell, action: action}
}

// Handle executa o comando Rake.
func (h *RakeHandler) Handle(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	projects, ok := request.Params.Arguments["projects"].(string)
	if !ok {
		return mcp.NewToolResultError("Param 'projects' must be a string"), nil
	}

	projects = strings.ToLower(projects)
	command := exec.Command(h.shell.Shell(), h.shell.ShellArgs(), fmt.Sprintf("rake %s projects=%s auto=true", h.action, projects))
	command.Dir = h.shell.WorkingDir()

	stdout, err := command.StdoutPipe()
	if err != nil {
		return mcp.NewToolResultError("Error creating stdout pipe: " + err.Error()), nil
	}
	stderr, err := command.StderrPipe()
	if err != nil {
		return mcp.NewToolResultError("Error creating stderr pipe: " + err.Error()), nil
	}

	var outBuffer, errBuffer strings.Builder
	var wg sync.WaitGroup

	streamReader := func(r io.Reader, buffer *strings.Builder) {
		defer wg.Done()
		if _, err := io.Copy(buffer, r); err != nil {
			buffer.WriteString(fmt.Sprintf("\nError reading stream: %v", err))
		}
	}

	wg.Add(2)
	go streamReader(stdout, &outBuffer)
	go streamReader(stderr, &errBuffer)

	if err := command.Start(); err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Error during command execution (start)\n%s", err.Error())), nil
	}

	wg.Wait()
	if err := command.Wait(); err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Error during command execution (wait)\n%s\n\n%s", err.Error(), errBuffer.String())), nil
	}

	return mcp.NewToolResultText(fmt.Sprintf("Command completed successfully!\n\n%s", outBuffer.String())), nil
}
