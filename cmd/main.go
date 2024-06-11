package main

import (
	"github.com/lazygophers/log"
	"github.com/lazygophers/lrpc"
	"github.com/lazygophers/maids/internal/state"
)

func main() {
	err := state.Load()
	if err != nil {
		log.Errorf("err:%v", err)
		return
	}

	app := lrpc.NewApp(&lrpc.Config{
		Name: state.State.Config.Name,
	})

	app.AddRoutes(Routes)

	err = app.ListenAndServe(state.State.Config.Port)
	if err != nil {
		log.Errorf("err:%v", err)
		return
	}
}
