package confs

type DebugConf struct {
	ListenConf `yaml:",inline"`
	Enabled    bool `env:"ENABLED" envDefault:"true" yaml:"enabled,omitempty"`
}
