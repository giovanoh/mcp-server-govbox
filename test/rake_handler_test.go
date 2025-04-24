package test

import (
	"context"
	"testing"

	"github.com/giovanoh/mcp-server-govbox/internal/config"
	"github.com/giovanoh/mcp-server-govbox/internal/web/handlers"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Should_Return_Error_When_Action_Is_Empty(t *testing.T) {
	shellCfg, _ := config.NewShellConfiguration("/bin/bash", "-c", "/valid/path")
	_, err := handlers.NewRakeHandler(shellCfg, "")

	require.Error(t, err)
	require.ErrorIs(t, err, handlers.ErrInvalidAction)
}

func Test_Handle_Should_Return_Error_When_Projects_Param_Is_Missing(t *testing.T) {
	shellCfg, _ := config.NewShellConfiguration("/bin/bash", "-c", "/valid/path")
	handler, _ := handlers.NewRakeHandler(shellCfg, "namespace:action")

	request := mcp.CallToolRequest{}
	request.Params.Arguments = map[string]interface{}{}

	result, err := handler.Handle(context.Background(), request)

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.True(t, result.IsError)
	assert.Contains(t, result.Content[0].(mcp.TextContent).Text, "Param 'projects' must be a string")
}

func Test_Handle_Should_Return_Error_When_Projects_Param_Is_Not_String(t *testing.T) {
	shellCfg, _ := config.NewShellConfiguration("/bin/bash", "-c", "/valid/path")
	handler, _ := handlers.NewRakeHandler(shellCfg, "namespace:action")

	request := mcp.CallToolRequest{}
	request.Params.Arguments = map[string]interface{}{
		"projects": 123,
	}

	result, err := handler.Handle(context.Background(), request)

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.True(t, result.IsError)
	assert.Contains(t, result.Content[0].(mcp.TextContent).Text, "Param 'projects' must be a string")
}

func Test_Should_Create_RakeHandler_Successfully(t *testing.T) {
	shellCfg, _ := config.NewShellConfiguration("/bin/bash", "-c", "/valid/path")
	action := "namespace:action"
	handler, err := handlers.NewRakeHandler(shellCfg, action)

	require.NoError(t, err)
	require.NotNil(t, handler)

	assert.Equal(t, action, handler.Action())
	assert.Equal(t, shellCfg, handler.Shell())
}
