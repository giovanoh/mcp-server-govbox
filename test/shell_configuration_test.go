package test

import (
	"testing"

	"github.com/giovanoh/mcp-server-govbox/internal/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Should_Return_Error_When_Shell_Commnand_Is_Empty(t *testing.T) {
	_, err := config.NewShellConfiguration("", "-c", "/some/dir")

	require.Error(t, err)

	var expectedErr *config.ErrInvalidShell
	assert.ErrorAs(t, err, &expectedErr)
}

func Test_Should_Return_Error_When_Shell_Working_Directory_Is_Empty(t *testing.T) {
	_, err := config.NewShellConfiguration("/bin/bash", "-c", "")

	require.Error(t, err)

	var expectedErr *config.ErrInvalidShell
	assert.ErrorAs(t, err, &expectedErr)
}

func Test_Should_Create_Configuration_Successfully(t *testing.T) {
	shellCmd := "/bin/zsh"
	shellArgs := "-c"
	workingDir := "/path/to/project"
	cfg, err := config.NewShellConfiguration(shellCmd, shellArgs, workingDir)

	require.NoError(t, err)

	assert.Equal(t, shellCmd, cfg.Shell())
	assert.Equal(t, shellArgs, cfg.ShellArgs())
	assert.Equal(t, workingDir, cfg.WorkingDir())
}
