package confs

type GRPCConf struct {
	ListenConf `yaml:",inline"`
	Enabled    bool            `env:"ENABLED" envDefault:"true" yaml:"enabled,omitempty"`
	Gateway    GRPCGatewayConf `envPrefix:"GATEWAY_" yaml:"gateway,omitempty"`
}

type GRPCGatewayConf struct {
	ListenConf `yaml:",inline"`
	Enabled    bool   `env:"ENABLED" envDefault:"true" yaml:"enabled,omitempty"`
	Prefix     string `env:"PREFIX" yaml:"prefix,omitempty"`
}
