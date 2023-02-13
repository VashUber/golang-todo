package infrastructure

type Size struct {
	W int
	H int
}

type Config struct {
	Size Size
}

func NewConfig() *Config {
	cfg := &Config{
		Size: Size{
			W: 450,
			H: 600,
		},
	}

	return cfg
}
