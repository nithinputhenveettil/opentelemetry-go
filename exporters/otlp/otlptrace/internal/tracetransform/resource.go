// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package tracetransform // import "github.com/nithinputhenveettil/opentelemetry-go/exporters/otlp/otlptrace/internal/tracetransform"

import (
	resourcepb "go.opentelemetry.io/proto/otlp/resource/v1"

	"github.com/nithinputhenveettil/opentelemetry-go/sdk/resource"
)

// Resource transforms a Resource into an OTLP Resource.
func Resource(r *resource.Resource) *resourcepb.Resource {
	if r == nil {
		return nil
	}
	return &resourcepb.Resource{Attributes: ResourceAttributes(r)}
}
