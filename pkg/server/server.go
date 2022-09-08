package server

import (
	"context"
	"deall/pkg/router"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"golang.org/x/net/netutil"
)

// Options for the web Handler.
type Options struct {
	ListenAddress string        `envconfig:"LISTEN_ADDRESS"`
	MaxConnection int           `envconfig:"MAX_CONNECTION"`
	ReadTimeout   time.Duration `envconfig:"READ_TIMEOUT"`
	WriteTimeout  time.Duration `envconfig:"WRITE_TIMEOUT"`
	Timeout       time.Duration `envconfig:"TIMEOUT"`
}

// Server serves various HTTP endpoints of the application server
type Server interface {
	Run() chan error
	Router() *router.Router
	Stop() error
}

type server struct {
	router  *router.Router
	server  *http.Server
	options *Options
}

// New initialize server
func New(options *Options) Server {
	return &server{
		router:  router.New(),
		options: options,
	}
}

func (s *server) Router() *router.Router {
	return s.router
}

func (s *server) Run() chan error {
	var ch = make(chan error)

	go s.run(ch)
	return ch
}

func (s *server) run(ch chan error) {
	listener, err := net.Listen("tcp", s.options.ListenAddress)
	if err != nil {
		ch <- err
		return
	}

	if s.options.MaxConnection > 0 {
		listener = netutil.LimitListener(listener, s.options.MaxConnection)
	}

	s.server = &http.Server{
		Handler:      s.router,
		ReadTimeout:  s.options.ReadTimeout,
		WriteTimeout: s.options.WriteTimeout,
	}
	log.Println(fmt.Sprintf("Http Server running on Port %s", s.options.ListenAddress))
	ch <- s.server.Serve(listener)
}

func (s *server) Stop() error {
	if s.server == nil {
		return nil
	}

	timeout := s.options.Timeout
	if timeout <= 0 {
		timeout = time.Second * 20
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	return s.server.Shutdown(ctx)
}
