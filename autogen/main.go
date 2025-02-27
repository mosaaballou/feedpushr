//go:generate goagen bootstrap -d github.com/ncarlier/feedpushr/design

package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/ncarlier/feedpushr/autogen/app"
)

func main() {
	// Create service
	service := goa.New("feedpushr")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "feed" controller
	c := NewFeedController(service)
	app.MountFeedController(service, c)
	// Mount "filter" controller
	c2 := NewFilterController(service)
	app.MountFilterController(service, c2)
	// Mount "health" controller
	c3 := NewHealthController(service)
	app.MountHealthController(service, c3)
	// Mount "opml" controller
	c4 := NewOpmlController(service)
	app.MountOpmlController(service, c4)
	// Mount "output" controller
	c5 := NewOutputController(service)
	app.MountOutputController(service, c5)
	// Mount "pshb" controller
	c6 := NewPshbController(service)
	app.MountPshbController(service, c6)
	// Mount "swagger" controller
	c7 := NewSwaggerController(service)
	app.MountSwaggerController(service, c7)
	// Mount "vars" controller
	c8 := NewVarsController(service)
	app.MountVarsController(service, c8)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}
}
