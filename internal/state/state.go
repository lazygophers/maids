package state

import (
	"github.com/lazygophers/log"
	"github.com/lazygophers/lrpc"
	"github.com/lazygophers/lrpc/middleware/service_discovery/ldiscovery"
	"github.com/lazygophers/lrpc/middleware/storage/etcd"
	"github.com/lazygophers/utils/app"
)

type state struct {
	Config *Config
	Etcd   *etcd.Client

	// NOTE: Please fill in the state below
	Watcher *Watcher
}

var State = new(state)

func Load() (err error) {
	log.SetPrefixMsg(app.Name)

	err = LoadConfig()
	if err != nil {
		log.Errorf("err:%v", err)
		return err
	}

	State.Etcd, err = etcd.ConnectWithLazy()
	if err != nil {
		log.Errorf("err:%v", err)
		return err
	}
	ldiscovery.SetEtcd(State.Etcd)

	State.Watcher = NewWatcher()

	return nil
}

func AfterListen(listen lrpc.ListenData) (err error) {
	err = State.Watcher.Start()
	if err != nil {
		log.Errorf("err:%v", err)
		return err
	}

	return nil
}
