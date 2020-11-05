package yunke

import (
	`github.com/class100/core`
)

type (
	// Option 配置选项
	Option func(*options)

	options struct {
		class100.ClientOptions

		// Endpoint 服务端点
		Endpoint string
	}
)

func defaultOptions() options {
	return options{}
}

// WithSecret 配置应用授权
func WithSecret(secret class100.Secret) Option {
	return func(options *options) {
		options.ClientOptions.Secret = secret
	}
}

// WithEndpoint 配置服务端点
func WithEndpoint(endpoint string) Option {
	return func(options *options) {
		options.Endpoint = endpoint
	}
}
