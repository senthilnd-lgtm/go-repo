package telemetry

// import (
// 	"context"

// 	"go.opentelemetry.io/otel"
// 	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
// 	"go.opentelemetry.io/otel/sdk/resource"
// 	sdktrace "go.opentelemetry.io/otel/sdk/trace"
// 	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
// )

// func InitTracer() func(context.Context) error {
// 	exporter, _ := stdouttrace.New(stdouttrace.WithPrettyPrint())
// 	tp := sdktrace.NewTracerProvider(
// 		sdktrace.WithBatcher(exporter),
// 		sdktrace.WithResource(resource.NewWithAttributes(
// 			semconv.SchemaURL,
// 			semconv.ServiceName("pet-api"),
// 		)),
// 	)
// 	otel.SetTracerProvider(tp)
// 	return tp.Shutdown
// }

import (
	"context"
	"log"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
)

func InitTracer() func(context.Context) error {
	exporter, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint("http://localhost:14268/api/traces")))
	if err != nil {
		log.Fatal("Jaeger exporter error:", err)
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName("pet-api"),
		)),
	)
	otel.SetTracerProvider(tp)
	return tp.Shutdown
}
