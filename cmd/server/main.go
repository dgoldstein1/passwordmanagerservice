package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"google.golang.org/grpc"
	"google.golang.org/grpc/stats"

	"github.com/deciphernow/gm-fabric-go/metrics/grpcmetrics"
	ms "github.com/deciphernow/gm-fabric-go/metrics/metricsserver"
	pm "github.com/deciphernow/gm-fabric-go/metrics/prometheus"

	"github.com/dgoldstein1/passwordservice/cmd/server/config"
	"github.com/dgoldstein1/passwordservice/cmd/server/methods"

	// we don't use this directly, but need it in vendor for gateway grpc plugin
	_ "github.com/golang/glog"
	_ "github.com/grpc-ecosystem/grpc-gateway/runtime"
)

func main() {
	var tlsMetricsConf *tls.Config
	var tlsServerConf *tls.Config
	var err error
	var zkCancels []zkCancelFunc

	logger := zerolog.New(os.Stderr).With().Timestamp().Logger().
		Output(zerolog.ConsoleWriter{Out: os.Stderr})

	logger.Info().Str("service", "passwordservice").Msg("starting")

	ctx, cancelFunc := context.WithCancel(context.Background())

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	defer func() {
		for _, f := range zkCancels {
			f()
		}
	}()
	
	logger.Debug().Str("service", "passwordservice").Msg("initializing config")
	if err = config.Initialize(); err != nil {
		logger.Fatal().AnErr("config.Initialize()", err).Msg("")
	}

	if tlsMetricsConf, err = buildMetricsTLSConfigIfNeeded(logger); err != nil {
		logger.Fatal().AnErr("buildMetricsTLSConfigIfNeeded", err).Msg("")
	}

	if tlsServerConf, err = buildServerTLSConfigIfNeeded(logger); err != nil {
		logger.Fatal().AnErr("buildServerTLSConfigIfNeeded", err).Msg("")
	}

	logger.Debug().Str("service", "passwordservice").
		Str("host", viper.GetString("grpc_server_host")).
		Int("port", viper.GetInt("grpc_server_port")).
		Msg("creating listener")

	lis, err := net.Listen(
		"tcp",
		fmt.Sprintf(
			"%s:%d",
			viper.GetString("grpc_server_host"),
			viper.GetInt("grpc_server_port"),
		),
	)
	if err != nil {
		logger.Fatal().AnErr("net.Listen", err).Msg("")
	}

	ctx = putOauthInCtxIfNeeded(ctx)
	mux := http.NewServeMux()
	var statsHandler stats.Handler
	var opts []grpc.ServerOption

	logger.Debug().Str("metrics", "Internal").Msg("preparing metrics")
	if statsHandler, err = prepareInternalMetrics(ctx, mux); err != nil {
		logger.Fatal().AnErr("preparePrometheusMetrics", err).Msg("")
	}
	if viper.GetBool("report_prometheus") {
		logger.Debug().Str("metrics", "Prometheus").Msg("preparing metrics")
		pmStatsHandler, err := pm.NewStatsHandler()
		if err != nil {
			logger.Fatal().AnErr("pm.NewStatsHandler", err).Msg("")
		}
		statsHandler = grpcmetrics.NewFanoutHandler(statsHandler, pmStatsHandler)

		mux.Handle(
			viper.GetString("metrics_prometheus_uri_path"),
			promhttp.Handler(),
		)
	}
	opts = append(opts, grpc.StatsHandler(statsHandler))
	
	logger.Debug().Str("service", "passwordservice").
		Str("host", viper.GetString("metrics_server_host")).
		Int("port", viper.GetInt("metrics_server_port")).
		Msg("starting metrics server")

	mServer := ms.NewMetricsServer(
		fmt.Sprintf(
			"%s:%d",
			viper.GetString("metrics_server_host"),
			viper.GetInt("metrics_server_port"),
		),		
		tlsMetricsConf,
	)
	mServer.Handler = mux
	
	if mServer.TLSConfig == nil {
		go mServer.ListenAndServe()
	} else {
		go mServer.ListenAndServeTLS("", "")
	}
	
	cancels, err := notifyZkOfMetricsIfNeeded(logger)
	if err != nil {
		logger.Fatal().Err(err).Msg("zk metrics announcement")
	}

	zkCancels = append(
		zkCancels,
		cancels...,
	)

	opts = append(opts, getTLSOptsIfNeeded(tlsServerConf)...)

	oauthOpts, err := getOauthOptsIfNeeded(logger)
	if err != nil {
		logger.Fatal().AnErr("getOauthOptsIfNeeded", err).Msg("")
	}
	opts = append(opts, oauthOpts...)

	grpcServer := grpc.NewServer(opts...)
	methods.CreateAndRegisterServer(logger, grpcServer)

	logger.Debug().Str("service", "passwordservice").
		Msg("starting grpc server")
	go grpcServer.Serve(lis)

	cancels, err = notifyZkOfRPCServerIfNeeded(logger)
	if err != nil {
		logger.Fatal().Err(err).Msg("")
	}

	zkCancels = append(
		zkCancels,
		cancels...,
	)

	if viper.GetBool("use_gateway_proxy") {
		logger.Debug().Str("service", "passwordservice").
			Msg("starting gateway proxy")
		if err = startGatewayProxy(ctx, logger); err != nil {
			logger.Fatal().AnErr("startGatewayProxy", err).Msg("")
		}
	}

	cancels, err = notifyZkOfGatewayEndpointIfNeeded(logger)
	if err != nil {
		logger.Fatal().Err(err).Msg("")
	}

	zkCancels = append(
		zkCancels,
		cancels...,
	)

	s := <- sigChan
	logger.Info().Str("service", "passwordservice") .
		Str("signal", s.String()).
		Msg("shutting down")
	cancelFunc()
	grpcServer.Stop()
}
