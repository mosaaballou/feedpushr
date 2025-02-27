package config

import (
	"expvar"
)

var configMap = expvar.NewMap("config")

func exportConfigVar(key, value string) {
	configMap.Set(key, new(expvar.String))
	configMap.Get(key).(*expvar.String).Set(value)
}

// ExportConfigVars export some configuration variables to expvar
func ExportConfigVars() {
	conf := Get()
	exportConfigVar("addr", conf.ListenAddr)
	exportConfigVar("db", conf.DB)
	exportConfigVar("public-url", conf.PublicURL)
	exportConfigVar("delay", conf.Delay.String())
	exportConfigVar("timeout", conf.Timeout.String())
	exportConfigVar("cache-retention", conf.CacheRetention.String())
}
