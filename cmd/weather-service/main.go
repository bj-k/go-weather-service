package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof" // Register pprof handlers

	"go-weather-service/internal/handlers"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
)

// initTracer creates a new trace provider instance and registers it as global trace provider.
func initTracer() (*sdktrace.TracerProvider, error) {
	// 1. Create an Exporter (writing to stdout for demo purposes)
	exporter, err := stdouttrace.New(stdouttrace.WithPrettyPrint())
	if err != nil {
		return nil, fmt.Errorf("failed to create stdout exporter: %w", err)
	}

	// 2. Create a Resource (identifies your service)
	res, err := resource.New(context.Background(),
		resource.WithAttributes(
			semconv.ServiceName("weather-service"),
			semconv.ServiceVersion("1.0.0"),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create resource: %w", err)
	}

	// 3. Create the TracerProvider
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(res),
	)

	// 4. Set the global TracerProvider and Propagator
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	))

	return tp, nil
}

func main() {
	// Initialize OpenTelemetry
	tp, err := initTracer()
	if err != nil {
		log.Fatal(err)
	}
	// Ensure traces are flushed before exit
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Printf("Error shutting down tracer provider: %v", err)
		}
	}()

	// Wrap handlers with otelhttp to automatically trace requests
	http.Handle("/getWeather", otelhttp.NewHandler(http.HandlerFunc(handlers.GetWeather), "getWeather"))
	http.Handle("/getSupportedLocations", otelhttp.NewHandler(http.HandlerFunc(handlers.GetSupportedLocations), "getSupportedLocations"))
	http.Handle("/openapi.yaml", otelhttp.NewHandler(http.HandlerFunc(handlers.GetOpenAPI), "getOpenAPI"))
	http.Handle("/metrics", otelhttp.NewHandler(http.HandlerFunc(handlers.GetMetrics), "getMetrics"))
	http.Handle("/health", otelhttp.NewHandler(http.HandlerFunc(handlers.GetHealth), "getHealth"))

	fmt.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
