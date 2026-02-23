// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Package internal provides internal functionality for the otlploghttp
// package.
package internal // import "github.com/nithinputhenveettil/opentelemetry-go/exporters/otlp/otlplog/otlploghttp/internal"

//go:generate  gotmpl --body=../../../../../internal/shared/x/x.go.tmpl "--data={ \"pkg\": \"github.com/nithinputhenveettil/opentelemetry-go/exporters/otlp/otlplog/otlploghttp\" }"  --out=x/x.go
//go:generate gotmpl --body=../../../../../internal/shared/x/x_test.go.tmpl "--data={}" --out=x/x_test.go

//go:generate gotmpl --body=../../../../../internal/shared/otlp/retry/retry.go.tmpl "--data={}" --out=retry/retry.go
//go:generate gotmpl --body=../../../../../internal/shared/otlp/retry/retry_test.go.tmpl "--data={}" --out=retry/retry_test.go

//go:generate gotmpl --body=../../../../../internal/shared/otlp/otlplog/transform/attr_test.go.tmpl "--data={}" --out=transform/attr_test.go
//go:generate gotmpl --body=../../../../../internal/shared/otlp/otlplog/transform/log.go.tmpl "--data={}" --out=transform/log.go
//go:generate gotmpl --body=../../../../../internal/shared/otlp/otlplog/transform/log_attr_test.go.tmpl "--data={}" --out=transform/log_attr_test.go
//go:generate gotmpl --body=../../../../../internal/shared/otlp/otlplog/transform/log_test.go.tmpl "--data={}" --out=transform/log_test.go
