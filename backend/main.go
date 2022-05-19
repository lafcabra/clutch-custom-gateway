package main

import (
	"github.com/lyft/clutch/backend/gateway"

	"github.com/lafcabra/clutch-custom-gateway/backend/cmd/assets"
	amiibomod "github.com/lafcabra/clutch-custom-gateway/backend/module/amiibo"
	"github.com/lafcabra/clutch-custom-gateway/backend/module/echo"
	amiiboservice "github.com/lafcabra/clutch-custom-gateway/backend/service/amiibo"
)

func main() {
	flags := gateway.ParseFlags()

	components := gateway.CoreComponentFactory
	components.Modules[amiibomod.Name] = amiibomod.New
	components.Services[amiiboservice.Name] = amiiboservice.New
	// Add custom components.
	components.Modules[echo.Name] = echo.New

	gateway.Run(flags, components, assets.VirtualFS)
}
