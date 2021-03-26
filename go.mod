module github.com/ukinhappy/ukin

go 1.14

replace gopkg.in/resty.v1 => github.com/go-resty/resty/v2 v2.4.0

require (
	github.com/go-resty/resty/v2 v2.5.0 // indirect
	github.com/mitchellh/go-homedir v1.1.0
	github.com/spf13/cobra v1.1.1
	github.com/spf13/viper v1.7.1
	github.com/ukinhappy/go-utils v1.0.0
	gopkg.in/resty.v1 v1.12.0
)
