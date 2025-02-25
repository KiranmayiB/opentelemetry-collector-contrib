// Code generated by mdatagen. DO NOT EDIT.

package metadatatest

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/otel/sdk/metric/metricdata"
	"go.opentelemetry.io/otel/sdk/metric/metricdata/metricdatatest"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/otelarrow/internal/metadata"
)

func TestSetupTelemetry(t *testing.T) {
	testTel := SetupTelemetry()
	tb, err := metadata.NewTelemetryBuilder(
		testTel.NewTelemetrySettings(),
		metadata.WithOtelarrowAdmissionInFlightBytesCallback(func() int64 { return 1 }),
		metadata.WithOtelarrowAdmissionWaitingBytesCallback(func() int64 { return 1 }),
	)
	require.NoError(t, err)
	require.NotNil(t, tb)

	testTel.AssertMetrics(t, []metricdata.Metrics{
		{
			Name:        "otelcol_otelarrow_admission_in_flight_bytes",
			Description: "Number of bytes that have started processing but are not finished.",
			Unit:        "By",
			Data: metricdata.Sum[int64]{
				Temporality: metricdata.CumulativeTemporality,
				IsMonotonic: false,
				DataPoints: []metricdata.DataPoint[int64]{
					{},
				},
			},
		},
		{
			Name:        "otelcol_otelarrow_admission_waiting_bytes",
			Description: "Number of items waiting to start processing.",
			Unit:        "By",
			Data: metricdata.Sum[int64]{
				Temporality: metricdata.CumulativeTemporality,
				IsMonotonic: false,
				DataPoints: []metricdata.DataPoint[int64]{
					{},
				},
			},
		},
	}, metricdatatest.IgnoreTimestamp(), metricdatatest.IgnoreValue())
	require.NoError(t, testTel.Shutdown(context.Background()))
}
