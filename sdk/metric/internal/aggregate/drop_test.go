// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package aggregate

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/nithinputhenveettil/opentelemetry-go/attribute"
	"github.com/nithinputhenveettil/opentelemetry-go/sdk/metric/exemplar"
)

func TestDrop(t *testing.T) {
	t.Run("Int64", testDropFiltered[int64])
	t.Run("Float64", testDropFiltered[float64])
}

func testDropFiltered[N int64 | float64](t *testing.T) {
	r := dropReservoir[N](*attribute.EmptySet())

	var dest []exemplar.Exemplar
	r.Collect(&dest)

	assert.Empty(t, dest, "non-sampled context should not be offered")
}
