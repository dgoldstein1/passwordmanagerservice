package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"google.golang.org/grpc/stats"

	"github.com/spf13/viper"

	gometrics "github.com/armon/go-metrics"

	"github.com/deciphernow/gm-fabric-go/metrics/gmfabricsink"
	"github.com/deciphernow/gm-fabric-go/metrics/gometricsobserver"
	"github.com/deciphernow/gm-fabric-go/metrics/grpcmetrics"
	"github.com/deciphernow/gm-fabric-go/metrics/grpcobserver"
	ms "github.com/deciphernow/gm-fabric-go/metrics/metricsserver"
	"github.com/deciphernow/gm-fabric-go/metrics/subject"
)

func prepareInternalMetrics(ctx context.Context, mux *http.ServeMux) (stats.Handler, error) {
	var err error

	grpcObserver := grpcobserver.New(viper.GetInt("metrics_cache_size"))
	miscReporter := ms.NewMiscReporter()
	goMetObserver := gometricsobserver.New()
	observers := []subject.Observer{grpcObserver, goMetObserver}

	metricsChan := subject.New(ctx, observers...)

	sink := gmfabricsink.New(metricsChan)
	gmConfig := gometrics.DefaultConfig("passwordservice")
	if gmConfig.ServiceName == "" || gmConfig.HostName == "" {
		return nil, fmt.Errorf("Invalid gometrics.DefaultConfig %+v", gmConfig)
	}
	if _, err = gometrics.NewGlobal(gmConfig, sink); err != nil {
		return nil, fmt.Errorf("gometrics.NewGlobal failed: %v", err)
	}

	mux.Handle(
		viper.GetString("metrics_dashboard_uri_path"),
		ms.NewDashboardHandler(grpcObserver.Report, miscReporter.Report, goMetObserver.Report),
	)

	hostName, err := os.Hostname()
	if err != nil {
		return nil, fmt.Errorf("unable to determine host name: %s", err)
	}
	statsTags := []string{
		subject.JoinTag("service", "passwordservice"),
		subject.JoinTag("host", hostName),
	}

	return grpcmetrics.NewStatsHandlerWithTags(
		metricsChan,
		statsTags,
	), nil
}
