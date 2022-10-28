package internal

import (
	"go-quick-template/internal/migration"
	"go-quick-template/internal/routers/api"
	"go-quick-template/internal/service"
)

func Initialize() {
	// migrate database if needed
	migration.Run()

	// initialize service
	service.Initialize()
	api.Initialize()
}
