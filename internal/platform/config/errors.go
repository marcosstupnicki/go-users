package config

type ErrorConfig struct {
	LowLevel error
	HumanReadable string
}

func (e ErrorConfig) Error() string {
	return e.HumanReadable
}

func (e ErrorConfig) Unwrap() error {
	return e.LowLevel
}

func NewErrorConfig(inner error, outer string) *ErrorConfig {
	return &ErrorConfig{
		LowLevel: inner,
		HumanReadable: outer,
	}
}

