//go:generate goa gen github.com/pastelnetwork/gonode/walletnode/api/design

package api

import (
	"context"
	"net"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/pastelnetwork/gonode/common/errors"
	"github.com/pastelnetwork/gonode/common/log"
	"github.com/pastelnetwork/gonode/walletnode/api/docs"

	goahttp "goa.design/goa/v3/http"
	goahttpmiddleware "goa.design/goa/v3/http/middleware"
)

const (
	defaultShutdownTimeout = time.Second * 30
	logPrefix              = "api"
)

type service interface {
	Mount(ctx context.Context, mux goahttp.Muxer) goahttp.Server
}

// API represents RESTAPI service.
type API struct {
	config          *Config
	shutdownTimeout time.Duration
	services        []service
}

// Run startworks RESTAPI service.
func (api *API) Run(ctx context.Context) error {
	ctx = context.WithValue(ctx, log.PrefixKey, logPrefix)

	apiHTTP := api.handler(ctx)

	mux := http.NewServeMux()
	mux.Handle("/", apiHTTP)
	mux.Handle("/swagger/swagger.json", apiHTTP)

	if api.config.Swagger {
		mux.Handle("/swagger/", http.FileServer(http.FS(docs.SwaggerContent)))
	}

	addr := net.JoinHostPort(api.config.Hostname, strconv.Itoa(api.config.Port))
	srv := &http.Server{Addr: addr, Handler: mux}

	errCh := make(chan error, 1)
	go func() {
		<-ctx.Done()
		log.WithContext(ctx).Infof("Server is shutting down...")

		ctx, cancel := context.WithTimeout(ctx, api.shutdownTimeout)
		defer cancel()

		srv.SetKeepAlivesEnabled(false)
		if err := srv.Shutdown(ctx); err != nil {
			errCh <- errors.Errorf("gracefully shutdown the server: %w", err)

		}
		close(errCh)
	}()

	log.WithContext(ctx).Infof("Server is ready to handle requests at %q", addr)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return errors.Errorf("error starting server: %w", err)
	}
	defer log.WithContext(ctx).Infof("Server stoped")

	err := <-errCh
	return err
}

func (api *API) handler(ctx context.Context) http.Handler {
	mux := goahttp.NewMuxer()

	var servers goahttp.Servers
	for _, service := range api.services {
		servers = append(servers, service.Mount(ctx, mux))
	}
	servers.Use(goahttpmiddleware.Debug(mux, os.Stdout))

	var handler http.Handler = mux

	handler = Log(ctx)(handler)
	handler = goahttpmiddleware.RequestID()(handler)

	return handler
}

// New returns a new API instance.
func New(config *Config, services ...service) *API {
	return &API{
		config:          config,
		shutdownTimeout: defaultShutdownTimeout,
		services:        services,
	}
}
