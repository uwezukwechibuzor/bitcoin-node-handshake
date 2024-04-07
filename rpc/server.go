package rpc

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"
)

// Server represents a JSON-RPC server.
type Server struct {
	Port   int
	RPC    *rpc.Server
	logger *zap.Logger
}

// NewServer creates a new JSON-RPC server.
func NewServer(port int, node Node) (*Server, error) {
	rpcs := rpc.NewServer()

	handlers := RPC{node: node}
	if err := rpcs.Register(handlers); err != nil {
		return nil, err
	}

	// Initialize Zap logger
	logger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}

	s := Server{
		Port:   port,
		RPC:    rpcs,
		logger: logger,
	}

	return &s, nil
}

// Run starts the JSON-RPC server.
func (s *Server) Run() {
	// Initialize signal handler for graceful shutdown
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", s.Port))
	if err != nil {
		s.logger.Error("failed to start JSON-RPC server", zap.Error(err))
		return
	}

	s.logger.Info("JSON-RPC server started successfully", zap.Int("port", s.Port))
	defer func() {
		_ = listener.Close()
		s.logger.Sync()
	}()

	// Handle graceful shutdown
	go func() {
		<-sig
		s.logger.Info("Received interrupt signal. Shutting down...")
		_ = listener.Close()
		os.Exit(0)
	}()

	for {
		conn, err := listener.Accept()
		if err != nil {
			s.logger.Error("JSON-RPC connection failed", zap.Error(err))
			// Continue listening for new connections
			continue
		}

		go s.RPC.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
