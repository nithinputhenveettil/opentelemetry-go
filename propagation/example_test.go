// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package propagation_test

import (
	"github.com/nithinputhenveettil/opentelemetry-go/propagation"
	"go.opentelemetry.io/otel"
)

func ExampleSetTextMapPropagator() {
	// Create a new composite text map propagator.
	propagator := propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	)

	// Set it as the global text map propagator.
	otel.SetTextMapPropagator(propagator)
}
