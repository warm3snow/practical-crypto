/**
 * @Author: xueyanghan
 * @File: metrics.go
 * @Version: 1.0.0
 * @Description: desc.
 * @Date: 2024/5/27 17:15
 */

package main

import (
	"context"
	"log"
	"math/rand"
	"runtime"
	"time"

	api "go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/sdk/metric"

	"go.opentelemetry.io/otel/exporters/prometheus"
)

var Meter api.Meter

func init() {

	// The exporter embeds a default OpenTelemetry Reader and
	// implements prometheus.Collector, allowing it to be used as
	// both a Reader and Collector.
	exporter, err := prometheus.New()
	if err != nil {
		log.Fatal(err)
	}
	provider := metric.NewMeterProvider(metric.WithReader(exporter))
	Meter = provider.Meter(meterName)
}

// counter demonstrates how to measure non-decreasing numbers, for example,
// number of requests or connections.
func counter(ctx context.Context) {
	counter, _ := Meter.Int64Counter(
		"some.prefix.counter",
		api.WithUnit("1"),
		api.WithDescription("TODO"),
	)

	for {
		counter.Add(ctx, 1)
		time.Sleep(time.Millisecond)
	}
}

// upDownCounter demonstrates how to measure numbers that can go up and down, for example,
// number of goroutines or customers.
func upDownCounter(ctx context.Context) {
	counter, _ := Meter.Int64UpDownCounter(
		"some.prefix.up_down_counter",
		api.WithUnit("1"),
		api.WithDescription("TODO"),
	)
	for {
		if rand.Float64() >= 0.5 {
			counter.Add(ctx, +1)
		} else {
			counter.Add(ctx, -1)
		}

		time.Sleep(time.Second)
	}
}

// histogram demonstrates how to record a distribution of individual values, for example,
// request or query timings. With this instrument you get total number of records,
// avg/min/max values, and heatmaps/percentiles.
func histogram(ctx context.Context) {
	durRecorder, _ := Meter.Int64Histogram(
		"some.prefix.histogram",
		api.WithUnit("microseconds"),
		api.WithDescription("TODO"),
	)

	for {
		dur := time.Duration(rand.NormFloat64()*5000000) * time.Microsecond
		durRecorder.Record(ctx, dur.Microseconds())

		time.Sleep(time.Millisecond)
	}
}

// counterObserver demonstrates how to measure monotonic (non-decreasing) numbers,
// for example, number of requests or connections.
func counterObserver(ctx context.Context) {
	counter, _ := Meter.Int64ObservableCounter(
		"some.prefix.counter_observer",
		api.WithUnit("1"),
		api.WithDescription("TODO"),
	)

	var number int64
	if _, err := Meter.RegisterCallback(
		// SDK periodically calls this function to collect data.
		func(ctx context.Context, o api.Observer) error {
			number++
			o.ObserveInt64(counter, number)
			return nil
		},
	); err != nil {
		panic(err)
	}
}

// upDownCounterObserver demonstrates how to measure numbers that can go up and down,
// for example, number of goroutines or customers.
func upDownCounterObserver(ctx context.Context) {
	counter, err := Meter.Int64ObservableUpDownCounter(
		"some.prefix.up_down_counter_async",
		api.WithUnit("1"),
		api.WithDescription("TODO"),
	)
	if err != nil {
		panic(err)
	}

	if _, err := Meter.RegisterCallback(
		func(ctx context.Context, o api.Observer) error {
			num := runtime.NumGoroutine()
			o.ObserveInt64(counter, int64(num))
			return nil
		},
		counter,
	); err != nil {
		panic(err)
	}
}

// gaugeObserver demonstrates how to measure non-additive numbers that can go up and down,
// for example, cache hit rate or memory utilization.
func gaugeObserver(ctx context.Context) {
	gauge, _ := Meter.Float64ObservableGauge(
		"some.prefix.gauge_observer",
		api.WithUnit("1"),
		api.WithDescription("TODO"),
	)

	if _, err := Meter.RegisterCallback(
		func(ctx context.Context, o api.Observer) error {
			o.ObserveFloat64(gauge, rand.Float64())
			return nil
		},
		gauge,
	); err != nil {
		panic(err)
	}
}
