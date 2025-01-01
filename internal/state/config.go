package state

import (
	"fmt"
	"github.com/lazygophers/log"
	"github.com/lazygophers/utils/app"
	"github.com/lazygophers/utils/config"
	"github.com/lazygophers/utils/runtime"
	"path/filepath"
)

type CfgWatchType uint8

const (
	CfgWatchTypeEtcd CfgWatchType = iota + 1
)

type CfgWatch struct {
	Name string       `json:"name,omitempty" yaml:"name,omitempty" toml:"name,omitempty"`
	Type CfgWatchType `json:"type,omitempty" yaml:"type,omitempty" toml:"type,omitempty"`

	Prefix    string `json:"prefix,omitempty" yaml:"prefix,omitempty" toml:"prefix,omitempty"`
	Directory string `json:"directory,omitempty" yaml:"directory,omitempty" toml:"directory,omitempty"`
}

type Config struct {
	Name string `json:"name,omitempty" yaml:"name,omitempty" toml:"name,omitempty"`

	Port int    `json:"port,omitempty" yaml:"port,omitempty" toml:"port,omitempty"`
	Host string `json:"host,omitempty" yaml:"host,omitempty" toml:"host,omitempty"`

	// NOTE: Please fill in the configuration below
	// 开启对服务发现的同步
	Watchers []*CfgWatch `json:"watchers,omitempty" yaml:"watchers,omitempty" toml:"watchers,omitempty"`
}

func (p *Config) apply() {
	for _, watcher := range p.Watchers {
		if watcher.Prefix == "" {
			if watcher.Name != "" {
				watcher.Prefix = fmt.Sprintf("/%s/%s/", app.Organization, watcher.Name)
				log.Infof("watcher prefix is empty,use name to set default:%s", watcher.Prefix)
			} else {
				log.Fatalf("watcher prefix is empty")
			}
		}

		if watcher.Directory == "" {
			if watcher.Name != "" {
				watcher.Directory = filepath.Join(runtime.LazyConfigDir(), watcher.Name)
				log.Infof("watcher directory is empty,use name to set default:%s", watcher.Directory)
			} else if watcher.Prefix == "" {
				watcher.Directory = filepath.Join(runtime.LazyConfigDir(), watcher.Prefix)
				log.Infof("watcher directory is empty,use prefix to set default:%s", watcher.Directory)
			} else {
				log.Fatalf("watcher directory is empty")
			}
		}
	}
}

func LoadConfig() (err error) {
	State.Config = new(Config)
	err = config.LoadConfig(State.Config)
	if err != nil {
		log.Errorf("err:%v", err)
		return err
	}

	if app.Name == "" {
		app.Name = State.Config.Name
		log.SetPrefixMsg(app.Name)
	}

	State.Config.apply()

	log.Info(State.Config.Watchers)

	return nil
}
