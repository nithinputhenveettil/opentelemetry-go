// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package tracetransform // import "github.com/nithinputhenveettil/opentelemetry-go/exporters/otlp/otlptrace/internal/tracetransform"

import (
	commonpb "go.opentelemetry.io/proto/otlp/common/v1"

	"github.com/nithinputhenveettil/opentelemetry-go/sdk/instrumentation"
)

func InstrumentationScope(il instrumentation.Scope) *commonpb.InstrumentationScope {
	if il == (instrumentation.Scope{}) {
		return nil
	}
	return &commonpb.InstrumentationScope{
		Name:       il.Name,
		Version:    il.Version,
		Attributes: Iterator(il.Attributes.Iter()),
	}
}
