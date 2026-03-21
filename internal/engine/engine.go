package engine

import (
	"github.com/kaluginivann/Aegis/internal/configs"
	"github.com/kaluginivann/Aegis/internal/files"
)

type Engine struct {
	conf *configs.Config
}

func NewEngine(conf *configs.Config) *Engine {
	return &Engine{conf: conf}
}

func (e *Engine) Run() {
	if err := files.CheckExistsFile(e.conf); err != nil {
		panic(err)
	}
}
