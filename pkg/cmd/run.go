package cmd

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/bnassif/jquery-proxy/pkg/cmdutil"
	"github.com/bnassif/jquery-proxy/pkg/server"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var defaultDir = "/opt/jquery-proxy/"
var defaultFile = ".jquery-proxy.env"

var runOpts *viper.Viper

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run the server",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// Build config from flags & defaults
		config := cmdutil.BuildConfig(runOpts)
		// Create server object
		svr := server.NewServer(config)

		// Create signal channel for clean shutdown
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
		go func() {
			sig := <-sigCh

			svr.Logger.Info(
				"shutdown signal received",
				slog.String("signal", sig.String()),
			)

			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			if err := svr.Shutdown(ctx); err != nil {
				svr.Logger.Error(
					"graceful shutdown failed",
					slog.Any("error", err),
				)
			}
		}()

		// Run the server
		err := svr.Run()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			svr.Logger.Error(
				"server exited unexpectedly",
				slog.Any("error", err),
			)
		}

		svr.Logger.Info("server shutdown")
		fmt.Println("Server shutdown")
		return
	},
}

func init() {
	runOpts = viper.New()

	// Flag Definitions
	// Server Flags
	runCmd.PersistentFlags().StringP("address", "a", "0.0.0.0", "The address the server listens on")
	runCmd.PersistentFlags().StringP("port", "p", "8080", "The port the server listens on")
	runCmd.PersistentFlags().String("read-timeout", "5s", "The read timeout for the server")
	runCmd.PersistentFlags().String("write-timeout", "10s", "The write timeout for the server")
	runCmd.PersistentFlags().String("idle-timeout", "15s", "The idle timeout for the server")
	// Client Flags
	runCmd.PersistentFlags().BoolP("persistent-cookies", "C", false, "Whether to use persistent cookies")
	runCmd.PersistentFlags().String("client-timeout", "10s", "The timeout for proxied requests")
	// Cache Flags
	runCmd.PersistentFlags().Bool("redis-enable", false, "Whether to enable redis caching")
	runCmd.PersistentFlags().String("redis-url", "", "The url of the Redis instance. This should include authentication details.")
	runCmd.PersistentFlags().String("redis-prefix", "jquery", "The prefix to assign keys in Redis")
	runCmd.PersistentFlags().String("redis-ttl", "30m", "The default TTL for cached responses")
	runCmd.PersistentFlags().String("redis-connect-timeout", "5s", "The connection timeout for redis")
	runCmd.PersistentFlags().String("redis-read-timeout", "5s", "The read timeout for redis")
	runCmd.PersistentFlags().String("redis-write-timeout", "5s", "The write timeout for redis")
	runCmd.PersistentFlags().Int("redis-pool-size", 8, "The number of connection pools to reserve")
	runCmd.PersistentFlags().Int("redis-min-idle-connections", 0, "The number of minimum idle connections to maintain")
	// Proxy Flags
	runCmd.PersistentFlags().Bool("proxy-enable", false, "Whether to enable outbound proxying")
	runCmd.PersistentFlags().String("proxy-url", "", "The URL of the outbound proxy to use")
	// Logging Flags
	runCmd.PersistentFlags().String("log-level", "info", "The logging level to choose. One of: 'error', 'warning', 'debug'")
	runCmd.PersistentFlags().String("log-format", "json", "The output format to use. One of: 'text', 'json'")
	runCmd.PersistentFlags().Bool("log-add-source", false, "Output the source of the logged entry")

	// Persistent Flag Bindings to Cobra
	_ = runOpts.BindPFlag("address", runCmd.PersistentFlags().Lookup("address"))
	_ = runOpts.BindPFlag("port", runCmd.PersistentFlags().Lookup("port"))
	_ = runOpts.BindPFlag("read-timeout", runCmd.PersistentFlags().Lookup("read-timeout"))
	_ = runOpts.BindPFlag("write-timeout", runCmd.PersistentFlags().Lookup("write-timeout"))
	_ = runOpts.BindPFlag("idle-timeout", runCmd.PersistentFlags().Lookup("idle-timeout"))
	_ = runOpts.BindPFlag("persistent-cookies", runCmd.PersistentFlags().Lookup("persistent-cookies"))
	_ = runOpts.BindPFlag("client-timeout", runCmd.PersistentFlags().Lookup("client-timeout"))
	_ = runOpts.BindPFlag("redis-enable", runCmd.PersistentFlags().Lookup("redis-enable"))
	_ = runOpts.BindPFlag("redis-url", runCmd.PersistentFlags().Lookup("redis-url"))
	_ = runOpts.BindPFlag("redis-prefix", runCmd.PersistentFlags().Lookup("redis-prefix"))
	_ = runOpts.BindPFlag("redis-ttl", runCmd.PersistentFlags().Lookup("redis-ttl"))
	_ = runOpts.BindPFlag("redis-connect-timeout", runCmd.PersistentFlags().Lookup("redis-connect-timeout"))
	_ = runOpts.BindPFlag("redis-read-timeout", runCmd.PersistentFlags().Lookup("redis-read-timeout"))
	_ = runOpts.BindPFlag("redis-write-timeout", runCmd.PersistentFlags().Lookup("redis-write-timeout"))
	_ = runOpts.BindPFlag("redis-pool-size", runCmd.PersistentFlags().Lookup("redis-pool-size"))
	_ = runOpts.BindPFlag("redis-min-idle-connections", runCmd.PersistentFlags().Lookup("redis-min-idle-connections"))
	_ = runOpts.BindPFlag("proxy-enable", runCmd.PersistentFlags().Lookup("proxy-enable"))
	_ = runOpts.BindPFlag("proxy-url", runCmd.PersistentFlags().Lookup("proxy-url"))
	_ = runOpts.BindPFlag("log-level", runCmd.PersistentFlags().Lookup("log-level"))
	_ = runOpts.BindPFlag("log-format", runCmd.PersistentFlags().Lookup("log-format"))
	_ = runOpts.BindPFlag("log-add-source", runCmd.PersistentFlags().Lookup("log-add-source"))

	runOpts.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	runOpts.AutomaticEnv()
}
