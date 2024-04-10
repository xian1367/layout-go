package main

import (
	"context"
	"github.com/xian1367/layout-go/config"
	"github.com/xian1367/layout-go/pkg/logger"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	sdkTrace "go.opentelemetry.io/otel/sdk/trace"
	"google.golang.org/grpc/credentials"
)

func initTracer() func(context.Context) error {
	secureOption := otlptracegrpc.WithTLSCredentials(credentials.NewClientTLSFromCert(nil, ""))
	secureOption = otlptracegrpc.WithInsecure()

	exporter, err := otlptrace.New(
		context.Background(),
		otlptracegrpc.NewClient(
			secureOption,
			otlptracegrpc.WithEndpoint(config.Get().Http[0].OpenTelemetryCollectorUrl),
		),
	)
	logger.FatalIf(err)

	resources, err := resource.New(
		context.Background(),
		resource.WithAttributes(
			attribute.String("service.name", config.Get().Http[0].Name),
			attribute.String("library.language", "go"),
		),
	)
	logger.FatalIf(err)

	otel.SetTracerProvider(
		sdkTrace.NewTracerProvider(
			sdkTrace.WithSampler(sdkTrace.AlwaysSample()),
			sdkTrace.WithBatcher(exporter),
			sdkTrace.WithResource(resources),
		),
	)
	return exporter.Shutdown
}
