module go.opentelemetry.io/otel/trace

go 1.20

replace go.opentelemetry.io/otel => ../

require (
	github.com/google/go-cmp v0.6.0
	github.com/stretchr/testify v1.8.4
	go.opentelemetry.io/otel v1.24.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace go.opentelemetry.io/otel/metric => ../metric
