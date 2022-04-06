package confs

import (
	"fmt"
	"net"
	"os"
	"path/filepath"

	"go.uber.org/multierr"

	"github.com/adrg/xdg"
)

//go:generate gomodifytags -file=conf.go -w -all -add-tags yaml -transform snakecase --skip-unexported -add-options yaml=omitempty

type ListenConf struct {
	Addr string `env:"ADDR" yaml:"addr,omitempty"`
	Port int    `env:"PORT" yaml:"port,omitempty"`
}

func (c ListenConf) GetAddr() string {
	addr := c.Addr
	if addr == "" && c.Port != 0 {
		addr = fmt.Sprintf(":%v", c.Port)
	}
	return addr
}

func (c ListenConf) Serve(ss ServeService) error {
	listener, err := c.Listen()
	if err != nil {
		return err
	}
	return ss.Serve(listener)
}

type ServeService interface {
	Serve(l net.Listener) error
}

func (c ListenConf) Listen() (net.Listener, error) {
	addr := c.GetAddr()
	if addr == "" {
		return nil, fmt.Errorf("no address or port")
	}
	return net.Listen("tcp", addr)
}

type DirConf struct {
	Name      string `yaml:"name,omitempty"` // application name
	RootDir   string `env:"ROOT_DIR" yaml:"root_dir,omitempty"`
	DataDir   string `env:"DATA_DIR" yaml:"data_dir,omitempty"`
	CacheDir  string `env:"CACHE_DIR" yaml:"cache_dir,omitempty"`
	ConfigDir string `env:"CONFIG_DIR" yaml:"config_dir,omitempty"`
}

func (conf *DirConf) SetDirEnv() error {
	return multierr.Combine(
		os.Setenv("DATA_DIR", conf.DataDir),
		os.Setenv("CACHE_DIR", conf.CacheDir),
		os.Setenv("CONFIG_DIR", conf.ConfigDir),
	)
}

func (conf *DirConf) InitDirConf(name string) {
	if conf.RootDir != "" {
		conf.DataDir = defaultTo(conf.DataDir, filepath.Join(conf.RootDir, "data"))
		conf.CacheDir = defaultTo(conf.CacheDir, filepath.Join(conf.RootDir, "cache"))
		conf.ConfigDir = defaultTo(conf.ConfigDir, filepath.Join(conf.RootDir, "config"))
	} else {
		conf.DataDir = defaultTo(conf.DataDir, filepath.Join(xdg.DataHome, name))
		conf.CacheDir = defaultTo(conf.CacheDir, filepath.Join(xdg.CacheHome, name))
		conf.ConfigDir = defaultTo(conf.ConfigDir, filepath.Join(xdg.ConfigHome, name))
	}
}

func defaultTo(v string, def string) string {
	if v == "" {
		return def
	}
	return v
}
