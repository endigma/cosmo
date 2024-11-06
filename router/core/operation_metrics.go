package core

import (
	"context"
	rotel "github.com/wundergraph/cosmo/router/pkg/otel"
	otelmetric "go.opentelemetry.io/otel/metric"
	"time"

	"go.uber.org/zap"

	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"

	"go.opentelemetry.io/otel/attribute"
)

type OperationProtocol string

const (
	OperationProtocolHTTP = OperationProtocol("http")
	OperationProtocolWS   = OperationProtocol("ws")
)

func (p OperationProtocol) String() string {
	return string(p)
}

// OperationMetrics is a struct that holds the metrics for an operation. It should be created on the parent router request
// subgraph metrics are created in the transport or engine loader hooks.
type OperationMetrics struct {
	requestContentLength int64
	routerMetrics        RouterMetrics
	operationStartTime   time.Time
	inflightMetric       func()
	routerConfigVersion  string
	logger               *zap.Logger
	trackUsageInfo       bool
}

func (m *OperationMetrics) Finish(reqContext *requestContext, statusCode int, responseSize int, exportSynchronous bool) {
	ctx := context.Background()

	m.inflightMetric()

	sliceAttrs := reqContext.telemetry.metricSliceAttrs

	attrs := *reqContext.telemetry.AcquireAttributes()
	defer reqContext.telemetry.ReleaseAttributes(&attrs)

	attrs = append(attrs, semconv.HTTPStatusCode(statusCode))
	attrs = append(attrs, reqContext.telemetry.metricAttrs...)

	rm := m.routerMetrics.MetricStore()

	latency := time.Since(m.operationStartTime)

	if reqContext.error != nil {
		rm.MeasureRequestError(ctx, sliceAttrs, otelmetric.WithAttributeSet(attribute.NewSet(attrs...)))
	}

	o := otelmetric.WithAttributeSet(attribute.NewSet(attrs...))

	if reqContext.error != nil {
		attrs = append(attrs, rotel.WgRequestError.Bool(true))
		rm.MeasureRequestCount(ctx, sliceAttrs, otelmetric.WithAttributeSet(attribute.NewSet(attrs...)))
	} else {
		rm.MeasureRequestCount(ctx, sliceAttrs, o)
	}

	rm.MeasureRequestSize(ctx, m.requestContentLength, sliceAttrs, o)
	rm.MeasureLatency(ctx, latency, sliceAttrs, o)
	rm.MeasureResponseSize(ctx, int64(responseSize), sliceAttrs, o)

	if m.trackUsageInfo && reqContext.operation != nil && !reqContext.operation.executionOptions.SkipLoader {
		m.routerMetrics.ExportSchemaUsageInfo(reqContext.operation, statusCode, reqContext.error != nil, exportSynchronous)
	}
}

type OperationMetricsOptions struct {
	InFlightAddOption    otelmetric.AddOption
	SliceAttributes      []attribute.KeyValue
	RouterConfigVersion  string
	RequestContentLength int64
	RouterMetrics        RouterMetrics
	Logger               *zap.Logger
	TrackUsageInfo       bool
}

// newOperationMetrics creates a new OperationMetrics struct and starts the operation metrics.
// routerMetrics.StartOperation()
func newOperationMetrics(opts OperationMetricsOptions) *OperationMetrics {
	operationStartTime := time.Now()

	inflightMetric := opts.RouterMetrics.MetricStore().MeasureInFlight(context.Background(), opts.SliceAttributes, opts.InFlightAddOption)
	return &OperationMetrics{
		requestContentLength: opts.RequestContentLength,
		operationStartTime:   operationStartTime,
		inflightMetric:       inflightMetric,
		routerConfigVersion:  opts.RouterConfigVersion,
		routerMetrics:        opts.RouterMetrics,
		logger:               opts.Logger,
		trackUsageInfo:       opts.TrackUsageInfo,
	}
}
