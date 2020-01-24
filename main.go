package main

import (
	"context"
	"net/http"

	"github.com/fguy/scooters-api/config"
	handler "github.com/fguy/scooters-api/handlers/scooters"
	"github.com/fguy/scooters-api/repositories"
	repo "github.com/fguy/scooters-api/repositories/scooters"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

var app = fx.New(
	// Provide all the constructors we need, which teaches Fx how we'd like to
	// construct the *config.AppConfig, *zap.Logger and *http.Server types.
	// Remember that constructors are called lazily, so this block doesn't do
	// much on its own.
	fx.Provide(
		config.NewAppConfig,
		NewLogger,
		handler.NewAvailableHandler,
		handler.NewReserveHandler,
		handler.NewHTTPServer,
		repositories.NewDB,
		repo.New,
	),
	// Since constructors are called lazily, we need some invocations to
	// kick-start our application. In this case, we'll use Register. Since it
	// depends on an *grpc.Server and *http.Server, calling it requires Fx
	// to build those types using the constructors above. Since we call
	// NewGRPCServer and NewHTTPServer, we also register Lifecycle hooks to start and stop both
	// servers.
	fx.Invoke(
		RegisterHTTP,
	),
)

// RegisterHTTP starts our HTTP server.
//
// Register is a typical top-level application function: it takes a generic
// type like Server, which typically comes from a third-party library, and
// introduces it to a type that contains our application logic. In this case,
// that introduction consists of registering an HTTP server. Other typical
// examples include registering RPC procedures and starting queue consumers.
//
// Fx calls these functions invocations, and they're treated differently from
// the constructor functions above. Their arguments are still supplied via
// dependency injection and they may still return an error to indicate
// failure, but any other return values are ignored.
//
// Unlike constructors, invocations are called eagerly. See the main function
// below for details.
func RegisterHTTP(cfg *config.AppConfig, lc fx.Lifecycle, logger *zap.Logger, httpServer *http.Server) {
	// If NewHTTPServer is called, we know that another function is using the server. In
	// that case, we'll use the Lifecycle type to register a Hook that starts
	// and stops our HTTP server.
	//
	// Hooks are executed in dependency order. At startup, NewLogger's hooks run
	// before NewHTTPServer's. On shutdown, the order is reversed.
	//
	// Returning an error from OnStart hooks interrupts application startup. Fx
	// immediately runs the OnStop portions of any successfully-executed OnStart
	// hooks (so that types which started cleanly can also shut down cleanly),
	// then exits.
	//
	// Returning an error from OnStop hooks logs a warning, but Fx continues to
	// run the remaining hooks.
	lc.Append(fx.Hook{
		// To mitigate the impact of deadlocks in application startup and
		// shutdown, Fx imposes a time limit on OnStart and OnStop hooks. By
		// default, hooks have a total of 30 seconds to complete. Timeouts are
		// passed via Go's usual context.Context.
		OnStart: func(context.Context) error {
			logger.Info("Starting HTTP server", zap.String("addr", cfg.Addr))
			// In production, we'd want to separate the Listen and Serve phases for
			// better error-handling.
			go httpServer.ListenAndServe()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Info("Stopping HTTP server.")
			return httpServer.Shutdown(ctx)
		},
	})
}

func main() {
	app.Run()
}
