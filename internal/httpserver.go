package internal

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/venuyeredla/pan-services/internal/config"
)

const Timeout = 5 * time.Second

type WaitFunc func()

type HttpServer struct {
	httpConfig *config.Config
	server     *http.Server
	waitGroup  *sync.WaitGroup
}

func NewHttpServer(config *config.Config, handler http.Handler, waitGroup *sync.WaitGroup) *HttpServer {

	server := &http.Server{
		Addr:         fmt.Sprintf(":%v", config.Http.Port),
		Handler:      handler,
		ReadTimeout:  time.Duration(config.Http.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(config.Http.WriteTimeout) * time.Second,
	}

	return &HttpServer{
		httpConfig: config,
		server:     server,
		waitGroup:  waitGroup,
	}
}

func (s *HttpServer) StartUp(quit <-chan struct{}) {

	s.waitGroup.Add(1)

	go func() {
		defer s.waitGroup.Done()
		// Wait for interrupt signal to gracefully shutdown the server with
		// a timeout of 5 seconds.
		// kill (no param) default send syscall.SIGTERM
		// kill -2 is syscall.SIGINT
		// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it

		//signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit
		log.Println("Shutdown Server ...")
		fmt.Println("Starting HTTP server...")
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := s.server.Shutdown(ctx); err != nil {
			log.Fatal("Server Shutdown:", err)
		}
		// catching ctx.Done(). timeout of 5 seconds.

		//d := <-ctx.Done()
		log.Println("Server exiting")
	}()

	s.waitGroup.Add(1)
	go func() {
		defer s.waitGroup.Done()
		if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("Error starting HTTP server: %v\n", err)
			log.Fatalf("listen error : %s\n", err)
		} else {
			fmt.Printf("HTTP server started successfully. Listening on %s\n", s.server.Addr)
		}
	}()
}

func ShutdownGracefully() (context.Context, *sync.WaitGroup, WaitFunc) {
	ctx, cancel := context.WithCancel(context.Background())
	waitGroup := new(sync.WaitGroup)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigChan
		log.Printf("Received signal: %s. Beginning  Shutdown ...\n", sig.String())
		cancel()
		sig = <-sigChan
		log.Printf("Received signal: %s. Shutting down immediately ...\n", sig.String())
	}()
	return ctx, waitGroup, func() { waitFunc(ctx, waitGroup) }
}
func waitFunc(ctx context.Context, wg *sync.WaitGroup) {
	<-ctx.Done()
	start := time.Now()
	waited := make(chan struct{})

	go func() {
		defer close(waited)
		wg.Wait()
	}()

	select {
	case <-waited:
		log.Printf("Waited for %v seconds before exiting waitFunc", time.Since(start).Seconds())
	case <-time.After(Timeout):
		log.Println("Timeout reached, exiting waitFunc")
	}
}
