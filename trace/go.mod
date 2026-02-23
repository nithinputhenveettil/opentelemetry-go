module github.com/nithinputhenveettil/opentelemetry-go/trace

go 1.24.0

replace github.com/nithinputhenveettil/opentelemetry-go => ../

require (
	github.com/google/go-cmp v0.7.0
	github.com/stretchr/testify v1.11.1
	go.opentelemetry.io/otel v1.40.0
)

require (
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/nithinputhenveettil/opentelemetry-go/metric => ../metric
