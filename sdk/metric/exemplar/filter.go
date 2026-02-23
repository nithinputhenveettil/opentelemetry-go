// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package exemplar // import "github.com/nithinputhenveettil/opentelemetry-go/sdk/metric/exemplar"

import (
	"context"

	"github.com/nithinputhenveettil/opentelemetry-go/trace"
)

// Filter determines if a measurement should be offered.
//
// The passed ctx needs to contain any baggage or span that were active
// when the measurement was made. This information may be used by the
// Reservoir in making a sampling decision.
type Filter func(context.Context) bool

// TraceBasedFilter is a [Filter] that will only offer measurements
// if the passed context associated with the measurement contains a sampled
// [github.com/nithinputhenveettil/opentelemetry-go/trace.SpanContext].
func TraceBasedFilter(ctx context.Context) bool {
	return trace.SpanContextFromContext(ctx).IsSampled()
}

// AlwaysOnFilter is a [Filter] that always offers measurements.
func AlwaysOnFilter(context.Context) bool {
	return true
}

// AlwaysOffFilter is a [Filter] that never offers measurements.
func AlwaysOffFilter(context.Context) bool {
	return false
}
