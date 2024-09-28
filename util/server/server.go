package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

// Config represents server specific config
type Config struct {
	Port         int
	ReadTimeout  int
	WriteTimeout int
	AllowOrigins []string
}

// DefaultConfig for the API server
var DefaultConfig = Config{
	Port:         3000,
	ReadTimeout:  10,
	WriteTimeout: 5,
	AllowOrigins: []string{"*"},
}

func (c *Config) fillDefaults() {
	if c.Port == 0 {
		c.Port = DefaultConfig.Port
	}
	if c.ReadTimeout == 0 {
		c.ReadTimeout = DefaultConfig.ReadTimeout
	}
	if c.WriteTimeout == 0 {
		c.WriteTimeout = DefaultConfig.WriteTimeout
	}
	if c.AllowOrigins == nil && len(c.AllowOrigins) == 0 {
		c.AllowOrigins = DefaultConfig.AllowOrigins
	}
}

// New instance new Echo server
func New(cfg *Config) *echo.Echo {
	cfg.fillDefaults()
	e := echo.New()

	e.HTTPErrorHandler = NewErrorHandler(e).Handle
	e.Binder = NewBinder()
	e.Logger.SetLevel(log.DEBUG)

	e.Server.Addr = fmt.Sprintf(":%d", cfg.Port)
	e.Server.ReadTimeout = time.Duration(cfg.ReadTimeout) * time.Minute
	e.Server.WriteTimeout = time.Duration(cfg.WriteTimeout) * time.Minute
	return e
}

// Start starts echo server
func Start(e *echo.Echo) {
	// Start server
	go func() {
		if err := e.StartServer(e.Server); err != nil {
			if err == http.ErrServerClosed {
				fmt.Println("shutting down the server")
			} else {
				fmt.Printf("error shutting down the server: %v", err)
			}
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal, 1) // buffered channel
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
