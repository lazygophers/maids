package state

import (
	"github.com/lazygophers/log"
	"github.com/lazygophers/utils/config"
)

type Config struct {
	Name string `json:"name,omitempty" yaml:"name,omitempty" toml:"name,omitempty"`

	Port int    `json:"port,omitempty" yaml:"port,omitempty" toml:"port,omitempty"`
	Host string `json:"host,omitempty" yaml:"host,omitempty" toml:"host,omitempty"`

	// NOTE: Please fill in the configuration below
}

func LoadConfig() (err error) {
	err = config.LoadConfig(State.Config)
	if err != nil {
		log.Errorf("err:%v", err)
		return err
	}

	return nil
}
