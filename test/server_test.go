package test

import (
	"testing"

	"github.com/giovanoh/mcp-server-govbox/internal/web/handlers"
	"github.com/giovanoh/mcp-server-govbox/internal/web/server"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Should_Create_Server_Successfully(t *testing.T) {
	name := "test-server"
	version := "1.0.0"

	srv := server.NewServer(name, version)

	require.NotNil(t, srv, "NewServer should return a non-nil server instance")
	assert.Equal(t, name, srv.Name())
	assert.Equal(t, version, srv.Version())

	toolHandlers := srv.ToolHandlers()
	assert.NotNil(t, toolHandlers)

	srv.RegisterTool(&mcp.Tool{}, &handlers.RakeHandler{})
	assert.Equal(t, 1, len(srv.ToolHandlers()))
}
