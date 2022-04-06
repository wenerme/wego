package confs

type LogConf struct {
	Level string `env:"LEVEL" envDefault:"info" yaml:"level,omitempty"`
}
