// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package propagation_test

import (
	"github.com/nithinputhenveettil/opentelemetry-go/propagation"
	"go.opentelemetry.io/otel"
)

func ExampleTraceContext() {
	tc := propagation.TraceContext{}
	// Register the TraceContext propagator globally.
	otel.SetTextMapPropagator(tc)
}
