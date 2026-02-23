module github.com/nithinputhenveettil/opentelemetry-go/sdk

go 1.24.0

replace github.com/nithinputhenveettil/opentelemetry-go => ../

require (
	github.com/go-logr/logr v1.4.3
	github.com/google/go-cmp v0.7.0
	github.com/google/uuid v1.6.0
	github.com/stretchr/testify v1.11.1
	go.opentelemetry.io/otel v1.40.0
	go.opentelemetry.io/otel/metric v1.40.0
	go.opentelemetry.io/otel/sdk/metric v1.40.0
	go.opentelemetry.io/otel/trace v1.40.0
	go.uber.org/goleak v1.3.0
	golang.org/x/sys v0.41.0
)

require (
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	go.opentelemetry.io/auto/sdk v1.2.1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/nithinputhenveettil/opentelemetry-go/trace => ../trace

replace github.com/nithinputhenveettil/opentelemetry-go/metric => ../metric

replace github.com/nithinputhenveettil/opentelemetry-go/sdk/metric => ./metric
