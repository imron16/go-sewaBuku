package server

import "github.com/imron16/go-sewaBuku/interface/container"

type handler struct {
	bukuHandler *bukuHandler
	sewaHandler *termConditionHandler
}

func setupHandler(container container.Container) *handler {
	bukuHandler := newBukuHanlder(container.BukuService)
	sewaHandler := newTermConditionHanlder(container.SewaService)
	return &handler{bukuHandler: bukuHandler,
		sewaHandler: sewaHandler}
}
