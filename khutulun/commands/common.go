package commands

import (
	"github.com/tliron/kutil/logging"
)

const toolName = "khutulun"

var log = logging.GetLogger(toolName)

var clusterName string
var pseudoTerminal bool
var unpack string
