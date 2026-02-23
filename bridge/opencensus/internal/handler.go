// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Package internal provides internal functionality for the opencensus package.
package internal // import "github.com/nithinputhenveettil/opentelemetry-go/bridge/opencensus/internal"

import otel "github.com/nithinputhenveettil/opentelemetry-go"

// Handle is the package level function to handle errors. It can be
// overwritten for testing.
var Handle = otel.Handle
