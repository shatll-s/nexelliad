package main

import (
	"github.com/Nexellia-Network/nexelliad/infrastructure/logger"
	"github.com/Nexellia-Network/nexelliad/util/panics"
)

var (
	backendLog = logger.NewBackend()
	log        = backendLog.Logger("RPIC")
	spawn      = panics.GoroutineWrapperFunc(log)
)
