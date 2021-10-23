package main

import (
	"log"
	"os"

	// application
	"github.com/vrras/go-hexagonal-arch/internal/application/api"
	"github.com/vrras/go-hexagonal-arch/internal/application/core/arithmetic"

	// adapters
	gRPC "github.com/vrras/go-hexagonal-arch/internal/adapters/framework/left/grpc"
	"github.com/vrras/go-hexagonal-arch/internal/adapters/framework/right/db"
)

func main() {
	var err error

	dbasDriver := os.Getenv("DB_DRIVER")
	dsourceName := os.Getenv("DS_NAME")

	dbAdapter, err := db.NewAdapter(dbasDriver, dsourceName)
	if err != nil {
		log.Fatalf("failed to initiate dbase connection: %v", err)
	}
	defer dbAdapter.CloseDBConnection()

	// core
	core := arithmetic.New()

	// NOTE: The application's right side port for driven
	// adapters, in this case, a db adapter.
	// Therefore the type for the dbAdapter parameter
	// that is to be injected into the NewApplication will
	// be of type DbPort
	applicationAPI := api.NewApplication(dbAdapter, core)

	// NOTE: We use dependency injection to give the grpc
	// adapter access to the application, therefore
	// the location of the port is inverted. That is
	// the grpc adapter accesses the hexagon's driving port at the
	// application boundary via dependency injection,
	// therefore the type for the applicationAPI parameter
	// that is to be injected into the gRPC adapter will
	// be of type APIPort which is our hexagons left side
	// port for driving adapters
	gRPCAdapter := gRPC.NewAdapter(applicationAPI)
	gRPCAdapter.Run()
}
