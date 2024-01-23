package prefixmanager

import (
	"github.com/shatll-s/nexelliad/infrastructure/logger"
	"github.com/shatll-s/nexelliad/util/panics"
)

var log = logger.RegisterSubSystem("PRFX")
var spawn = panics.GoroutineWrapperFunc(log)
