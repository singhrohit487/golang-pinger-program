package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/spf13/viper"
)

const DefaultPingTimeout = 10 * time.Second
const DefaultTimeout = 5 * time.Second
const ExitCodeSuccess = 0
const ExitCodeInitFailed = 1
const ExitCodeMainFailed = 2

var config Config
var errorLogger = log.New(os.Stdout, "  error|", log.LstdFlags)
var serverLogger = log.New(os.Stdout, " server|", log.LstdFlags)
var serviceLogger = log.New(os.Stdout, "service|", log.LstdFlags)
var readiness = map[string]bool{
	"target_up": false,
}

func init() {
	defer func() {
		if r := recover(); r != nil {
			log.Fatal(r)
			os.Exit(ExitCodeInitFailed)
		}
	}()

	// initialise the configuration
	c := viper.New()
	c.SetDefault("interface", DefaultInterface)
	c.SetDefault("ping_timeout", DefaultPingTimeout)
	c.SetDefault("port", DefaultPort)
	c.SetDefault("target_proto", DefaultTargetProto)
	c.SetDefault("target_host", DefaultTargetHost)
	c.SetDefault("target_port", DefaultTargetPort)
	c.SetDefault("target_path", DefaultTargetPath)

	// draws in the configuration from the environment
	c.AutomaticEnv()

	config = Config{
		// from the INTERFACE environment variable
		Interface: c.GetString("interface"),
		// from the PORT environment variable
		Port: uint16(c.GetUint("port")),
		// from the TARGET_PROTO environment variable
		TargetProto: c.GetString("target_proto"),
		// from the TARGET_HOST environment variable
		TargetHost: c.GetString("target_host"),
		// from the TARGET_PORT environment variable
		TargetPort: uint16(c.GetUint("target_port")),
		// from the TARGET_PATH environment variable
		TargetPath: c.GetString("target_path"),
	}
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Fatal(r)
			os.Exit(ExitCodeMainFailed)
		}
	}()

	// setup
	serviceLogger.Printf("initialising service...\n")
	addr := fmt.Sprintf("%s:%v", config.Interface, config.Port)
	mux := createMux()
	server := createServer(addr, mux)
	tick := time.Tick(time.Second)
	done := make(chan bool, 1)
	ossig := make(chan os.Signal, 1)
	signal.Notify(ossig, syscall.SIGTERM, syscall.SIGINT)

	// listen
	serviceLogger.Printf("attempting to listen on '%s'...\n", server.Addr)
	go server.ListenAndServe()
	for {
		select {
		// pinger schedule
		case <-tick:
			client := http.Client{
				Timeout: DefaultPingTimeout,
			}
			req, err := http.NewRequest("GET", config.getTargetURL(), nil)
			handleError(err, errorLogger)
			response, err := client.Do(req)
			handleError(err, errorLogger)
			serviceLogger.Printf("> %s -> '%s'\n", config.getTargetURL(), response.Status)
		// handle os termination
		case sig := <-ossig:
			serviceLogger.Printf("received termination signal '%v', shutting down server now\n", sig)
			done <- true
		// handle graceful shutdown
		case <-done:
			server.Close()
			serviceLogger.Println("exiting now...")
			os.Exit(ExitCodeSuccess)
		}
	}
}

// createServer returns a http.Server
func createServer(addr string, handler http.Handler) http.Server {
	return http.Server{
		Addr:              addr,
		Handler:           requestLoggerMiddleware(handler),
		ReadTimeout:       DefaultTimeout,
		ReadHeaderTimeout: DefaultTimeout,
		WriteTimeout:      DefaultTimeout,
		ErrorLog:          errorLogger,
	}
}

// handleError is the generic error handler
func handleError(err error, log *log.Logger) {
	if err != nil {
		log.Println(err)
	}
}

// requestLoggerMiddleware is a middleware that logs incoming requests
func requestLoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
		serverLogger.Printf("< %s <- %s | %s %s %s \n", r.Host, r.RemoteAddr, strings.ToUpper(r.Proto), r.Method, r.URL.Path)
	})
}
