package server

import (
	"fmt"

	"github.com/giovanoh/mcp-server-govbox/internal/web/handlers"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

type ToolHandler struct {
	tool    *mcp.Tool
	handler *handlers.RakeHandler
}

type Server struct {
	name         string
	version      string
	server       *server.MCPServer
	toolHandlers []*ToolHandler
}

func NewServer(name, version string) *Server {
	return &Server{
		name:    name,
		version: version,
	}
}

func (s *Server) RegisterTool(tool *mcp.Tool, handler *handlers.RakeHandler) {
	s.toolHandlers = append(s.toolHandlers, &ToolHandler{
		tool:    tool,
		handler: handler,
	})
}

func (s *Server) Start() {
	s.server = server.NewMCPServer(s.name, s.version)
	for _, toolHandler := range s.toolHandlers {
		s.server.AddTool(*toolHandler.tool, toolHandler.handler.Handle)
	}
	if err := server.ServeStdio(s.server); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
