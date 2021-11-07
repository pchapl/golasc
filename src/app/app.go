package app

import (
	"golasc/src/api"
	"golasc/src/appconf"
)

func Run() {
	api.Run(appconf.Config.Port)
}
