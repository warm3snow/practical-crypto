/**
 * @Author: xueyanghan
 * @File: main.go
 * @Version: 1.0.0
 * @Description: desc.
 * @Date: 2024/5/27 17:56
 */

package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const meterName = "github.com/open-telemetry/prometheus"

func main() {
	ctx := context.Background()

	// Start the prometheus HTTP server and pass the exporter Collector to it
	go serveMetrics()

	// This is the equivalent of prometheus.NewCounterVec
	counter(ctx)

	// This is the equivalent of prometheus.NewCounterVec
	counterObserver(ctx)

	// This is the equivalent of prometheus.NewCounterVec
	upDownCounter(ctx)

	// This is the equivalent of prometheus.NewGaugeVec
	gaugeObserver(ctx)

	// This is the equivalent of prometheus.NewHistogramVec
	histogram(ctx)

	ctx, _ = signal.NotifyContext(ctx, os.Interrupt)
	<-ctx.Done()
}

func serveMetrics() {
	log.Printf("serving metrics at localhost:2223/metrics")
	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(":2223", nil) //nolint:gosec // Ignoring G114: Use of net/http serve function that has no support for setting timeouts.
	if err != nil {
		fmt.Printf("error serving http: %v", err)
		return
	}
}
