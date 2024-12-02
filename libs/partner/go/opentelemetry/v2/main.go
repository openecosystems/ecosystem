package opentelemetryv2

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploggrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	logger "go.opentelemetry.io/otel/log"
	"go.opentelemetry.io/otel/log/global"
	meter "go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/log"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/trace"
	tracer "go.opentelemetry.io/otel/trace"
	"libs/public/go/sdk/v2alpha"
)

// Binding struct that holds binding specific fields
type Binding struct {
	Propagator     *propagation.TextMapPropagator
	TraceProvider  *trace.TracerProvider
	Tracer         *tracer.Tracer
	MeterProvider  *metric.MeterProvider
	Meter          *meter.Meter
	LoggerProvider *log.LoggerProvider
	Logger         *logger.Logger

	configuration *Configuration
}

var (
	Bound       *Binding
	BindingName = "OPEN_TELEMETRY_BINDING"
)

func (b *Binding) Name() string {
	return BindingName
}

func (b *Binding) Validate(_ context.Context, _ *sdkv2alphalib.Bindings) error {
	// Verify any requirements

	return nil
}

func (b *Binding) Bind(ctx context.Context, bindings *sdkv2alphalib.Bindings) *sdkv2alphalib.Bindings {
	if Bound == nil {
		var once sync.Once
		once.Do(
			func() {
				Bound = &Binding{}

				// Set up propagator
				propagator := newPropagator()
				otel.SetTextMapPropagator(propagator)
				Bound.Propagator = &propagator

				if ResolvedConfiguration.Opentelemetry.TraceProviderEnabled {
					// Set up trace provider
					tracerProvider, err := newTraceProvider(ctx)
					if err != nil {
						panic(err)
					}
					otel.SetTracerProvider(tracerProvider)
					t := tracerProvider.Tracer("platform-server-go")
					Bound.TraceProvider = tracerProvider
					Bound.Tracer = &t
				}

				if ResolvedConfiguration.Opentelemetry.MeterProviderEnabled {
					// Set up meter provider
					meterProvider, err := newMeterProvider(ctx)
					if err != nil {
						panic(err)
					}
					otel.SetMeterProvider(meterProvider)
					m := meterProvider.Meter("platform-server-go")
					Bound.MeterProvider = meterProvider
					Bound.Meter = &m
				}

				if ResolvedConfiguration.Opentelemetry.LoggerProviderEnabled {
					// Set up logger provider.
					loggerProvider, err := newLoggerProvider(ctx)
					if err != nil {
						panic(err)
					}
					global.SetLoggerProvider(loggerProvider)
					l := loggerProvider.Logger("platform-server-go")
					Bound.LoggerProvider = loggerProvider
					Bound.Logger = &l
				}

				bindings.Registered[b.Name()] = Bound
			})
	} else {
		bindings.Registered[b.Name()] = Bound
		fmt.Println("Open Telemetry already bound")
	}

	return bindings
}

func (b *Binding) GetBinding() interface{} {
	return Bound
}

func (b *Binding) Close() error {
	var err error

	if ResolvedConfiguration.Opentelemetry.TraceProviderEnabled {
		t := b.TraceProvider.Shutdown(context.Background())
		if t != nil {
			err = errors.Join(err, t)
		}
		fmt.Println("Closing the Open telemetry TraceProvider Binding")
	}

	if ResolvedConfiguration.Opentelemetry.MeterProviderEnabled {
		m := b.MeterProvider.Shutdown(context.Background())
		if m != nil {
			err = errors.Join(err, m)
		}
		fmt.Println("Closing the Open telemetry MeterProvider Binding")
	}

	if ResolvedConfiguration.Opentelemetry.LoggerProviderEnabled {
		l := b.LoggerProvider.Shutdown(context.Background())
		if l != nil {
			err = errors.Join(err, l)
		}
		fmt.Println("Closing the Open telemetry LoggerProvider Binding")
	}

	return nil
}

func newPropagator() propagation.TextMapPropagator {
	return propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	)
}

func newTraceProvider(ctx context.Context) (*trace.TracerProvider, error) {
	traceExporter, err := otlptracegrpc.New(ctx, otlptracegrpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return trace.NewTracerProvider(
		trace.WithBatcher(traceExporter,
			trace.WithBatchTimeout(5*time.Second)),
	), nil
}

func newMeterProvider(ctx context.Context) (*metric.MeterProvider, error) {
	metricExporter, err := otlpmetricgrpc.New(ctx, otlpmetricgrpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return metric.NewMeterProvider(
		metric.WithReader(metric.NewPeriodicReader(metricExporter,
			metric.WithInterval(60*time.Second))),
	), nil
}

func newLoggerProvider(ctx context.Context) (*log.LoggerProvider, error) {
	logExporter, err := otlploggrpc.New(ctx, otlploggrpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return log.NewLoggerProvider(
		log.WithProcessor(log.NewBatchProcessor(logExporter)),
	), nil
}
