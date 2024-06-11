package state

import (
	"github.com/lazygophers/log"
	"github.com/lazygophers/lrpc/middleware/i18n"
)

type state struct {
	Config *Config

	// NOTE: Please fill in the state below
	I18n *i18n.I18n
}

var State = new(state)

func Load() (err error) {
	err = LoadConfig()
	if err != nil {
	    log.Errorf("err:%v", err)
	    return err
	}

	err = LoadI18n()
	if err != nil {
		log.Errorf("err:%v", err)
		return err
	}

	return nil
}
