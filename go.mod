module github.com/tokane888/go-sandbox

go 1.23.8

require (
	github.com/joho/godotenv v1.5.1
	github.com/tokane888/go_common_module/v2 v2.0.4
	go.uber.org/zap v1.27.0
)

require go.uber.org/multierr v1.10.0 // indirect

// replace github.com/tokane888/go_common_module/v2 => ../go_common_module
