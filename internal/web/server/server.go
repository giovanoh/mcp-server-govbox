package server

import (
	"fmt"

	"github.com/giovanoh/mcp-server-govbox/internal/web/handlers"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// ToolHandler representa um manipulador para um comando MCP.
type ToolHandler struct {
	tool    *mcp.Tool
	handler *handlers.RakeHandler
}

// Server representa um servidor MCP.
type Server struct {
	name         string
	version      string
	server       *server.MCPServer
	toolHandlers []*ToolHandler
}

// NewServer cria um novo servidor MCP.
func NewServer(name, version string) *Server {
	return &Server{
		name:         name,
		version:      version,
		toolHandlers: make([]*ToolHandler, 0),
	}
}

// RegisterTool registra um novo comando MCP.
func (s *Server) RegisterTool(tool *mcp.Tool, handler *handlers.RakeHandler) {
	s.toolHandlers = append(s.toolHandlers, &ToolHandler{
		tool:    tool,
		handler: handler,
	})
}

// Start inicia o servidor MCP.
func (s *Server) Start() {
	s.server = server.NewMCPServer(s.name, s.version)
	for _, toolHandler := range s.toolHandlers {
		s.server.AddTool(*toolHandler.tool, toolHandler.handler.Handle)
	}
	if err := server.ServeStdio(s.server); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}

// Name retorna o nome do servidor.
func (s *Server) Name() string {
	return s.name
}

// Version retorna a versão do servidor.
func (s *Server) Version() string {
	return s.version
}

// ToolHandlers retorna os manipuladores de comandos MCP.
func (s *Server) ToolHandlers() []*ToolHandler {
	return s.toolHandlers
}
