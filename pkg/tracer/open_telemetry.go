package tracer

import (
	"context"
	"github.com/xian1367/layout-go/config"
	"github.com/xian1367/layout-go/pkg/logger"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/semconv/v1.24.0"
	"time"
)

var TP *trace.TracerProvider

func InitTracer() {
	exp, _ := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(config.Get().Http[0].TelemetryEndpoint)))

	// 创建资源
	res, err := resource.New(context.Background(),
		resource.WithAttributes(
			semconv.ServiceNameKey.String(config.Get().Http[0].Name),
		),
	)
	logger.ErrorIf(err)

	// 创建 TracerProvider
	TP = trace.NewTracerProvider(
		trace.WithSampler(trace.AlwaysSample()),
		trace.WithBatcher(exp),
		trace.WithResource(res),
	)

	// 设置全局的 TracerProvider 和 Propagators
	otel.SetTracerProvider(TP)
	otel.SetTextMapPropagator(propagation.TraceContext{})
}

func Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	err := TP.Shutdown(ctx)
	logger.ErrorIf(err)
}
