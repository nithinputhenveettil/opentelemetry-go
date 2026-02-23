// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build linux

package resource // import "github.com/nithinputhenveettil/opentelemetry-go/sdk/resource"

var platformHostIDReader hostIDReader = &hostIDReaderLinux{
	readFile: readFile,
}
