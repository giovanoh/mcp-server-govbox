package main

import (
	"log"
	"os"

	"github.com/giovanoh/mcp-server-govbox/internal/domain/valueobject"
	"github.com/giovanoh/mcp-server-govbox/internal/web/handlers"
	"github.com/giovanoh/mcp-server-govbox/internal/web/server"

	"github.com/joho/godotenv"
	"github.com/mark3labs/mcp-go/mcp"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file\n", err)
	}

	shell, err := valueobject.NewShell(
		os.Getenv("MCP_SHELL_PATH"),
		os.Getenv("MCP_SHELL_OPTIONS"),
		os.Getenv("MCP_GOVBOX_PROJECT_PATH"),
	)
	if err != nil {
		log.Fatal("Invalid shell configuration\n", err)
	}

	toolBuild := mcp.NewTool("build_govbox",
		mcp.WithDescription("Build a project from govbox solution"),
		mcp.WithString("projects",
			mcp.Required(),
			mcp.Description("The projects to build, separated by comma"),
		),
	)

	toolUpload := mcp.NewTool("upload_govbox",
		mcp.WithDescription("Upload a project from govbox solution"),
		mcp.WithString("projects",
			mcp.Required(),
			mcp.Description("The projects to upload, separated by comma"),
		),
	)

	toolUpdateDb := mcp.NewTool("update_db_govbox",
		mcp.WithDescription("Update the database with the new version of the project"),
		mcp.WithString("projects",
			mcp.Required(),
			mcp.Description("The projects to update the database, separated by comma"),
		),
	)

	buildHandler := handlers.NewRakeHandler(shell, "build:trunk")
	uploadHandler := handlers.NewRakeHandler(shell, "build:upload_site")
	updateDbHandler := handlers.NewRakeHandler(shell, "build:update_db")

	server := server.NewServer("mcp-server-govbox", "1.0.0")
	server.RegisterTool(&toolBuild, buildHandler)
	server.RegisterTool(&toolUpload, uploadHandler)
	server.RegisterTool(&toolUpdateDb, updateDbHandler)
	server.Start()
}
